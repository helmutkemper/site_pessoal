package telerik

type KendoPosition int

var kendoPositions = [...]string{
	"",
	"bottom left",
	"bottom center",
	"bottom right",
	"center left",
	"center center",
	"center right",
	"top left",
	"top center",
	"top right",
}

func (el KendoPosition) String() string {
	return kendoPositions[el]
}

const (
	POSITION_BOTTOM_LEFT KendoPosition = iota + 1
	POSITION_BOTTOM_CENTER
	POSITION_BOTTOM_RIGHT
	POSITION_CENTER_LEFT
	POSITION_CENTER_CENTER
	POSITION_CENTER_RIGHT
	POSITION_TOP_LEFT
	POSITION_TOP_CENTER
	POSITION_TOP_RIGHT
)
