# Shared Color Palette

This directory contains shared configuration and styling for the GT application.

## Color Palette Usage

The color palette is defined in `colors.go` and follows the **Panda Syntax** theme with both dark and light variants. The colors are chosen for excellent readability and visual appeal in terminal environments.

### Dark Theme Colors (Default)

| Purpose | Color | Hex | Style Variable |
|---------|-------|-----|----------------|
| **Success** | Green | #A9DC52 | `config.SuccessStyle` |
| **Error** | Red | #FF6188 | `config.ErrorStyle` |
| **Warning** | Yellow | #FFD866 | `config.WarningStyle` |
| **Info** | Blue | #78DCE8 | `config.InfoStyle` |
| **Debug/Meta** | Purple | #AB9DF2 | `config.DebugStyle` |
| **Highlight** | Yellow | #FFD866 | `config.HighlightStyle` |
| **Command** | Blue | #78DCE8 | `config.CommandStyle` |
| **Branch** | Purple | #AB9DF2 | `config.BranchStyle` |
| **File** | Yellow (Italic) | #FFD866 | `config.FileStyle` |
| **Status** | Yellow | #FFD866 | `config.StatusStyle` |
| **Title** | Green + Yellow BG | #A9DC52 + #FFD866 | `config.TitleStyle` |

### Light Theme Colors

| Purpose | Color | Hex | Style Variable |
|---------|-------|-----|----------------|
| **Success** | Dark Green | #2D5016 | `config.LightSuccessStyle` |
| **Error** | Dark Red | #8B1538 | `config.LightErrorStyle` |
| **Warning** | Dark Orange | #B8860B | `config.LightWarningStyle` |
| **Info** | Dark Blue | #1E3A8A | `config.LightInfoStyle` |
| **Debug/Meta** | Dark Purple | #6B46C1 | `config.LightDebugStyle` |
| **Highlight** | Dark Yellow | #B8860B | `config.LightHighlightStyle` |
| **Command** | Dark Blue | #1E3A8A | `config.LightCommandStyle` |
| **Branch** | Dark Purple | #6B46C1 | `config.LightBranchStyle` |
| **File** | Dark Yellow (Italic) | #B8860B | `config.LightFileStyle` |
| **Status** | Dark Orange | #B8860B | `config.LightStatusStyle` |

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

// Status indicators with icons (automatically themed)
fmt.Println(config.SuccessIndicator("Operation completed successfully"))
fmt.Println(config.ErrorIndicator("Something went wrong"))

// Just the icons (automatically themed)
fmt.Printf("%s Operation completed\n", config.SuccessIconOnly())
fmt.Printf("%s Error occurred\n", config.ErrorIconOnly())

// Theme-aware styling (automatically switches between dark/light)
fmt.Println(config.GetInfoStyle().Render("This is informational"))
fmt.Println(config.GetWarningStyle().Render("Please be careful"))
fmt.Println(config.GetDebugStyle().Render("Debug info"))
fmt.Println(config.GetHighlightStyle().Render("Special highlight"))
fmt.Println(config.GetCommandStyle().Render("Command name"))
fmt.Println(config.GetBranchStyle().Render("branch-name"))
fmt.Println(config.GetFileStyle().Render("filename.txt"))
fmt.Println(config.GetStatusStyle().Render("Status message"))

// Direct styling (always uses dark theme colors)
fmt.Println(config.InfoStyle.Render("This is informational"))
fmt.Println(config.WarningStyle.Render("Please be careful"))
fmt.Println(config.DebugStyle.Render("Debug info"))
fmt.Println(config.HighlightStyle.Render("Special highlight"))

// Title with background (always uses dark theme)
fmt.Println(config.TitleStyle.Render("Section Title"))
```

### Best Practices

1. **Use status indicators**: Prefer `SuccessIndicator()` and `ErrorIndicator()` for consistent messaging
2. **Use theme-aware functions**: Prefer `Get*Style()` functions over direct style access for automatic dark/light theme support
3. **Use semantic colors**: Choose colors based on the message purpose, not just appearance
4. **Be consistent**: Use the same color for the same type of message across all commands
5. **Accessibility**: The colors are chosen to be readable in most terminal environments
6. **Theme support**: The palette automatically adapts to user's theme preference (dark/light)
7. **Import once**: Import the config package once at the top of your file

### Adding New Commands

When creating new commands, import the config package and use the shared styles:

```go
package commands

import (
    "fmt"
    "github.com/pavlovic265/265-gt/config"
)

func someCommand() {
    // Use theme-aware styling for better user experience
    fmt.Println(config.TitleStyle.Render("My Command"))
    fmt.Println(config.SuccessIndicator("Operation successful"))
    fmt.Printf("Status: %s\n", config.SuccessIconOnly())
    
    // Use semantic styling for different types of content
    fmt.Println(config.GetInfoStyle().Render("Information message"))
    fmt.Println(config.GetWarningStyle().Render("Warning message"))
    fmt.Println(config.GetCommandStyle().Render("command-name"))
    fmt.Println(config.GetBranchStyle().Render("branch-name"))
    fmt.Println(config.GetFileStyle().Render("filename.txt"))
}
``` 