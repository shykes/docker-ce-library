package supervisor

import (
	"time"

	"github.com/docker/containerd/runtime"
	"github.com/docker/containerd/specs"
)

// AddProcessTask holds everything necessary to add a process to a
// container
type AddProcessTask struct {
	baseTask
	ID            string
	PID           string
	Stdout        string
	Stderr        string
	Stdin         string
	ProcessSpec   *specs.ProcessSpec
	StartResponse chan StartResponse
}

func (s *Supervisor) addProcess(t *AddProcessTask) error {
	ci, ok := s.containers[t.ID]
	if !ok {
		return ErrContainerNotFound
	}
	process, err := ci.container.Exec(t.PID, *t.ProcessSpec, runtime.NewStdio(t.Stdin, t.Stdout, t.Stderr))
	if err != nil {
		return err
	}
	if err := s.monitor.Add(process); err != nil {
		return err
	}
	t.StartResponse <- StartResponse{}
	s.notifySubscribers(Event{
		Timestamp: time.Now(),
		Type:      StateStartProcess,
		PID:       t.PID,
		ID:        t.ID,
	})
	return nil
}
