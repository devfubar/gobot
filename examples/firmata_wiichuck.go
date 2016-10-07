package main

import (
	"fmt"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/i2c"
	"github.com/devfubar/gobot/platforms/firmata"
)

func main() {
	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
	wiichuck := i2c.NewWiichuckDriver(firmataAdaptor)

	work := func() {
		wiichuck.On(wiichuck.Event("joystick"), func(data interface{}) {
			fmt.Println("joystick", data)
		})

		wiichuck.On(wiichuck.Event("c"), func(data interface{}) {
			fmt.Println("c")
		})

		wiichuck.On(wiichuck.Event("z"), func(data interface{}) {
			fmt.Println("z")
		})

		wiichuck.On(wiichuck.Event("error"), func(data interface{}) {
			fmt.Println("Wiichuck error:", data)
		})
	}

	robot := gobot.NewRobot("chuck",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{wiichuck},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
