package containerd

import (
	"os"
	"path/filepath"
	goruntime "runtime"

	"github.com/Sirupsen/logrus"
	"github.com/docker/containerd/runtime"
	"github.com/opencontainers/runc/libcontainer"
)

// NewSupervisor returns an initialized Process supervisor.
func NewSupervisor(stateDir string, tasks chan *StartTask) (*Supervisor, error) {
	if err := os.MkdirAll(stateDir, 0755); err != nil {
		return nil, err
	}
	// register counters
	r, err := newRuntime(stateDir)
	if err != nil {
		return nil, err
	}
	j, err := newJournal(filepath.Join(stateDir, "journal.json"))
	if err != nil {
		return nil, err
	}
	s := &Supervisor{
		stateDir:   stateDir,
		containers: make(map[string]runtime.Container),
		processes:  make(map[int]runtime.Container),
		runtime:    r,
		journal:    j,
		tasks:      tasks,
	}
	// register default event handlers
	s.handlers = map[EventType]Handler{
		ExecExitEventType:        &ExecExitEvent{s},
		ExitEventType:            &ExitEvent{s},
		StartContainerEventType:  &StartEvent{s},
		DeleteEventType:          &DeleteEvent{s},
		GetContainerEventType:    &GetContainersEvent{s},
		SignalEventType:          &SignalEvent{s},
		AddProcessEventType:      &AddProcessEvent{s},
		UpdateContainerEventType: &UpdateEvent{s},
	}
	// start the container workers for concurrent container starts
	return s, nil
}

type Supervisor struct {
	// stateDir is the directory on the system to store container runtime state information.
	stateDir    string
	containers  map[string]runtime.Container
	processes   map[int]runtime.Container
	handlers    map[EventType]Handler
	runtime     runtime.Runtime
	journal     *journal
	events      chan *Event
	tasks       chan *StartTask
	subscribers map[subscriber]bool
}

type subscriber chan *Event

// need proper close logic for jobs and stuff so that sending to the channels dont panic
// but can complete jobs
func (s *Supervisor) Close() error {
	//TODO: unsubscribe all channels
	return s.journal.Close()
}

func (s *Supervisor) Events() subscriber {
	return subscriber(make(chan *Event))
}

func (s *Supervisor) Unsubscribe(sub subscriber) {
	delete(s.subscribers, sub)
}

func (s *Supervisor) NotifySubscribers(e *Event) {
	for sub := range s.subscribers {
		sub <- e
	}
}

// Start is a non-blocking call that runs the supervisor for monitoring contianer processes and
// executing new containers.
//
// This event loop is the only thing that is allowed to modify state of containers and processes.
func (s *Supervisor) Start(events chan *Event) error {
	if events == nil {
		return ErrEventChanNil
	}
	s.events = events
	go func() {
		// allocate an entire thread to this goroutine for the main event loop
		// so that nothing else is scheduled over the top of it.
		goruntime.LockOSThread()
		for e := range events {
			EventsCounter.Inc(1)
			s.journal.write(e)
			h, ok := s.handlers[e.Type]
			if !ok {
				e.Err <- ErrUnknownEvent
				continue
			}
			if err := h.Handle(e); err != nil {
				if err != errDeferedResponse {
					e.Err <- err
					close(e.Err)
				}
				continue
			}
			close(e.Err)
		}
	}()
	return nil
}

func (s *Supervisor) getContainerForPid(pid int) (runtime.Container, error) {
	for _, container := range s.containers {
		cpid, err := container.Pid()
		if err != nil {
			if lerr, ok := err.(libcontainer.Error); ok {
				if lerr.Code() == libcontainer.ProcessNotExecuted {
					continue
				}
			}
			logrus.WithField("error", err).Error("containerd: get container pid")
		}
		if pid == cpid {
			return container, nil
		}
	}
	return nil, errNoContainerForPid
}

func (s *Supervisor) SendEvent(evt *Event) {
	s.events <- evt
}
