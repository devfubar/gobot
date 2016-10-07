package joystick

import (
	"errors"
	"testing"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/gobottest"
)

var _ gobot.Adaptor = (*Adaptor)(nil)

func initTestAdaptor() *Adaptor {
	a := NewAdaptor()
	a.connect = func(j *Adaptor) (err error) {
		j.joystick = &testJoystick{}
		return nil
	}
	return a
}

func TestAdaptorConnect(t *testing.T) {
	a := initTestAdaptor()
	gobottest.Assert(t, len(a.Connect()), 0)

	a = NewAdaptor()
	gobottest.Assert(t, a.Connect()[0], errors.New("No joystick available"))
}

func TestAdaptorFinalize(t *testing.T) {
	a := initTestAdaptor()
	a.Connect()
	gobottest.Assert(t, len(a.Finalize()), 0)
}
