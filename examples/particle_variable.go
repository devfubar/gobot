package main

import (
	"fmt"
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/platforms/particle"
)

func main() {
	gbot := gobot.NewGobot()

	core := particle.NewAdaptor("DEVICE_ID", "ACCESS_TOKEN")

	work := func() {
		gobot.Every(1*time.Second, func() {
			if temp, err := core.Variable("temperature"); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("result from \"temperature\" is:", temp)
			}
		})
	}

	robot := gobot.NewRobot("spark",
		[]gobot.Connection{core},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
