package telerik

import "fmt"

func ExampleKendoUiAutoComplete_ToJavaScript_1() {
	html := KendoUiAutoComplete{
		Html: HtmlInputText{
			Global: HtmlGlobalAttributes{
				Id: "auto_complete",
			},
		},
		Animation: KendoAnimation{
			Open: KendoOpen{
				Duration: 300,
				Effects:  EFFECT_EXPAND_IN,
			},
			Close: KendoClose{
				Duration: 300,
				Effects:  EFFECT_EXPAND_OUT,
			},
		},
		AutoWidth:     FALSE,
		ClearButton:   TRUE,
		DataTextField: "name",
		DataSource: []map[string]interface{}{
			{
				"id":   1,
				"name": "Apples",
			},
			{
				"id":   2,
				"name": "Oranges",
			},
		},
		Delay:            500,
		Enable:           TRUE,
		EnforceMinLength: TRUE,
		Filter:           FILTER_CONTAINS,
		Height:           500,
		HighlightFirst:   TRUE,
		IgnoreCase:       TRUE,
		MinLength:        3,
		NoDataTemplate:   "No Data!",
		Placeholder:      "Enter value ...",
		Popup: KendoPopup{
			AppendTo: JavaScript{
				Code: "$('#container')",
			},
		},
		Separator:      ", ",
		Suggest:        TRUE,
		HeaderTemplate: "<div><h2>Fruits</h2></div>",
		Template: JavaScript{
			Code: "kendo.template($('#template').html())",
		},
		Value:          "Oranges",
		ValuePrimitive: FALSE,
	}

	fmt.Printf("%s", html.ToJavaScript())

	// Output:
	// $("#auto_complete").kendoAutoComplete({animation: { close: { effects: "expand:out",duration: 300,},open: { effects: "expand:in",duration: 300,},},autoWidth: false,dataSource: [{"id": 1,"name": "Apples",},{"id": 2,"name": "Oranges",},],clearButton: true,dataTextField: "name",delay: 500,enable: true,enforceMinLength: true,filter: "contains",height: 500,highlightFirst: true,ignoreCase: true,minLength: 3,noDataTemplate: "No Data!",placeholder: "Enter value ...",popup: { appendTo: $('#container'),},separator: ", ",suggest: true,headerTemplate: "<div><h2>Fruits</h2></div>",template: kendo.template($('#template').html()),value: "Oranges",valuePrimitive: false,});
}

func ExampleKendoUiAutoComplete_ToJavaScript_2() {
	html := KendoUiAutoComplete{
		Html: HtmlInputText{
			Global: HtmlGlobalAttributes{
				Id: "auto_complete",
			},
		},
		DataSource: []string{
			"Albania",
			"Andorra",
			"Armenia",
			"Austria",
			"Azerbaijan",
			"Belarus",
			"Belgium",
			"Bosnia & Herzegovina",
			"Bulgaria",
			"Croatia",
			"Cyprus",
			"Czech Republic",
			"Denmark",
			"Estonia",
			"Finland",
			"France",
			"Georgia",
			"Germany",
			"Greece",
			"Hungary",
			"Iceland",
			"Ireland",
			"Italy",
			"Kosovo",
			"Latvia",
			"Liechtenstein",
			"Lithuania",
			"Luxembourg",
			"Macedonia",
			"Malta",
			"Moldova",
			"Monaco",
			"Montenegro",
			"Netherlands",
			"Norway",
			"Poland",
			"Portugal",
			"Romania",
			"Russia",
			"San Marino",
			"Serbia",
			"Slovakia",
			"Slovenia",
			"Spain",
			"Sweden",
			"Switzerland",
			"Turkey",
			"Ukraine",
			"United Kingdom",
			"Vatican City",
		},
	}

	fmt.Printf("%s", html.ToJavaScript())

	// Output:
	// $("#auto_complete").kendoAutoComplete({dataSource: ["Albania","Andorra","Armenia","Austria","Azerbaijan","Belarus","Belgium","Bosnia & Herzegovina","Bulgaria","Croatia","Cyprus","Czech Republic","Denmark","Estonia","Finland","France","Georgia","Germany","Greece","Hungary","Iceland","Ireland","Italy","Kosovo","Latvia","Liechtenstein","Lithuania","Luxembourg","Macedonia","Malta","Moldova","Monaco","Montenegro","Netherlands","Norway","Poland","Portugal","Romania","Russia","San Marino","Serbia","Slovakia","Slovenia","Spain","Sweden","Switzerland","Turkey","Ukraine","United Kingdom","Vatican City",],});
}

func ExampleKendoUiAutoComplete_ToJavaScript_3() {
	html := KendoUiAutoComplete{
		Html: HtmlInputText{
			Global: HtmlGlobalAttributes{
				Id: "auto_complete",
			},
		},
		DataSource: &KendoDataSource{
			Aggregate: []KendoAggregates{
				{
					Aggregate: AGGREGATE_SUM,
					Field:     "age",
				},
				{
					Aggregate: AGGREGATE_MIN,
					Field:     "age",
				},
				{
					Aggregate: AGGREGATE_MAX,
					Field:     "age",
				},
			},
			AutoSync: TRUE,
			Batch:    TRUE,
			Data: []map[string]interface{}{
				{
					"name": "Jane Doe",
					"age":  30,
				},
				{
					"name": "Jane Doe",
					"age":  33,
				},
			},
			InPlaceSort:    FALSE,
			OfflineStorage: "products-offline",
		},
	}

	fmt.Printf("%s", html.ToJavaScript())

	// Output:
	// $("#auto_complete").kendoAutoComplete({dataSource:  new kendo.data.DataSource({aggregate: [{aggregate: "sum",field: "age",},{aggregate: "min",field: "age",},{aggregate: "max",field: "age",},],autoSync: true,batch: true,data: [{"name": "Jane Doe","age": 30,},{"age": 33,"name": "Jane Doe",},],inPlaceSort: false,offlineStorage: "products-offline",});});
}
