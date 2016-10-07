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
		l.On(leap.GestureEvent, func(data interface{}) {
			printGesture(data.(leap.Gesture))
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

func printGesture(gesture leap.Gesture) {
	fmt.Println("Gesture", gesture)
}
