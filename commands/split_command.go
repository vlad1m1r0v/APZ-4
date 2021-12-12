package commands

import (
	"github.com/vlad1m1r0v/APZ-4/engine"
	"regexp"
)

type SplitCommand struct {
	Str   string
	Delim string
}
func (p *SplitCommand) Execute(loop engine.Handler) {
	re := regexp.MustCompile(p.Delim)
	args := re.Split(p.Str, -1)
	for _, argument := range args {
		loop.Post(&PrintCommand{Arg: argument})
	}
}
