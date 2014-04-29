// +build linux

package diaboro

import (
	"fmt"
	"os/exec"
)

type XDoTool struct{}

func (x XDoTool) Click(b Button) Command {
	return Command(fmt.Sprintf("xdotool click %d", b))
}

func (x XDoTool) Press(k Key) Command {
	return Command(fmt.Sprintf("xdotool key %s", k))
}

func (x XDoTool) Move(c Coordinate) Command {
	return Command(fmt.Sprintf("xdotool mousemove %d %d", c.X, c.Y))
}

func init() {
	if _, err := exec.LookPath("xdotool"); err != nil {
		panic(err)
	}

	TheInputer = XDoTool{}
}
