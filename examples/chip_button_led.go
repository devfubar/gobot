package main

import (
	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/chip"
)

func main() {
	gbot := gobot.NewGobot()

	chipAdaptor := chip.NewAdaptor()
	button := gpio.NewButtonDriver(chipAdaptor, "XIO-P6")
	led := gpio.NewLedDriver(chipAdaptor, "XIO-P7")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			led.On()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			led.Off()
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{chipAdaptor},
		[]gobot.Device{button, led},
		work,
	)
	gbot.AddRobot(robot)
	gbot.Start()
}
