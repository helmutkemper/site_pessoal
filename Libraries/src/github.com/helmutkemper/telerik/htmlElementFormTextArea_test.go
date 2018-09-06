package telerik

import "fmt"

func ExampleHtmlElementFormTextArea_ToHtml() {
	el := HtmlElementFormTextArea{
		Global: HtmlGlobalAttributes{
			Id:    "textAreaId",
			Class: "className",
		},
		Rows: 20,
		Cols: 40,
		Name: NAMES_FOR_AUTOCOMPLETE_ORGANIZATION,
		Content: Content{
			"company name",
		},
		AutoCapitalize: AUTOCAPITALIZE_SENTENCES,
	}

	fmt.Printf("%s\n", el.ToHtml())

	// Output:
	// <textarea class="className" id="textAreaId" name="organization" autocapitalize="sentences" cols="40" rows="20">company name</textarea>
}
