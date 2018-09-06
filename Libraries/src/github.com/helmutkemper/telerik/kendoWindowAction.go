package telerik

type KendoWindowAction int

func (el KendoWindowAction) IsSet() bool {
	if el != 0 {
		return true
	}
	return false
}

func (el KendoWindowAction) String() string {
	return KendoWindowActions[el]
}

func (el KendoWindowAction) ToAttr(name string) string {
	var ret string
	if el != 0 {
		ret = ` ` + name + `="` + el.String() + `" `
	}

	return ret
}

var KendoWindowActions = [...]string{
	"",
	"Close",
	"Refresh",
	"Minimize",
	"Maximize",
}

const (
	KENDO_WINDOW_ACTIONS_CLOSE KendoWindowAction = iota + 1
	KENDO_WINDOW_ACTIONS_REFRESH
	KENDO_WINDOW_ACTIONS_MINIMIZE
	KENDO_WINDOW_ACTIONS_MAXIMIZE
)
