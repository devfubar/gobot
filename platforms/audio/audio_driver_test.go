// Based on aplay audio adaptor written by @colemanserious (https://github.com/colemanserious)

package audio

import (
	"os/exec"
	"testing"

	"github.com/devfubar/gobot"
	"github.com/devfubar/gobot/gobottest"
)

var _ gobot.Driver = (*Driver)(nil)

func TestAudioDriver(t *testing.T) {
	d := NewDriver(NewAdaptor(), "../../examples/laser.mp3")

	gobottest.Assert(t, d.Filename(), "../../examples/laser.mp3")

	gobottest.Refute(t, d.Connection(), nil)

	gobottest.Assert(t, len(d.Start()), 0)

	gobottest.Assert(t, len(d.Halt()), 0)
}

func TestAudioDriverSoundWithNoFilename(t *testing.T) {
	d := NewDriver(NewAdaptor(), "")

	errors := d.Sound("")
	gobottest.Assert(t, errors[0].Error(), "Requires filename for audio file.")
}

func TestAudioDriverSoundWithDefaultFilename(t *testing.T) {
	execCommand = gobottest.ExecCommand
	defer func() { execCommand = exec.Command }()

	d := NewDriver(NewAdaptor(), "../../examples/laser.mp3")

	errors := d.Play()
	gobottest.Assert(t, len(errors), 0)
}
