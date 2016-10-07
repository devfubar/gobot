/*
Package sphero provides the Gobot adaptor and driver for the Sphero.

Installing:

	go get github.com/devfubar/gobot/platforms/sphero

Example:

	package main

	import (
		"fmt"
		"time"

		"github.com/devfubar/gobot"
		"github.com/devfubar/gobot/platforms/sphero"
	)

	func main() {
		gbot := gobot.NewGobot()

		adaptor := sphero.NewAdaptor("/dev/rfcomm0")
		driver := sphero.NewSpheroDriver(adaptor)

		work := func() {
			gobot.Every(3*time.Second, func() {
				driver.Roll(30, uint16(gobot.Rand(360)))
			})
		}

		robot := gobot.NewRobot("sphero",
			[]gobot.Connection{adaptor},
			[]gobot.Device{driver},
			work,
		)

		gbot.AddRobot(robot)

		gbot.Start()
	}

For further information refer to sphero readme:
https://github.com/devfubar/gobot/blob/master/platforms/sphero/README.md
*/
package sphero
