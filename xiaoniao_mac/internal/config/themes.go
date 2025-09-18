package config

import (
	"github.com/charmbracelet/lipgloss"
)

// Theme 定义UI主题
type Theme struct {
	Name        string
	Description string
	Colors      ThemeColors
	Borders     ThemeBorders
}

// ThemeColors 定义主题颜色
type ThemeColors struct {
	Primary     lipgloss.Color
	Secondary   lipgloss.Color
	Accent      lipgloss.Color
	Text        lipgloss.Color
	Background  lipgloss.Color
	Border      lipgloss.Color
	Success     lipgloss.Color
	Warning     lipgloss.Color
	Error       lipgloss.Color
	Info        lipgloss.Color
}

// ThemeBorders 定义边框样式
type ThemeBorders struct {
	Normal lipgloss.Border
	Double lipgloss.Border
	Thick  lipgloss.Border
}

// Themes 预定义主题
var Themes = map[string]Theme{
	"default": {
		Name:        "默认",
		Description: "经典蓝色主题",
		Colors: ThemeColors{
			Primary:    lipgloss.Color("#3498db"),
			Secondary:  lipgloss.Color("#2ecc71"),
			Accent:     lipgloss.Color("#f39c12"),
			Text:       lipgloss.Color("#ffffff"),
			Background: lipgloss.Color("0"),
			Border:     lipgloss.Color("#3498db"),
			Success:    lipgloss.Color("#2ecc71"),
			Warning:    lipgloss.Color("#f39c12"),
			Error:      lipgloss.Color("#e74c3c"),
			Info:       lipgloss.Color("#3498db"),
		},
		Borders: ThemeBorders{
			Normal: lipgloss.NormalBorder(),
			Double: lipgloss.DoubleBorder(),
			Thick:  lipgloss.ThickBorder(),
		},
	},
	"tokyo-night": {
		Name:        "Tokyo Night",
		Description: "暗色主题，灵感来自东京夜景",
		Colors: ThemeColors{
			Primary:    lipgloss.Color("#7aa2f7"),
			Secondary:  lipgloss.Color("#9ece6a"),
			Accent:     lipgloss.Color("#ff9e64"),
			Text:       lipgloss.Color("#a9b1d6"),
			Background: lipgloss.Color("#1a1b26"),
			Border:     lipgloss.Color("#565f89"),
			Success:    lipgloss.Color("#9ece6a"),
			Warning:    lipgloss.Color("#e0af68"),
			Error:      lipgloss.Color("#f7768e"),
			Info:       lipgloss.Color("#7aa2f7"),
		},
		Borders: ThemeBorders{
			Normal: lipgloss.RoundedBorder(),
			Double: lipgloss.DoubleBorder(),
			Thick:  lipgloss.ThickBorder(),
		},
	},
	"catppuccin": {
		Name:        "Catppuccin",
		Description: "柔和的粉彩主题",
		Colors: ThemeColors{
			Primary:    lipgloss.Color("#89b4fa"),
			Secondary:  lipgloss.Color("#a6e3a1"),
			Accent:     lipgloss.Color("#fab387"),
			Text:       lipgloss.Color("#cdd6f4"),
			Background: lipgloss.Color("#1e1e2e"),
			Border:     lipgloss.Color("#45475a"),
			Success:    lipgloss.Color("#a6e3a1"),
			Warning:    lipgloss.Color("#f9e2af"),
			Error:      lipgloss.Color("#f38ba8"),
			Info:       lipgloss.Color("#89b4fa"),
		},
		Borders: ThemeBorders{
			Normal: lipgloss.RoundedBorder(),
			Double: lipgloss.DoubleBorder(),
			Thick:  lipgloss.ThickBorder(),
		},
	},
	"minimal": {
		Name:        "极简",
		Description: "简洁的黑白主题",
		Colors: ThemeColors{
			Primary:    lipgloss.Color("#ffffff"),
			Secondary:  lipgloss.Color("#888888"),
			Accent:     lipgloss.Color("#ffffff"),
			Text:       lipgloss.Color("#ffffff"),
			Background: lipgloss.Color("#000000"),
			Border:     lipgloss.Color("#888888"),
			Success:    lipgloss.Color("#ffffff"),
			Warning:    lipgloss.Color("#bbbbbb"),
			Error:      lipgloss.Color("#ffffff"),
			Info:       lipgloss.Color("#888888"),
		},
		Borders: ThemeBorders{
			Normal: lipgloss.NormalBorder(),
			Double: lipgloss.NormalBorder(),
			Thick:  lipgloss.NormalBorder(),
		},
	},
}

// GetTheme 获取主题
func GetTheme(name string) Theme {
	if theme, ok := Themes[name]; ok {
		return theme
	}
	return Themes["default"]
}

// GetThemeNames 获取所有主题名称
func GetThemeNames() []string {
	names := make([]string, 0, len(Themes))
	for name := range Themes {
		names = append(names, name)
	}
	return names
}