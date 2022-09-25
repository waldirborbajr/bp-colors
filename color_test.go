package bpcolors_test

import (
	"fmt"

color	"github.com/waldirborbajr/bp-colors
)

func PrintColors() {
	// color "gitlab.com/waldirborbajr/color"
	fmt.Printf("%q\n", color.Blue+"Blue"+color.Reset)
	fmt.Printf("%q\n", color.Red+"Red"+color.Reset)
	fmt.Printf("%q\n", color.Magenta+"Magenta"+color.Reset)
	fmt.Printf("%q\n", color.Cyan+"Cyan"+color.Reset)
}
