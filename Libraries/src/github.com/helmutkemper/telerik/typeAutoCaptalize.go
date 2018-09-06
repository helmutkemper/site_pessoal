package telerik

type AutoCapitalize int

func (el AutoCapitalize) IsSet() bool {
	if el != 0 {
		return true
	}
	return false
}

func (el AutoCapitalize) String() string {
	return autoCapitalizes[el]
}

func (el AutoCapitalize) ToAttr(name string) string {
	var ret string
	if el != 0 {
		ret = ` ` + name + `="` + el.String() + `" `
	}

	return ret
}

var autoCapitalizes = [...]string{
	"",
	"sentences",
	"words",
	"characters",
}

const (
	AUTOCAPITALIZE_SENTENCES AutoCapitalize = iota + 1
	AUTOCAPITALIZE_WORDS
	AUTOCAPITALIZE_CHARACTERS
)
