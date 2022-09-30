# Go-BF

An extensible [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) interpreter written in Go.

# Example

```go
package main

import (
	"strings"
	framework "github.com/wengchaoxi/gobf"
)

func UnixTidyPlugin(program string) string {
	return strings.ReplaceAll(program, "\n", "")
}

func WindowsTidyPlugin(program string) string {
	return strings.ReplaceAll(program, "\r\n", "")
}

func main() {
	t := framework.NewTape(1024)
	m := framework.NewMachine(t)

	// DO WHAT THE FUCK YOU WANT TO
	m.Use(UnixTidyPlugin, WindowsTidyPlugin)

	m.Run("+[,.]\n")
}
```
