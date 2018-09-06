package telerik

import "fmt"

func ExampleKendoUiGrid_ToHtml_1() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Selectable: KENDO_GRID_SELECTABLE_MULTIPLE_CELL,
		AllowCopy:  TRUE,
		Columns: []KendoGridColumns{
			{Field: "productName"},
			{Field: "category"},
		},
		DataSource: []map[string]interface{}{
			{"productName": "Tea", "category": "Beverages"},
			{"productName": "Coffee", "category": "Beverages"},
			{"productName": "Ham", "category": "Food"},
			{"productName": "Bread", "category": "Food"},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({allowCopy: true,columns: [{field: "productName",},{field: "category",},],dataSource: [{"category": "Beverages","productName": "Tea",},{"productName": "Coffee","category": "Beverages",},{"productName": "Ham","category": "Food",},{"productName": "Bread","category": "Food",},],selectable: "multiple, cell",});
}

func ExampleKendoUiGrid_ToHtml_2() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "name"},
			{Field: "age"},
		},
		ColumnMenu: TRUE,
		DataSource: []map[string]interface{}{
			{"name": "Jane Doe", "age": 30},
			{"name": "John Doe", "age": 33},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",},{field: "age",},],dataSource: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],});
}

func ExampleKendoUiGrid_ToHtml_3() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "name"},
			{Field: "age"},
		},
		ColumnMenu: KendoGridColumnMenu{
			Columns: FALSE,
		},
		Sortable: TRUE,
		DataSource: []map[string]interface{}{
			{"name": "Jane Doe", "age": 30},
			{"name": "John Doe", "age": 33},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: {columns: false,},columns: [{field: "name",},{field: "age",},],dataSource: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],sortable: true,});
}

func ExampleKendoUiGrid_ToHtml_4() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "name"},
			{Field: "age"},
		},
		ColumnMenu: KendoGridColumnMenu{
			Messages: KendoGridColumnMenuMessages{
				Columns:        "Choose columns",
				Filter:         "Apply filter",
				SortAscending:  "Sort (asc)",
				SortDescending: "Sort (desc)",
			},
		},
		Sortable:   TRUE,
		Filterable: TRUE,
		DataSource: []map[string]interface{}{
			{"name": "Jane Doe", "age": 30},
			{"name": "John Doe", "age": 33},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: {messages: {columns: "Choose columns",filter: "Apply filter",sortAscending: "Sort (asc)",sortDescending: "Sort (desc)",},},columns: [{field: "name",},{field: "age",},],dataSource: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],filterable: true,sortable: true,});
}

func ExampleKendoUiGrid_ToHtml_5() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "name", Title: "Name"},
			{Field: "age", Title: "Age"},
		},
		DataSource: []map[string]interface{}{
			{"name": "Jane Doe", "age": 30},
			{"name": "John Doe", "age": 33},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},],dataSource: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],});
}

func ExampleKendoUiGrid_ToHtml_6() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "name", Title: "Name"},
			{Field: "age", Title: "Age"},
			{Command: []TypeKendoGridColumnsCommand{COLUMNS_COMMAND_EDIT}},
		},

		DataSource: KendoDataSource{
			Data: []map[string]interface{}{
				{"name": "Jane Doe", "age": 30},
				{"name": "John Doe", "age": 33},
			},
			Schema: KendoSchema{
				Model: KendoDataModel{
					Id: "id",
					Fields: map[string]KendoField{
						"age": {
							Type: JAVASCRIPT_NUMBER,
						},
					},
				},
			},
		},
		Editable: KENDO_GRID_EDITOR_MODE_POPUP,
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: ["edit",],},],dataSource:  new kendo.data.DataSource({data: [{"age": 30,"name": "Jane Doe",},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",});
}

func ExampleKendoUiGrid_ToHtml_7() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "name", Title: "Name"},
			{Field: "age", Title: "Age"},
			{Command: []TypeKendoGridColumnsCommand{COLUMNS_COMMAND_EDIT}},
		},
		ColumnMenu: TRUE,
		Toolbar:    []KendoGridToolBarName{GRID_TOOLBAR_NAME_SAVE, GRID_TOOLBAR_NAME_EXCEL},
		Excel: KendoGridExcel{
			AllPages: TRUE,
		},
		DataSource: KendoDataSource{
			Data: []map[string]interface{}{
				{"name": "Jane Doe", "age": 30},
				{"name": "John Doe", "age": 33},
			},
			Schema: KendoSchema{
				Model: KendoDataModel{
					Id: "id",
					Fields: map[string]KendoField{
						"age": {
							Type: JAVASCRIPT_NUMBER,
						},
					},
				},
			},
		},
		Editable: KENDO_GRID_EDITOR_MODE_POPUP,
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: ["edit",],},],dataSource:  new kendo.data.DataSource({data: [{"age": 30,"name": "Jane Doe",},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},toolbar: ["save","excel",],});
}

func ExampleKendoUiGrid_ToHtml_8() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "name", Title: "Name"},
			{Field: "age", Title: "Age"},
			{Command: []TypeKendoGridColumnsCommand{COLUMNS_COMMAND_EDIT}},
		},
		ColumnMenu: TRUE,
		Toolbar: []KendoGridToolbar{
			{
				Name:      GRID_TOOLBAR_NAME_SAVE,
				IconClass: "k-icon k-i-copy",
			},
			{
				Name: GRID_TOOLBAR_NAME_EXCEL,
			},
		},
		Excel: KendoGridExcel{
			AllPages: TRUE,
		},
		DataSource: KendoDataSource{
			Data: []map[string]interface{}{
				{"name": "Jane Doe", "age": 30},
				{"name": "John Doe", "age": 33},
			},
			Schema: KendoSchema{
				Model: KendoDataModel{
					Id: "id",
					Fields: map[string]KendoField{
						"age": {
							Type: JAVASCRIPT_NUMBER,
						},
					},
				},
			},
		},
		Editable: KENDO_GRID_EDITOR_MODE_POPUP,
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: ["edit",],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"age": 33,"name": "John Doe",},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_9() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "name", Title: "Name"},
			{Field: "age", Title: "Age"},
			{Command: []TypeKendoGridColumnsCommand{COLUMNS_COMMAND_EDIT}},
		},
		ColumnMenu: TRUE,
		Toolbar: []KendoGridToolbar{
			{
				Name:      GRID_TOOLBAR_NAME_SAVE,
				IconClass: "k-icon k-i-copy",
			},
			{
				Name: GRID_TOOLBAR_NAME_EXCEL,
			},
		},
		Excel: KendoGridExcel{
			AllPages: TRUE,
		},
		DataSource: KendoDataSource{
			Data: []map[string]interface{}{
				{"name": "Jane Doe", "age": 30},
				{"name": "John Doe", "age": 33},
			},
			Schema: KendoSchema{
				Model: KendoDataModel{
					Id: "id",
					Fields: map[string]KendoField{
						"age": {
							Type: JAVASCRIPT_NUMBER,
						},
					},
				},
			},
		},
		Editable: KENDO_GRID_EDITOR_MODE_POPUP,
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: ["edit",],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_10() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "name", Title: "Name"},
			{Field: "age", Title: "Age"},
			{Command: []KendoGridColumnsCommand{
				{
					Name:      COLUMNS_COMMAND_CUSTOM,
					Text:      "details",
					Click:     JavaScript{Code: `function(e){ /* command button click handler */ }`},
					ClassName: "btn-destroy",
					IconClass: "k-icon k-i-copy",
				},
				{
					Name: COLUMNS_COMMAND_DESTROY,
				},
			}},
		},
		ColumnMenu: TRUE,
		Toolbar: []KendoGridToolbar{
			{
				Name:      GRID_TOOLBAR_NAME_SAVE,
				IconClass: "k-icon k-i-copy",
			},
			{
				Name: GRID_TOOLBAR_NAME_EXCEL,
			},
		},
		Excel: KendoGridExcel{
			AllPages: TRUE,
		},
		DataSource: KendoDataSource{
			Data: []map[string]interface{}{
				{"name": "Jane Doe", "age": 30},
				{"name": "John Doe", "age": 33},
			},
			Schema: KendoSchema{
				Model: KendoDataModel{
					Id: "id",
					Fields: map[string]KendoField{
						"age": {
							Type: JAVASCRIPT_NUMBER,
						},
					},
				},
			},
		},
		Editable: KENDO_GRID_EDITOR_MODE_POPUP,
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: [{className: "btn-destroy",click: function(e){ /* command button click handler */ },iconClass: "k-icon k-i-copy",name: "custom",text: "details",},{name: "destroy",},],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_11() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "name", Title: "Name"},
			{Field: "age", Title: "Age"},
			{
				Command: []KendoGridColumnsCommand{
					{
						Name:      COLUMNS_COMMAND_DESTROY,
						Text:      "remove",
						ClassName: "btn-destroy",
						IconClass: KendoGridColumnsIconClass{
							Edit:   "k-icon k-i-edit",
							Update: "k-icon k-i-copy",
							Cancel: "k-icon k-i-arrow-60-up",
						},
					},
					{
						Name: COLUMNS_COMMAND_EDIT,
					},
				},
			},
		},
		ColumnMenu: TRUE,
		Toolbar: []KendoGridToolbar{
			{
				Name:      GRID_TOOLBAR_NAME_SAVE,
				IconClass: "k-icon k-i-copy",
			},
			{
				Name: GRID_TOOLBAR_NAME_EXCEL,
			},
		},
		Excel: KendoGridExcel{
			AllPages: TRUE,
		},
		DataSource: KendoDataSource{
			Data: []map[string]interface{}{
				{"name": "Jane Doe", "age": 30},
				{"name": "John Doe", "age": 33},
			},
			Schema: KendoSchema{
				Model: KendoDataModel{
					Id: "id",
					Fields: map[string]KendoField{
						"age": {
							Type: JAVASCRIPT_NUMBER,
						},
					},
				},
			},
		},
		Editable: KENDO_GRID_EDITOR_MODE_POPUP,
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: [{className: "btn-destroy",iconClass: {edit: "k-icon k-i-edit",update: "k-icon k-i-copy",cancel: "k-icon k-i-arrow-60-up",},name: "destroy",text: "remove",},{name: "edit",},],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_12() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},

		Columns: []KendoGridColumns{
			{
				Field: "name",
				Title: "Name",
				Filterable: KendoGridColumnsFilterable{
					Cell: KendoGridColumnsCell{
						DataSource: KendoDataSource{
							Data: []map[string]interface{}{
								{"someField": "Jane"},
								{"someField": "Jake"},
								{"someField": "John"},
							},
						},
						DataTextField: "someField",
					},
				},
			},
			{
				Field: "age",
				Title: "Age",
			},
			{
				Command: []KendoGridColumnsCommand{
					{
						Name:      COLUMNS_COMMAND_DESTROY,
						Text:      "remove",
						ClassName: "btn-destroy",
					},
					{
						Name: COLUMNS_COMMAND_EDIT,
					},
				}},
		},
		Filterable: KendoGridFilterable{
			Mode: KENDO_GRID_FILTERABLE_MODE_ROW,
		},
		ColumnMenu: TRUE,
		Toolbar: []KendoGridToolbar{
			{
				Name:      GRID_TOOLBAR_NAME_SAVE,
				IconClass: "k-icon k-i-copy",
			},
			{
				Name: GRID_TOOLBAR_NAME_EXCEL,
			},
		},
		Excel: KendoGridExcel{
			AllPages: TRUE,
		},
		DataSource: KendoDataSource{
			Data: []map[string]interface{}{
				{"name": "Jane Doe", "age": 30},
				{"name": "John Doe", "age": 33},
			},
			Schema: KendoSchema{
				Model: KendoDataModel{
					Id: "id",
					Fields: map[string]KendoField{
						"age": {
							Type: JAVASCRIPT_NUMBER,
						},
					},
				},
			},
		},
		Editable: KENDO_GRID_EDITOR_MODE_POPUP,
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: true,columns: [{field: "name",filterable: {cell: {dataSource:  new kendo.data.DataSource({data: [{"someField": "Jane",},{"someField": "Jake",},{"someField": "John",},],}),dataTextField: "someField",},},title: "Name",},{field: "age",title: "Age",},{command: [{className: "btn-destroy",name: "destroy",text: "remove",},{name: "edit",},],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},filterable: {mode: "row",},toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_13() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},

		Columns: []KendoGridColumns{
			{
				Field: "name",
				Title: "Name",
			},
			{
				Field: "age",
				Title: "Age",
			},
			{
				Command: []KendoGridColumnsCommand{
					{
						Name:      COLUMNS_COMMAND_DESTROY,
						Text:      "remove",
						ClassName: "btn-destroy",
					},
					{
						Name: COLUMNS_COMMAND_EDIT,
					},
				}},
		},
		ColumnMenu: KendoGridColumnMenu{
			Columns: FALSE,
		},
		Sortable: TRUE,
		Toolbar: []KendoGridToolbar{
			{
				Name:      GRID_TOOLBAR_NAME_SAVE,
				IconClass: "k-icon k-i-copy",
			},
			{
				Name: GRID_TOOLBAR_NAME_EXCEL,
			},
		},
		Excel: KendoGridExcel{
			AllPages: TRUE,
		},
		DataSource: KendoDataSource{
			Data: []map[string]interface{}{
				{"name": "Jane Doe", "age": 30},
				{"name": "John Doe", "age": 33},
			},
			Schema: KendoSchema{
				Model: KendoDataModel{
					Id: "id",
					Fields: map[string]KendoField{
						"age": {
							Type: JAVASCRIPT_NUMBER,
						},
					},
				},
			},
		},
		Editable: KENDO_GRID_EDITOR_MODE_POPUP,
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: {columns: false,},columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: [{className: "btn-destroy",name: "destroy",text: "remove",},{name: "edit",},],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: "popup",excel: {allpages: true,},sortable: true,toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_14() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},

		Columns: []KendoGridColumns{
			{
				Field: "name",
				Title: "Name",
			},
			{
				Field: "age",
				Title: "Age",
			},
		},
		ColumnMenu: KendoGridColumnMenu{
			Columns: FALSE,
		},
		Sortable: TRUE,
		DataSource: KendoDataSource{
			Data: []map[string]interface{}{
				{"name": "Jane Doe", "age": 30},
				{"name": "John Doe", "age": 33},
			},
			Schema: KendoSchema{
				Model: KendoDataModel{
					Id: "id",
					Fields: map[string]KendoField{
						"age": {
							Type: JAVASCRIPT_NUMBER,
						},
					},
				},
			},
		},
		Editable: TRUE,
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: {columns: false,},columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: true,sortable: true,});
}

func ExampleKendoUiGrid_ToHtml_15() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},

		Columns: []KendoGridColumns{
			{
				Field: "name",
				Title: "Name",
			},
			{
				Field: "age",
				Title: "Age",
			},
			{
				Command: []KendoGridColumnsCommand{
					{
						Name:      COLUMNS_COMMAND_DESTROY,
						Text:      "remove",
						ClassName: "btn-destroy",
					},
				}},
		},
		ColumnMenu: KendoGridColumnMenu{
			Columns: FALSE,
		},
		Sortable: TRUE,
		Toolbar: []KendoGridToolbar{
			{
				Name:      GRID_TOOLBAR_NAME_SAVE,
				IconClass: "k-icon k-i-copy",
			},
			{
				Name: GRID_TOOLBAR_NAME_EXCEL,
			},
		},
		Excel: KendoGridExcel{
			AllPages: TRUE,
		},
		DataSource: KendoDataSource{
			Data: []map[string]interface{}{
				{"name": "Jane Doe", "age": 30},
				{"name": "John Doe", "age": 33},
			},
			Schema: KendoSchema{
				Model: KendoDataModel{
					Id: "id",
					Fields: map[string]KendoField{
						"age": {
							Type: JAVASCRIPT_NUMBER,
						},
					},
				},
			},
		},
		Editable: KendoGridEditable{
			Confirmation: "Are you sure that you want to delete this record?",
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: {columns: false,},columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: [{className: "btn-destroy",name: "destroy",text: "remove",},],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"age": 33,"name": "John Doe",},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: {confirmation: "Are you sure that you want to delete this record?",},excel: {allpages: true,},sortable: true,toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_16() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},

		Columns: []KendoGridColumns{
			{
				Field: "name",
				Title: "Name",
			},
			{
				Field: "age",
				Title: "Age",
			},
			{
				Command: []KendoGridColumnsCommand{
					{
						Name:      COLUMNS_COMMAND_DESTROY,
						Text:      "remove",
						ClassName: "btn-destroy",
					},
				}},
		},
		ColumnMenu: KendoGridColumnMenu{
			Columns: FALSE,
		},
		Sortable: TRUE,
		Toolbar: []KendoGridToolbar{
			{
				Name:      GRID_TOOLBAR_NAME_SAVE,
				IconClass: "k-icon k-i-copy",
			},
			{
				Name: GRID_TOOLBAR_NAME_EXCEL,
			},
		},
		Excel: KendoGridExcel{
			AllPages: TRUE,
		},
		DataSource: KendoDataSource{
			Data: []map[string]interface{}{
				{"name": "Jane Doe", "age": 30},
				{"name": "John Doe", "age": 33},
			},
			Schema: KendoSchema{
				Model: KendoDataModel{
					Id: "id",
					Fields: map[string]KendoField{
						"age": {
							Type: JAVASCRIPT_NUMBER,
						},
					},
				},
			},
		},
		Editable: KendoGridEditable{
			Confirmation: JavaScript{
				Code: `function(e) { return  "Are you sure that you want to delete record for " + e.name + "?"; }`,
			},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: {columns: false,},columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: [{className: "btn-destroy",name: "destroy",text: "remove",},],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: {confirmation: function(e) { return  "Are you sure that you want to delete record for " + e.name + "?"; },},excel: {allpages: true,},sortable: true,toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_17() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},

		Columns: []KendoGridColumns{
			{
				Field: "name",
				Title: "Name",
			},
			{
				Field: "age",
				Title: "Age",
			},
			{
				Command: []KendoGridColumnsCommand{
					{
						Name:      COLUMNS_COMMAND_DESTROY,
						Text:      "remove",
						ClassName: "btn-destroy",
					},
				}},
		},
		ColumnMenu: KendoGridColumnMenu{
			Columns: FALSE,
		},
		Filterable: KendoGridFilterable{
			Extra: FALSE,
		},
		Sortable: TRUE,
		Toolbar: []KendoGridToolbar{
			{
				Name:      GRID_TOOLBAR_NAME_SAVE,
				IconClass: "k-icon k-i-copy",
			},
			{
				Name: GRID_TOOLBAR_NAME_EXCEL,
			},
		},
		Excel: KendoGridExcel{
			AllPages: TRUE,
		},
		DataSource: KendoDataSource{
			Data: []map[string]interface{}{
				{"name": "Jane Doe", "age": 30},
				{"name": "John Doe", "age": 33},
			},
			Schema: KendoSchema{
				Model: KendoDataModel{
					Id: "id",
					Fields: map[string]KendoField{
						"age": {
							Type: JAVASCRIPT_NUMBER,
						},
					},
				},
			},
		},
		Editable: KendoGridEditable{
			Confirmation: JavaScript{
				Code: `function(e) { return  "Are you sure that you want to delete record for " + e.name + "?"; }`,
			},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: {columns: false,},columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: [{className: "btn-destroy",name: "destroy",text: "remove",},],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: {confirmation: function(e) { return  "Are you sure that you want to delete record for " + e.name + "?"; },},excel: {allpages: true,},sortable: true,toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_18() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{
				Field: "name",
				Title: "Name",
			},
			{
				Field: "age",
				Title: "Age",
			},
			{
				Command: []KendoGridColumnsCommand{
					{
						Name:      COLUMNS_COMMAND_DESTROY,
						Text:      "remove",
						ClassName: "btn-destroy",
					},
				}},
		},
		ColumnMenu: KendoGridColumnMenu{
			Columns: FALSE,
		},
		Filterable: KendoGridFilterable{
			Messages: KendoGridFilterableMessages{
				And: "and",
				Or:  "or",
			},
		},
		Sortable: TRUE,
		Toolbar: []KendoGridToolbar{
			{
				Name:      GRID_TOOLBAR_NAME_SAVE,
				IconClass: "k-icon k-i-copy",
			},
			{
				Name: GRID_TOOLBAR_NAME_EXCEL,
			},
		},
		Excel: KendoGridExcel{
			AllPages: TRUE,
		},
		DataSource: KendoDataSource{
			Data: []map[string]interface{}{
				{"name": "Jane Doe", "age": 30},
				{"name": "John Doe", "age": 33},
			},
			Schema: KendoSchema{
				Model: KendoDataModel{
					Id: "id",
					Fields: map[string]KendoField{
						"age": {
							Type: JAVASCRIPT_NUMBER,
						},
					},
				},
			},
		},
		Editable: KendoGridEditable{
			Confirmation: JavaScript{
				Code: `function(e) { return  "Are you sure that you want to delete record for " + e.name + "?"; }`,
			},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: {columns: false,},columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: [{className: "btn-destroy",name: "destroy",text: "remove",},],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: {confirmation: function(e) { return  "Are you sure that you want to delete record for " + e.name + "?"; },},excel: {allpages: true,},filterable: {messages: {and: "and",or: "or",},},sortable: true,toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_19() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{
				Field: "name",
				Title: "Name",
			},
			{
				Field: "age",
				Title: "Age",
			},
			{
				Command: []KendoGridColumnsCommand{
					{
						Name:      COLUMNS_COMMAND_DESTROY,
						Text:      "remove",
						ClassName: "btn-destroy",
					},
					{
						Name:      COLUMNS_COMMAND_EDIT,
						Text:      "edit",
						ClassName: "btn-edit",
					},
				},
			},
		},
		ColumnMenu: KendoGridColumnMenu{
			Columns: FALSE,
		},
		Filterable: KendoGridFilterable{
			Operators: KendoGridFilterableOperators{
				String: KendoGridFilterableOperatorsString{
					IsNotEmpty:     "Is not empty",
					IsNull:         "Is null",
					IsEmpty:        "Is empty",
					DoesNotContain: "Does not contain",
					Contains:       "Contains",
					EndsWith:       "Ends with",
					StartsWith:     "Starts with",
					Eq:             "Is equal to",
					Neq:            "Is not equal to",
				},
				Date: KendoGridFilterableOperatorsDate{
					Neq:    "Is not equal to",
					IsNull: "Is null",
					Eq:     "Is equal to",
					Gt:     "Is after",
					Lte:    "Is before or equal to",
					Lt:     "Is before",
					Gte:    "Is after or equal to",
				},
				Number: KendoGridFilterableOperatorsNumber{
					Neq:       "Is not equal to",
					Lt:        "Is less than",
					Eq:        "Is equal to",
					IsNull:    "Is null",
					Lte:       "Is less than or equal to",
					Gt:        "Is greater than",
					Gte:       "Is greater than or equal to",
					IsNotNull: "Is not null",
				},
				Enums: KendoGridFilterableOperatorsEnums{
					Eq:        "Is equal to",
					IsNull:    "Is null",
					IsNotNull: "Is not null",
					Neq:       "Is not equal to",
				},
			},
		},
		Sortable: TRUE,
		Toolbar: []KendoGridToolbar{
			{
				Name:      GRID_TOOLBAR_NAME_SAVE,
				IconClass: "k-icon k-i-copy",
			},
			{
				Name: GRID_TOOLBAR_NAME_EXCEL,
			},
		},
		Excel: KendoGridExcel{
			AllPages: TRUE,
		},
		DataSource: KendoDataSource{
			Data: []map[string]interface{}{
				{"name": "Jane Doe", "age": 30},
				{"name": "John Doe", "age": 33},
			},
			Schema: KendoSchema{
				Model: KendoDataModel{
					Id: "id",
					Fields: map[string]KendoField{
						"age": {
							Type: JAVASCRIPT_NUMBER,
						},
					},
				},
			},
		},
		Editable: KendoGridEditable{
			Mode: KENDO_GRID_EDITOR_MODE_POPUP,
			Confirmation: JavaScript{
				Code: `function(e) { return  "Are you sure that you want to delete record for " + e.name + "?"; }`,
			},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columnMenu: {columns: false,},columns: [{field: "name",title: "Name",},{field: "age",title: "Age",},{command: [{className: "btn-destroy",name: "destroy",text: "remove",},{className: "btn-edit",name: "edit",text: "edit",},],},],dataSource:  new kendo.data.DataSource({data: [{"name": "Jane Doe","age": 30,},{"name": "John Doe","age": 33,},],schema: {model: {id: "id",fields: {"age": {type: "number",}, },},},}),editable: {confirmation: function(e) { return  "Are you sure that you want to delete record for " + e.name + "?"; },mode: "popup",},excel: {allpages: true,},filterable: {operators: {enums: {eq: "Is equal to",isnull: "Is null",isnotnull: "Is not null",neq: "Is not equal to",},date: {gt: "Is after",lte: "Is before or equal to",neq: "Is not equal to",isnull: "Is null",lt: "Is before",gte: "Is after or equal to",eq: "Is equal to",},string: {isnotempty: "Is not empty",isnull: "Is null",isempty: "Is empty",doesnotcontain: "Does not contain",contains: "Contains",endswith: "Ends with",startswith: "Starts with",eq: "Is equal to",neq: "Is not equal to",},number: {neq: "Is not equal to",lt: "Is less than",eq: "Is equal to",isnull: "Is null",lte: "Is less than or equal to",gt: "Is greater than",gte: "Is greater than or equal to",isnotnull: "Is not null",},},},sortable: true,toolbar: [{iconClass: "k-icon k-i-copy",name: "save",},{name: "excel",},],});
}

func ExampleKendoUiGrid_ToHtml_20() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "productName"},
			{Field: "category"},
		},
		Sortable: TRUE,
		Groupable: KendoGridGroupable{
			ShowFooter: TRUE,
			Enabled:    TRUE,
			Messages: KendoGridGroupableMessages{
				Empty: "Drop columns here",
			},
		},
		DataSource: []map[string]interface{}{
			{"productName": "Tea", "category": "Beverages"},
			{"productName": "Coffee", "category": "Beverages"},
			{"productName": "Ham", "category": "Food"},
			{"productName": "Bread", "category": "Food"},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columns: [{field: "productName",},{field: "category",},],dataSource: [{"productName": "Tea","category": "Beverages",},{"productName": "Coffee","category": "Beverages",},{"productName": "Ham","category": "Food",},{"productName": "Bread","category": "Food",},],groupable: {showFooter: true,enabled: true,messages: {empty: "Drop columns here",},},sortable: true,});
}

func ExampleKendoUiGrid_ToHtml_21() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "productName"},
			{Field: "category"},
		},
		Sortable:  TRUE,
		Groupable: TRUE,
		DataSource: []map[string]interface{}{
			{"productName": "Tea", "category": "Beverages"},
			{"productName": "Coffee", "category": "Beverages"},
			{"productName": "Ham", "category": "Food"},
			{"productName": "Bread", "category": "Food"},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columns: [{field: "productName",},{field: "category",},],dataSource: [{"productName": "Tea","category": "Beverages",},{"category": "Beverages","productName": "Coffee",},{"productName": "Ham","category": "Food",},{"productName": "Bread","category": "Food",},],groupable: true,sortable: true,});
}

func ExampleKendoUiGrid_ToHtml_22() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "productName"},
			{Field: "category"},
		},
		Sortable:  TRUE,
		Groupable: TRUE,
		Toolbar:   []KendoGridToolBarName{GRID_TOOLBAR_NAME_CREATE, GRID_TOOLBAR_NAME_SAVE, GRID_TOOLBAR_NAME_CANCEL},
		Messages: KendoGridMessages{
			Commands: KendoGridMessagesCommands{
				Cancel:     "Cancel changes",
				CancelEdit: "Cancel",
				Create:     "Add new record",
				Destroy:    "Delete",
				Edit:       "Edit",
				Save:       "Save changes",
				Select:     "Select",
				Update:     "Update",
			},
		},
		DataSource: []map[string]interface{}{
			{"productName": "Tea", "category": "Beverages"},
			{"productName": "Coffee", "category": "Beverages"},
			{"productName": "Ham", "category": "Food"},
			{"productName": "Bread", "category": "Food"},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columns: [{field: "productName",},{field: "category",},],dataSource: [{"productName": "Tea","category": "Beverages",},{"productName": "Coffee","category": "Beverages",},{"productName": "Ham","category": "Food",},{"productName": "Bread","category": "Food",},],groupable: true,messages: {commands: {select: "Select",save: "Save changes",destroy: "Delete",update: "Update",cancel: "Cancel changes",edit: "Edit",create: "Add new record",canceledit: "Cancel",},},sortable: true,toolbar: ["create","save","cancel",],});
}

func ExampleKendoUiGrid_ToHtml_23() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "productName"},
			{Field: "category"},
		},
		NoRecords: TRUE,
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columns: [{field: "productName",},{field: "category",},],dataSource: [{"productName": "Tea","category": "Beverages",},{"productName": "Coffee","category": "Beverages",},{"productName": "Ham","category": "Food",},{"productName": "Bread","category": "Food",},],groupable: true,messages: {commands: {select: "Select",save: "Save changes",destroy: "Delete",update: "Update",cancel: "Cancel changes",edit: "Edit",create: "Add new record",canceledit: "Cancel",},},sortable: true,toolbar: ["create","save","cancel",],});
}

func ExampleKendoUiGrid_ToHtml_24() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "productName"},
			{Field: "category"},
		},
		NoRecords: KendoGridNoRecords{
			Template: "No data available on current page. Current page is: #=this.dataSource.page()#",
		},
		Pageable: KendoGridPageable{
			PageSize:      10,
			ButtonCount:   2,
			Refresh:       TRUE,
			Numeric:       TRUE,
			AlwaysVisible: TRUE,
			Info:          FALSE,
			Messages: KendoGridPageableMessages{
				Display:      "Showing {0}-{1} from {2} data items",
				Empty:        "No data",
				Page:         "Enter page",
				Of:           "from {0}",
				ItemsPerPage: "data items per page",
				First:        "First page",
				Last:         "Last page",
				Next:         "Next page",
				Previous:     "Previous page",
				Refresh:      "Refresh the grid",
				MorePages:    "More pages",
			},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columns: [{field: "productName",},{field: "category",},],noRecords: { template: "No data available on current page. Current page is: #=this.dataSource.page()#",},pageable: { numeric: true,refresh: true,alwaysVisible: true,messages: { display: "Showing {0}-{1} from {2} data items",empty: "No data",page: "Enter page",of: "from {0}",itemsPerPage: "data items per page",first: "First page",last: "Last page",next: "Next page",previous: "Previous page",refresh: "Refresh the grid",morePages: "More pages",},buttonCount: 2,pageSize: 10,info: false,},});
}

func ExampleKendoUiGrid_ToHtml_25() {
	el := KendoUiGrid{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "grid",
			},
		},
		Columns: []KendoGridColumns{
			{Field: "productName"},
			{Field: "category"},
		},
		DataSource: []map[string]interface{}{
			{"productName": "Tea", "category": "Beverages"},
			{"productName": "Coffee", "category": "Beverages"},
			{"productName": "Ham", "category": "Food"},
			{"productName": "Bread", "category": "Food"},
		},
		Toolbar: []KendoGridToolBarName{GRID_TOOLBAR_NAME_PDF},
		NoRecords: KendoGridNoRecords{
			Template: "No data available on current page. Current page is: #=this.dataSource.page()#",
		},
		Pdf: KendoGridPdf{
			Title: "Drinks",
		},
		Pageable: KendoGridPageable{
			PageSize:      10,
			ButtonCount:   2,
			Refresh:       TRUE,
			Numeric:       TRUE,
			AlwaysVisible: TRUE,
			Info:          FALSE,
			Messages: KendoGridPageableMessages{
				Display:      "Showing {0}-{1} from {2} data items",
				Empty:        "No data",
				Page:         "Enter page",
				Of:           "from {0}",
				ItemsPerPage: "data items per page",
				First:        "First page",
				Last:         "Last page",
				Next:         "Next page",
				Previous:     "Previous page",
				Refresh:      "Refresh the grid",
				MorePages:    "More pages",
			},
		},
	}

	fmt.Printf("%s\n", el.ToHtml())
	fmt.Printf("%s", el.ToJavaScript())

	//Output:
	//<div id="grid"></div>
	//$("#grid").kendoGrid({columns: [{field: "productName",},{field: "category",},],noRecords: { template: "No data available on current page. Current page is: #=this.dataSource.page()#",},pageable: { numeric: true,refresh: true,alwaysVisible: true,messages: { display: "Showing {0}-{1} from {2} data items",empty: "No data",page: "Enter page",of: "from {0}",itemsPerPage: "data items per page",first: "First page",last: "Last page",next: "Next page",previous: "Previous page",refresh: "Refresh the grid",morePages: "More pages",},buttonCount: 2,pageSize: 10,info: false,},});
}
