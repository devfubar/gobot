package main

import (
	"fmt"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/platforms/particle"
)

func main() {
	gbot := gobot.NewGobot()

	core := particle.NewAdaptor("DEVICE_ID", "ACCESS_TOKEN")

	work := func() {
		if result, err := core.Function("brew", "202,230"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("result from \"brew\":", result)
		}
	}

	robot := gobot.NewRobot("spark",
		[]gobot.Connection{core},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
