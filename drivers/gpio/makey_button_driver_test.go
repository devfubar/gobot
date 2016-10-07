package gpio

import (
	"errors"
	"testing"
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/gobottest"
)

var _ gobot.Driver = (*MakeyButtonDriver)(nil)

const MAKEY_TEST_DELAY = 30

func initTestMakeyButtonDriver() *MakeyButtonDriver {
	return NewMakeyButtonDriver(newGpioTestAdaptor(), "1")
}

func TestMakeyButtonDriverHalt(t *testing.T) {
	d := initTestMakeyButtonDriver()
	go func() {
		<-d.halt
	}()
	gobottest.Assert(t, len(d.Halt()), 0)
}

func TestMakeyButtonDriver(t *testing.T) {
	d := NewMakeyButtonDriver(newGpioTestAdaptor(), "1")
	gobottest.Assert(t, d.Pin(), "1")
	gobottest.Refute(t, d.Connection(), nil)

	d = NewMakeyButtonDriver(newGpioTestAdaptor(), "1", 30*time.Second)
	gobottest.Assert(t, d.interval, MAKEY_TEST_DELAY*time.Second)
}

func TestMakeyButtonDriverStart(t *testing.T) {
	sem := make(chan bool, 0)
	d := initTestMakeyButtonDriver()
	gobottest.Assert(t, len(d.Start()), 0)

	testAdaptorDigitalRead = func() (val int, err error) {
		val = 0
		return
	}

	d.Once(ButtonPush, func(data interface{}) {
		gobottest.Assert(t, d.Active, true)
		sem <- true
	})

	select {
	case <-sem:
	case <-time.After(MAKEY_TEST_DELAY * time.Millisecond):
		t.Errorf("MakeyButton Event \"Push\" was not published")
	}

	testAdaptorDigitalRead = func() (val int, err error) {
		val = 1
		return
	}

	d.Once(ButtonRelease, func(data interface{}) {
		gobottest.Assert(t, d.Active, false)
		sem <- true
	})

	select {
	case <-sem:
	case <-time.After(MAKEY_TEST_DELAY * time.Millisecond):
		t.Errorf("MakeyButton Event \"Release\" was not published")
	}

	testAdaptorDigitalRead = func() (val int, err error) {
		err = errors.New("digital read error")
		return
	}

	d.Once(Error, func(data interface{}) {
		sem <- true
	})

	select {
	case <-sem:
	case <-time.After(MAKEY_TEST_DELAY * time.Millisecond):
		t.Errorf("MakeyButton Event \"Error\" was not published")
	}

	d.Once(ButtonRelease, func(data interface{}) {
		sem <- true
	})
	testAdaptorDigitalRead = func() (val int, err error) {
		val = 1
		return
	}

	d.halt <- true

	select {
	case <-sem:
		t.Errorf("MakeyButton Event should not have been published")
	case <-time.After(MAKEY_TEST_DELAY * time.Millisecond):
	}
}
