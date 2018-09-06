package telerik

type Warp int

func (el Warp) IsSet() bool {
	if el != 0 {
		return true
	}
	return false
}

func (el Warp) String() string {
	return warps[el]
}

func (el Warp) ToAttr(name string) string {
	var ret string
	if el != 0 {
		ret = ` ` + name + `="` + el.String() + `" `
	}

	return ret
}

var warps = [...]string{
	"",
	"hard",
	"soft",
	"off",
}

const (
	HARD Warp = iota + 1
	SOFT
	OFF
)
