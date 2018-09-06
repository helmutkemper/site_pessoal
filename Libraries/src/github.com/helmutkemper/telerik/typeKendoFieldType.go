package telerik

type JavaScriptType int

func (el JavaScriptType) IsSet() bool {
	if el != 0 {
		return true
	}
	return false
}

func (el JavaScriptType) String() string {
	return JavaScriptTypes[el]
}

var JavaScriptTypes = [...]string{
	"",
	"number",
	"string",
	"enums",
	"date",
	"boolean",
}

const (
	JAVASCRIPT_NUMBER JavaScriptType = iota + 1
	JAVASCRIPT_STRING
	JAVASCRIPT_ENUMS
	JAVASCRIPT_DATE
	JAVASCRIPT_BOOLEAN
)
