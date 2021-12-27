package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"testing"
)

func ExtendSplitCommand(times int) {
    for i := 0; i < times; i++ {
        input,_  := ioutil.ReadFile("./example/example-test.txt")
        parts := strings.Fields(string(input))
        cmd := parts[0]
        params := parts[1]
        delimiter := parts[2]
        params = fmt.Sprintf("%s-%s", params, params)
        newCommand := fmt.Sprintf("%s %s %s", cmd, params, delimiter)
        ioutil.WriteFile("./example/example-test.txt", []byte(newCommand), 0644)
    }
}

func SetupFile() {
    ioutil.WriteFile("./example/example-test.txt", []byte("split uno-dos-tres-cuatro-cinco -"), 0644)
}

func BenchmarkParser(b *testing.B) {
    for i := 1; i <= 20; i++ {
        SetupFile()
        ExtendSplitCommand(i)
        command,_  := ioutil.ReadFile("./example/example-test.txt")
        b.Run(fmt.Sprintf("%d elements", int(math.Pow(2, float64(i)))), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                parse(string(command))
            }
        })
    }
}