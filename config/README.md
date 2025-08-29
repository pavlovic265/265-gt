# Shared Color Palette

This directory contains shared configuration and styling for the GT application.

## Color Palette Usage

The color palette is defined in `colors.go` and follows a semantic purpose-based approach:

### Available Styles

| Purpose | Color | Hex | Style Variable |
|---------|-------|-----|----------------|
| **Success** | Turquoise-Green | #19F9D8 | `config.SuccessStyle` |
| **Error** | Pink | #FF76B5 | `config.ErrorStyle` |
| **Warning** | Orange | #F7B36A | `config.WarningStyle` |
| **Info** | Light Blue | #6FC1FF | `config.InfoStyle` |
| **Debug/Meta** | Light Purple | #B180D7 | `config.DebugStyle` |
| **Highlight** | Light Pink | #FF87B4 | `config.HighlightStyle` |
| **Title** | Success + Highlight BG | #19F9D8 + #FF87B4 | `config.TitleStyle` |

### Status Indicators

For consistent status messages, use the built-in indicator functions with ASCII icons:

| Function | Icon | Usage | Example |
|----------|------|-------|---------|
| `SuccessIndicator(message)` | ✓ | Success messages | `✓ Operation completed` |
| `ErrorIndicator(message)` | ✗ | Error messages | `✗ Something went wrong` |
| `SuccessIconOnly()` | ✓ | Just the success icon | `✓` |
| `ErrorIconOnly()` | ✗ | Just the error icon | `✗` |

### Usage Example

```go
import "github.com/pavlovic265/265-gt/config"

// Status indicators with icons
fmt.Println(config.SuccessIndicator("Operation completed successfully"))
fmt.Println(config.ErrorIndicator("Something went wrong"))

// Just the icons
fmt.Printf("%s Operation completed\n", config.SuccessIconOnly())
fmt.Printf("%s Error occurred\n", config.ErrorIconOnly())

// Info message
fmt.Println(config.InfoStyle.Render("This is informational"))

// Warning message
fmt.Println(config.WarningStyle.Render("Please be careful"))

// Debug/Meta information
fmt.Println(config.DebugStyle.Render("Debug info"))

// Highlighted text
fmt.Println(config.HighlightStyle.Render("Special highlight"))

// Title with background
fmt.Println(config.TitleStyle.Render("Section Title"))
```

### Best Practices

1. **Use status indicators**: Prefer `SuccessIndicator()` and `ErrorIndicator()` for consistent messaging
2. **Use semantic colors**: Choose colors based on the message purpose, not just appearance
3. **Be consistent**: Use the same color for the same type of message across all commands
4. **Accessibility**: The colors are chosen to be readable in most terminal environments
5. **Import once**: Import the config package once at the top of your file

### Adding New Commands

When creating new commands, import the config package and use the shared styles:

```go
package commands

import (
    "fmt"
    "github.com/pavlovic265/265-gt/config"
)

func someCommand() {
    fmt.Println(config.TitleStyle.Render("My Command"))
    fmt.Println(config.SuccessIndicator("Operation successful"))
    fmt.Printf("Status: %s\n", config.SuccessIconOnly())
}
``` 