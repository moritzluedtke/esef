package util

type FormatOptions struct {
	ShowCompactFunction         bool
	ShowVariableNamesInFormular bool
	UseTreeFormat               bool
}

func (fo *FormatOptions) SetShowCompactFunction(newValue bool) {
	fo.ShowCompactFunction = newValue
}

func (fo *FormatOptions) SetShowVariableNamesInFunction(newValue bool) {
	fo.ShowVariableNamesInFormular = newValue
}

func (fo *FormatOptions) SetUseTreeFormat(newValue bool) {
	fo.UseTreeFormat = newValue
}
