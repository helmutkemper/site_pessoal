package telerik

type KendoLogic int

var kendoLogics = [...]string{
	"",
	"and",
	"or",
}

func (el KendoLogic) String() string {
	return kendoLogics[el]
}

const (
	LOGIC_AND KendoLogic = iota + 1
	LOGIC_OR
)
