package console

// #include <termios.h>
import "C"

import (
	"os"
	"syscall"
	"unsafe"
)

// NewPty creates a new pty pair
// The master is returned as the first console and a string
// with the path to the pty slave is returned as the second
func NewPty() (Console, string, error) {
	f, err := os.OpenFile("/dev/ptmx", syscall.O_RDWR|syscall.O_NOCTTY|syscall.O_CLOEXEC, 0)
	if err != nil {
		return nil, "", err
	}
	if err := saneTerminal(f); err != nil {
		return nil, "", err
	}
	slave, err := ptsname(f)
	if err != nil {
		return nil, "", err
	}
	if err := unlockpt(f); err != nil {
		return nil, "", err
	}
	return &master{
		f: f,
	}, slave, nil
}

type master struct {
	f       *os.File
	termios *syscall.Termios
}

func (m *master) Read(b []byte) (int, error) {
	return m.f.Read(b)
}

func (m *master) Write(b []byte) (int, error) {
	return m.f.Write(b)
}

func (m *master) Close() error {
	return m.f.Close()
}

func (m *master) Resize(ws WinSize) error {
	return ioctl(
		m.f.Fd(),
		uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&ws)),
	)
}

func (m *master) ResizeFrom(c Console) error {
	ws, err := c.Size()
	if err != nil {
		return err
	}
	return m.Resize(ws)
}

func (m *master) Reset() error {
	if m.termios == nil {
		return nil
	}
	return tcset(m.f.Fd(), m.termios)
}

func (m *master) SetRaw() error {
	m.termios = &syscall.Termios{}
	if err := tcget(m.f.Fd(), m.termios); err != nil {
		return err
	}
	rawState := *m.termios
	C.cfmakeraw((*C.struct_termios)(unsafe.Pointer(&rawState)))
	rawState.Oflag = rawState.Oflag | C.OPOST
	return tcset(m.f.Fd(), &rawState)
}

func (m *master) Size() (WinSize, error) {
	var ws WinSize
	if err := ioctl(
		m.f.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&ws)),
	); err != nil {
		return ws, err
	}
	return ws, nil
}

// checkConsole checks if the provided file is a console
func checkConsole(f *os.File) error {
	var termios syscall.Termios
	if tcget(f.Fd(), &termios) != nil {
		return ErrNotAConsole
	}
	return nil
}
