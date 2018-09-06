package telerik

type relationStt struct {
	Type     string
	Owner    string
	Receiver string
}

type relationArr []relationStt

func (el *relationArr) Add(rType, owner, receiver string) {
	*el = append(*el, relationStt{Type: rType, Owner: owner, Receiver: receiver})
}

var relation relationArr

func init() {
	relation = make(relationArr, 0)
}
