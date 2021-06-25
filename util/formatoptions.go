package util

type FormatOptions struct {
	ShowCompactFormulars         bool
	ShowVariableNamesInFormulars bool
	HideFormulars                bool
	UseTreeFormat                bool
}

func (fo *FormatOptions) SetShowCompactFormulars(newValue bool) {
	fo.ShowCompactFormulars = newValue
}

func (fo *FormatOptions) SetShowVariableNamesInFormulars(newValue bool) {
	fo.ShowVariableNamesInFormulars = newValue
}

func (fo *FormatOptions) SetHideFormulars(newValue bool) {
	fo.HideFormulars = newValue
}

func (fo *FormatOptions) SetUseTreeFormat(newValue bool) {
	fo.UseTreeFormat = newValue
}
