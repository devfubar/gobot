package main

import (
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/beaglebone"
)

func main() {
	gbot := gobot.NewGobot()

	beagleboneAdaptor := beaglebone.NewAdaptor()
	led := gpio.NewDirectPinDriver(beagleboneAdaptor, "P8_10")
	button := gpio.NewDirectPinDriver(beagleboneAdaptor, "P8_9")

	work := func() {
		gobot.Every(500*time.Millisecond, func() {
			val, _ := button.DigitalRead()
			if val == 1 {
				led.DigitalWrite(1)
			} else {
				led.DigitalWrite(0)
			}
		})
	}

	robot := gobot.NewRobot("pinBot",
		[]gobot.Connection{beagleboneAdaptor},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
