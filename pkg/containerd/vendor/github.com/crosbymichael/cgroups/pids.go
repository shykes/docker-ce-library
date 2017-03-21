package cgroups

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	specs "github.com/opencontainers/runtime-spec/specs-go"
)

func NewPids(root string) *pidsController {
	return &pidsController{
		root: filepath.Join(root, string(Pids)),
	}
}

type pidsController struct {
	root string
}

func (p *pidsController) Name() Name {
	return Pids
}

func (p *pidsController) Path(path string) string {
	return filepath.Join(p.root, path)
}

func (p *pidsController) Create(path string, resources *specs.LinuxResources) error {
	if err := os.MkdirAll(p.Path(path), defaultDirPerm); err != nil {
		return err
	}
	if resources.Pids != nil && resources.Pids.Limit > 0 {
		return ioutil.WriteFile(
			filepath.Join(p.Path(path), "pids.max"),
			[]byte(strconv.FormatInt(resources.Pids.Limit, 10)),
			defaultFilePerm,
		)
	}
	return nil
}

func (p *pidsController) Update(path string, resources *specs.LinuxResources) error {
	return p.Create(path, resources)
}

func (p *pidsController) Stat(path string, stats *Stats) error {
	current, err := readUint(filepath.Join(p.Path(path), "pids.current"))
	if err != nil {
		return err
	}
	var max uint64
	maxData, err := ioutil.ReadFile(filepath.Join(p.Path(path), "pids.max"))
	if err != nil {
		return err
	}
	if maxS := strings.TrimSpace(string(maxData)); maxS != "max" {
		if max, err = parseUint(maxS, 10, 64); err != nil {
			return err
		}
	}
	stats.Pids = &PidsStat{
		Current: current,
		Limit:   max,
	}
	return nil
}
