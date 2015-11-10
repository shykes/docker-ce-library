package containerd

import (
	"os"
	"time"
)

type EventType string

const (
	ExitEventType                EventType = "exit"
	StartContainerEventType      EventType = "startContainer"
	ContainerStartErrorEventType EventType = "startContainerError"
	GetContainerEventType        EventType = "getContainer"
	SignalEventType              EventType = "signal"
)

func NewEvent(t EventType) *Event {
	return &Event{
		Type:      t,
		Timestamp: time.Now(),
		Err:       make(chan error, 1),
	}
}

type Event struct {
	Type       EventType   `json:"type"`
	Timestamp  time.Time   `json:"timestamp"`
	ID         string      `json:"id,omitempty"`
	BundlePath string      `json:"bundlePath,omitempty"`
	Pid        int         `json:"pid,omitempty"`
	Status     int         `json:"status,omitempty"`
	Signal     os.Signal   `json:"signal,omitempty"`
	Containers []Container `json:"-"`
	Err        chan error  `json:"-"`
}
