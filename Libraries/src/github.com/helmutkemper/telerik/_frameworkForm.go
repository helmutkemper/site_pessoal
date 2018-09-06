package telerik

import (
	"bytes"
	"fmt"
	"html/template"
)

type elementProcessed struct {
	Label string
	Input string
}

type FrameworkInput struct {
	Id                string
	Label             string
	Title             string
	PlaceHolder       string
	ValidationMessage string
	Input             interface{}
}

type FrameworkForm struct {
	Template          string
	Class             string
	Form              HtmlElementForm
	Elements          []FrameworkInput
	elementsProcessed []elementProcessed
}

func (el *FrameworkForm) ToScript() string {
	var out bytes.Buffer
	for i := 0; i != len(el.Elements); i += 1 {
		switch elementCast := el.Elements[i].Input.(type) {
		case KendoUiAutoComplete:
			elementCast.HtmlId = el.Elements[i].Id
			out.WriteString(elementCast.String())
		}
	}
	return out.String()
}
func (el *FrameworkForm) ToForm() string {
	var out bytes.Buffer

	el.elementsProcessed = make([]elementProcessed, 0)

	for i := 0; i != len(el.Elements); i += 1 {
		element := elementProcessed{}

		label := HtmlElementFormLabel{
			Content: Content{
				el.Elements[i].Label,
			},
			For: String(el.Elements[i].Id),
			Global: HtmlGlobalAttributes{
				Title: String(el.Elements[i].Title),
			},
		}
		switch elementCast := el.Elements[i].Input.(type) {
		case KendoUiAutoComplete:
			elementNew := HtmlInputSearch{}
			elementNew.Global.Id = String(el.Elements[i].Id)
			elementNew.Global.Class = "k-textbox"
			elementNew.Name = String(el.Elements[i].Id)
			elementNew.PlaceHolder = String(el.Elements[i].PlaceHolder)

			if el.Elements[i].ValidationMessage != "" {
				label.Global.Class += " required"

				elementNew.Required = TRUE
				elementNew.Global.Extra = map[string]string{
					"validationMessage": el.Elements[i].ValidationMessage,
				}
			}

			element.Input = elementNew.String()
		case HtmlInputText:
			elementCast.Global.Id = String(el.Elements[i].Id)
			elementCast.Global.Class = "k-textbox"
			elementCast.Name = String(el.Elements[i].Id)
			elementCast.PlaceHolder = String(el.Elements[i].PlaceHolder)

			if el.Elements[i].ValidationMessage != "" {
				label.Global.Class += " required"

				elementCast.Required = TRUE
				elementCast.Global.Extra = map[string]string{
					"validationMessage": el.Elements[i].ValidationMessage,
				}
			}

			element.Input = elementCast.String()
		}

		element.Label = label.String()

		el.elementsProcessed = append(el.elementsProcessed, element)
	}

	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"safeHTML": func(s interface{}) template.HTML {
			return template.HTML(s.(string))
		},
	}).Parse(el.Template))
	err := tmpl.Execute(&out, el.elementsProcessed)
	if err != nil {
		fmt.Println(err.Error())
	}

	form := el.Form
	form.Content = Content{
		out.String(),
	}

	return form.String()
}
