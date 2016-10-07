package ardrone

import (
	"testing"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/gobottest"
)

var _ gobot.Driver = (*Driver)(nil)

func initTestArdroneDriver() *Driver {
	a := NewAdaptor()
	a.connect = func(a *Adaptor) (drone, error) {
		return &testDrone{}, nil
	}
	d := NewDriver(a)
	d.SetName("mydrone")
	a.Connect()
	return d
}

func TestArdroneDriver(t *testing.T) {
	d := initTestArdroneDriver()
	gobottest.Assert(t, d.Name(), "mydrone")
}

func TestArdroneDriverStart(t *testing.T) {
	d := initTestArdroneDriver()
	gobottest.Assert(t, len(d.Start()), 0)
}

func TestArdroneDriverHalt(t *testing.T) {
	d := initTestArdroneDriver()
	gobottest.Assert(t, len(d.Halt()), 0)
}
func TestArdroneDriverTakeOff(t *testing.T) {
	d := initTestArdroneDriver()
	d.TakeOff()
}

func TestArdroneDriverand(t *testing.T) {
	d := initTestArdroneDriver()
	d.Land()
}

func TestArdroneDriverUp(t *testing.T) {
	d := initTestArdroneDriver()
	d.Up(1)
}

func TestArdroneDriverDown(t *testing.T) {
	d := initTestArdroneDriver()
	d.Down(1)
}

func TestArdroneDriverLeft(t *testing.T) {
	d := initTestArdroneDriver()
	d.Left(1)
}

func TestArdroneDriverRight(t *testing.T) {
	d := initTestArdroneDriver()
	d.Right(1)
}

func TestArdroneDriverForward(t *testing.T) {
	d := initTestArdroneDriver()
	d.Forward(1)
}

func TestArdroneDriverackward(t *testing.T) {
	d := initTestArdroneDriver()
	d.Backward(1)
}

func TestArdroneDriverClockwise(t *testing.T) {
	d := initTestArdroneDriver()
	d.Clockwise(1)
}

func TestArdroneDriverCounterClockwise(t *testing.T) {
	d := initTestArdroneDriver()
	d.CounterClockwise(1)
}

func TestArdroneDriverHover(t *testing.T) {
	d := initTestArdroneDriver()
	d.Hover()
}
