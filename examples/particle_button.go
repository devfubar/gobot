package main

import (
	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/particle"
)

func main() {
	gbot := gobot.NewGobot()

	core := particle.NewAdaptor("device_id", "access_token")
	led := gpio.NewLedDriver(core, "D7")
	button := gpio.NewButtonDriver(core, "D5")

	work := func() {
		button.On(button.Event("push"), func(data interface{}) {
			led.On()
		})

		button.On(button.Event("release"), func(data interface{}) {
			led.Off()
		})
	}

	robot := gobot.NewRobot("spark",
		[]gobot.Connection{core},
		[]gobot.Device{button, led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
