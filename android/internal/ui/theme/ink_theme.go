package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// InkTheme 水墨风主题
type InkTheme struct{}

// 水墨风配色方案
var (
	// 基础色
	InkBlack   = color.NRGBA{R: 28, G: 28, B: 28, A: 255}    // 墨黑
	PaperWhite = color.NRGBA{R: 245, G: 245, B: 240, A: 255} // 宣纸白

	// 灰度层次（水墨渐变）
	Ink90 = color.NRGBA{R: 40, G: 40, B: 40, A: 230}    // 浓墨
	Ink70 = color.NRGBA{R: 60, G: 60, B: 60, A: 180}    // 中墨
	Ink50 = color.NRGBA{R: 100, G: 100, B: 100, A: 130} // 淡墨
	Ink30 = color.NRGBA{R: 150, G: 150, B: 150, A: 80}  // 轻墨
	Ink10 = color.NRGBA{R: 200, G: 200, B: 200, A: 50}  // 水痕

	// 功能色（极简）
	ErrorRed     = color.NRGBA{R: 139, G: 0, B: 0, A: 255}  // 朱砂红
	SuccessGreen = color.NRGBA{R: 34, G: 87, B: 34, A: 255} // 竹青

	// 小鸟图标状态色（复用桌面版）
	BirdBlue  = color.NRGBA{R: 30, G: 144, B: 255, A: 255} // 道奇蓝
	BirdGreen = color.NRGBA{R: 50, G: 205, B: 50, A: 255}  // 石灰绿
	BirdRed   = color.NRGBA{R: 255, G: 69, B: 0, A: 255}   // 橙红色
)

// Color 返回主题颜色
func (t InkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return PaperWhite
	case theme.ColorNameForeground:
		return InkBlack
	case theme.ColorNameButton:
		return Ink10
	case theme.ColorNameDisabledButton:
		return Ink30
	case theme.ColorNamePrimary:
		return Ink90
	case theme.ColorNameHover:
		return Ink30
	case theme.ColorNameFocus:
		return Ink70
	case theme.ColorNameSelection:
		return Ink50
	case theme.ColorNameBorder:
		return Ink30
	case theme.ColorNameShadow:
		return Ink10
	case theme.ColorNameError:
		return ErrorRed
	case theme.ColorNameSuccess:
		return SuccessGreen
	case theme.ColorNameWarning:
		return color.NRGBA{R: 255, G: 165, B: 0, A: 255} // 橙色
	case theme.ColorNameDisabled:
		return Ink30
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	case theme.ColorNameInputBorder:
		return Ink50
	case theme.ColorNamePlaceHolder:
		return Ink30
	case theme.ColorNamePressed:
		return Ink70
	case theme.ColorNameScrollBar:
		return Ink30
	case theme.ColorNameSeparator:
		return Ink10
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

// Font 返回字体
func (t InkTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

// Icon 返回图标
func (t InkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// Size 返回尺寸
func (t InkTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 8
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNameScrollBar:
		return 12
	case theme.SizeNameScrollBarSmall:
		return 6
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameText:
		return 14
	case theme.SizeNameHeadingText:
		return 24
	case theme.SizeNameSubHeadingText:
		return 18
	case theme.SizeNameCaptionText:
		return 12
	case theme.SizeNameInputBorder:
		return 2
	default:
		return theme.DefaultTheme().Size(name)
	}
}