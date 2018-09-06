package telerik

type KendoTagMode int

var kendoTagModes = [...]string{
	"",
	"multiple",
	"single",
}

func (el KendoTagMode) String() string {
	return kendoTagModes[el]
}

const (
	TAG_MODE_MULTIPLE KendoTagMode = iota + 1
	TAG_MODE_SINGLE
)
