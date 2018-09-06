package telerik

type KendoPdfTemplate int

var KendoPdfTemplates = [...]string{
	"",
	"pageNum",
	"totalPages",
}

func (el KendoPdfTemplate) String() string {
	return KendoPdfTemplates[el]
}

const (
	PDF_TEMPLATE_PAGE_NUM KendoPdfTemplate = iota + 1
	PDF_TEMPLATE_TOTAL_PAGES
)
