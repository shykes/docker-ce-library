// +build linux

package overlay

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/log"
	"github.com/containerd/containerd/plugin"
	"github.com/containerd/containerd/snapshot"
	"github.com/containerd/containerd/snapshot/storage"
	"github.com/pkg/errors"
)

func init() {
	plugin.Register("snapshot-overlay", &plugin.Registration{
		Type: plugin.SnapshotPlugin,
		Init: func(ic *plugin.InitContext) (interface{}, error) {
			return NewSnapshotter(filepath.Join(ic.Root, "snapshot", "overlay"))
		},
	})
}

type snapshotter struct {
	root string
	ms   *storage.MetaStore
}

type activeSnapshot struct {
	id       string
	name     string
	parentID interface{}
	readonly bool
}

// NewSnapshotter returns a Snapshotter which uses overlayfs. The overlayfs
// diffs are stored under the provided root. A metadata file is stored under
// the root.
func NewSnapshotter(root string) (snapshot.Snapshotter, error) {
	if err := os.MkdirAll(root, 0700); err != nil {
		return nil, err
	}
	ms, err := storage.NewMetaStore(filepath.Join(root, "metadata.db"))
	if err != nil {
		return nil, err
	}

	if err := os.Mkdir(filepath.Join(root, "snapshots"), 0700); err != nil && !os.IsExist(err) {
		return nil, err
	}

	return &snapshotter{
		root: root,
		ms:   ms,
	}, nil
}

// Stat returns the info for an active or committed snapshot by name or
// key.
//
// Should be used for parent resolution, existence checks and to discern
// the kind of snapshot.
func (o *snapshotter) Stat(ctx context.Context, key string) (snapshot.Info, error) {
	ctx, t, err := o.ms.TransactionContext(ctx, false)
	if err != nil {
		return snapshot.Info{}, err
	}
	defer t.Rollback()
	return storage.GetInfo(ctx, key)
}

func (o *snapshotter) Prepare(ctx context.Context, key, parent string) ([]containerd.Mount, error) {
	return o.createActive(ctx, key, parent, false)
}

func (o *snapshotter) View(ctx context.Context, key, parent string) ([]containerd.Mount, error) {
	return o.createActive(ctx, key, parent, true)
}

// Mounts returns the mounts for the transaction identified by key. Can be
// called on an read-write or readonly transaction.
//
// This can be used to recover mounts after calling View or Prepare.
func (o *snapshotter) Mounts(ctx context.Context, key string) ([]containerd.Mount, error) {
	ctx, t, err := o.ms.TransactionContext(ctx, false)
	if err != nil {
		return nil, err
	}
	active, err := storage.GetActive(ctx, key)
	t.Rollback()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get active mount")
	}
	return o.mounts(active), nil
}

func (o *snapshotter) Commit(ctx context.Context, name, key string) error {
	ctx, t, err := o.ms.TransactionContext(ctx, true)
	if err != nil {
		return err
	}
	if _, err := storage.CommitActive(ctx, key, name); err != nil {
		if rerr := t.Rollback(); rerr != nil {
			log.G(ctx).WithError(rerr).Warn("Failure rolling back transaction")
		}
		return errors.Wrap(err, "failed to commit snapshot")
	}
	return t.Commit()
}

// Remove abandons the transaction identified by key. All resources
// associated with the key will be removed.
func (o *snapshotter) Remove(ctx context.Context, key string) (err error) {
	ctx, t, err := o.ms.TransactionContext(ctx, true)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil && t != nil {
			if rerr := t.Rollback(); rerr != nil {
				log.G(ctx).WithError(rerr).Warn("Failure rolling back transaction")
			}
		}
	}()

	id, _, err := storage.Remove(ctx, key)
	if err != nil {
		return errors.Wrap(err, "failed to remove")
	}

	path := filepath.Join(o.root, "snapshots", id)
	renamed := filepath.Join(o.root, "snapshots", "rm-"+id)
	if err := os.Rename(path, renamed); err != nil {
		return errors.Wrap(err, "failed to rename")
	}

	err = t.Commit()
	t = nil
	if err != nil {
		if err1 := os.Rename(renamed, path); err1 != nil {
			// May cause inconsistent data on disk
			log.G(ctx).WithError(err1).WithField("path", renamed).Errorf("Failed to rename after failed commit")
		}
		return errors.Wrap(err, "failed to commit")
	}
	if err := os.RemoveAll(renamed); err != nil {
		// Must be cleaned up, any "rm-*" could be removed if no active transactions
		log.G(ctx).WithError(err).WithField("path", renamed).Warnf("Failed to remove root filesystem")
	}

	return nil
}

// Walk the committed snapshots.
func (o *snapshotter) Walk(ctx context.Context, fn func(context.Context, snapshot.Info) error) error {
	ctx, t, err := o.ms.TransactionContext(ctx, false)
	if err != nil {
		return err
	}
	defer t.Rollback()
	return storage.WalkInfo(ctx, fn)
}

func (o *snapshotter) createActive(ctx context.Context, key, parent string, readonly bool) ([]containerd.Mount, error) {
	var (
		path        string
		snapshotDir = filepath.Join(o.root, "snapshots")
	)

	td, err := ioutil.TempDir(snapshotDir, "new-")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create temp dir")
	}
	defer func() {
		if err != nil {
			if td != "" {
				if err1 := os.RemoveAll(td); err1 != nil {
					err = errors.Wrapf(err, "remove failed: %v", err1)
				}
			}
			if path != "" {
				if err1 := os.RemoveAll(path); err1 != nil {
					err = errors.Wrapf(err, "failed to remove path: %v", err1)
				}
			}
		}
	}()

	if err = os.MkdirAll(filepath.Join(td, "fs"), 0711); err != nil {
		return nil, err
	}
	if !readonly {
		if err = os.MkdirAll(filepath.Join(td, "work"), 0700); err != nil {
			return nil, err
		}
	}

	ctx, t, err := o.ms.TransactionContext(ctx, true)
	if err != nil {
		return nil, err
	}

	active, err := storage.CreateActive(ctx, key, parent, readonly)
	if err != nil {
		if rerr := t.Rollback(); rerr != nil {
			log.G(ctx).WithError(rerr).Warn("Failure rolling back transaction")
		}
		return nil, errors.Wrap(err, "failed to create active")
	}

	path = filepath.Join(snapshotDir, active.ID)
	if err = os.Rename(td, path); err != nil {
		if rerr := t.Rollback(); rerr != nil {
			log.G(ctx).WithError(rerr).Warn("Failure rolling back transaction")
		}
		return nil, errors.Wrap(err, "failed to rename")
	}
	td = ""

	if err = t.Commit(); err != nil {
		return nil, errors.Wrap(err, "commit failed")
	}

	return o.mounts(active), nil
}

func (o *snapshotter) mounts(active storage.Active) []containerd.Mount {
	if len(active.ParentIDs) == 0 {
		// if we only have one layer/no parents then just return a bind mount as overlay
		// will not work
		roFlag := "rw"
		if active.Readonly {
			roFlag = "ro"
		}

		return []containerd.Mount{
			{
				Source: o.upperPath(active.ID),
				Type:   "bind",
				Options: []string{
					roFlag,
					"rbind",
				},
			},
		}
	}
	var options []string

	if !active.Readonly {
		options = append(options,
			fmt.Sprintf("workdir=%s", o.workPath(active.ID)),
			fmt.Sprintf("upperdir=%s", o.upperPath(active.ID)),
		)
	} else if len(active.ParentIDs) == 1 {
		return []containerd.Mount{
			{
				Source: o.upperPath(active.ParentIDs[0]),
				Type:   "bind",
				Options: []string{
					"ro",
					"rbind",
				},
			},
		}
	}

	parentPaths := make([]string, len(active.ParentIDs))
	for i := range active.ParentIDs {
		parentPaths[i] = o.upperPath(active.ParentIDs[i])
	}

	options = append(options, fmt.Sprintf("lowerdir=%s", strings.Join(parentPaths, ":")))
	return []containerd.Mount{
		{
			Type:    "overlay",
			Source:  "overlay",
			Options: options,
		},
	}

}

func (o *snapshotter) upperPath(id string) string {
	return filepath.Join(o.root, "snapshots", id, "fs")
}

func (o *snapshotter) workPath(id string) string {
	return filepath.Join(o.root, "snapshots", id, "work")
}
