package telerik

import "fmt"

func ExampleHtmlElementFormLabel_String() {
	el := HtmlElementFormLabel{
		Form: "name",
		Content: Content{
			"Name",
		},
	}

	fmt.Printf("%s\n", el.ToHtml())

	// Output:
	// <label form="name">Name</label>
}
