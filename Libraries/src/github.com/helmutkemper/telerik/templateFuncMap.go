package telerik

import (
	"html/template"
)

var templateFuncMap template.FuncMap

func init() {
	templateFuncMap = template.FuncMap{
		"safeHTML": func(s interface{}) template.HTML {
			return template.HTML(s.(string))
		},
	}
}
