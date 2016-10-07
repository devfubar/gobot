package main

import (
	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/api"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/digispark"
)

func main() {
	gbot := gobot.NewGobot()

	api.NewAPI(gbot).Start()

	digisparkAdaptor := digispark.NewAdaptor()
	led := gpio.NewLedDriver(digisparkAdaptor, "0")

	robot := gobot.NewRobot("digispark",
		[]gobot.Connection{digisparkAdaptor},
		[]gobot.Device{led},
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
