package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"moritzluedtke/ESEF/util"
)

const AppTitle = "ESEF v0.3"
const CopiedOutputToClipboardMessageText = "Copied output to clipboard!"
const InputEntryPlaceholder = "Enter a valid explain API response (json)"
const InputLabelText = "Input"
const OutputLabelText = "Output"
const SimpleFormatButtonText = "Simple Format"
const TreeFormatButtonText = "Tree Format"
const CopyButtonText = "Copy to Clipboard"
const ClearButtonText = "Clear"
const PasteFromClipboardButtonText = "Paste from Clipboard"

var WindowSize = fyne.NewSize(900, 600)
var Window fyne.Window

var InputEntry *widget.Entry
var InputLabel *widget.Label
var OutputEntry *widget.Entry
var OutputLabel *widget.Label
var CopyButton *widget.Button

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

func buildInputArea() *widget.Card {
	InputEntry = widget.NewMultiLineEntry()
	InputEntry.SetPlaceHolder(InputEntryPlaceholder)

	InputLabel = widget.NewLabel(InputLabelText)
	InputLabel.Alignment = fyne.TextAlignCenter

	return widget.NewCard(
		InputLabelText,
		"",
		container.NewBorder(
			nil,
			container.NewAdaptiveGrid(
				2,
				widget.NewButton(PasteFromClipboardButtonText, func() {
					handlePasteFromClipboardButtonClick()
				}),
				widget.NewButton(ClearButtonText, func() {
					handleClearInputButtonClick()
				}),
			),
			nil,
			nil,
			InputEntry,
		),
	)
}

func buildOutputArea() *widget.Card {
	OutputEntry = widget.NewMultiLineEntry()
	OutputEntry.OnChanged = handleOnChangeOfOutputEntry

	OutputLabel = widget.NewLabel(OutputLabelText)
	OutputLabel.Alignment = fyne.TextAlignCenter

	CopyButton = widget.NewButton(CopyButtonText, func() {
		handleCopyOutputButtonClick()
	})
	CopyButton.Disable()

	return widget.NewCard(
		OutputLabelText,
		"",
		container.NewBorder(
			nil,
			CopyButton,
			nil,
			nil,
			OutputEntry,
		),
	)
}

func handlePasteFromClipboardButtonClick() {
	InputEntry.SetText(Window.Clipboard().Content())
}

func handleClearInputButtonClick() {
	InputEntry.SetText("")
}

func handleCopyOutputButtonClick() {
	Window.Clipboard().SetContent(OutputEntry.Text)
	sendOSNotification(CopiedOutputToClipboardMessageText)
}

func handleOnChangeOfOutputEntry(newText string) {
	changeDisabledStateOfCopyButton(newText)
}

func changeDisabledStateOfCopyButton(newText string) {
	if newText == "" {
		CopyButton.Disable()
	} else {
		CopyButton.Enable()
	}
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
