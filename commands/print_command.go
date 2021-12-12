package commands

import (
	"fmt"
	"github.com/vlad1m1r0v/APZ-4/engine"
)

type PrintCommand struct {
	Arg string
}
func (p *PrintCommand) Execute(loop engine.Handler) {
	fmt.Println(p.Arg)
}

