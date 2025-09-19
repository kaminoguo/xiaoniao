package settings

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/kaminoguo/xiaoniao-android/internal/config"
	"github.com/kaminoguo/xiaoniao/internal/i18n"
	"github.com/kaminoguo/xiaoniao/internal/translator"
)

// Screen is the settings screen
type Screen struct {
	config     *config.Config
	translator *translator.Translator
	window     fyne.Window
}

// NewScreen creates a new settings screen
func NewScreen(cfg *config.Config, trans *translator.Translator, window fyne.Window) *Screen {
	return &Screen{
		config:     cfg,
		translator: trans,
		window:     window,
	}
}

// CreateUI creates the UI for the settings screen
func (s *Screen) CreateUI() fyne.CanvasObject {
	// Title
	title := widget.NewLabelWithStyle(
		i18n.T("settings.title"),
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	// Settings list
	apiItem := widget.NewButtonWithIcon(i18n.T("settings.api"), theme.DocumentSaveIcon(), func() {
		s.showAPISettings()
	})
	apiItem.Alignment = widget.ButtonAlignLeading

	modelItem := widget.NewButtonWithIcon(i18n.T("settings.model"), theme.ComputerIcon(), func() {
		s.showModelSettings()
	})
	modelItem.Alignment = widget.ButtonAlignLeading

	promptItem := widget.NewButtonWithIcon(i18n.T("settings.prompts"), theme.DocumentCreateIcon(), func() {
		s.showPromptSettings()
	})
	promptItem.Alignment = widget.ButtonAlignLeading

	permItem := widget.NewButtonWithIcon(i18n.T("settings.permissions"), theme.WarningIcon(), func() {
		s.showPermissionSettings()
	})
	permItem.Alignment = widget.ButtonAlignLeading

	// Main content
	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		apiItem,
		widget.NewSeparator(),
		modelItem,
		widget.NewSeparator(),
		promptItem,
		widget.NewSeparator(),
		permItem,
	)

	// Add scrolling
	scroll := container.NewVScroll(content)
	return scroll
}

// showAPISettings shows the API configuration dialog
func (s *Screen) showAPISettings() {
	// API key entry
	apiKeyEntry := widget.NewPasswordEntry()
	apiKeyEntry.SetPlaceHolder("sk-...")
	apiKeyEntry.SetText(s.config.GetAPIKey())

	// API URL entry (optional)
	apiURLEntry := widget.NewEntry()
	apiURLEntry.SetPlaceHolder("https://api.openai.com (optional)")
	apiURLEntry.SetText(s.config.APIUrl)

	// Test button
	testBtn := widget.NewButton(i18n.T("btn.test_connection"), func() {
		if apiKeyEntry.Text == "" {
			dialog.ShowError(fmt.Errorf(i18n.T("error.api_key_required")), s.window)
			return
		}

		// Test the API connection
		s.translator.SetAPIKey(apiKeyEntry.Text)
		if apiURLEntry.Text != "" {
			// TODO: Set custom API URL
		}

		ctx := context.Background()
		_, err := s.translator.Translate(ctx, "Hello")
		if err != nil {
			dialog.ShowError(err, s.window)
		} else {
			dialog.ShowInformation(i18n.T("info.success"), i18n.T("info.api_test_success"), s.window)
		}
	})

	// Create form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: i18n.T("settings.api_key"), Widget: apiKeyEntry},
			{Text: i18n.T("settings.api_url"), Widget: apiURLEntry},
		},
	}

	// Show dialog
	dlg := dialog.NewCustomConfirm(
		i18n.T("settings.api"),
		i18n.T("btn.save"),
		i18n.T("btn.cancel"),
		container.NewVBox(form, testBtn),
		func(save bool) {
			if save {
				s.config.SetAPIKey(apiKeyEntry.Text)
				s.config.APIUrl = apiURLEntry.Text
				s.config.Save()
			}
		},
		s.window,
	)
	dlg.Resize(fyne.NewSize(300, 250))
	dlg.Show()
}

// showModelSettings shows the model selection dialog
func (s *Screen) showModelSettings() {
	// Model selection
	models := []string{
		"gpt-4o",
		"gpt-4o-mini",
		"gpt-3.5-turbo",
		"claude-3-5-sonnet-latest",
		"claude-3-5-haiku-latest",
		"gemini-2.0-flash-exp",
		"llama-3.3-70b-versatile",
	}

	modelRadio := widget.NewRadioGroup(models, nil)
	modelRadio.SetSelected(s.config.GetModel())

	// Test translation
	testInput := widget.NewEntry()
	testInput.SetPlaceHolder(i18n.T("test.input_text"))
	testInput.SetText("Hello, world!")

	testOutput := widget.NewMultiLineEntry()
	testOutput.SetPlaceHolder(i18n.T("test.output_text"))
	testOutput.Resize(fyne.NewSize(0, 100))
	testOutput.Disable()

	testBtn := widget.NewButton(i18n.T("btn.test_translate"), func() {
		if modelRadio.Selected == "" || testInput.Text == "" {
			return
		}

		testOutput.SetText(i18n.T("translate.translating"))

		// Test translation with selected model
		go func() {
			s.translator.SetAPIKey(s.config.GetAPIKey())
			s.translator.SetModel(modelRadio.Selected)

			ctx := context.Background()
			result, err := s.translator.Translate(ctx, testInput.Text)

			if err != nil {
				testOutput.SetText(fmt.Sprintf("Error: %v", err))
			} else {
				testOutput.SetText(result)
			}
		}()
	})

	// Show dialog
	content := container.NewVBox(
		widget.NewLabel(i18n.T("settings.select_model")),
		modelRadio,
		widget.NewSeparator(),
		widget.NewLabel(i18n.T("test.title")),
		testInput,
		testBtn,
		testOutput,
	)

	dlg := dialog.NewCustomConfirm(
		i18n.T("settings.model"),
		i18n.T("btn.save"),
		i18n.T("btn.cancel"),
		container.NewVScroll(content),
		func(save bool) {
			if save && modelRadio.Selected != "" {
				s.config.SetModel(modelRadio.Selected)
				s.config.Save()
			}
		},
		s.window,
	)
	dlg.Resize(fyne.NewSize(350, 400))
	dlg.Show()
}

// showPromptSettings shows the prompt management dialog
func (s *Screen) showPromptSettings() {
	// Prompt list
	prompts := []string{
		i18n.T("prompt.simple"),
		i18n.T("prompt.detailed"),
		i18n.T("prompt.academic"),
		i18n.T("prompt.casual"),
	}

	// Add custom prompts from config
	prompts = append(prompts, s.config.UserPrompts...)

	promptList := widget.NewList(
		func() int { return len(prompts) },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(prompts[i])
		},
	)

	// Add button
	addBtn := widget.NewButton(i18n.T("btn.add_prompt"), func() {
		// TODO: Show add prompt dialog
		dialog.ShowInformation(i18n.T("info.coming_soon"), i18n.T("info.feature_coming_soon"), s.window)
	})

	// Show dialog
	content := container.NewBorder(
		nil,
		addBtn,
		nil,
		nil,
		promptList,
	)

	dlg := dialog.NewCustom(
		i18n.T("settings.prompts"),
		i18n.T("btn.close"),
		content,
		s.window,
	)
	dlg.Resize(fyne.NewSize(300, 400))
	dlg.Show()
}

// showPermissionSettings shows the permission guide
func (s *Screen) showPermissionSettings() {
	// Permission status
	permissions := []struct {
		name   string
		status bool
		action func()
	}{
		{
			name:   i18n.T("perm.notification"),
			status: true, // TODO: Check actual permission
			action: nil,
		},
		{
			name:   i18n.T("perm.accessibility"),
			status: false, // TODO: Check actual permission
			action: func() {
				// TODO: Open system settings
				dialog.ShowInformation(i18n.T("info.open_settings"), i18n.T("info.open_accessibility_settings"), s.window)
			},
		},
		{
			name:   i18n.T("perm.overlay"),
			status: false, // TODO: Check actual permission
			action: func() {
				// TODO: Open system settings
				dialog.ShowInformation(i18n.T("info.open_settings"), i18n.T("info.open_overlay_settings"), s.window)
			},
		},
		{
			name:   i18n.T("perm.battery"),
			status: false, // TODO: Check actual permission
			action: func() {
				// TODO: Open system settings
				dialog.ShowInformation(i18n.T("info.open_settings"), i18n.T("info.open_battery_settings"), s.window)
			},
		},
	}

	// Create permission list
	content := container.NewVBox()
	for _, perm := range permissions {
		var statusLabel *widget.Label
		if perm.status {
			statusLabel = widget.NewLabel("✓ " + i18n.T("perm.granted"))
			statusLabel.TextStyle = fyne.TextStyle{Bold: true}
		} else {
			statusLabel = widget.NewLabel("✗ " + i18n.T("perm.not_granted"))
		}

		permCard := widget.NewCard(perm.name, "", statusLabel)

		if !perm.status && perm.action != nil {
			btn := widget.NewButton(i18n.T("btn.go_to_settings"), perm.action)
			permCard.SetContent(container.NewVBox(statusLabel, btn))
		}

		content.Add(permCard)
	}

	// Show dialog
	dlg := dialog.NewCustom(
		i18n.T("settings.permissions"),
		i18n.T("btn.close"),
		container.NewVScroll(content),
		s.window,
	)
	dlg.Resize(fyne.NewSize(350, 400))
	dlg.Show()
}