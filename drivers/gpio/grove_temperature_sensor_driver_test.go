package gpio

import (
	"github.com/devfubar/gobot"
)

var _ gobot.Driver = (*GroveTemperatureSensorDriver)(nil)
