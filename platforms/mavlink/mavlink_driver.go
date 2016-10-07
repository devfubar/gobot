package mavlink

import (
	"time"

	"github.com/devfubar/gobot"
	common "github.com/devfubar/gobot/platforms/mavlink/common"
)

const (
	// PacketEvent event
	PacketEvent = "packet"
	// MessageEvent event
	MessageEvent = "message"
	// ErrorIOE event
	ErrorIOEvent = "errorIO"
	// ErrorMAVLinkEvent event
	ErrorMAVLinkEvent = "errorMAVLink"
)

type Driver struct {
	name       string
	connection gobot.Connection
	interval   time.Duration
	gobot.Eventer
}

type MavlinkInterface interface {
}

// NewDriver creates a new mavlink driver.
//
// It add the following events:
//	"packet" - triggered when a new packet is read
//	"message" - triggered when a new valid message is processed
func NewDriver(a *Adaptor, v ...time.Duration) *Driver {
	m := &Driver{
		name:       "Mavlink",
		connection: a,
		Eventer:    gobot.NewEventer(),
		interval:   10 * time.Millisecond,
	}

	if len(v) > 0 {
		m.interval = v[0]
	}

	m.AddEvent(PacketEvent)
	m.AddEvent(MessageEvent)
	m.AddEvent(ErrorIOEvent)
	m.AddEvent(ErrorMAVLinkEvent)

	return m
}

func (m *Driver) Connection() gobot.Connection { return m.connection }
func (m *Driver) Name() string                 { return m.name }
func (m *Driver) SetName(n string)             { m.name = n }

// adaptor returns driver associated adaptor
func (m *Driver) adaptor() *Adaptor {
	return m.Connection().(*Adaptor)
}

// Start begins process to read mavlink packets every m.Interval
// and process them
func (m *Driver) Start() (errs []error) {
	go func() {
		for {
			packet, err := common.ReadMAVLinkPacket(m.adaptor().sp)
			if err != nil {
				m.Publish(ErrorIOEvent, err)
				continue
			}
			m.Publish(PacketEvent, packet)
			message, err := packet.MAVLinkMessage()
			if err != nil {
				m.Publish(ErrorMAVLinkEvent, err)
				continue
			}
			m.Publish(MessageEvent, message)
			<-time.After(m.interval)
		}
	}()
	return
}

// Halt returns true if device is halted successfully
func (m *Driver) Halt() (errs []error) { return }

// SendPacket sends a packet to mavlink device
func (m *Driver) SendPacket(packet *common.MAVLinkPacket) (err error) {
	_, err = m.adaptor().sp.Write(packet.Pack())
	return err
}
