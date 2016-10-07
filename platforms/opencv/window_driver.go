package opencv

import (
	"github.com/devfubar/gobot"
	cv "github.com/lazywei/go-opencv/opencv"
)

type window interface {
	ShowImage(*cv.IplImage)
}

type WindowDriver struct {
	name   string
	window window
	start  func(*WindowDriver)
}

// NewWindowDriver creates a new window driver.
// It adds an start function to initialize window
func NewWindowDriver() *WindowDriver {
	return &WindowDriver{
		name: "Window",
		start: func(w *WindowDriver) {
			w.window = cv.NewWindow(w.Name(), cv.CV_WINDOW_NORMAL)
		},
	}
}

func (w *WindowDriver) Name() string                 { return w.name }
func (w *WindowDriver) SetName(n string)             { w.name = n }
func (w *WindowDriver) Connection() gobot.Connection { return nil }

// Start starts window thread and driver
func (w *WindowDriver) Start() (errs []error) {
	cv.StartWindowThread()
	w.start(w)
	return
}

// Halt returns true if camera is halted successfully
func (w *WindowDriver) Halt() (errs []error) { return }

// ShowImage displays image in window
func (w *WindowDriver) ShowImage(image *cv.IplImage) {
	w.window.ShowImage(image)
}
