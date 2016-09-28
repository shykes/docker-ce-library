package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/docker/containerkit"
	"github.com/docker/containerkit/osutils"
	"github.com/docker/containerkit/runc"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

var RWM = "rwm"

// "Hooks do optional work. Drivers do mandatory work"
func main() {
	if err := osutils.SetSubreaper(1); err != nil {
		logrus.Fatal(err)
	}
	if err := runTest(); err != nil {
		logrus.Fatal(err)
	}
}

func runTest() error {
	// create a new runc runtime that implements the ExecutionDriver interface
	driver, err := runc.New("/run/runc", "/tmp/runc")
	if err != nil {
		return err
	}
	// create a new container
	container, err := containerkit.NewContainer(
		"/var/lib/containerkit", /* container root */
		"test",                  /* container id */
		containerkit.Mount{
			Type:   "bind",
			Source: "/containers/redis/rootfs",
			Options: []string{
				"rbind",
				"rw",
			},
		}, /* mount from the graph subsystem for the container */
		spec("test"), /* the spec for the container */
		driver,       /* the exec driver to use for the container */
	)
	if err != nil {
		return err
	}
	// setup some stdio for our container
	container.Stdin = os.Stdin
	container.Stdout = os.Stdout
	container.Stderr = os.Stderr

	// go ahead and set the container in the create state and have it ready to start
	logrus.Info("create container")
	if err := container.Create(); err != nil {
		return err
	}

	// start the user defined process in the container
	logrus.Info("start container")
	if err := container.Start(); err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		process, err := container.NewProcess(&specs.Process{
			Args: []string{
				"echo", fmt.Sprintf("sup from itteration %d", i),
			},
			Env:             []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			Terminal:        false,
			Cwd:             "/",
			NoNewPrivileges: true,
			Capabilities: []string{
				"CAP_AUDIT_WRITE",
				"CAP_KILL",
				"CAP_FOWNER",
				"CAP_CHOWN",
				"CAP_MKNOD",
				"CAP_FSETID",
				"CAP_DAC_OVERRIDE",
				"CAP_SETFCAP",
				"CAP_SETPCAP",
				"CAP_SETGID",
				"CAP_SETUID",
				"CAP_NET_BIND_SERVICE",
			},
		})
		process.Stdin = os.Stdin
		process.Stdout = os.Stdout
		process.Stderr = os.Stderr
		if err := process.Start(); err != nil {
			return err
		}
		procStatus, err := process.Wait()
		if err != nil {
			return err
		}
		logrus.Infof("process %d returned with %d", i, procStatus)
	}

	// wait for it to exit and get the exit status
	logrus.Info("wait container")
	status, err := container.Wait()
	if err != nil {
		return err
	}

	// delete the container after it is done
	logrus.Info("delete container")
	if container.Delete(); err != nil {
		return err
	}
	logrus.Infof("exit status %d", status)
	return nil
}

// bla bla bla spec stuff
func spec(id string) *specs.Spec {
	cgpath := filepath.Join("/containerkit", id)
	return &specs.Spec{
		Version: specs.Version,
		Platform: specs.Platform{
			OS:   runtime.GOOS,
			Arch: runtime.GOARCH,
		},
		Root: specs.Root{
			Path:     "rootfs",
			Readonly: false,
		},
		Process: specs.Process{
			Env:             []string{"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"},
			Args:            []string{"sleep", "10"},
			Terminal:        false,
			Cwd:             "/",
			NoNewPrivileges: true,
			Capabilities: []string{
				"CAP_AUDIT_WRITE",
				"CAP_KILL",
				"CAP_FOWNER",
				"CAP_CHOWN",
				"CAP_MKNOD",
				"CAP_FSETID",
				"CAP_DAC_OVERRIDE",
				"CAP_SETFCAP",
				"CAP_SETPCAP",
				"CAP_SETGID",
				"CAP_SETUID",
				"CAP_NET_BIND_SERVICE",
			},
		},
		Hostname: "containerkit",
		Mounts: []specs.Mount{
			{
				Destination: "/proc",
				Type:        "proc",
				Source:      "proc",
			},
			{
				Destination: "/dev",
				Type:        "tmpfs",
				Source:      "tmpfs",
				Options:     []string{"nosuid", "strictatime", "mode=755", "size=65536k"},
			},
			{
				Destination: "/dev/pts",
				Type:        "devpts",
				Source:      "devpts",
				Options:     []string{"nosuid", "noexec", "newinstance", "ptmxmode=0666", "mode=0620", "gid=5"},
			},
			{
				Destination: "/dev/shm",
				Type:        "tmpfs",
				Source:      "shm",
				Options:     []string{"nosuid", "noexec", "nodev", "mode=1777", "size=65536k"},
			},
			{
				Destination: "/dev/mqueue",
				Type:        "mqueue",
				Source:      "mqueue",
				Options:     []string{"nosuid", "noexec", "nodev"},
			},
			{
				Destination: "/sys",
				Type:        "sysfs",
				Source:      "sysfs",
				Options:     []string{"nosuid", "noexec", "nodev"},
			},
			{
				Destination: "/run",
				Type:        "tmpfs",
				Source:      "tmpfs",
				Options:     []string{"nosuid", "strictatime", "mode=755", "size=65536k"},
			},
			{
				Destination: "/etc/resolv.conf",
				Type:        "bind",
				Source:      "/etc/resolv.conf",
				Options:     []string{"rbind", "ro"},
			},
			{
				Destination: "/etc/hosts",
				Type:        "bind",
				Source:      "/etc/hosts",
				Options:     []string{"rbind", "ro"},
			},
			{
				Destination: "/etc/localtime",
				Type:        "bind",
				Source:      "/etc/localtime",
				Options:     []string{"rbind", "ro"},
			},
		},
		Linux: &specs.Linux{
			CgroupsPath: &cgpath,
			Resources: &specs.Resources{
				Devices: []specs.DeviceCgroup{
					{
						Allow:  false,
						Access: &RWM,
					},
				},
			},
			Namespaces: []specs.Namespace{
				{
					Type: "pid",
				},
				{
					Type: "ipc",
				},
				{
					Type: "uts",
				},
				{
					Type: "mount",
				},
			},
		},
	}

}
