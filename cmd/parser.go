package main

import (
	"github.com/vlad1m1r0v/APZ-4/commands"
	"github.com/vlad1m1r0v/APZ-4/engine"
	"strings"
)

func parse(line string) engine.Command {
	parts := strings.Fields(line)
	cmd := parts[0]

	if cmd == "print" {
		if len(parts[1:]) != 1 {
			 return &commands.PrintCommand{Arg: "SYNTAX ERROR: INAPPROPRIATE AMOUNT OF ARGUMENTS"}
		}
		return &commands.PrintCommand{Arg: parts[1]}
	}
	if cmd == "split" {
		if len(parts[1:]) != 2 {
			return &commands.PrintCommand{Arg: "SYNTAX ERROR: INAPPROPRIATE AMOUNT OF ARGUMENTS"}
		}
		return &commands.SplitCommand{Str: parts[1], Delim: parts[2]}
	}
	return &commands.PrintCommand{Arg: "SYNTAX ERROR: UNKNOWN COMMAND"}
}
