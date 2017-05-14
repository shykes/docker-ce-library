// +build darwin freebsd

package main

import (
	"os"

	"golang.org/x/sys/unix"

	"github.com/containerd/containerd/log"
	"github.com/containerd/containerd/reaper"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

const (
	defaultConfigPath = "/etc/containerd/config.toml"
)

var (
	handledSignals = []os.Signal{unix.SIGTERM, unix.SIGINT, unix.SIGUSR1, unix.SIGCHLD}
)

func platformInit(context *cli.Context) error {
	if err := os.MkdirAll(conf.State, 0750); err != nil {
		return err
	}
	if err := os.Chown(conf.State, conf.GRPC.Uid, conf.GRPC.Gid); err != nil {
		return err
	}
	return nil
}

func handleSignals(signals chan os.Signal, server *grpc.Server) error {
	for s := range signals {
		log.G(global).WithField("signal", s).Debug("received signal")
		switch s {
		case unix.SIGCHLD:
			if err := reaper.Reap(); err != nil {
				log.G(global).WithError(err).Error("reap containerd processes")
			}
		default:
			server.Stop()
			return nil
		}
	}
	return nil
}
