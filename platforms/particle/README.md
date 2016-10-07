# Particle

The Particle Photon is a Wi-Fi connected microcontroller from Particle (http://particle.io), the company formerly known as Spark Devices. Once it connects to a Wi-Fi network, it automatically connects with a central server (the "Particle Cloud") and stays connected so it can be controlled from external systems, such as a Gobot program. To run gobot programs please make sure you are running default tinker firmware on the Photon.

For more info about the Particle platform go to https://www.particle.io/

## How to Install

Installing Gobot with Particle support is pretty easy.

```
go get -d -u github.com/devfubar/gobot/... && go install github.com/devfubar/gobot/platforms/particle
```

## How to Use

```go
package main

import (
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/drivers/gpio"
	"github.com/devfubar/gobot/platforms/particle"
)

func main() {
	gbot := gobot.NewGobot()

	core := particle.NewAdaptor("device_id", "access_token")
	led := gpio.NewLedDriver(core, "D7")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("spark",
		[]gobot.Connection{sparkCore},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
```
