// +build linux

package diaboro

import (
	"fmt"
	"os/exec"

	"github.com/dagoof/diaboro/res"
)

const (
	xdotool = "xdotool"
)

type Cmd struct {
	Name string
	Args []string
}

func NewCmd(name string, args ...string) *Cmd {
	return &Cmd{name, args}
}

func (c *Cmd) Execute() error {
	return exec.Command(c.Name, c.Args...).Run()
}

type XDoTool struct{}

func (i XDoTool) Click(b Button) Command {
	return NewCmd(xdotool, "click", fmt.Sprintf("%d", b))
}

func (i XDoTool) Press(k Key) Command {
	return NewCmd(xdotool, "key", string(k))
}

func (i XDoTool) Move(p res.Pointer) Command {
	c := p.Point(i.Resolution())
	x := fmt.Sprintf("%d", c.X)
	y := fmt.Sprintf("%d", c.Y)
	return NewCmd(xdotool, "mousemove", x, y)
}

func (i XDoTool) Resolution() res.Resolution {
	return res.Resolution{1920, 1200}
}

func init() {
	if _, err := exec.LookPath("xdotool"); err != nil {
		panic(err)
	}

	TheInputer = XDoTool{}
}
