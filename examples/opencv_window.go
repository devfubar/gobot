package main

import (
	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/platforms/opencv"
	cv "github.com/lazywei/go-opencv/opencv"
)

func main() {
	gbot := gobot.NewGobot()

	window := opencv.NewWindowDriver()
	camera := opencv.NewCameraDriver(0)

	work := func() {
		camera.On(camera.Event("frame"), func(data interface{}) {
			window.ShowImage(data.(*cv.IplImage))
		})
	}

	robot := gobot.NewRobot("cameraBot",
		[]gobot.Device{window, camera},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
