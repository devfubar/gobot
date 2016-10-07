/*
Package firmata provides the Gobot adaptor for microcontrollers that support the Firmata protocol.

Installing:

	go get -d -u github.com/devfubar/gobot/... && go get github.com/devfubar/gobot/platforms/firmata

Example:

	package main

	import (
		"time"

		"github.com/devfubar/gobot"
		"github.com/devfubar/gobot/platforms/firmata"
		"github.com/devfubar/gobot/drivers/gpio"
	)

	func main() {
		gbot := gobot.NewGobot()

		firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
		led := gpio.NewLedDriver(firmataAdaptor, "13")

		work := func() {
			gobot.Every(1*time.Second, func() {
				led.Toggle()
			})
		}

		robot := gobot.NewRobot("bot",
			[]gobot.Connection{firmataAdaptor},
			[]gobot.Device{led},
			work,
		)

		gbot.AddRobot(robot)

		gbot.Start()
	}

For further information refer to firmata readme:
https://github.com/devfubar/gobot/blob/master/platforms/firmata/README.md
*/
package firmata
