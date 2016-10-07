package main

import (
	"fmt"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/raspi"
)

func main() {
	gbot := gobot.NewGobot()

	r := raspi.NewAdaptor()
	button := gpio.NewButtonDriver(r, "11")
	led := gpio.NewLedDriver(r, "7")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("button pressed")
			led.On()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("button released")
			led.Off()
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{r},
		[]gobot.Device{button, led},
		work,
	)
	gbot.AddRobot(robot)
	gbot.Start()
}
