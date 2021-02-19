package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var WindowSize = fyne.NewSize(900, 600)
var Window fyne.Window

const APP_TITLE = "ESEF v0.1-alpha"

func main() {
	buildMainWindow().ShowAndRun()
}

func buildMainWindow() fyne.Window {
	application := app.New()
	Window = application.NewWindow(APP_TITLE)

	// Refactor this stuff into their methods and start global references for use in other funcs
	inputEntry := widget.NewMultiLineEntry()
	inputEntry.Text = GetExplainApiOutputExample()

	outputEntry := widget.NewMultiLineEntry()

	inputLabel := widget.NewLabel("Input")
	inputLabel.Alignment = fyne.TextAlignCenter

	outputLabel := widget.NewLabel("Output")
	outputLabel.Alignment = fyne.TextAlignCenter

	Window.SetContent(container.NewBorder(
		nil,
		container.NewAdaptiveGrid(
			2,
			widget.NewButton("Simple Format", func() {
				handleSimpleFormatButtonClick(inputEntry, outputEntry)
			}),
			widget.NewButton("Tree Format", func() {
				handleTreeFormatButtonClick(inputEntry, outputEntry)
			}),
		),
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

func handleSimpleFormatButtonClick(inputEntry *widget.Entry, outputEntry *widget.Entry) {
	outputEntry.SetText(formatInput(inputEntry.Text, false))
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

func handleTreeFormatButtonClick(inputEntry *widget.Entry, outputEntry *widget.Entry) {
	outputEntry.SetText(formatInput(inputEntry.Text, true))
}

func handleCopyOutputButtonClick(output string) {
	Window.Clipboard().SetContent(output)
	sendOSNotification("Copied output to clipboard!")
}

func formatInput(input string, useTreeFormat bool) string {
	explainApiDocument := ExtractDataFromExplainAPI(input)
	if useTreeFormat {
		document := FormatExplainApiDocument(explainApiDocument, useTreeFormat)
		fmt.Println(document)
		return document
	} else {
		document := FormatExplainApiDocument(explainApiDocument, useTreeFormat)
		fmt.Println(document)
		return document
	}

}

func sendOSNotification(message string) {
	fyne.CurrentApp().SendNotification(fyne.NewNotification(APP_TITLE, message))
}
