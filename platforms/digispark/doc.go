/*
Package digispark provides the Gobot adaptor for the Digispark ATTiny-based USB development board.

Installing:

This package requires installing `libusb`.
Then you can install the package with:

	go get github.com/devfubar/gobot/platforms/digispark

Example:

	package main

	import (
		"time"

		"github.com/devfubar/gobot"
		"github.com/devfubar/gobot/platforms/digispark"
		"github.com/devfubar/gobot/drivers/gpio"
	)

	func main() {
		gbot := gobot.NewGobot()

		digisparkAdaptor := digispark.NewAdaptor()
		led := gpio.NewLedDriver(digisparkAdaptor, "0")

		work := func() {
			gobot.Every(1*time.Second, func() {
				led.Toggle()
			})
		}

		robot := gobot.NewRobot("blinkBot",
			[]gobot.Connection{digisparkAdaptor},
			[]gobot.Device{led},
			work,
		)

		gbot.AddRobot(robot)

		gbot.Start()
	}

For further information refer to digispark README:
https://github.com/devfubar/gobot/blob/master/platforms/digispark/README.md
*/
package digispark
