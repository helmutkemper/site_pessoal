package telerik

import "strconv"

type Int int

func (el Int) ToAttr(name string) string {
	var ret string
	if el != 0 {
		ret = ` ` + name + `=` + strconv.Itoa(int(el)) + ` `
	}

	return ret
}
