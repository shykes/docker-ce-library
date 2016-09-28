package runc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/docker/containerkit"
)

func New(root, log string) (*Runc, error) {
	if err := os.MkdirAll(root, 0711); err != nil {
		return nil, err
	}
	return &Runc{
		root: root,
		log:  log,
	}, nil
}

type Runc struct {
	root string
	log  string
}

func (r *Runc) Create(c *containerkit.Container) (containerkit.ProcessDelegate, error) {
	pidFile := fmt.Sprintf("%s/%s.pid", filepath.Join(r.root, c.ID()), "init")
	cmd := r.command("create", "--pid-file", pidFile, "--bundle", c.Path(), c.ID())
	cmd.Stdin, cmd.Stdout, cmd.Stderr = c.Stdin, c.Stdout, c.Stderr
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return nil, err
	}
	i, err := strconv.Atoi(string(data))
	if err != nil {
		return nil, err
	}
	return newProcess(i)
}

func (r *Runc) Start(c *containerkit.Container) error {
	return r.command("start", c.ID()).Run()
}

func (r *Runc) Delete(c *containerkit.Container) error {
	return r.command("delete", c.ID()).Run()
}

func (r *Runc) Exec(c *containerkit.Container, p *containerkit.Process) (containerkit.ProcessDelegate, error) {
	f, err := ioutil.TempFile(filepath.Join(r.root, c.ID()), "process")
	if err != nil {
		return nil, err
	}
	path := f.Name()
	pidFile := fmt.Sprintf("%s/%s.pid", filepath.Join(r.root, c.ID()), filepath.Base(path))
	err = json.NewEncoder(f).Encode(p.Spec())
	f.Close()
	if err != nil {
		return nil, err
	}
	cmd := r.command("exec", "--process", path, "--pid-file", pidFile, c.ID())
	cmd.Stdin, cmd.Stdout, cmd.Stderr = p.Stdin, p.Stdout, p.Stderr
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return nil, err
	}
	i, err := strconv.Atoi(string(data))
	if err != nil {
		return nil, err
	}
	return newProcess(i)
}

func (r *Runc) command(args ...string) *exec.Cmd {
	return exec.Command("runc", append([]string{
		"--root", r.root,
		"--log", r.log,
	}, args...)...)
}

func newProcess(pid int) (*process, error) {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}
	return &process{
		proc: proc,
	}, nil
}

type process struct {
	proc *os.Process
}

func (p *process) Pid() int {
	return p.proc.Pid
}

func (p *process) Wait() (uint32, error) {
	state, err := p.proc.Wait()
	if err != nil {
		return 0, nil
	}
	return uint32(state.Sys().(syscall.WaitStatus).ExitStatus()), nil
}

func (p *process) Signal(s os.Signal) error {
	return p.proc.Signal(s)
}
