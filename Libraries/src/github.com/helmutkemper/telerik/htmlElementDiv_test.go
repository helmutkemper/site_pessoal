package telerik

import "fmt"

func ExampleHtmlElementDiv_ToHtml() {
	el := HtmlElementDiv{
		Global: HtmlGlobalAttributes{
			Id:    "divId",
			Class: "test",
		},
		Content: Content{
			HtmlElementFormLabel{
				Form: "name",
				Content: Content{
					"label_1",
				},
			},
			HtmlInputText{
				Global: HtmlGlobalAttributes{
					Id: "name",
				},
				Name: NAMES_FOR_AUTOCOMPLETE_NAME,
			},
		},
	}

	fmt.Printf("%s", el.ToHtml())

	// Output:
	// <div class="test" id="divId"><label form="name">label_1</label><input type="text" id="name" name="name"></div>
}
