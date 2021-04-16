package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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
const ShowCompactFormularsLabel = "Show compact TF/IDF formulars"
const ShowVariableNamesInFormularsLabel = "Show variable names in formulars"
const HideDetailedFormularsCheckLabel = "Hide detailed TF/IDF formulars"
const DarkThemeButtonLabel = "Dark"
const LightThemeButtonLabel = "Light"
const SettingsMenuLabel = "Settings" // If changed the settings menu will move into it's own menu tab instead of being under "ESEF"
const SettingsLabel = SettingsMenuLabel

const SplitContainerOffset = 0.35

var FormatOptions = new(util.FormatOptions)

var MainWindowSize = fyne.NewSize(1400, 800)
var SettingsWindowSize = fyne.NewSize(400, 200)
var window fyne.Window

var InputEntry *widget.Entry
var InputLabel *widget.Label
var OutputEntry *widget.Entry
var OutputLabel *widget.Label
var CopyButton *widget.Button
var HideFormularsCheck *widget.Check
var ShowCompactFormularsCheck *widget.Check
var ShowVariableNamesCheck *widget.Check

func main() {
	buildMainWindow().ShowAndRun()
}

func buildMainWindow() fyne.Window {
	application := app.New()
	window = application.NewWindow(AppTitle)

	buildSettingsMenu(application, window)

	window.SetContent(container.NewBorder(
		nil,
		buildFormatArea(),
		nil,
		nil,
		buildInputOutputArea(),
	))

	window.Resize(MainWindowSize)
	window.CenterOnScreen()

	return window
}

func buildSettingsMenu(application fyne.App, mainWindow fyne.Window) {
	settingsMenuItem := fyne.NewMenuItem(SettingsMenuLabel, func() {
		buildSettingsWindow(application)
	})

	mainMenu := fyne.NewMainMenu(fyne.NewMenu(SettingsLabel, settingsMenuItem))
	mainWindow.SetMainMenu(mainMenu)
	mainWindow.SetMaster()
}

func buildSettingsWindow(application fyne.App) {
	settingsWindow := application.NewWindow(SettingsLabel)
	settingsWindow.SetContent(buildSettingsArea(application))
	settingsWindow.Resize(SettingsWindowSize)
	settingsWindow.CenterOnScreen()
	settingsWindow.Show()
}

func buildSettingsArea(app fyne.App) fyne.CanvasObject {
	darkThemeButton := widget.NewButton(DarkThemeButtonLabel, func() {
		app.Settings().SetTheme(theme.DarkTheme())
	})

	lightThemeButton := widget.NewButton(LightThemeButtonLabel, func() {
		app.Settings().SetTheme(theme.LightTheme())
	})

	return widget.NewCard(SettingsLabel, "", container.NewVBox(darkThemeButton, lightThemeButton))
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
	ShowVariableNamesCheck = widget.NewCheck(ShowVariableNamesInFormularsLabel, func(newValue bool) {
		FormatOptions.SetShowVariableNamesInFormular(newValue)
	})
	ShowVariableNamesCheck.DisableableWidget = widget.DisableableWidget{} // !this NEEDs to be first, then .Disable()!
	ShowVariableNamesCheck.Disable()

	ShowCompactFormularsCheck = widget.NewCheck(ShowCompactFormularsLabel, func(newValue bool) {
		handleShowCompactFormularClick(newValue)
	})
	ShowCompactFormularsCheck.DisableableWidget = widget.DisableableWidget{}
	ShowCompactFormularsCheck.Disable()

	HideFormularsCheck = widget.NewCheck(HideDetailedFormularsCheckLabel, func(newValue bool) {
		handleHideFormularClick(newValue)
	})
	HideFormularsCheck.Checked = true

	FormatOptions.SetHideFormular(true)

	return container.NewVBox(
		HideFormularsCheck,
		ShowCompactFormularsCheck,
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
	FormatOptions.SetUseTreeFormat(false)
	OutputEntry.SetText(formatInput())
}

func handleTreeFormatButtonClick() {
	FormatOptions.SetUseTreeFormat(true)
	OutputEntry.SetText(formatInput())
}

func handlePasteFromClipboardButtonClick() {
	InputEntry.SetText(window.Clipboard().Content())
}

func handleClearInputButtonClick() {
	InputEntry.SetText("")
}

func handleCopyOutputButtonClick() {
	window.Clipboard().SetContent(OutputEntry.Text)
	sendOSNotification(CopiedOutputToClipboardMessageText)
}

func handleOnChangeOfOutputEntry(newText string) {
	changeDisabledStateOfCopyButton(newText)
}

func handleHideFormularClick(newValue bool) {
	FormatOptions.SetHideFormular(newValue)

	changeDisableStateOfCheckWidget(!newValue, ShowCompactFormularsCheck)

	if ShowCompactFormularsCheck.Checked {
		changeDisableStateOfCheckWidget(!newValue, ShowVariableNamesCheck)
	}
}

func handleShowCompactFormularClick(newValue bool) {
	FormatOptions.SetShowCompactFormular(newValue)

	changeDisableStateOfCheckWidget(newValue, ShowVariableNamesCheck)
}

func changeDisableStateOfCheckWidget(newValue bool, checkWidgetToChange *widget.Check) {
	if newValue {
		checkWidgetToChange.Enable()
	} else {
		checkWidgetToChange.Disable()
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

	if FormatOptions.UseTreeFormat {
		formattedString = util.FormatExplainApiDocument(explainApiDocument, FormatOptions)
	} else {
		formattedString = util.FormatExplainApiDocument(explainApiDocument, FormatOptions)
	}

	return formattedString
}

func sendOSNotification(message string) {
	fyne.CurrentApp().SendNotification(fyne.NewNotification(AppTitle, message))
}
