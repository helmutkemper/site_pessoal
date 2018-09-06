package telerik

type TypeHtmlDropZone int

var typeHtmlDropZones = [...]string{
	"",
	"copy",
	"move",
	"link",
}

func (el TypeHtmlDropZone) String() string {
	return typeHtmlDropZones[el]
}

const (
	COPY TypeHtmlDropZone = iota + 1
	MOVE
	LINK
)
