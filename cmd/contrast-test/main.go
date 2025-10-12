package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/constants"
)

func main() {
	fmt.Println("🎨 Dark Theme Contrast Test")
	fmt.Println(strings.Repeat("=", 50))

	// Show background vs foreground contrast
	fmt.Println("\n📋 Background vs Foreground:")

	// Background color with white text
	bgStyle := lipgloss.NewStyle().
		Background(constants.GetBackgroundColor()).
		Foreground(lipgloss.Color("#ffffff")). // Pure white text
		Padding(1, 2)

	fmt.Println("Background (#1d1e20):", bgStyle.Render("  Background Color  "))

	// Foreground color with black background
	fgStyle := lipgloss.NewStyle().
		Foreground(constants.GetForegroundColor()).
		Background(lipgloss.Color("#000000")). // Pure black background
		Padding(1, 2)

	fmt.Println("Foreground (#e6e6e6):", fgStyle.Render("  Foreground Color  "))

	// Show ANSI 0 (black) vs ANSI 7 (white)
	fmt.Println("\n🌈 ANSI Black vs White:")

	ansi0Style := lipgloss.NewStyle().
		Foreground(constants.Black).           // Black
		Background(lipgloss.Color("#ffffff")). // White background
		Padding(1, 2)

	ansi7Style := lipgloss.NewStyle().
		Foreground(constants.White).           // White
		Background(lipgloss.Color("#000000")). // Black background
		Padding(1, 2)

	fmt.Println("ANSI 0 (Black):", ansi0Style.Render("  Black Text  "))
	fmt.Println("ANSI 7 (White):", ansi7Style.Render("  White Text  "))

	// Show the actual color values from constants
	fmt.Println("\n📊 Color Values (from ~/.gtconfig.yaml):")
	fmt.Println("Background: #1d1e20 (distinct dark bg)")
	fmt.Println("Foreground: #e6e6e6 (distinct light fg)")
	fmt.Println("ANSI 0:     #1d1e20 (black)")
	fmt.Println("ANSI 7:     #e6e6e6 (white)")
	fmt.Println("ANSI 1:     #ff5c57 (red)")
	fmt.Println("ANSI 2:     #A9DC52 (green)")
	fmt.Println("ANSI 3:     #f3f99d (yellow)")
	fmt.Println("ANSI 4:     #57c7ff (blue)")
	fmt.Println("ANSI 5:     #ff6ac1 (magenta)")
	fmt.Println("ANSI 6:     #9aedfe (cyan)")

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Now background/foreground and ANSI 0/7 should be clearly different!")
}
