package main

import (
	contextpkg "context"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/docker/containerd"
	"github.com/urfave/cli"
)

var (
	background = contextpkg.Background()
)

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(os.Args[0], containerd.Package, containerd.Version)
	}
	app := cli.NewApp()
	app.Name = "dist"
	app.Version = containerd.Version
	app.Usage = `
        ___      __
   ____/ (_)____/ /_
  / __  / / ___/ __/
 / /_/ / (__  ) /_
 \__,_/_/____/\__/  

distribution tool
`
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug output in logs",
		},
		cli.DurationFlag{
			Name:   "timeout",
			Usage:  "total timeout for fetch",
			EnvVar: "CONTAINERD_FETCH_TIMEOUT",
		},
		cli.StringFlag{
			// TODO(stevvooe): for now, we allow circumventing the GRPC. Once
			// we have clear separation, this will likely go away.
			Name:  "root",
			Usage: "path to content store root",
			Value: "/var/lib/containerd",
		},
		cli.StringFlag{
			Name:  "socket, s",
			Usage: "socket path for containerd's GRPC server",
			Value: "/run/containerd/containerd.sock",
		},
	}
	app.Commands = []cli.Command{
		fetchCommand,
		ingestCommand,
		activeCommand,
		getCommand,
		deleteCommand,
		listCommand,
		applyCommand,
	}
	app.Before = func(context *cli.Context) error {
		var (
			debug   = context.GlobalBool("debug")
			timeout = context.GlobalDuration("timeout")
		)
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		if timeout > 0 {
			background, _ = contextpkg.WithTimeout(background, timeout)
		}
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "dist: %s\n", err)
		os.Exit(1)
	}
}
