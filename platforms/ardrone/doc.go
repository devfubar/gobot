/*
Package ardrone provides the Gobot adaptor and driver for the Parrot Ardrone.

Installing:

	go get -d -u github.com/devfubar/gobot/... && go install github.com/devfubar/gobot/platforms/ardrone

Example:

	package main

	import (
		"time"

		"github.com/devfubar/gobot"
		"github.com/devfubar/gobot/platforms/ardrone"
	)

	func main() {
		gbot := gobot.NewGobot()

		ardroneAdaptor := ardrone.NewAdaptor()
		drone := ardrone.NewDriver(ardroneAdaptor)

		work := func() {
			drone.TakeOff()
			gobot.On(drone.Event("flying"), func(data interface{}) {
				gobot.After(3*time.Second, func() {
					drone.Land()
				})
			})
		}

		robot := gobot.NewRobot("drone",
			[]gobot.Connection{ardroneAdaptor},
			[]gobot.Device{drone},
			work,
		)
		gbot.AddRobot(robot)

		gbot.Start()
	}

For more information refer to the ardrone README:
https://github.com/devfubar/gobot/tree/master/platforms/ardrone
*/
package ardrone
