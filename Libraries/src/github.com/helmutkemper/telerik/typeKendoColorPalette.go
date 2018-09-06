package telerik

type KendoColorPalette int

func (el KendoColorPalette) IsSet() bool {
	if el != 0 {
		return true
	}
	return false
}

func (el KendoColorPalette) String() string {
	return kendoColorPalettes[el]
}

func (el KendoColorPalette) ToAttr(name string) string {
	var ret string
	if el != 0 {
		ret = ` ` + name + `="` + el.String() + `" `
	}

	return ret
}

var kendoColorPalettes = [...]string{
	"",
	"basic",
	"websafe",
}

const (
	COLOR_PALETTE_BASIC KendoColorPalette = iota + 1
	COLOR_PALETTE_WEBSAFE
)
