package main

import (
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/firmata"
)

func main() {
	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
	motor := gpio.NewMotorDriver(firmataAdaptor, "3")

	work := func() {
		speed := byte(0)
		fadeAmount := byte(15)

		gobot.Every(100*time.Millisecond, func() {
			motor.Speed(speed)
			speed = speed + fadeAmount
			if speed == 0 || speed == 255 {
				fadeAmount = -fadeAmount
			}
		})
	}

	robot := gobot.NewRobot("motorBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{motor},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
