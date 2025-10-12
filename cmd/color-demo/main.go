package main

import (
	"fmt"
	"strings"

	"github.com/pavlovic265/265-gt/constants"
)

func main() {
	fmt.Println("🎨 265-gt Color Palette Demo")
	fmt.Println("=============================")

	// Run the example usage
	constants.ExampleUsage()

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Theme switching example:")
	constants.ExampleThemeSwitch()
}
