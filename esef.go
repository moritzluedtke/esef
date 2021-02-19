package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"./util"
)

const AppTitle = "ESEF v0.2"
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
	Window = application.NewWindow(AppTitle)

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

func buildFormatArea() *widget.Card {
	return widget.NewCard("Format", "",
		container.NewAdaptiveGrid(
			2,
			widget.NewButton(SimpleFormatButtonText, func() {
				handleSimpleFormatButtonClick()
			}),
			widget.NewButton(TreeFormatButtonText, func() {
				handleTreeFormatButtonClick()
			}),
		),
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

func buildOutputArea() *widget.Card {
	OutputEntry = widget.NewMultiLineEntry()

	OutputLabel = widget.NewLabel(OutputLabelText)
	OutputLabel.Alignment = fyne.TextAlignCenter

	return widget.NewCard(
		OutputLabelText,
		"",
		container.NewBorder(
			nil,
			widget.NewButton(CopyButtonText, func() {
				handleCopyOutputButtonClick()
			}),
			nil,
			nil,
			OutputEntry,
		),
	)
}

func buildInputArea() *widget.Card {
	InputEntry = widget.NewMultiLineEntry()

	InputLabel = widget.NewLabel(InputLabelText)
	InputLabel.Alignment = fyne.TextAlignCenter

	return widget.NewCard(
		InputLabelText,
		"",
		InputEntry,
	)
}

func handleCopyOutputButtonClick() {
	Window.Clipboard().SetContent(OutputEntry.Text)
	sendOSNotification(CopiedOutputToClipboardMessageText)
}

func formatInput(useTreeFormat bool) string {
	var formattedString string
	explainApiDocument, err := util.ExtractDataFromExplainAPI(InputEntry.Text)

	if err != nil {
		return err.Error()
	}

	if useTreeFormat {
		formattedString = util.FormatExplainApiDocument(explainApiDocument, useTreeFormat)
	} else {
		formattedString = util.FormatExplainApiDocument(explainApiDocument, useTreeFormat)
	}

	return formattedString
}

func sendOSNotification(message string) {
	fyne.CurrentApp().SendNotification(fyne.NewNotification(AppTitle, message))
}
