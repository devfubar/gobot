package main

import (
	"fmt"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/beaglebone"
)

func main() {
	gbot := gobot.NewGobot()

	beagleboneAdaptor := beaglebone.NewAdaptor()
	button := gpio.NewMakeyButtonDriver(beagleboneAdaptor, "P8_9")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("button pressed")
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("button released")
		})
	}

	robot := gobot.NewRobot("makeyBot",
		[]gobot.Connection{beagleboneAdaptor},
		[]gobot.Device{button},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
