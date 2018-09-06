package telerik

type TypeMenuOrientation int

var TypeMenuOrientations = [...]string{
	"",
	"horizontal",
	"vertical",
}

func (el TypeMenuOrientation) String() string {
	return TypeMenuOrientations[el]
}

const (
	MENU_ORIENTATION_HORIZONTAL TypeMenuOrientation = iota + 1
	MENU_ORIENTATION_VERTICAL
)
