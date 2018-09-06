package telerik

type HtmlMethod int

func (el HtmlMethod) IsSet() bool {
	if el != 0 {
		return true
	}
	return false
}

func (el HtmlMethod) String() string {
	return htmlMethods[el]
}

func (el HtmlMethod) ToAttr(name string) string {
	var ret string
	if el != 0 {
		ret = ` ` + name + `="` + el.String() + `" `
	}

	return ret
}

var htmlMethods = [...]string{
	"",
	"POST",
	"GET",
	"PUT",
	"DELETE",
}

const (
	HTML_METHOD_POST HtmlMethod = iota + 1
	HTML_METHOD_GET
	HTML_METHOD_PUT
	HTML_METHOD_DELETE
)
