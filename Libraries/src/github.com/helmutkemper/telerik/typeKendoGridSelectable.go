package telerik

type KendoGridSelectable int

func (el KendoGridSelectable) String() string {
	return KendoGridSelectables[el]
}

var KendoGridSelectables = [...]string{
	"",
	"row",
	"cell",
	"multiple, row",
	"multiple, cell",
}

const (
	KENDO_GRID_SELECTABLE_ROW KendoGridSelectable = iota + 1
	KENDO_GRID_SELECTABLE_CELL
	KENDO_GRID_SELECTABLE_MULTIPLE_ROW
	KENDO_GRID_SELECTABLE_MULTIPLE_CELL
)
