package telerik

type KendoButtonLayout int

var kendoButtonLayouts = [...]string{
	"",
	"normal",
	"stretched",
}

func (el KendoButtonLayout) String() string {
	return kendoButtonLayouts[el]
}

const (
	BUTTON_LAYOUT_NORMAL KendoButtonLayout = iota + 1
	BUTTON_LAYOUT_STRETCHED
)
