package main

import (
	"strings"

	"github.com/vlad1m1r0v/APZ-4/commands"
	"github.com/vlad1m1r0v/APZ-4/engine"
)

func parsePrint(args ...string) engine.Command{
		if len(args) != 1 {
			 return &commands.PrintCommand{Arg: "SYNTAX ERROR: INAPPROPRIATE AMOUNT OF ARGUMENTS"}
		}
		return &commands.PrintCommand{Arg: args[0]}

}

func parseSplit(args ...string) engine.Command{
		if len(args) != 2 {
			return &commands.PrintCommand{Arg: "SYNTAX ERROR: INAPPROPRIATE AMOUNT OF ARGUMENTS"}
		}
		return &commands.SplitCommand{Str: args[0], Delim: args[1]}
}

type fn func(args ...string) engine.Command

var parseMap = map[string]fn{
	"print": parsePrint,
	"split": parseSplit,

}

func parse(line string) engine.Command {
	parts := strings.Fields(line)
	field := parts[0]
	args := parts[1:]
	if fn, ok := parseMap[field]; ok {
		return fn(args...)
	}
	return &commands.PrintCommand{Arg: "SYNTAX ERROR: INVALID FUNCTION NAME"}
}
