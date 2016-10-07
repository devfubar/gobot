package main

import (
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/platforms/audio"
)

func main() {
	gbot := gobot.NewGobot()

	e := audio.NewAdaptor()
	laser := audio.NewDriver(e, "./examples/laser.mp3")

	work := func() {
		gobot.Every(2*time.Second, func() {
			laser.Play()
		})
	}

	robot := gobot.NewRobot("soundBot",
		[]gobot.Connection{e},
		[]gobot.Device{laser},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
