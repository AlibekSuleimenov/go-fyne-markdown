package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// config represents the configuration settings for the application.
type config struct {
	EditWidget    *widget.Entry    // EditWidget holds the widget for editing text input.
	PreviewWidget *widget.RichText // PreviewWidget holds the widget for displaying rich text preview.
	CurrentFile   fyne.URI         // CurrentFile represents the URI of the currently opened file.
	SaveMenuItem  *fyne.MenuItem   // SaveMenuItem holds the menu item for saving the current file.
}

// cfg represents the global configuration instance.
var cfg config

// Main function serves as the entry point for the application.
// Application logic will be implemented here.
func main() {
	// create a fyne app
	a := app.New()

	// create window
	win := a.NewWindow("Markdown")

	// get user interface
	edit, preview := cfg.makeUI()
	cfg.createMenuItems(win)
	// set the content of the window
	win.SetContent(container.NewHSplit(edit, preview))

	// show window and run the app
	win.Resize(fyne.Size{Width: 800, Height: 500})
	win.CenterOnScreen()
	win.ShowAndRun()
}

// makeUI creates and initializes the user interface components.
// It returns pointers to the Entry widget for text input and the RichText widget for displaying rich text.
func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")

	app.EditWidget = edit
	app.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

// createMenuItems creates and adds menu items to the main window.
func (app *config) createMenuItems(window fyne.Window) {
	openMenuItem := fyne.NewMenuItem("Open...", func() {})

	saveMenuItem := fyne.NewMenuItem("Save", func() {})

	saveAsMenuItem := fyne.NewMenuItem("Save as...", func() {})

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	menu := fyne.NewMainMenu(fileMenu)

	window.SetMainMenu(menu)
}
