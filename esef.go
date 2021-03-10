package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"ESEF/util"
)

const AppTitle = "ESEF"
const CopiedOutputToClipboardMessageText = "Copied output to clipboard!"
const InputEntryPlaceholder = "Enter a valid explain API response (json)"
const InputLabelText = "Input"
const OutputLabelText = "Output"
const SimpleFormatButtonLabel = "Simple Format"
const TreeFormatButtonLabel = "Tree Format"
const CopyButtonLabel = "Copy to Clipboard"
const ClearButtonLabel = "Clear"
const PasteFromClipboardButtonLabel = "Paste from Clipboard"
const ShowCompactFunctionLabel = "Show compact TF & IDF functions"
const ShowVariableNamesInFunctionLabel = "Show variable names in functions"

const SplitContainerOffset = 0.35

var formatOptions = new(util.FormatOptions)

var WindowSize = fyne.NewSize(1400, 800)
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
		container.NewVBox(
			buildFormatOptions(),
			container.NewAdaptiveGrid(
				2,
				widget.NewButton(SimpleFormatButtonLabel, func() {
					handleSimpleFormatButtonClick()
				}),
				widget.NewButton(TreeFormatButtonLabel, func() {
					handleTreeFormatButtonClick()
				}),
			),
		),
	)
}

func buildFormatOptions() fyne.CanvasObject {
	ShowVariableNamesCheck := widget.NewCheck(ShowVariableNamesInFunctionLabel, func(newValue bool) {
		formatOptions.SetShowVariableNamesInFunction(newValue)
	})
	ShowVariableNamesCheck.Disable()
	ShowVariableNamesCheck.DisableableWidget = widget.DisableableWidget{}

	ShowCompactFunctionCheck := widget.NewCheck(ShowCompactFunctionLabel, func(newValue bool) {
		handleShowCompactFunctionClick(newValue, ShowVariableNamesCheck)
	})

	return container.NewVBox(
		ShowCompactFunctionCheck,
		ShowVariableNamesCheck,
	)
}

func buildInputOutputArea() *container.Split {
	splitContainer := container.NewHSplit(
		buildInputArea(),
		buildOutputArea(),
	)

	splitContainer.SetOffset(SplitContainerOffset)

	return splitContainer
}

func buildInputArea() *widget.Card {
	InputEntry = widget.NewMultiLineEntry()
	InputEntry.SetPlaceHolder(InputEntryPlaceholder)
	InputEntry.TextStyle = fyne.TextStyle{Monospace: true}

	InputLabel = widget.NewLabel(InputLabelText)
	InputLabel.Alignment = fyne.TextAlignCenter

	return widget.NewCard(
		InputLabelText,
		"",
		container.NewBorder(
			nil,
			container.NewAdaptiveGrid(
				2,
				widget.NewButton(PasteFromClipboardButtonLabel, func() {
					handlePasteFromClipboardButtonClick()
				}),
				widget.NewButton(ClearButtonLabel, func() {
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
	OutputEntry.TextStyle = fyne.TextStyle{Monospace: true}

	OutputLabel = widget.NewLabel(OutputLabelText)
	OutputLabel.Alignment = fyne.TextAlignCenter

	CopyButton = widget.NewButton(CopyButtonLabel, func() {
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

func handleSimpleFormatButtonClick() {
	formatOptions.SetUseTreeFormat(false)
	OutputEntry.SetText(formatInput())
}

func handleTreeFormatButtonClick() {
	formatOptions.SetUseTreeFormat(true)
	OutputEntry.SetText(formatInput())
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

func handleShowCompactFunctionClick(newValue bool, showVariableNamesCheck *widget.Check) {
	formatOptions.SetShowCompactFunction(newValue)

	changeDisableStateOfShowVariableNamesCheck(newValue, showVariableNamesCheck)
}

func changeDisableStateOfShowVariableNamesCheck(newValue bool, ShowVariableNamesCheck *widget.Check) {
	if newValue {
		ShowVariableNamesCheck.Enable()
	} else {
		ShowVariableNamesCheck.Disable()
	}
}

func changeDisabledStateOfCopyButton(newText string) {
	if newText == "" {
		CopyButton.Disable()
	} else {
		CopyButton.Enable()
	}
}

func formatInput() string {
	var formattedString string
	explainApiDocument, err := util.ExtractDataFromExplainAPI(InputEntry.Text)

	if err != nil {
		return err.Error()
	}

	if formatOptions.UseTreeFormat {
		formattedString = util.FormatExplainApiDocument(explainApiDocument, formatOptions)
	} else {
		formattedString = util.FormatExplainApiDocument(explainApiDocument, formatOptions)
	}

	return formattedString
}

func sendOSNotification(message string) {
	fyne.CurrentApp().SendNotification(fyne.NewNotification(AppTitle, message))
}
