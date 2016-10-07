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
		ollie.SetRGB(255, 0, 255)
		gobot.Every(3*time.Second, func() {
			ollie.Roll(40, uint16(gobot.Rand(360)))
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
