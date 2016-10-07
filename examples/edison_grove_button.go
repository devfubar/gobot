package main

import (
	"fmt"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/intel-iot/edison"
)

func main() {
	gbot := gobot.NewGobot()

	e := edison.NewAdaptor()
	button := gpio.NewGroveButtonDriver(e, "2")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("On!")
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("Off!")
		})

	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{e},
		[]gobot.Device{button},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
