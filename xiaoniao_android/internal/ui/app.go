package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"github.com/kaminoguo/xiaoniao-android/internal/config"
	"github.com/kaminoguo/xiaoniao-android/internal/ui/components"
	"github.com/kaminoguo/xiaoniao-android/internal/ui/screens/about"
	"github.com/kaminoguo/xiaoniao-android/internal/ui/screens/settings"
	"github.com/kaminoguo/xiaoniao-android/internal/ui/screens/status"
	"github.com/kaminoguo/xiaoniao-android/internal/ui/screens/translate"
	"github.com/kaminoguo/xiaoniao/internal/translator"
)

// MainApp is the main application structure
type MainApp struct {
	window        fyne.Window
	version       string
	config        *config.Config
	translator    *translator.Translator
	currentScreen int

	// Screens
	translateScreen *translate.Screen
	statusScreen    *status.Screen
	settingsScreen  *settings.Screen
	aboutScreen     *about.Screen

	// UI components
	bottomNav   *components.BottomNav
	contentArea *fyne.Container
}

// NewMainApp creates a new main application
func NewMainApp(window fyne.Window, version string) *MainApp {
	app := &MainApp{
		window:        window,
		version:       version,
		currentScreen: 0,
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Printf("Failed to load config: %v, using defaults", err)
		cfg = config.NewConfig()
	}
	app.config = cfg

	// Initialize translator
	app.translator = translator.New()

	// Initialize screens
	app.translateScreen = translate.NewScreen(app.config, app.translator, app.window)
	app.statusScreen = status.NewScreen(app.config, app.window)
	app.settingsScreen = settings.NewScreen(app.config, app.translator, app.window)
	app.aboutScreen = about.NewScreen(version, app.window)

	return app
}

// CreateUI creates the main UI
func (app *MainApp) CreateUI() fyne.CanvasObject {
	// Create content area container
	app.contentArea = container.NewMax()

	// Create bottom navigation
	app.bottomNav = components.NewBottomNav(func(index int) {
		app.switchScreen(index)
	})

	// Initialize with translate screen (first tab)
	app.contentArea.Objects = []fyne.CanvasObject{app.translateScreen.CreateUI()}

	// Create main container with bottom navigation
	mainContent := container.NewBorder(
		nil,               // top
		app.bottomNav,     // bottom
		nil,               // left
		nil,               // right
		app.contentArea,   // center
	)

	return mainContent
}

// switchScreen switches between screens
func (app *MainApp) switchScreen(index int) {
	if index == app.currentScreen {
		return
	}

	app.currentScreen = index

	var newScreen fyne.CanvasObject
	switch index {
	case 0: // Translate
		newScreen = app.translateScreen.CreateUI()
	case 1: // Status
		newScreen = app.statusScreen.CreateUI()
	case 2: // Settings
		newScreen = app.settingsScreen.CreateUI()
	case 3: // About
		newScreen = app.aboutScreen.CreateUI()
	default:
		return
	}

	// Update content area
	app.contentArea.Objects = []fyne.CanvasObject{newScreen}
	app.contentArea.Refresh()
}

// SaveConfig saves the configuration
func (app *MainApp) SaveConfig() {
	if app.config != nil {
		if err := app.config.Save(); err != nil {
			log.Printf("Failed to save config: %v", err)
		}
	}
}

// Cleanup performs cleanup operations
func (app *MainApp) Cleanup() {
	// Stop any running services
	if app.statusScreen != nil {
		app.statusScreen.StopMonitoring()
	}

	// Save configuration
	app.SaveConfig()

	log.Println("Application cleanup completed")
}