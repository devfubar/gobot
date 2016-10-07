package main

import (
	"fmt"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/intel-iot/edison"
)

func main() {
	gbot := gobot.NewGobot()

	e := edison.NewAdaptor()
	sensor := gpio.NewAnalogSensorDriver(e, "0")
	led := gpio.NewLedDriver(e, "3")

	work := func() {
		sensor.On(gpio.Data, func(data interface{}) {
			brightness := uint8(
				gobot.ToScale(gobot.FromScale(float64(data.(int)), 0, 4096), 0, 255),
			)
			fmt.Println("sensor", data)
			fmt.Println("brightness", brightness)
			led.Brightness(brightness)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{e},
		[]gobot.Device{sensor, led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
