# Color Palette Documentation

This directory contains the color palette definitions for the 265-gt tool.

## Files

- `colors.go` - Main color definitions and styling functions
- `platform.go` - Platform constants (GitHub, GitLab)
- `example_usage.go` - Examples of how to use the color palette

## Color Palettes

### 1. Panda Syntax Colors (Original)
The original Panda Syntax color scheme with dark and light variants.

### 2. 16-Color ANSI Palette (New)
A comprehensive 16-color ANSI palette with both dark and light themes.

#### Dark Theme Colors
- **Background**: `#1d1e20` (dark gray)
- **Foreground**: `#e6e6e6` (light gray)
- **Cursor**: `#ffb86c` (orange)

**ANSI Colors (0-15):**
- `0`: `#1d1e20` (black)
- `1`: `#ff5c57` (red)
- `2`: `#A9DC52` (green - Panda success)
- `3`: `#f3f99d` (yellow)
- `4`: `#57c7ff` (blue)
- `5`: `#ff6ac1` (magenta)
- `6`: `#9aedfe` (cyan)
- `7`: `#e6e6e6` (white)
- `8`: `#555555` (bright black)
- `9`: `#ff7a90` (bright red)
- `10`: `#69ff94` (bright green)
- `11`: `#ffffa5` (bright yellow)
- `12`: `#9aedfe` (bright blue)
- `13`: `#ff92d0` (bright magenta)
- `14`: `#c8ffff` (bright cyan)
- `15`: `#ffffff` (bright white)

#### Light Theme Colors
- **Background**: `#fafafa` (light gray)
- **Foreground**: `#2e2e2e` (dark gray)
- **Cursor**: `#ff5c57` (red)

**ANSI Colors (0-15):**
- `0`: `#2e2e2e` (black)
- `1`: `#ff5c57` (red)
- `2`: `#3bb273` (green)
- `3`: `#d4b106` (yellow)
- `4`: `#268bd2` (blue)
- `5`: `#af5fff` (magenta)
- `6`: `#00bcd4` (cyan)
- `7`: `#fafafa` (white)
- `8`: `#999999` (bright black)
- `9`: `#ff7b72` (bright red)
- `10`: `#44d88d` (bright green)
- `11`: `#ffe36e` (bright yellow)
- `12`: `#4fc3f7` (bright blue)
- `13`: `#c586c0` (bright magenta)
- `14`: `#4dd0e1` (bright cyan)
- `15`: `#ffffff` (bright white)

## Usage Examples

### Basic ANSI Colors
```go
import "github.com/pavlovic265/265-gt/constants"

// Get a specific ANSI color
style := constants.GetAnsiStyle(2) // Green
fmt.Println(style.Render("Success message"))
```

### Convenience Functions
```go
// Common use cases
fmt.Println(constants.GetSuccessAnsiStyle().Render("✓ Success"))
fmt.Println(constants.GetErrorAnsiStyle().Render("✗ Error"))
fmt.Println(constants.GetWarningAnsiStyle().Render("⚠ Warning"))
fmt.Println(constants.GetInfoAnsiStyle().Render("ℹ Info"))
fmt.Println(constants.GetDebugAnsiStyle().Render("[D] Debug"))
```

### Background and Foreground
```go
// Get theme-aware colors
bgColor := constants.Background
fgColor := constants.Foreground

style := lipgloss.NewStyle().
    Background(bgColor).
    Foreground(fgColor)
```

### Custom Combinations
```go
// Create custom styles using named colors
customStyle := lipgloss.NewStyle().
    Foreground(constants.Green).  // Green
    Background(constants.Black).  // Black
    Bold(true)
```

## Theme Switching

The color palette automatically adapts to the current theme. To switch themes:

1. Update the `isLightTheme()` function in `colors.go` to read from config
2. All colors will automatically switch between dark and light variants

## Integration with Config

The color palette is designed to work with the config system:

```yaml
# ~/.gtconfig.yaml
theme:
  type: "dark"   # or "light"
```

When theme switching is implemented, all colors will automatically adapt to the selected theme.
