package status

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/kaminoguo/xiaoniao-android/internal/config"
	"github.com/kaminoguo/xiaoniao/internal/i18n"
)

// Screen is the status screen
type Screen struct {
	config  *config.Config
	window  fyne.Window
	running bool

	// UI components
	statusLabel  *widget.Label
	countLabel   *widget.Label
	modelLabel   *widget.Label
	modeRadio    *widget.RadioGroup
	startBtn     *widget.Button
	stopBtn      *widget.Button

	// Statistics
	todayCount int
}

// NewScreen creates a new status screen
func NewScreen(cfg *config.Config, window fyne.Window) *Screen {
	return &Screen{
		config:     cfg,
		window:     window,
		running:    false,
		todayCount: 0,
	}
}

// CreateUI creates the UI for the status screen
func (s *Screen) CreateUI() fyne.CanvasObject {
	// Title
	title := widget.NewLabelWithStyle(
		i18n.T("status.title"),
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	// Monitor mode selection
	modeLabel := widget.NewLabel(i18n.T("status.monitor_mode"))
	s.modeRadio = widget.NewRadioGroup([]string{
		i18n.T("mode.off"),
		i18n.T("mode.text_menu"),
		i18n.T("mode.clipboard"),
	}, func(selected string) {
		// Map display name to mode
		switch selected {
		case i18n.T("mode.off"):
			s.config.SetMonitorMode("off")
			s.updateStatus(false)
		case i18n.T("mode.text_menu"):
			s.config.SetMonitorMode("text_menu")
			s.updateStatus(false)
		case i18n.T("mode.clipboard"):
			s.config.SetMonitorMode("clipboard")
		}
	})

	// Set current mode
	switch s.config.GetMonitorMode() {
	case "off":
		s.modeRadio.SetSelected(i18n.T("mode.off"))
	case "text_menu":
		s.modeRadio.SetSelected(i18n.T("mode.text_menu"))
	case "clipboard":
		s.modeRadio.SetSelected(i18n.T("mode.clipboard"))
	}

	// Status information card
	s.statusLabel = widget.NewLabel(s.getStatusText())
	s.statusLabel.TextStyle = fyne.TextStyle{Bold: true}

	s.countLabel = widget.NewLabel(fmt.Sprintf("%s: %d %s",
		i18n.T("status.today"),
		s.todayCount,
		i18n.T("status.times")))

	s.modelLabel = widget.NewLabel(fmt.Sprintf("%s: %s",
		i18n.T("status.model"),
		s.config.GetModel()))

	statusCard := widget.NewCard("", "", container.NewVBox(
		s.statusLabel,
		widget.NewSeparator(),
		s.countLabel,
		s.modelLabel,
	))

	// Control buttons for clipboard mode
	s.startBtn = widget.NewButtonWithIcon(i18n.T("btn.start_monitor"), theme.MediaPlayIcon(), func() {
		s.startMonitoring()
	})
	s.startBtn.Importance = widget.HighImportance

	s.stopBtn = widget.NewButtonWithIcon(i18n.T("btn.stop_monitor"), theme.MediaStopIcon(), func() {
		s.StopMonitoring()
	})
	s.stopBtn.Importance = widget.DangerImportance
	s.stopBtn.Disable()

	controlButtons := container.New(layout.NewGridLayout(2),
		s.startBtn,
		s.stopBtn,
	)

	// Only show control buttons for clipboard mode
	s.updateControlButtons()

	// Main content
	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		modeLabel,
		s.modeRadio,
		widget.NewSeparator(),
		statusCard,
		controlButtons,
	)

	// Add scrolling
	scroll := container.NewVScroll(content)
	return scroll
}

// startMonitoring starts clipboard monitoring
func (s *Screen) startMonitoring() {
	if s.config.GetMonitorMode() != "clipboard" {
		return
	}

	s.running = true
	s.updateStatus(true)
	s.startBtn.Disable()
	s.stopBtn.Enable()

	log.Println("Started clipboard monitoring")

	// TODO: Start actual clipboard monitoring service
}

// StopMonitoring stops clipboard monitoring
func (s *Screen) StopMonitoring() {
	s.running = false
	s.updateStatus(false)
	s.startBtn.Enable()
	s.stopBtn.Disable()

	log.Println("Stopped clipboard monitoring")

	// TODO: Stop actual clipboard monitoring service
}

// updateStatus updates the status display
func (s *Screen) updateStatus(running bool) {
	s.running = running
	if s.statusLabel != nil {
		s.statusLabel.SetText(s.getStatusText())
	}
}

// getStatusText returns the current status text
func (s *Screen) getStatusText() string {
	mode := s.config.GetMonitorMode()

	switch mode {
	case "off":
		return fmt.Sprintf("● %s", i18n.T("status.disabled"))
	case "text_menu":
		return fmt.Sprintf("● %s", i18n.T("status.text_menu_enabled"))
	case "clipboard":
		if s.running {
			return fmt.Sprintf("● %s", i18n.T("status.monitoring"))
		}
		return fmt.Sprintf("● %s", i18n.T("status.stopped"))
	default:
		return fmt.Sprintf("● %s", i18n.T("status.unknown"))
	}
}

// updateControlButtons updates the visibility of control buttons
func (s *Screen) updateControlButtons() {
	if s.config.GetMonitorMode() == "clipboard" {
		if s.startBtn != nil {
			s.startBtn.Show()
		}
		if s.stopBtn != nil {
			s.stopBtn.Show()
		}
	} else {
		if s.startBtn != nil {
			s.startBtn.Hide()
		}
		if s.stopBtn != nil {
			s.stopBtn.Hide()
		}
	}
}

// IncrementCount increments the translation count
func (s *Screen) IncrementCount() {
	s.todayCount++
	if s.countLabel != nil {
		s.countLabel.SetText(fmt.Sprintf("%s: %d %s",
			i18n.T("status.today"),
			s.todayCount,
			i18n.T("status.times")))
	}
}