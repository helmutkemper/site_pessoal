package telerik

type Boolean int8

func (el Boolean) IsSet() bool {
	if el != 0 {
		return true
	}
	return false
}

func (el Boolean) String() string {
	switch el {
	case -1:
		return "false"
	case 0:
		return ""
	case 1:
		return "true"
	}

	return "false"
}

func (el Boolean) OnOff() string {
	if el == 1 {
		return "on"
	}

	return "off"
}

func (el Boolean) ToAttr(name string) string {
	var ret string
	if el != 0 {
		ret = ` ` + name + `=` + el.String() + ` `
	}

	return ret
}

func (el Boolean) ToAttrOnOff(name string) string {
	var ret string
	if el != 0 {
		ret = ` ` + name + `="` + el.OnOff() + `" `
	}

	return ret
}

func (el Boolean) ToAttrSet(name string) string {
	var ret string
	if el == 1 {
		ret = ` ` + name + ` `
	}

	return ret
}

const (
	FALSE   Boolean = -1
	NOT_SET         = 0
	TRUE            = 1
)
