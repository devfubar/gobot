package main

import (
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/api"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/firmata"
)

func main() {
	gbot := gobot.NewGobot()

	a := api.NewAPI(gbot)
	a.Start()

	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
