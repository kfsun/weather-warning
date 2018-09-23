package main

import (
	"fmt"
	"github.com/kfsworks/weather-warning/processor"
)

func main() {
	processor.Process()

	var input string
	fmt.Scanln(&input)
}
