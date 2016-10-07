package main

import (
	"fmt"
	"os"
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/platforms/ble"
)

func main() {
	gbot := gobot.NewGobot()

	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	battery := ble.NewBatteryDriver(bleAdaptor)

	work := func() {
		gobot.Every(5*time.Second, func() {
			fmt.Println("Battery level:", battery.GetBatteryLevel())
		})
	}

	robot := gobot.NewRobot("bleBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{battery},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
