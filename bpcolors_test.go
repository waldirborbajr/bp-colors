package bpcolors_test

import (
	"fmt"
	"testing"

	color "github.com/waldirborbajr/bp-colors"
)

func TestMain(m *testing.M) {
	// color "gitlab.com/waldirborbajr/color"
	fmt.Printf("%q\n", color.Blue+"Blue"+color.Reset)
	fmt.Printf("%q\n", color.Red+"Red"+color.Reset)
	fmt.Printf("%q\n", color.Magenta+"Magenta"+color.Reset)
	fmt.Printf("%q\n", color.Cyan+"Cyan"+color.Reset)

	fmt.Println(color.Blue, "Blue", color.Reset)
	fmt.Println(color.Red, "Red", color.Reset)
	fmt.Println(color.Magenta, "Magenta", color.Reset)
	fmt.Println(color.Cyan, "Cyan", color.Reset)
}
