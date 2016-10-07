/*
Package beaglebone provides the Gobot adaptor for the Beaglebone Black.

Installing:

	go get github.com/hybridgroup/platforms/gobot/beaglebone

Example:

	package main

	import (
		"time"

		"github.com/devfubar/gobot"
		"github.com/devfubar/gobot/platforms/beaglebone"
		"github.com/devfubar/gobot/drivers/gpio"
	)

	func main() {
		gbot := gobot.NewGobot()

		beagleboneAdaptor := beaglebone.NewAdaptor()
		led := gpio.NewLedDriver(beagleboneAdaptor, "P9_12")

		work := func() {
			gobot.Every(1*time.Second, func() {
				led.Toggle()
			})
		}

		robot := gobot.NewRobot("blinkBot",
			[]gobot.Connection{beagleboneAdaptor},
			[]gobot.Device{led},
			work,
		)

		gbot.AddRobot(robot)

		gbot.Start()
	}

For more information refer to the beaglebone README:
https://github.com/devfubar/gobot/blob/master/platforms/beaglebone/README.md
*/
package beaglebone
