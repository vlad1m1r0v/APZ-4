package main

import (
	"bufio"
	"github.com/vlad1m1r0v/APZ-4/engine"
	"log"
	"os"
)

func main () {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	inputFile := "./cmd/example/example.txt"
	input, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		commandLine := scanner.Text()
		cmd := parse(commandLine) // parse the line to get an instance of Command
		eventLoop.Post(cmd)
	}
	eventLoop.AwaitFinish()
}