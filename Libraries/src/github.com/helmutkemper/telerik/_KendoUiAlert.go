package telerik

import (
	"bytes"
	"fmt"
	"html/template"
)

type KendoUiAlert struct {
	HtmlId String

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/alert#configuration-messages

	  Defines the text of the labels that are shown within the alert dialog. Used primarily for localization.

	  Example
	   <div id="alert"></div>
	   <script>
	   $("#alert").kendoAlert({
	     messages:{
	       okText: "OK"
	     }
	   }).data("kendoAlert").open();
	   </script>
	*/

	Messages interface{}
}

func (el *KendoUiAlert) IsSet() bool {
	return el != nil
}
func (el *KendoUiAlert) String() string {
	var buffer bytes.Buffer
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"safeHTML": func(s interface{}) template.HTML {
			return template.HTML(fmt.Sprint(s))
		},
	}).Parse(GetTemplate()))
	err := tmpl.ExecuteTemplate(&buffer, "KendoUiAlert : KendoUiDialog", *(el))
	if err != nil {
		fmt.Println(err.Error())
	}

	return buffer.String()
}
