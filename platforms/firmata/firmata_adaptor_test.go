package firmata

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/drivers/i2c"
	"github.com/devfubar/gobot/gobottest"
	"github.com/devfubar/gobot/platforms/firmata/client"
)

var _ gobot.Adaptor = (*Adaptor)(nil)

var _ gpio.DigitalReader = (*Adaptor)(nil)
var _ gpio.DigitalWriter = (*Adaptor)(nil)
var _ gpio.AnalogReader = (*Adaptor)(nil)
var _ gpio.PwmWriter = (*Adaptor)(nil)
var _ gpio.ServoWriter = (*Adaptor)(nil)

var _ i2c.I2c = (*Adaptor)(nil)

type readWriteCloser struct{}

func (readWriteCloser) Write(p []byte) (int, error) {
	return testWriteData.Write(p)
}

var testReadData = []byte{}
var testWriteData = bytes.Buffer{}

func (readWriteCloser) Read(b []byte) (int, error) {
	size := len(b)
	if len(testReadData) < size {
		size = len(testReadData)
	}
	copy(b, []byte(testReadData)[:size])
	testReadData = testReadData[size:]

	return size, nil
}

func (readWriteCloser) Close() error {
	return nil
}

type mockFirmataBoard struct {
	disconnectError error
	gobot.Eventer
	pins []client.Pin
}

func newMockFirmataBoard() *mockFirmataBoard {
	m := &mockFirmataBoard{
		Eventer:         gobot.NewEventer(),
		disconnectError: nil,
		pins:            make([]client.Pin, 100),
	}

	m.pins[1].Value = 1
	m.pins[15].Value = 133

	m.AddEvent("I2cReply")
	return m
}

func (mockFirmataBoard) Connect(io.ReadWriteCloser) error { return nil }
func (m mockFirmataBoard) Disconnect() error {
	return m.disconnectError
}
func (m mockFirmataBoard) Pins() []client.Pin {
	return m.pins
}
func (mockFirmataBoard) AnalogWrite(int, int) error      { return nil }
func (mockFirmataBoard) SetPinMode(int, int) error       { return nil }
func (mockFirmataBoard) ReportAnalog(int, int) error     { return nil }
func (mockFirmataBoard) ReportDigital(int, int) error    { return nil }
func (mockFirmataBoard) DigitalWrite(int, int) error     { return nil }
func (mockFirmataBoard) I2cRead(int, int) error          { return nil }
func (mockFirmataBoard) I2cWrite(int, []byte) error      { return nil }
func (mockFirmataBoard) I2cConfig(int) error             { return nil }
func (mockFirmataBoard) ServoConfig(int, int, int) error { return nil }

func initTestAdaptor() *Adaptor {
	a := NewAdaptor("/dev/null")
	a.board = newMockFirmataBoard()
	a.openSP = func(port string) (io.ReadWriteCloser, error) {
		return &readWriteCloser{}, nil
	}
	a.Connect()
	return a
}

func TestAdaptor(t *testing.T) {
	a := initTestAdaptor()
	gobottest.Assert(t, a.Port(), "/dev/null")
}

func TestAdaptorFinalize(t *testing.T) {
	a := initTestAdaptor()
	gobottest.Assert(t, len(a.Finalize()), 0)

	a = initTestAdaptor()
	a.board.(*mockFirmataBoard).disconnectError = errors.New("close error")
	gobottest.Assert(t, a.Finalize()[0], errors.New("close error"))
}

func TestAdaptorConnect(t *testing.T) {
	var openSP = func(port string) (io.ReadWriteCloser, error) {
		return &readWriteCloser{}, nil
	}
	a := NewAdaptor("/dev/null")
	a.openSP = openSP
	a.board = newMockFirmataBoard()
	gobottest.Assert(t, len(a.Connect()), 0)

	a = NewAdaptor("/dev/null")
	a.board = newMockFirmataBoard()
	a.openSP = func(port string) (io.ReadWriteCloser, error) {
		return nil, errors.New("connect error")
	}
	gobottest.Assert(t, a.Connect()[0], errors.New("connect error"))

	a = NewAdaptor(&readWriteCloser{})
	a.board = newMockFirmataBoard()
	gobottest.Assert(t, len(a.Connect()), 0)

}

func TestAdaptorServoWrite(t *testing.T) {
	a := initTestAdaptor()
	a.ServoWrite("1", 50)
}

func TestAdaptorPwmWrite(t *testing.T) {
	a := initTestAdaptor()
	a.PwmWrite("1", 50)
}

func TestAdaptorDigitalWrite(t *testing.T) {
	a := initTestAdaptor()
	a.DigitalWrite("1", 1)
}

func TestAdaptorDigitalRead(t *testing.T) {
	a := initTestAdaptor()
	val, err := a.DigitalRead("1")
	gobottest.Assert(t, err, nil)
	gobottest.Assert(t, val, 1)
}

func TestAdaptorAnalogRead(t *testing.T) {
	a := initTestAdaptor()
	val, err := a.AnalogRead("1")
	gobottest.Assert(t, val, 133)
	gobottest.Assert(t, err, nil)
}

func TestAdaptorI2cStart(t *testing.T) {
	a := initTestAdaptor()
	a.I2cStart(0x00)
}
func TestAdaptorI2cRead(t *testing.T) {
	a := initTestAdaptor()
	i := []byte{100}
	i2cReply := client.I2cReply{Data: i}
	go func() {
		<-time.After(10 * time.Millisecond)
		a.Publish(a.board.Event("I2cReply"), i2cReply)
	}()
	data, err := a.I2cRead(0x00, 1)
	gobottest.Assert(t, err, nil)
	gobottest.Assert(t, data, i)
}
func TestAdaptorI2cWrite(t *testing.T) {
	a := initTestAdaptor()
	a.I2cWrite(0x00, []byte{0x00, 0x01})
}

func TestServoConfig(t *testing.T) {
	a := initTestAdaptor()
	err := a.ServoConfig("9", 0, 0)
	gobottest.Assert(t, err, nil)

	// test atoi error
	err = a.ServoConfig("a", 0, 0)
	gobottest.Assert(t, true, strings.Contains(fmt.Sprintf("%v", err), "invalid syntax"))
}
