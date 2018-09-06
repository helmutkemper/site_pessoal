package telerik

type TypeMenuDirection int

var TypeMenuDirections = [...]string{
	"",
	"top",
	"bottom",
	"left",
	"right",
	"top left",
	"top right",
	"bottom left",
	"bottom right",
}

func (el TypeMenuDirection) String() string {
	return TypeMenuDirections[el]
}

const (
	MENU_DIRECTION_TOP TypeMenuDirection = iota + 1
	MENU_DIRECTION_BOTTOM
	MENU_DIRECTION_LEFT
	MENU_DIRECTION_RIGHT
	MENU_DIRECTION_TOP_LEFT
	MENU_DIRECTION_TOP_RIGHT
	MENU_DIRECTION_BOTTOM_LEFT
	MENU_DIRECTION_BOTTOM_RIGHT
)
