package main

import (
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/api"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/particle"
)

func main() {
	gbot := gobot.NewGobot()
	api.NewAPI(gbot).Start()

	core := particle.NewAdaptor("device_id", "access_token")
	led := gpio.NewLedDriver(core, "D7")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("spark",
		[]gobot.Connection{core},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
