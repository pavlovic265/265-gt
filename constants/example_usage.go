package constants

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// ExampleUsage demonstrates how to use the new 16-color ANSI palette
func ExampleUsage() {
	// Example 1: Using individual ANSI colors
	fmt.Println("=== Panda Terminal Theme Example ===")
	fmt.Println("Based on: https://github.com/PandaTheme/panda-terminal")

	// Dark theme colors (0-15)
	for i := 0; i < 16; i++ {
		style := GetAnsiStyle(i)
		fmt.Printf("ANSI %2d: %s\n", i, style.Render("Sample text"))
	}

	fmt.Println("\n=== Common Use Cases ===")

	// Example 2: Using convenience functions
	fmt.Println(GetSuccessAnsiStyle().Render("✓ Success message"))
	fmt.Println(GetErrorAnsiStyle().Render("✗ Error message"))
	fmt.Println(GetWarningAnsiStyle().Render("⚠ Warning message"))
	fmt.Println(GetInfoAnsiStyle().Render("ℹ Info message"))
	fmt.Println(GetDebugAnsiStyle().Render("[D] Debug message"))

	// Example 3: Using background and foreground colors
	bgStyle := lipgloss.NewStyle().Background(GetBackgroundColor())
	fgStyle := lipgloss.NewStyle().Foreground(GetForegroundColor())

	fmt.Println(bgStyle.Render("Background color"))
	fmt.Println(fgStyle.Render("Foreground color"))

	// Example 4: Custom combinations
	customStyle := lipgloss.NewStyle().
		Foreground(Green). // Green
		Background(Black). // Black
		Bold(true)

	fmt.Println(customStyle.Render("Custom styled text"))
}

// ExampleThemeSwitch shows how theme switching would work
func ExampleThemeSwitch() {
	fmt.Println("=== Theme Switching Example ===")

	// This would be controlled by config in real usage
	// For now, we default to dark theme

	fmt.Println("Current theme: Dark")
	fmt.Println("Success color:", GetSuccessAnsiStyle().Render("Green"))
	fmt.Println("Error color:", GetErrorAnsiStyle().Render("Red"))

	// In a real implementation, you would:
	// 1. Read theme from config
	// 2. Update isLightTheme() function to check config
	// 3. All colors would automatically switch
}
