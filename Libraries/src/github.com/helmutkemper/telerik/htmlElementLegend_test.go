package telerik

import "fmt"

func ExampleHtmlElementFormLegend_ToHtml() {
	el := HtmlElementFormLegend{
		Name: "legend",
		Global: HtmlGlobalAttributes{
			Id: "legend",
		},
	}

	fmt.Printf("%s", el.ToHtml())

	// Output:
	// <legend id="legend" name="legend"></legend>
}
