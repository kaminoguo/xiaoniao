package translate

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/kaminoguo/xiaoniao-android/internal/config"
	"github.com/kaminoguo/xiaoniao/internal/i18n"
	"github.com/kaminoguo/xiaoniao/internal/translator"
)

// Screen is the translation screen (Google Translate-like)
type Screen struct {
	config     *config.Config
	translator *translator.Translator
	window     fyne.Window

	// UI components
	promptSelect  *widget.Select
	inputEntry    *widget.Entry
	outputEntry   *widget.Entry
	translateBtn  *widget.Button
	clearBtn      *widget.Button
	copyBtn       *widget.Button
}

// NewScreen creates a new translation screen
func NewScreen(cfg *config.Config, trans *translator.Translator, window fyne.Window) *Screen {
	return &Screen{
		config:     cfg,
		translator: trans,
		window:     window,
	}
}

// CreateUI creates the UI for the translation screen
func (s *Screen) CreateUI() fyne.CanvasObject {
	// Title
	title := widget.NewLabelWithStyle(
		"xiaoniao",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	// Prompt selector
	prompts := []string{
		i18n.T("prompt.simple"),
		i18n.T("prompt.detailed"),
		i18n.T("prompt.academic"),
		i18n.T("prompt.casual"),
	}
	s.promptSelect = widget.NewSelect(prompts, func(selected string) {
		// Map display name to actual prompt
		switch selected {
		case i18n.T("prompt.simple"):
			s.config.SetCurrentPrompt(translator.SimplePrompt)
		case i18n.T("prompt.detailed"):
			s.config.SetCurrentPrompt(translator.DetailPrompt)
		case i18n.T("prompt.academic"):
			s.config.SetCurrentPrompt(translator.AcademicPrompt)
		case i18n.T("prompt.casual"):
			s.config.SetCurrentPrompt(translator.CasualPrompt)
		}
	})
	s.promptSelect.SetSelected(i18n.T("prompt.simple"))

	// Input text area
	s.inputEntry = widget.NewMultiLineEntry()
	s.inputEntry.SetPlaceHolder(i18n.T("translate.input_placeholder"))
	s.inputEntry.Resize(fyne.NewSize(0, 150))

	// Clear and translate buttons
	s.clearBtn = widget.NewButtonWithIcon(i18n.T("btn.clear"), theme.ContentClearIcon(), func() {
		s.inputEntry.SetText("")
		s.outputEntry.SetText("")
	})

	s.translateBtn = widget.NewButtonWithIcon(i18n.T("btn.translate"), theme.NavigateNextIcon(), func() {
		s.doTranslation()
	})
	s.translateBtn.Importance = widget.HighImportance

	inputButtons := container.New(layout.NewGridLayout(2),
		s.clearBtn,
		s.translateBtn,
	)

	// Output text area
	s.outputEntry = widget.NewMultiLineEntry()
	s.outputEntry.SetPlaceHolder(i18n.T("translate.output_placeholder"))
	s.outputEntry.Resize(fyne.NewSize(0, 150))
	s.outputEntry.Disable() // Read-only

	// Copy button
	s.copyBtn = widget.NewButtonWithIcon(i18n.T("btn.copy"), theme.ContentCopyIcon(), func() {
		if s.outputEntry.Text != "" {
			s.window.Clipboard().SetContent(s.outputEntry.Text)
			dialog.ShowInformation(i18n.T("info.copied"), i18n.T("info.translation_copied"), s.window)
		}
	})

	// Main content
	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		s.promptSelect,
		s.inputEntry,
		inputButtons,
		widget.NewSeparator(),
		s.outputEntry,
		s.copyBtn,
	)

	// Add scrolling
	scroll := container.NewVScroll(content)
	return scroll
}

// doTranslation performs the translation
func (s *Screen) doTranslation() {
	text := s.inputEntry.Text
	if text == "" {
		return
	}

	// Show loading
	s.outputEntry.SetText(i18n.T("translate.translating"))
	s.translateBtn.Disable()

	// Configure translator
	s.translator.SetAPIKey(s.config.GetAPIKey())
	s.translator.SetModel(s.config.GetModel())
	s.translator.SetPrompt(s.config.GetCurrentPrompt())

	// Perform translation in background
	go func() {
		ctx := context.Background()
		result, err := s.translator.Translate(ctx, text)

		// Update UI on main thread
		s.window.Canvas().Content().Refresh()

		if err != nil {
			s.outputEntry.SetText(fmt.Sprintf("%s: %v", i18n.T("error.translation_failed"), err))
		} else {
			s.outputEntry.SetText(result)
		}

		s.translateBtn.Enable()
	}()
}