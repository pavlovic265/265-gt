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

### Usage Example

```go
import "github.com/pavlovic265/265-gt/config"

// Success message
fmt.Println(config.SuccessStyle.Render("✅ Operation completed successfully"))

// Error message
fmt.Println(config.ErrorStyle.Render("❌ Something went wrong"))

// Info message
fmt.Println(config.InfoStyle.Render("ℹ️  This is informational"))

// Warning message
fmt.Println(config.WarningStyle.Render("⚠️  Please be careful"))

// Debug/Meta information
fmt.Println(config.DebugStyle.Render("🔍 Debug info"))

// Highlighted text
fmt.Println(config.HighlightStyle.Render("✨ Special highlight"))

// Title with background
fmt.Println(config.TitleStyle.Render("📋 Section Title"))
```

### Best Practices

1. **Use semantic colors**: Choose colors based on the message purpose, not just appearance
2. **Be consistent**: Use the same color for the same type of message across all commands
3. **Accessibility**: The colors are chosen to be readable in most terminal environments
4. **Import once**: Import the config package once at the top of your file

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
    fmt.Printf("Status: %s\n", config.SuccessStyle.Render("OK"))
}
``` 