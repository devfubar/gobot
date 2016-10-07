package main

import (
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/intel-iot/joule"
)

func main() {
	gbot := gobot.NewGobot()

	e := joule.NewAdaptor()
	led := gpio.NewRgbLedDriver(e, "25", "27", "29")

	work := func() {
		gobot.Every(1*time.Second, func() {
			r := uint8(gobot.Rand(255))
			g := uint8(gobot.Rand(255))
			b := uint8(gobot.Rand(255))
			led.SetRGB(r, g, b)
		})
	}

	robot := gobot.NewRobot("rgbBot",
		[]gobot.Connection{e},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
