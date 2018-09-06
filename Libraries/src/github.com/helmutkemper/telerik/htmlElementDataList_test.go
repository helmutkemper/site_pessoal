package telerik

import "fmt"

func ExampleHtmlElementFormDataList_ToHtml_1() {
	el := HtmlElementFormDataList{
		Global: HtmlGlobalAttributes{
			Id: "dataListId",
		},
		Name: "dataListName",
		List: "dataListName",
		Options: []HtmlOptions{
			{
				Label: "Label 1",
				Key:   "value_1",
			},
			{
				Label: "Label 2",
				Key:   "value_2",
			},
		},
		Value: "value_2",
	}

	fmt.Printf("%s", el.ToHtml())

	// Output:
	// <input id="dataListId" name="dataListName" list="dataListName" value="value_2"><datalist id="dataListName"><option value="value_1">Label 1</option><option value="value_2">Label 2</option></datalist>
}

func ExampleHtmlElementFormDataList_ToHtml_2() {
	el := HtmlElementFormDataList{
		Global: HtmlGlobalAttributes{
			Id: "dataListId",
		},
		Name: "dataListName",
		List: "dataListName",
		Options: []HtmlOptions{
			{
				Key: "value_1",
			},
			{
				Key: "value_2",
			},
		},
		Value: "value_1",
	}

	fmt.Printf("%s", el.ToHtml())

	// Output:
	// <input id="dataListId" name="dataListName" list="dataListName" value="value_1"><datalist id="dataListName"><option value="value_1"/><option value="value_2"/></datalist>
}
