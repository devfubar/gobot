package main

import (
	"math"
	"time"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/platforms/bebop"
	"github.com/devfubar/gobot/platforms/joystick"
)

type pair struct {
	x float64
	y float64
}

func main() {
	gbot := gobot.NewGobot()

	joystickAdaptor := joystick.NewAdaptor()
	stick := joystick.NewDriver(joystickAdaptor,
		"./platforms/joystick/configs/dualshock3.json",
	)

	bebopAdaptor := bebop.NewAdaptor()
	drone := bebop.NewDriver(bebopAdaptor)

	work := func() {
		offset := 32767.0
		rightStick := pair{x: 0, y: 0}
		leftStick := pair{x: 0, y: 0}

		recording := false

		stick.On(joystick.CirclePress, func(data interface{}) {
			if recording {
				drone.StopRecording()
			} else {
				drone.StartRecording()
			}
			recording = !recording
		})

		stick.On(joystick.SquarePress, func(data interface{}) {
			drone.HullProtection(true)
			drone.TakeOff()
		})
		stick.On(joystick.TrianglePress, func(data interface{}) {
			drone.Stop()
		})
		stick.On(joystick.XPress, func(data interface{}) {
			drone.Land()
		})
		stick.On(joystick.LeftX, func(data interface{}) {
			val := float64(data.(int16))
			if leftStick.x != val {
				leftStick.x = val
			}
		})
		stick.On(joystick.LeftY, func(data interface{}) {
			val := float64(data.(int16))
			if leftStick.y != val {
				leftStick.y = val
			}
		})
		stick.On(joystick.RightX, func(data interface{}) {
			val := float64(data.(int16))
			if rightStick.x != val {
				rightStick.x = val
			}
		})
		stick.On(joystick.RightY, func(data interface{}) {
			val := float64(data.(int16))
			if rightStick.y != val {
				rightStick.y = val
			}
		})

		gobot.Every(10*time.Millisecond, func() {
			pair := leftStick
			if pair.y < -10 {
				drone.Forward(validatePitch(pair.y, offset))
			} else if pair.y > 10 {
				drone.Backward(validatePitch(pair.y, offset))
			} else {
				drone.Forward(0)
			}

			if pair.x > 10 {
				drone.Right(validatePitch(pair.x, offset))
			} else if pair.x < -10 {
				drone.Left(validatePitch(pair.x, offset))
			} else {
				drone.Right(0)
			}
		})

		gobot.Every(10*time.Millisecond, func() {
			pair := rightStick
			if pair.y < -10 {
				drone.Up(validatePitch(pair.y, offset))
			} else if pair.y > 10 {
				drone.Down(validatePitch(pair.y, offset))
			} else {
				drone.Up(0)
			}

			if pair.x > 20 {
				drone.Clockwise(validatePitch(pair.x, offset))
			} else if pair.x < -20 {
				drone.CounterClockwise(validatePitch(pair.x, offset))
			} else {
				drone.Clockwise(0)
			}
		})
	}

	robot := gobot.NewRobot("bebop",
		[]gobot.Connection{joystickAdaptor, bebopAdaptor},
		[]gobot.Device{stick, drone},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}

func validatePitch(data float64, offset float64) int {
	value := math.Abs(data) / offset
	if value >= 0.1 {
		if value <= 1.0 {
			return int((float64(int(value*100)) / 100) * 100)
		}
		return 100
	}
	return 0
}
