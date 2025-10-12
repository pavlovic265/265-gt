package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/pavlovic265/265-gt/constants"
)

func main() {
	fmt.Println("🎨 Theme Comparison Demo")
	fmt.Println(strings.Repeat("=", 50))

	// Show current theme
	fmt.Println("\n📋 Current Theme:")
	fmt.Printf("Theme: %s\n", getCurrentTheme())

	// Show colors for current theme
	fmt.Println("\n🌈 Current Theme Colors:")
	showThemeColors()

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("To switch themes, change 'return false' to 'return true' in isLightTheme() function")
	fmt.Println("Location: constants/colors.go line 220")
}

func getCurrentTheme() string {
	// Check which theme is currently active
	if constants.GetBackgroundColor() == constants.LightBackground {
		return "Light"
	}
	return "Dark"
}

func showThemeColors() {
	// Show background and foreground
	bgStyle := lipgloss.NewStyle().
		Background(constants.GetBackgroundColor()).
		Foreground(constants.GetAnsiColor(7)). // White text
		Padding(1, 2)

	fgStyle := lipgloss.NewStyle().
		Foreground(constants.GetForegroundColor()).
		Background(constants.GetAnsiColor(0)). // Black background
		Padding(1, 2)

	fmt.Println("Background:", bgStyle.Render("  Background  "))
	fmt.Println("Foreground:", fgStyle.Render("  Foreground  "))

	// Show common colors
	fmt.Println("\nCommon Colors:")
	fmt.Println("Success:", constants.GetSuccessAnsiStyle().Render("✓ Success"))
	fmt.Println("Error:", constants.GetErrorAnsiStyle().Render("✗ Error"))
	fmt.Println("Warning:", constants.GetWarningAnsiStyle().Render("⚠ Warning"))
	fmt.Println("Info:", constants.GetInfoAnsiStyle().Render("ℹ Info"))
	fmt.Println("Debug:", constants.GetDebugAnsiStyle().Render("[D] Debug"))
}
