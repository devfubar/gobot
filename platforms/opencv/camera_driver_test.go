package opencv

import (
	"testing"
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/gobottest"
)

var _ gobot.Driver = (*CameraDriver)(nil)

func initTestCameraDriver() *CameraDriver {
	d := NewCameraDriver("")
	d.start = func(c *CameraDriver) (err error) {
		d.camera = &testCapture{}
		return nil
	}
	return d
}

func TestCameraDriver(t *testing.T) {
	d := initTestCameraDriver()
	gobottest.Assert(t, d.Name(), "Camera")
	gobottest.Assert(t, d.Connection(), (gobot.Connection)(nil))
}

func TestCameraDriverStart(t *testing.T) {
	sem := make(chan bool)
	d := initTestCameraDriver()
	gobottest.Assert(t, len(d.Start()), 0)
	d.On(d.Event("frame"), func(data interface{}) {
		sem <- true
	})
	select {
	case <-sem:
	case <-time.After(100 * time.Millisecond):
		t.Errorf("Event \"frame\" was not published")
	}

	d = NewCameraDriver("")
	gobottest.Assert(t, len(d.Start()), 0)

	d = NewCameraDriver(true)
	gobottest.Refute(t, len(d.Start()), 0)
}

func TestCameraDriverHalt(t *testing.T) {
	d := initTestCameraDriver()
	gobottest.Assert(t, len(d.Halt()), 0)
}
