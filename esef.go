package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const APP_TITLE = "ESEF v0.1-alpha"
const CopiedOutputToClipboardMessageText = "Copied output to clipboard!"
const InputLabelText = "Input"
const OutputLabelText = "Output"
const SimpleFormatButtonText = "Simple Format"
const TreeFormatButtonText = "Tree Format"
const CopyButtonText = "Copy"

var WindowSize = fyne.NewSize(900, 600)
var Window fyne.Window

var InputEntry *widget.Entry
var InputLabel *widget.Label
var OutputEntry *widget.Entry
var OutputLabel *widget.Label

func main() {
	buildMainWindow().ShowAndRun()
}

func buildMainWindow() fyne.Window {
	application := app.New()
	Window = application.NewWindow(APP_TITLE)

	Window.SetContent(container.NewBorder(
		nil,
		buildFormatArea(),
		nil,
		nil,
		buildInputOutputArea(),
	))

	Window.Resize(WindowSize)
	Window.CenterOnScreen()

	return Window
}

func buildFormatArea() *fyne.Container {
	return container.NewAdaptiveGrid(
		2,
		widget.NewButton(SimpleFormatButtonText, func() {
			handleSimpleFormatButtonClick()
		}),
		widget.NewButton(TreeFormatButtonText, func() {
			handleTreeFormatButtonClick()
		}),
	)
}

func buildInputOutputArea() *fyne.Container {
	return container.NewAdaptiveGrid(
		2,
		buildInputArea(),
		buildOutputArea(),
	)
}

func handleSimpleFormatButtonClick() {
	OutputEntry.SetText(formatInput(false))
}

func handleTreeFormatButtonClick() {
	OutputEntry.SetText(formatInput(true))
}

func buildOutputArea() *fyne.Container {
	OutputEntry = widget.NewMultiLineEntry()

	OutputLabel = widget.NewLabel(OutputLabelText)
	OutputLabel.Alignment = fyne.TextAlignCenter

	return container.NewBorder(
		OutputLabel,
		widget.NewButton(CopyButtonText, func() {
			handleCopyOutputButtonClick()
		}),
		nil,
		nil,
		OutputEntry,
	)
}

func buildInputArea() *fyne.Container {
	InputEntry = widget.NewMultiLineEntry()
	InputEntry.Text = GetExplainApiOutputExample()

	InputLabel = widget.NewLabel(InputLabelText)
	InputLabel.Alignment = fyne.TextAlignCenter

	return container.NewBorder(
		InputLabel,
		nil,
		nil,
		nil,
		InputEntry,
	)
}

func handleCopyOutputButtonClick() {
	Window.Clipboard().SetContent(OutputEntry.Text)
	sendOSNotification(CopiedOutputToClipboardMessageText)
}

func formatInput(useTreeFormat bool) string {
	explainApiDocument := ExtractDataFromExplainAPI(InputEntry.Text)
	var formattedString string

	if useTreeFormat {
		formattedString = FormatExplainApiDocument(explainApiDocument, useTreeFormat)
	} else {
		formattedString = FormatExplainApiDocument(explainApiDocument, useTreeFormat)
	}

	fmt.Println(formattedString)
	return formattedString
}

func sendOSNotification(message string) {
	fyne.CurrentApp().SendNotification(fyne.NewNotification(APP_TITLE, message))
}
