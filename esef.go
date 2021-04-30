package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"ESEF/util"
)

const (
	AppTitle                           = "ESEF"
	InputEntryPlaceholder              = "Enter a valid Elasticsearch Explain API response (json)"
	CopiedOutputToClipboardMessageText = "Copied output to clipboard!"
	InputLabelText                     = "Input"
	OutputLabelText                    = "Output"
	CopyButtonLabel                    = "Copy to Clipboard"
	ClearButtonLabel                   = "Clear All"
	PasteFromClipboardButtonLabel      = "Paste from Clipboard"
	ShowCompactFormularsLabel          = "Show compact TF/IDF formulars"
	ShowVariableNamesInFormularsLabel  = "Show variable names in formulars"
	HideDetailedFormularsCheckLabel    = "Hide detailed TF/IDF formulars"
	DarkThemeButtonLabel               = "Dark Theme"
	LightThemeButtonLabel              = "Light Theme"
	SettingsMenuLabel                  = "Settings" // If changed the settings menu will move into it's own menu tab instead of being under "ESEF"
	SettingsLabel                      = SettingsMenuLabel
	FormatCardLabel                    = "Format"
	ShowGuiOutputText                  = "Show GUI Output"
	ShowTextOutputText                 = "Show Text Output"
	SimpleFormatText                   = "Simple Format"
	TreeFormatText                     = "Tree Format"

	SplitContainerOffset   = 0.35
	SystemDefaultThemeName = "system default"
	LightThemeName         = "light"
	DarkThemeName          = "dark"

	IndexNameFormat  = "Index: %s"
	DocumentIdFormat = "Document ID: %s"
	MatchedFormat    = "Matched: %t"
)

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
var FormatRadioGroup *widget.RadioGroup
var OutputStyleRadioGroup *widget.RadioGroup
var TextOutputArea fyne.CanvasObject
var GuiOutputArea fyne.CanvasObject
var TreeOutput fyne.CanvasObject
var TreeOutputContainer *fyne.Container
var TreeIndexnameLabel = widget.NewLabel(IndexNameFormat)
var TreeDocumentIdLabel = widget.NewLabel(DocumentIdFormat)
var TreeMatchedLabel = widget.NewLabel(MatchedFormat)

func main() {
	buildMainWindow().ShowAndRun()
}

func buildMainWindow() fyne.Window {
	application := app.New()
	window = application.NewWindow(AppTitle)

	buildSettingsMenu(application, window)

	window.SetContent(container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		buildInputOutputArea(),
	))

	window.Resize(MainWindowSize)
	window.CenterOnScreen()

	return window
}

func buildSettingsMenu(application fyne.App, mainWindow fyne.Window) {
	settings := util.NewSettings()
	settingsMenuItem := fyne.NewMenuItem(SettingsMenuLabel, func() {
		buildSettingsWindow(application, settings)
	})

	mainMenu := fyne.NewMainMenu(fyne.NewMenu(SettingsLabel, settingsMenuItem))
	mainWindow.SetMainMenu(mainMenu)
	mainWindow.SetMaster()
}

func buildSettingsWindow(application fyne.App, settings *util.Settings) {
	settingsWindow := application.NewWindow(SettingsLabel)
	settingsWindow.SetContent(buildSettingsArea(application, settings))
	settingsWindow.Resize(SettingsWindowSize)
	settingsWindow.CenterOnScreen()
	settingsWindow.Show()
}

func buildSettingsArea(app fyne.App, settings *util.Settings) fyne.CanvasObject {
	if settings.FyneSettings.ThemeName != "" {
		settings.FyneSettings.ThemeName = SystemDefaultThemeName
	}

	darkThemeButton := widget.NewButton(DarkThemeButtonLabel, func() {
		setTheme(DarkThemeName, app, settings)
	})

	lightThemeButton := widget.NewButton(LightThemeButtonLabel, func() {
		setTheme(LightThemeName, app, settings)
	})

	return widget.NewCard(SettingsLabel, "", container.NewVBox(darkThemeButton, lightThemeButton))
}

func setTheme(inputTheme string, app fyne.App, settings *util.Settings) {
	if inputTheme == LightThemeName {
		app.Settings().SetTheme(theme.LightTheme())
	} else {
		app.Settings().SetTheme(theme.DarkTheme())
	}

	settings.FyneSettings.ThemeName = inputTheme
	err := settings.Save()
	if err != nil {
		return
	}
}

func buildFormatArea() *widget.Card {
	return widget.NewCard(FormatCardLabel, "",
		buildFormatOptions(),
	)
}

func buildFormatOptions() fyne.CanvasObject {
	ShowVariableNamesCheck = widget.NewCheck(ShowVariableNamesInFormularsLabel, func(newValue bool) {
		handleShowVariableNamesInFormularClick(newValue)
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

	FormatRadioGroup = widget.NewRadioGroup([]string{TreeFormatText, SimpleFormatText}, handleFormatRadioGroupToggle) // this needs to be first
	FormatRadioGroup.SetSelected(TreeFormatText)
	FormatRadioGroup.Required = true
	FormatRadioGroup.Horizontal = true

	OutputStyleRadioGroup = widget.NewRadioGroup([]string{ShowTextOutputText, ShowGuiOutputText}, handleGuiTextOutputRadioGroupToggle) // then comes this
	OutputStyleRadioGroup.SetSelected(ShowTextOutputText)
	OutputStyleRadioGroup.Required = true
	OutputStyleRadioGroup.Horizontal = true

	return container.NewVBox(
		OutputStyleRadioGroup,
		FormatRadioGroup,
		HideFormularsCheck,
		ShowCompactFormularsCheck,
		ShowVariableNamesCheck,
	)
}

func handleShowVariableNamesInFormularClick(newValue bool) {
	FormatOptions.SetShowVariableNamesInFormular(newValue)
	formatInput()
}

func handleFormatRadioGroupToggle(newValue string) {
	if newValue == SimpleFormatText {
		FormatOptions.SetUseTreeFormat(false)
	} else {
		FormatOptions.SetUseTreeFormat(true)
	}

	formatInput()
}

func handleGuiTextOutputRadioGroupToggle(newValue string) {
	if newValue == ShowTextOutputText {
		TextOutputArea.Show()
		GuiOutputArea.Hide()

		FormatRadioGroup.Show()
		ShowCompactFormularsCheck.Show()
		ShowVariableNamesCheck.Show()
		HideFormularsCheck.Show()
	} else {
		GuiOutputArea.Show()
		TextOutputArea.Hide()

		FormatRadioGroup.Hide()
		ShowCompactFormularsCheck.Hide()
		ShowVariableNamesCheck.Hide()
		HideFormularsCheck.Hide()
	}
}

func buildInputOutputArea() *container.Split {
	TextOutputArea = buildTextOutputArea()
	GuiOutputArea = buildTreeOutputArea()

	splitContainer := container.NewHSplit(
		container.NewAdaptiveGrid(
			1,
			buildInputArea(),
			buildFormatArea(),
		),
		container.NewMax(
			TextOutputArea,
			GuiOutputArea,
		),
	)

	splitContainer.SetOffset(SplitContainerOffset)

	return splitContainer
}

func buildInputArea() *widget.Card {
	InputEntry = widget.NewMultiLineEntry()
	InputEntry.SetPlaceHolder(InputEntryPlaceholder)
	InputEntry.OnChanged = handleOnChangeOfInputEntry
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

func buildTextOutputArea() *widget.Card {
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

func buildTreeOutputArea() *widget.Card {
	TreeOutput = buildBlankTree()
	TreeOutputContainer = container.NewMax(TreeOutput)
	return widget.NewCard(
		OutputLabelText,
		"",
		container.NewBorder(
			container.NewVBox(
				TreeIndexnameLabel,
				TreeDocumentIdLabel,
				TreeMatchedLabel,
			),
			nil,
			nil,
			nil,
			TreeOutputContainer,
		),
	)
}

func buildBlankTree() fyne.CanvasObject {
	return widget.NewTreeWithStrings(map[string][]string{"": {"Nothing formatted yet."}})
}

func handlePasteFromClipboardButtonClick() {
	InputEntry.SetText(window.Clipboard().Content())
}

func handleClearInputButtonClick() {
	InputEntry.SetText("")
	OutputEntry.SetText("")
}

func handleCopyOutputButtonClick() {
	window.Clipboard().SetContent(OutputEntry.Text)
	sendOSNotification(CopiedOutputToClipboardMessageText)
}

func handleOnChangeOfOutputEntry(newText string) {
	changeDisabledStateOfCopyButton(newText)
}

func handleOnChangeOfInputEntry(_ string) {
	formatInput()
}

func handleHideFormularClick(newValue bool) {
	FormatOptions.SetHideFormular(newValue)

	changeDisableStateOfCheckWidget(!newValue, ShowCompactFormularsCheck)

	if ShowCompactFormularsCheck.Checked {
		changeDisableStateOfCheckWidget(!newValue, ShowVariableNamesCheck)
	}

	formatInput()
}

func handleShowCompactFormularClick(newValue bool) {
	FormatOptions.SetShowCompactFormular(newValue)

	changeDisableStateOfCheckWidget(newValue, ShowVariableNamesCheck)

	formatInput()
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

func formatInput() {
	if InputEntry.Text != "" {
		var newTextOutput string

		explainApiDocument, err := util.ExtractDataFromExplainAPI(InputEntry.Text)
		if err != nil {
			newTextOutput = err.Error()
		}

		newTextOutput = util.FormatExplainApiDocument(explainApiDocument, FormatOptions)
		OutputEntry.SetText(newTextOutput)

		updateTreeOutputInGui(explainApiDocument, util.FormatExplainApiDocumentAsGuiTree(explainApiDocument))
	}
}

func updateTreeOutputInGui(explainApiDocument util.ExplainAPIDocument, newTree fyne.CanvasObject) {
	TreeIndexnameLabel.SetText(fmt.Sprintf(IndexNameFormat, explainApiDocument.Indexname))
	TreeDocumentIdLabel.SetText(fmt.Sprintf(DocumentIdFormat, explainApiDocument.DocumentId))
	TreeMatchedLabel.SetText(fmt.Sprintf(MatchedFormat, explainApiDocument.Matched))

	TreeOutputContainer.Remove(TreeOutput)
	TreeOutputContainer.Add(newTree)
	TreeOutputContainer.Refresh()

	TreeOutput = newTree
}

func sendOSNotification(message string) {
	fyne.CurrentApp().SendNotification(fyne.NewNotification(AppTitle, message))
}
