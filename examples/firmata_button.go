package main

import (
	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/firmata"
)

func main() {
	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")

	button := gpio.NewButtonDriver(firmataAdaptor, "2")
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			led.On()
		})
		button.On(gpio.ButtonRelease, func(data interface{}) {
			led.Off()
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{button, led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
