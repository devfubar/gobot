package main

import (
	"fmt"
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/i2c"
	"github.com/devfubar/gobot/platforms/intel-iot/edison"
)

func main() {
	gbot := gobot.NewGobot()

	board := edison.NewAdaptor()
	accel := i2c.NewGroveAccelerometerDriver(board)

	work := func() {
		gobot.Every(500*time.Millisecond, func() {
			if x, y, z, err := accel.XYZ(); err == nil {
				fmt.Println(x, y, z)
				fmt.Println(accel.Acceleration(x, y, z))
			} else {
				fmt.Println(err)
			}
		})
	}

	robot := gobot.NewRobot("accelBot",
		[]gobot.Connection{board},
		[]gobot.Device{accel},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
