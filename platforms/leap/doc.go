/*
Package leap provides the Gobot adaptor and driver for the Leap Motion.

Installing:

* First install the [Leap Motion Software](https://www.leapmotion.com/setup).
* Then install the package:

	go get github.com/devfubar/gobot/platforms/leap

Example:

	package main

	import (
		"fmt"

		"github.com/devfubar/gobot"
		"github.com/devfubar/gobot/platforms/leap"
	)

	func main() {
		gbot := gobot.NewGobot()

		leapMotionAdaptor := leap.NewAdaptor("127.0.0.1:6437")
		l := leap.NewDriver(leapMotionAdaptor)

		work := func() {
			gobot.On(l.Event("message"), func(data interface{}) {
				fmt.Println(data.(leap.Frame))
			})
		}

		robot := gobot.NewRobot("leapBot",
			[]gobot.Connection{leapMotionAdaptor},
			[]gobot.Device{l},
			work,
		)

		gbot.AddRobot(robot)

		gbot.Start()
	}

For more information refer to the leap README:
https://github.com/devfubar/gobot/blob/master/platforms/leap/README.md
*/
package leap
