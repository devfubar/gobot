package main

import (
	"fmt"
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/i2c"
	"github.com/devfubar/gobot/platforms/intel-iot/joule"
)

func main() {
	gbot := gobot.NewGobot()

	e := joule.NewAdaptor()
	blinkm := i2c.NewBlinkMDriver(e)

	work := func() {
		gobot.Every(3*time.Second, func() {
			r := byte(gobot.Rand(255))
			g := byte(gobot.Rand(255))
			b := byte(gobot.Rand(255))
			blinkm.Rgb(r, g, b)
			color, _ := blinkm.Color()
			fmt.Println("color", color)
		})
	}

	robot := gobot.NewRobot("blinkmBot",
		[]gobot.Connection{e},
		[]gobot.Device{blinkm},
		work,
	)

	gbot.AddRobot(robot)
	gbot.Start()
}
