package telerik

type HtmlTarget int

func (el HtmlTarget) String() string {
	return HtmlTargets[el]
}

var HtmlTargets = [...]string{
	"",
	"_blank",
	"_self",
	"_parent",
	"_top",
}

const (
	HTML_TARGET_BLANK HtmlTarget = iota + 1
	HTML_TARGET_SELF
	HTML_TARGET_PARENT
	HTML_TARGET_TOP
)
