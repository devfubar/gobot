package nats

import (
	"fmt"
	"testing"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/gobottest"
)

var _ gobot.Adaptor = (*Adaptor)(nil)

func TestNatsAdaptorReturnsHost(t *testing.T) {
	a := NewAdaptor("localhost:4222", 9999)
	gobottest.Assert(t, a.Host, "localhost:4222")
}

func TestNatsAdaptorPublishWhenConnected(t *testing.T) {
	a := NewAdaptor("localhost:4222", 9999)
	a.Connect()
	data := []byte("o")
	gobottest.Assert(t, a.Publish("test", data), true)
}

func TestNatsAdaptorOnWhenConnected(t *testing.T) {
	a := NewAdaptor("localhost:4222", 9999)
	a.Connect()
	gobottest.Assert(t, a.On("hola", func(data []byte) {
		fmt.Println("hola")
	}), true)
}

func TestNatsAdaptorPublishWhenConnectedWithAuth(t *testing.T) {
	a := NewAdaptorWithAuth("localhost:4222", 9999, "test", "testwd")
	a.Connect()
	data := []byte("o")
	gobottest.Assert(t, a.Publish("test", data), true)
}

func TestNatsAdaptorOnWhenConnectedWithAuth(t *testing.T) {
	a := NewAdaptorWithAuth("localhost:4222", 9999, "test", "testwd")
	a.Connect()
	gobottest.Assert(t, a.On("hola", func(data []byte) {
		fmt.Println("hola")
	}), true)
}

func TestNatsAdaptorConnect(t *testing.T) {
	a := NewAdaptor("localhost:9999", 9999)
	gobottest.Assert(t, a.Connect()[0].Error(), "nats: no servers available for connection")
}

func TestNatsAdaptorFinalize(t *testing.T) {
	a := NewAdaptor("localhost:9999", 9999)
	gobottest.Assert(t, len(a.Finalize()), 0)
}

func TestNatsAdaptorCannotPublishUnlessConnected(t *testing.T) {
	a := NewAdaptor("localhost:9999", 9999)
	data := []byte("o")
	gobottest.Assert(t, a.Publish("test", data), false)
}

func TestNatsAdaptorCannotOnUnlessConnected(t *testing.T) {
	a := NewAdaptor("localhost:9999", 9999)
	gobottest.Assert(t, a.On("hola", func(data []byte) {
		fmt.Println("hola")
	}), false)
}
