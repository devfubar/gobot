package nats

import (
	"github.com/nats-io/nats"
)

// Adaptor is a configuration struct for interacting with a NATS server.
// Name is a logical name for the adaptor/nats server connection.
// Host is in the form "localhost:4222" which is the hostname/ip and port of the nats server.
// ClientID is a unique identifier integer that specifies the identity of the client.
type Adaptor struct {
	name     string
	Host     string
	clientID int
	username string
	password string
	client   *nats.Conn
}

// NewAdaptor populates a new NATS Adaptor.
func NewAdaptor(host string, clientID int) *Adaptor {
	return &Adaptor{
		name:     "NATS",
		Host:     host,
		clientID: clientID,
	}
}

// NewAdaptorWithAuth populates a NATS Adaptor including username and password.
func NewAdaptorWithAuth(host string, clientID int, username string, password string) *Adaptor {
	return &Adaptor{
		Host:     host,
		clientID: clientID,
		username: username,
		password: password,
	}
}

// Name returns the logical client name.
func (a *Adaptor) Name() string { return a.name }

// SetName sets the logical client name.
func (a *Adaptor) SetName(n string) { a.name = n }

// Connect makes a connection to the Nats server.
func (a *Adaptor) Connect() (errs []error) {

	auth := ""
	if a.username != "" && a.password != "" {
		auth = a.username + ":" + a.password + "@"
	}

	defaultURL := "nats://" + auth + a.Host

	var err error
	a.client, err = nats.Connect(defaultURL)
	if err != nil {
		return append(errs, err)
	}
	return
}

// Disconnect from the nats server. Returns an error if the client doesn't exist.
func (a *Adaptor) Disconnect() (err error) {
	if a.client != nil {
		a.client.Close()
	}
	return
}

// Finalize is simply a helper method for the disconnect.
func (a *Adaptor) Finalize() (errs []error) {
	a.Disconnect()
	return
}

// Publish sends a message with the particular topic to the nats server.
func (a *Adaptor) Publish(topic string, message []byte) bool {
	if a.client == nil {
		return false
	}
	a.client.Publish(topic, message)
	return true
}

// On is an event-handler style subscriber to a particular topic (named event).
// Supply a handler function to use the bytes returned by the server.
func (a *Adaptor) On(event string, f func(s []byte)) bool {
	if a.client == nil {
		return false
	}
	a.client.Subscribe(event, func(msg *nats.Msg) {
		f(msg.Data)
	})
	return true
}
