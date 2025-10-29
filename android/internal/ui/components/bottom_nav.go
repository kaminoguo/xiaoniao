package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/kaminoguo/xiaoniao/internal/i18n"
)

// BottomNav is the bottom navigation component
type BottomNav struct {
	widget.BaseWidget

	currentTab    int
	onTabSelected func(int)

	// Tab buttons
	translateBtn *widget.Button
	statusBtn    *widget.Button
	settingsBtn  *widget.Button
	aboutBtn     *widget.Button

	container *fyne.Container
}

// NewBottomNav creates a new bottom navigation
func NewBottomNav(onTabSelected func(int)) *BottomNav {
	nav := &BottomNav{
		currentTab:    0,
		onTabSelected: onTabSelected,
	}

	// Create tab buttons with icons
	nav.translateBtn = widget.NewButtonWithIcon(i18n.T("nav.translate"), theme.DocumentIcon(), func() {
		nav.selectTab(0)
	})

	nav.statusBtn = widget.NewButtonWithIcon(i18n.T("nav.status"), theme.InfoIcon(), func() {
		nav.selectTab(1)
	})

	nav.settingsBtn = widget.NewButtonWithIcon(i18n.T("nav.settings"), theme.SettingsIcon(), func() {
		nav.selectTab(2)
	})

	nav.aboutBtn = widget.NewButtonWithIcon(i18n.T("nav.about"), theme.HelpIcon(), func() {
		nav.selectTab(3)
	})

	// Set initial selection style
	nav.updateButtonStyles()

	nav.ExtendBaseWidget(nav)
	return nav
}

// CreateRenderer creates the renderer for the bottom navigation
func (nav *BottomNav) CreateRenderer() fyne.WidgetRenderer {
	// Create buttons container
	buttonsContainer := container.NewGridWithColumns(4,
		nav.translateBtn,
		nav.statusBtn,
		nav.settingsBtn,
		nav.aboutBtn,
	)

	nav.container = buttonsContainer

	return widget.NewSimpleRenderer(nav.container)
}

// selectTab selects a tab
func (nav *BottomNav) selectTab(index int) {
	if index == nav.currentTab {
		return
	}

	nav.currentTab = index
	nav.updateButtonStyles()

	// Trigger callback
	if nav.onTabSelected != nil {
		nav.onTabSelected(index)
	}

	nav.Refresh()
}

// updateButtonStyles updates the button styles based on selection
func (nav *BottomNav) updateButtonStyles() {
	// Reset all buttons to low importance
	nav.translateBtn.Importance = widget.LowImportance
	nav.statusBtn.Importance = widget.LowImportance
	nav.settingsBtn.Importance = widget.LowImportance
	nav.aboutBtn.Importance = widget.LowImportance

	// Set selected button to high importance
	switch nav.currentTab {
	case 0:
		nav.translateBtn.Importance = widget.HighImportance
	case 1:
		nav.statusBtn.Importance = widget.HighImportance
	case 2:
		nav.settingsBtn.Importance = widget.HighImportance
	case 3:
		nav.aboutBtn.Importance = widget.HighImportance
	}

	// Refresh buttons
	nav.translateBtn.Refresh()
	nav.statusBtn.Refresh()
	nav.settingsBtn.Refresh()
	nav.aboutBtn.Refresh()
}