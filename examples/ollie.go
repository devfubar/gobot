package main

import (
	"os"
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/platforms/ble"
)

func main() {
	gbot := gobot.NewGobot()

	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ollie := ble.NewSpheroOllieDriver(bleAdaptor)

	work := func() {
		gobot.Every(1*time.Second, func() {
			r := uint8(gobot.Rand(255))
			g := uint8(gobot.Rand(255))
			b := uint8(gobot.Rand(255))
			ollie.SetRGB(r, g, b)
		})
	}

	robot := gobot.NewRobot("ollieBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ollie},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
