package pebble

import (
	"testing"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/gobottest"
)

var _ gobot.Adaptor = (*Adaptor)(nil)

func initTestAdaptor() *Adaptor {
	return NewAdaptor()
}

func TestAdaptor(t *testing.T) {
	a := initTestAdaptor()
	gobottest.Assert(t, a.Name(), "Pebble")
}
func TestAdaptorConnect(t *testing.T) {
	a := initTestAdaptor()
	gobottest.Assert(t, len(a.Connect()), 0)
}

func TestAdaptorFinalize(t *testing.T) {
	a := initTestAdaptor()
	gobottest.Assert(t, len(a.Finalize()), 0)
}
