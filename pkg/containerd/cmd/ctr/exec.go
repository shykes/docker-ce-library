package main

import (
	gocontext "context"
	"os"

	"github.com/containerd/containerd/api/services/execution"
	"github.com/crosbymichael/console"
	"github.com/urfave/cli"
)

var execCommand = cli.Command{
	Name:  "exec",
	Usage: "execute additional processes in an existing container",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "id",
			Usage: "id of the container",
		},
		cli.StringFlag{
			Name:  "cwd",
			Usage: "working directory of the new process",
		},
		cli.BoolFlag{
			Name:  "tty,t",
			Usage: "allocate a TTY for the container",
		},
	},
	Action: func(context *cli.Context) error {
		var (
			id  = context.String("id")
			ctx = gocontext.Background()
		)

		containers, err := getExecutionService(context)
		if err != nil {
			return err
		}
		events, err := containers.Events(ctx, &execution.EventsRequest{})
		if err != nil {
			return err
		}
		tmpDir, err := getTempDir(id)
		if err != nil {
			return err
		}
		defer os.RemoveAll(tmpDir)
		request, err := newExecRequest(context, tmpDir, id)
		if err != nil {
			return err
		}
		if request.Terminal {
			con := console.Current()
			defer con.Reset()
			if err := con.SetRaw(); err != nil {
				return err
			}
		}
		fwg, err := prepareStdio(request.Stdin, request.Stdout, request.Stderr, request.Terminal)
		if err != nil {
			return err
		}
		response, err := containers.Exec(ctx, request)
		if err != nil {
			return err
		}
		// Ensure we read all io only if container started successfully.
		defer fwg.Wait()

		status, err := waitContainer(events, id, response.Pid)
		if err != nil {
			return err
		}
		if status != 0 {
			return cli.NewExitError("", int(status))
		}
		return nil
	},
}
