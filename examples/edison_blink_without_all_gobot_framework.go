package main

import (
	"time"

	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/intel-iot/edison"
)

// Example of a simple led toggle without the initialization of
// the entire gobot framework.
// This might be useful if you want to use gobot as another
// golang library to interact with sensors and other devices.
func main() {
	e := edison.NewAdaptor()
	led := gpio.NewLedDriver(e, "13")
	e.Connect()
	led.Start()
	for {
		led.Toggle()
		time.Sleep(1000 * time.Millisecond)
	}
}
