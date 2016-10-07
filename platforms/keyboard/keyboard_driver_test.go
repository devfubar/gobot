package keyboard

import (
	"os"
	"testing"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/gobottest"
)

var _ gobot.Driver = (*Driver)(nil)

func initTestKeyboardDriver() *Driver {
	d := NewDriver()
	d.connect = func(k *Driver) (err error) {
		k.stdin = &os.File{}
		return nil
	}
	d.listen = func(k *Driver) {}
	return d
}

func TestKeyboardDriver(t *testing.T) {
	d := initTestKeyboardDriver()
	gobottest.Assert(t, d.Connection(), (gobot.Connection)(nil))
}

func TestKeyboardDriverStart(t *testing.T) {
	d := initTestKeyboardDriver()
	gobottest.Assert(t, len(d.Start()), 0)
}

func TestKeyboardDriverHalt(t *testing.T) {
	d := initTestKeyboardDriver()
	gobottest.Assert(t, len(d.Halt()), 0)
}
