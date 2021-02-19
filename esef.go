package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var WindowSize = fyne.NewSize(900, 600)
var Window fyne.Window

func main() {
	buildMainWindow().ShowAndRun()
}

func buildMainWindow() fyne.Window {
	application := app.New()
	Window = application.NewWindow("ESEF v0.1 alpha - by Moritz Lüdtke")

	inputEntry := widget.NewMultiLineEntry()
	inputEntry.Text = GetExplainApiOutputExample()

	outputEntry := widget.NewMultiLineEntry()

	inputLabel := widget.NewLabel("Input")
	inputLabel.Alignment = fyne.TextAlignCenter

	outputLabel := widget.NewLabel("Output")
	outputLabel.Alignment = fyne.TextAlignCenter

	Window.SetContent(container.NewBorder(
		nil,
		widget.NewButton("Format", func() {
			handleFormatButtonClick(inputEntry, outputEntry)
		}),
		nil,
		nil,
		container.NewAdaptiveGrid(
			2,
			buildInputArea(inputLabel, inputEntry),
			buildOutputArea(outputLabel, outputEntry),
		),
	),
	)

	Window.Resize(WindowSize)
	Window.CenterOnScreen()

	return Window
}

func buildOutputArea(outputLabel *widget.Label, outputEntry *widget.Entry) *fyne.Container {
	return container.NewBorder(
		outputLabel,
		widget.NewButton("Copy", func() {
			handleCopyOutputButtonClick(outputEntry.Text)
		}),
		nil,
		nil,
		outputEntry,
	)
}

func buildInputArea(inputLabel *widget.Label, inputEntry *widget.Entry) *fyne.Container {
	return container.NewBorder(
		inputLabel,
		nil,
		nil,
		nil,
		inputEntry,
	)
}

func handleFormatButtonClick(inputEntry *widget.Entry, outputEntry *widget.Entry) {
	outputEntry.SetText(formatInputString(inputEntry.Text))
}

func handleCopyOutputButtonClick(output string) {
	Window.Clipboard().SetContent(output)

	fyne.CurrentApp().SendNotification(fyne.NewNotification("ESEF", "Copied output to clipboard!"))
}

func formatInputString(input string) string {
	return ExtractDataFromExplainAPI(input)
}
