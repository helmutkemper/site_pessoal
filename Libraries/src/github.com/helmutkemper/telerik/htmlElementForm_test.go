package telerik

import "fmt"

func ExampleHtmlElementForm_ToHtml() {
	el := HtmlElementForm{
		Global: HtmlGlobalAttributes{
			Id: "from",
		},
		Name:   "form",
		Method: "get",
		Action: "./index.cpp",
		Content: Content{

			&HtmlElementFormLabel{
				Global: HtmlGlobalAttributes{
					Id: "-",
				},
				Form: "name",
				Content: Content{
					"label_1",
				},
			},
			&HtmlInputText{
				Global: HtmlGlobalAttributes{
					Id: "name",
				},
				Name: NAMES_FOR_AUTOCOMPLETE_NAME,
			},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())

	// Output:
	// <form id="from" name="form" action="./index.cpp" method="get"><label form="name">label_1</label><input type="text" id="name" name="name"></form>
}
