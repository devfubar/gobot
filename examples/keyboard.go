package main

import (
	"fmt"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/platforms/keyboard"
)

func main() {
	gbot := gobot.NewGobot()

	keys := keyboard.NewDriver()

	work := func() {
		keys.On(keyboard.Key, func(data interface{}) {
			key := data.(keyboard.KeyEvent)

			if key.Key == keyboard.A {
				fmt.Println("A pressed!")
			} else {
				fmt.Println("keyboard event!", key, key.Char)
			}
		})
	}

	robot := gobot.NewRobot("keyboardbot",
		[]gobot.Connection{},
		[]gobot.Device{keys},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
