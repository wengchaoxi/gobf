package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	framework "github.com/wengchaoxi/gobf"
)

func TidyPlugin(program string) string {
	program = strings.ReplaceAll(program, "\n", "")
	program = strings.ReplaceAll(program, "\r", "")
	program = strings.ReplaceAll(program, " ", "")
	return program
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// +++++++++ => +9
// --------- => -9
func MyPlugin(program string) string {
	target := ""

	pi := 0

	for pi < len(program) {
		ti := pi
		if IsDigit(program[ti]) {
			tmp := ""
			for ti < len(program) && IsDigit(program[ti]) {
				tmp += string(program[ti])
				ti++
			}
			count, _ := strconv.Atoi(tmp)
			target += strings.Repeat(string(program[pi-1]), count-1)

			pi = ti
		}
		target += string(program[pi])

		pi++
	}
	return target
}

// h3110 w0r1d\n
// 104 51 49 49 48 32 119 48 114 49 100 10
// +104.>+51.>+49.>+49.>+48.>+32.>+119.>+48.>+114.>+49.>+100.>+10.

func main() {
	t := framework.NewTape(1024)
	m := framework.NewMachine(t)
	m.Use(TidyPlugin, MyPlugin)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [file.bf]\n", os.Args[0])
		os.Exit(1)
	}

	program, _ := ioutil.ReadFile(os.Args[1])
	m.Run(string(program))
}
