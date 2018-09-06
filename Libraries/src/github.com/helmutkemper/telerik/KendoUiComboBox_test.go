package telerik

import "fmt"

func ExampleKendoUiComboBox_ToHtml() {
	html := KendoUiComboBox{
		Html: HtmlElementFormSelect{
			Global: HtmlGlobalAttributes{
				Id: "combobox",
			},
		},
		Animation: KendoAnimation{
			Open: KendoOpen{
				Effects:  EFFECT_EXPAND_OUT,
				Duration: 300,
			},
			Close: KendoClose{
				Effects:  EFFECT_EXPAND_IN,
				Duration: 300,
			},
		},
		AutoBind:         FALSE,
		AutoWidth:        TRUE,
		CascadeFrom:      "parent",
		CascadeFromField: "id",
		ClearButton:      FALSE,
		DataSource: []map[string]interface{}{
			{"name": "Child1", "id": 1, "parentId": 1},
			{"name": "Child2", "id": 2, "parentId": 2},
			{"name": "Child3", "id": 3, "parentId": 1},
			{"name": "Child4", "id": 4, "parentId": 2},
		},
		DataTextField:    "name",
		DataValueField:   "id",
		Delay:            500,
		Enable:           TRUE,
		EnforceMinLength: TRUE,
		Filter:           FILTER_CONTAINS,
		Height:           500,
		HighlightFirst:   TRUE,
		IgnoreCase:       TRUE,
		Index:            0,
		MinLength:        1,
		NoDataTemplate:   "No Data!",
		Placeholder:      "Select...",
		Suggest:          TRUE,
		SyncValueAndText: FALSE,
		Text:             "Chai",
		Value:            "Child1",
		ValuePrimitive:   TRUE,
	}

	fmt.Printf("%s", html.ToJavaScript())

	// Output:
	// $("#combobox").kendoComboBox({animation: { close: { effects: "expand:in",duration: 300,},open: { effects: "expand:out",duration: 300,},},autoBind: false,autoWidth: true,cascadeFrom: "parent",cascadeFromField: "id",clearButton: false,dataSource: [{"name": "Child1","id": 1,"parentId": 1,},{"name": "Child2","id": 2,"parentId": 2,},{"name": "Child3","id": 3,"parentId": 1,},{"name": "Child4","id": 4,"parentId": 2,},],dataTextField: "name",dataValueField: "id",delay: 500,enable: true,enforceMinLength: true,filter: "contains",height: 500,highlightFirst: true,ignoreCase: true,minLength: 1,noDataTemplate: "No Data!",placeholder: "Select...",suggest: true,syncValueAndText: false,text: "Chai",value: "Child1",valuePrimitive: true,});
}
