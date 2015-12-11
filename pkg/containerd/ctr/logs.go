package main

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/docker/containerd"
)

var LogsCommand = cli.Command{
	Name:  "logs",
	Usage: "view binary container logs generated by containerd",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "follow,f",
			Usage: "follow/tail the logs",
		},
	},
	Action: func(context *cli.Context) {
		path := context.Args().First()
		if path == "" {
			fatal("path to the log cannot be empty", 1)
		}
		if err := readLogs(path, context.Bool("follow")); err != nil {
			fatal(err.Error(), 1)
		}
	},
}

func readLogs(path string, follow bool) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	for {
		var msg *containerd.Message
		if err := dec.Decode(&msg); err != nil {
			if err == io.EOF {
				if follow {
					time.Sleep(100 * time.Millisecond)
					continue
				}
				return nil
			}
			return err
		}
		switch msg.Stream {
		case "stdout":
			os.Stdout.Write(msg.Data)
		case "stderr":
			os.Stderr.Write(msg.Data)
		}
	}
	return nil
}
