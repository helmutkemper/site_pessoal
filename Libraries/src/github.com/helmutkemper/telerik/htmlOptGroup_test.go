package telerik

import "fmt"

func ExampleHtmlOptGroup_ToHtml() {
	el := HtmlOptGroup{
		Global: HtmlGlobalAttributes{
			Id: "-",
		},
		Label: "Group Label",
		Options: []HtmlOptions{
			{
				Label: "Label 1",
				Key:   "label_1",
			},
			{
				Label: "Label 2",
				Key:   "label_2",
			},
		},
	}

	fmt.Printf("%s", el.ToHtml())

	// Output:
	// <optgroup label="Group Label"><option value="label_1">Label 1</option><option value="label_2">Label 2</option></optgroup>
}
