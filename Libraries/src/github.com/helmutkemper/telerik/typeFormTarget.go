package telerik

type FormTarget int

func (el FormTarget) IsSet() bool {
	if el != 0 {
		return true
	}
	return false
}

func (el FormTarget) String() string {
	return formTargets[el]
}

func (el FormTarget) ToAttr(name string) string {
	var ret string
	if el != 0 {
		ret = ` ` + name + `="` + el.String() + `" `
	}

	return ret
}

var formTargets = [...]string{
	"",
	"_self",
	"_blank",
	"_parent",
	"_top",
}

const (
	FORM_TARGET_SELF FormTarget = iota + 1
	FORM_TARGET_BLANK
	FORM_TARGET_PARENT
	FORM_TARGET_TOP
)
