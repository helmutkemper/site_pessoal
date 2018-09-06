package telerik

type TypeMenuCollision int

var TypeMenuCollisions = [...]string{
	"",
	"fit",
	"fit flip",
}

func (el TypeMenuCollision) String() string {
	return TypeMenuCollisions[el]
}

const (
	MENU_COLLISION_FIT TypeMenuCollision = iota + 1
	MENU_COLLISION_FIT_FLIP
)
