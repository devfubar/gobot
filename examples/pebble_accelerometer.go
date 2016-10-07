package main

import (
	"fmt"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/api"
	"github.com/devfubar/gobot/platforms/pebble"
)

func main() {
	gbot := gobot.NewGobot()
	a := api.NewAPI(gbot)
	a.Port = "8080"
	a.Start()

	pebbleAdaptor := pebble.NewAdaptor()
	pebbleDriver := pebble.NewDriver(pebbleAdaptor)

	work := func() {
		pebbleDriver.On(pebbleDriver.Event("accel"), func(data interface{}) {
			fmt.Println(data.(string))
		})
	}

	robot := gobot.NewRobot("pebble",
		[]gobot.Connection{pebbleAdaptor},
		[]gobot.Device{pebbleDriver},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
