package main

import (
	"fmt"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/intel-iot/edison"
)

func main() {
	gbot := gobot.NewGobot()

	board := edison.NewAdaptor()
	sensor := gpio.NewGrovePiezoVibrationSensorDriver(board, "0")

	work := func() {
		sensor.On(gpio.Vibration, func(data interface{}) {
			fmt.Println("got one!")
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{board},
		[]gobot.Device{sensor},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
