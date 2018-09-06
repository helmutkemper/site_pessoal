package telerik

type SupportCustomButtonType int

const (
	BUTTON_TYPE_ADD_AND_CLOSE SupportCustomButtonType = iota + 1
	BUTTON_TYPE_ADD
	BUTTON_TYPE_ADD_IN_TEMPLATE
	BUTTON_TYPE_CLOSE
)
