/*
Package spark provides the Gobot adaptor for the Spark Core.

Installing:

	go get github.com/devfubar/gobot && go install github.com/devfubar/gobot/platforms/spark

Example:

	package main

	import (
		"time"

		"github.com/devfubar/gobot"
		"github.com/devfubar/gobot/drivers/gpio"
		"github.com/devfubar/gobot/platforms/particle"
	)

	func main() {
		gbot := gobot.NewGobot()

		core := paticle.NewAdaptor("device_id", "access_token")
		led := gpio.NewLedDriver(core, "D7")

		work := func() {
			gobot.Every(1*time.Second, func() {
				led.Toggle()
			})
		}

		robot := gobot.NewRobot("spark",
			[]gobot.Connection{core},
			[]gobot.Device{led},
			work,
		)

		gbot.AddRobot(robot)

		gbot.Start()
	}

For further information refer to Particle readme:
https://github.com/devfubar/gobot/blob/master/platforms/particle/README.md
*/
package particle
