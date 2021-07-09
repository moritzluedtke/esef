package util

type FormatOptions struct {
	_ShowCompactFormulars         bool
	_ShowVariableNamesInFormulars bool
	_HideFormulars                bool
	_UseTreeFormat                bool
}

func (fo *FormatOptions) SetShowCompactFormulars(newValue bool) {
	fo._ShowCompactFormulars = newValue
}

func (fo *FormatOptions) SetShowVariableNamesInFormulars(newValue bool) {
	fo._ShowVariableNamesInFormulars = newValue
}

func (fo *FormatOptions) SetHideFormulars(newValue bool) {
	fo._HideFormulars = newValue
}

func (fo *FormatOptions) SetUseTreeFormat(newValue bool) {
	fo._UseTreeFormat = newValue
}
