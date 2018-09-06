package telerik

type KendoCollision int

var kendoCollisions = [...]string{
	"",
	"fit",
	"flip",
	"flip fit",
	"fit flip",
}

func (el KendoCollision) String() string {
	return kendoCollisions[el]
}

const (
	COLLISION_FIT KendoCollision = iota + 1
	COLLISION_FLIP
	COLLISION_FLIP_FIT
	COLLISION_FIT_FLIP
)
