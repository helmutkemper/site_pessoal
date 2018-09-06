package telerik

import "fmt"

func ExampleHtmlElementFormSelect_ToHtml() {
	el := HtmlElementFormSelect{
		Name: "selectName",
		Options: []HtmlOptions{
			{
				Key:   "key_1",
				Label: "value_1",
			},
			{
				Key:   "key_2",
				Label: "value_2",
			},
			{
				Key:   "key_3",
				Label: "value_3",
			},
		},
		Value: "key_2",
	}

	fmt.Printf("%s\n", el.ToHtml())

	// Output:
	// <select name="selectName"><option value="key_1">value_1</option><option selected value="key_2">value_2</option><option value="key_3">value_3</option></select>
}
