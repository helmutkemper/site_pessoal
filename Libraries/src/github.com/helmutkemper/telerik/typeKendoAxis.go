package telerik

type KendoAxis int

var kendoAxes = [...]string{
	"",
	"x",
	"y",
}

func (el KendoAxis) String() string {
	return kendoAxes[el]
}

const (
	AXIS_X KendoAxis = iota + 1
	AXIS_Y
)
