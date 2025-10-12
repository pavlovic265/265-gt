package constants

type Theme string

var (
	DarkTheme  Theme = "dark"
	LightTheme Theme = "light"
)

func (t Theme) String() string {
	switch t {
	case DarkTheme:
		return "dark"
	case LightTheme:
		return "light"
	default:
		return ""
	}
}
