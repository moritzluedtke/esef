package util

type FormatOptions struct {
	ShowCompactFormular         bool
	ShowVariableNamesInFormular bool
	HideFormular                bool
	UseTreeFormat               bool
}

func (fo *FormatOptions) SetShowCompactFormular(newValue bool) {
	fo.ShowCompactFormular = newValue
}

func (fo *FormatOptions) SetShowVariableNamesInFormular(newValue bool) {
	fo.ShowVariableNamesInFormular = newValue
}

func (fo *FormatOptions) SetHideFormular(newValue bool) {
	fo.HideFormular = newValue
}

func (fo *FormatOptions) SetUseTreeFormat(newValue bool) {
	fo.UseTreeFormat = newValue
}
