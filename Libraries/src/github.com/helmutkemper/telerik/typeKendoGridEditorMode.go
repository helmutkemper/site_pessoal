package telerik

type KendoGridEditorMode int

func (el KendoGridEditorMode) IsSet() bool {
	if el != 0 {
		return true
	}
	return false
}

func (el KendoGridEditorMode) String() string {
	return KendoGridEditorModes[el]
}

func (el KendoGridEditorMode) ToAttr(name string) string {
	var ret string
	if el != 0 {
		ret = ` ` + name + `="` + el.String() + `" `
	}

	return ret
}

var KendoGridEditorModes = [...]string{
	"",
	"incell",
	"inline",
	"popup",
}

const (
	KENDO_GRID_EDITOR_MODE_IN_CELL KendoGridEditorMode = iota + 1
	KENDO_GRID_EDITOR_MODE_IN_LINE
	KENDO_GRID_EDITOR_MODE_POPUP
)
