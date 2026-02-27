# Constants Package

This package contains domain constants used by core application logic, such as:

- Platforms (`GitHub`, `GitLab`)
- Branch relationship keys
- Config file names
- Theme enum values (`dark`, `light`)

UI styling values (colors, icons, and lipgloss style helpers) were moved to:

- `ui/theme`

Use `constants` for non-UI/domain values and `ui/theme` for presentation concerns.
