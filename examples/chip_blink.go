package main

import (
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/chip"
)

func main() {
	gbot := gobot.NewGobot()

	chipAdaptor := chip.NewAdaptor()
	led := gpio.NewLedDriver(chipAdaptor, "XIO-P0")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{chipAdaptor},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
