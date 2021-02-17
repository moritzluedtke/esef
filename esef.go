package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var WindowSize = fyne.NewSize(900, 600)
var DialogSize = fyne.NewSize(800, 500)

func main() {
	buildMainWindow().ShowAndRun()
}

func buildMainWindow() fyne.Window {
	application := app.New()
	window := application.NewWindow("ESEF - by Moritz Lüdtke")

	inputEntry := widget.NewMultiLineEntry()
	outputEntry := widget.NewMultiLineEntry()

	inputLabel := widget.NewLabel("Input")
	inputLabel.Alignment = fyne.TextAlignCenter

	outputLabel := widget.NewLabel("Output")
	outputLabel.Alignment = fyne.TextAlignCenter

	window.SetContent(container.NewBorder(
		nil,
		widget.NewButtonWithIcon("Format", theme.FileIcon(), func() {
			formatInput(inputEntry.Text)
		}),
		nil,
		nil,
		container.NewAdaptiveGrid(
			2,
			container.NewBorder(
				inputLabel,
				nil,
				nil,
				nil,
				inputEntry,
			),
			container.NewBorder(
				outputLabel,
				widget.NewButton("Copy", func() {

				}),
				nil,
				nil,
				outputEntry,
			),
		),
	),
	)

	application.Settings().SetTheme(theme.LightTheme())

	window.Resize(WindowSize)
	window.CenterOnScreen()

	return window
}

func formatInput(input string) {
	fmt.Println(input)
}
