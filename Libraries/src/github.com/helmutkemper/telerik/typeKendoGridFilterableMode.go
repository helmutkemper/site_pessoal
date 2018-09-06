package telerik

type KendoGridFilterableMode int

var KendoGridFilterableModes = [...]string{
	"",
	"row",
	"menu",
	"menu, row",
}

func (el KendoGridFilterableMode) String() string {
	return KendoGridFilterableModes[el]
}

const (
	KENDO_GRID_FILTERABLE_MODE_ROW KendoGridFilterableMode = iota + 1
	KENDO_GRID_FILTERABLE_MODE_MENU
	KENDO_GRID_FILTERABLE_MODE_MENU_ROW
)
