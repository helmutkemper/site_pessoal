package telerik

type KendoOrientation int

var kendoOrientations = [...]string{
	"",
	"horizontal",
	"vertical",
}

func (el KendoOrientation) String() string {
	return kendoOrientations[el]
}

const (
	ORIENTATION_HORIZONTAL KendoOrientation = iota + 1
	ORIENTATION_VERTICAL
)
