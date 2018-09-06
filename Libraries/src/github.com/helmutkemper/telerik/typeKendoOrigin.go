package telerik

type KendoOrigin int

var kendoOrigins = [...]string{
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

func (el KendoOrigin) String() string {
	return kendoOrigins[el]
}

const (
	ORIGIN_BOTTOM_LEFT KendoOrigin = iota + 1
	ORIGIN_BOTTOM_CENTER
	ORIGIN_BOTTOM_RIGHT
	ORIGIN_CENTER_LEFT
	ORIGIN_CENTER_CENTER
	ORIGIN_CENTER_RIGHT
	ORIGIN_TOP_LEFT
	ORIGIN_TOP_CENTER
	ORIGIN_TOP_RIGHT
)
