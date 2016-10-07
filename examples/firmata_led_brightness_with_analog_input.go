package main

import (
	"fmt"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/firmata"
)

func main() {
	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
	sensor := gpio.NewAnalogSensorDriver(firmataAdaptor, "0")
	led := gpio.NewLedDriver(firmataAdaptor, "3")

	work := func() {
		sensor.On(gpio.Data, func(data interface{}) {
			brightness := uint8(
				gobot.ToScale(gobot.FromScale(float64(data.(int)), 0, 1024), 0, 255),
			)
			fmt.Println("sensor", data)
			fmt.Println("brightness", brightness)
			led.Brightness(brightness)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{sensor, led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
