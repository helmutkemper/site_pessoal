package telerik

type JavaScript struct {
	Code string `jsObject:"value"`
}

func (el *JavaScript) String() string {
	return el.Code
}
