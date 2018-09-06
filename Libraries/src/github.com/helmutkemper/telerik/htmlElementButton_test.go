package telerik

import "fmt"

func ExampleHtmlElementFormButton_ToHtml() {
	el := HtmlElementFormButton{
		Global: HtmlGlobalAttributes{
			Id: "button",
		},
		Name: "button",
		Content: Content{
			"click-me",
		},
		Disabled: TRUE,
	}

	fmt.Printf("%s", el.ToHtml())

	// Output:
	// <button id="button" name="button" disabled>click-me</button>
}
