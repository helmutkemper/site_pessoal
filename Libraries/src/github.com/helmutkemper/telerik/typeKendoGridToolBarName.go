package telerik

type KendoGridToolBarName int

var KendoGridToolBarNames = [...]string{
	"",
	"cancel",
	"create",
	"save",
	"excel",
	"pdf",
	"custom",
}

func (el KendoGridToolBarName) String() string {
	return KendoGridToolBarNames[el]
}

const (
	GRID_TOOLBAR_NAME_CANCEL KendoGridToolBarName = iota + 1
	GRID_TOOLBAR_NAME_CREATE
	GRID_TOOLBAR_NAME_SAVE
	GRID_TOOLBAR_NAME_EXCEL
	GRID_TOOLBAR_NAME_PDF
	GRID_TOOLBAR_NAME_CUSTOM
)
