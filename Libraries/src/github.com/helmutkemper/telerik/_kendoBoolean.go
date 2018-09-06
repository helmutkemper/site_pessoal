package telerik

type Boolean int

var booleans = [...]int{
	-1,
	0,
	+1,
}

func (el Boolean) IsSet() int {
	return booleans[el]
}

const (
	FALSE Boolean = iota
	NOT_SET
	TRUE
)
