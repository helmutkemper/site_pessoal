package telerik

import "fmt"

func ExampleKendoDataSource_ToJavaScript_Animation_Object() {
	javaScript := KendoUiMultiSelect{
		Html: HtmlElementFormSelect{
			Global: HtmlGlobalAttributes{
				Id: "multiselect",
			},
		},
		Animation: KendoAnimation{
			Open: KendoOpen{
				Effects:  EFFECT_ZOOM_IN,
				Duration: 300,
			},
			Close: KendoClose{
				Effects:  EFFECT_ZOOM_OUT,
				Duration: 300,
			},
		},
		AutoBind:           FALSE,
		AutoClose:          FALSE,
		AutoWidth:          TRUE,
		ClearButton:        FALSE,
		DataSource:         "selectDataSource",
		DataTextField:      "Name",
		DataValueField:     "Id",
		Delay:              100,
		Enable:             TRUE,
		EnforceMinLength:   TRUE,
		Filter:             FILTER_CONTAINS,
		FixedGroupTemplate: "Fixed header: #: data #",
		FooterTemplate:     "Total <strong>#: instance.dataSource.total() #</strong> items found",
		GroupTemplate:      "Group template: #: data #",
		Height:             500,
		HighlightFirst:     FALSE,
		IgnoreCase:         FALSE,
		MinLength:          3,
		MaxSelectedItems:   3,
		NoDataTemplate:     "No Data!",
		Placeholder:        "Select...",
		Popup: KendoPopup{
			AppendTo: JavaScript{
				Code: "$('#container')",
			},
			Origin:   ORIGIN_TOP_LEFT,
			Position: POSITION_BOTTOM_RIGHT,
		},
		HeaderTemplate: "<div><h2>Fruits</h2></div>",
		ItemTemplate:   "<span><img src='/img/#: id #.png' alt='#: name #' />#: name #</span>",
		TagTemplate: JavaScript{
			Code: "kendo.template($('#tagTemplate').html())",
		},
		TagMode: TAG_MODE_MULTIPLE,
		Value: []map[string]interface{}{
			{"productName": "Item 1", "productId": "1"},
			{"productName": "Item 2", "productId": "2"},
			{"productName": "Item 3", "productId": "3"},
		},
		ValuePrimitive: TRUE,
		Virtual: KendoVirtual{
			ItemHeight: 26,
			ValueMapper: JavaScript{
				Code: `function(options) {
	$.ajax({
		url: 'https://demos.telerik.com/kendo-ui/service/Orders/ValueMapper',
		type: 'GET',
		dataType: 'jsonp',
		data: convertValues(options.value),
		success: function (data) {
			//the **data** is either index or array of indices.
			//Example:
			// 10258 -> 10 (index in the Orders collection)
			// [10258, 10261] -> [10, 14] (indices in the Orders collection)
			options.success(data);
		}
	})
}`,
			},
			MapValueTo: MAP_VALUE_TO_INDEX,
		},
	}

	fmt.Printf("%s", javaScript.ToJavaScript())

	// Output:
	// $("#multiselect").kendoMultiSelect({animation: { close: { effects: "zoom:out",duration: 300,},open: { effects: "zoom:in",duration: 300,},},autoBind: false,autoClose: false,autoWidth: true,clearButton: false,dataSource: "selectDataSource",dataTextField: "Name",dataValueField: "Id",delay: 100,enable: true,enforceMinLength: true,filter: "contains",fixedGroupTemplate: "Fixed header: #: data #",footerTemplate: "Total <strong>#: instance.dataSource.total() #</strong> items found",groupTemplate: "Group template: #: data #",height: 500,highlightFirst: false,ignoreCase: false,minLength: 3,maxSelectedItems: 3,noDataTemplate: "No Data!",placeholder: "Select...",popup: { appendTo: $('#container'),origin: "top left",position: "bottom right",},headerTemplate: "<div><h2>Fruits</h2></div>",itemTemplate: "<span><img src='/img/#: id #.png' alt='#: name #' />#: name #</span>",tagTemplate: kendo.template($('#tagTemplate').html()),tagMode: "multiple",value: [{"productName": "Item 1","productId": "1",},{"productName": "Item 2","productId": "2",},{"productName": "Item 3","productId": "3",},],valuePrimitive: true,virtual: { itemHeight: 26,mapValueTo: "index",valueMapper: function(options) {
	//	$.ajax({
	//		url: 'https://demos.telerik.com/kendo-ui/service/Orders/ValueMapper',
	//		type: 'GET',
	//		dataType: 'jsonp',
	//		data: convertValues(options.value),
	//		success: function (data) {
	//			//the **data** is either index or array of indices.
	//			//Example:
	//			// 10258 -> 10 (index in the Orders collection)
	//			// [10258, 10261] -> [10, 14] (indices in the Orders collection)
	//			options.success(data);
	//		}
	//	})
	//},},});
}

func ExampleKendoDataSource_ToJavaScript_Animation_Boolean_FALSE() {
	javaScript := KendoUiMultiSelect{
		Html: HtmlElementFormSelect{
			Global: HtmlGlobalAttributes{
				Id: "multiselect",
			},
		},
		Animation:          FALSE,
		AutoBind:           FALSE,
		AutoClose:          FALSE,
		AutoWidth:          TRUE,
		ClearButton:        FALSE,
		DataSource:         "selectDataSource",
		DataTextField:      "Name",
		DataValueField:     "Id",
		Delay:              100,
		Enable:             TRUE,
		EnforceMinLength:   TRUE,
		Filter:             FILTER_CONTAINS,
		FixedGroupTemplate: "Fixed header: #: data #",
		FooterTemplate:     "Total <strong>#: instance.dataSource.total() #</strong> items found",
		GroupTemplate:      "Group template: #: data #",
		Height:             500,
		HighlightFirst:     FALSE,
		IgnoreCase:         FALSE,
		MinLength:          3,
		MaxSelectedItems:   3,
		NoDataTemplate:     "No Data!",
		Placeholder:        "Select...",
		Popup: KendoPopup{
			AppendTo: JavaScript{
				Code: "$('#container')",
			},
			Origin:   ORIGIN_TOP_LEFT,
			Position: POSITION_BOTTOM_RIGHT,
		},
		HeaderTemplate: "<div><h2>Fruits</h2></div>",
		ItemTemplate:   "<span><img src='/img/#: id #.png' alt='#: name #' />#: name #</span>",
		TagTemplate: JavaScript{
			Code: "kendo.template($('#tagTemplate').html())",
		},
		TagMode: TAG_MODE_MULTIPLE,
		Value: []map[string]interface{}{
			{"productName": "Item 1", "productId": "1"},
			{"productName": "Item 2", "productId": "2"},
			{"productName": "Item 3", "productId": "3"},
		},
		ValuePrimitive: TRUE,
		Virtual: KendoVirtual{
			ItemHeight: 26,
			ValueMapper: JavaScript{
				Code: `function(options) {
	$.ajax({
		url: 'https://demos.telerik.com/kendo-ui/service/Orders/ValueMapper',
		type: 'GET',
		dataType: 'jsonp',
		data: convertValues(options.value),
		success: function (data) {
			//the **data** is either index or array of indices.
			//Example:
			// 10258 -> 10 (index in the Orders collection)
			// [10258, 10261] -> [10, 14] (indices in the Orders collection)

			options.success(data);
		}
	})
}`,
			},
			MapValueTo: MAP_VALUE_TO_INDEX,
		},
	}

	fmt.Printf("%s", javaScript.ToJavaScript())

	// Output:
	// $("multiselect")kendoMultiSelect({
	//   animation: false,
	//   autoBind: false,
	//   autoClose: false,
	//   autoWidth: true,
	//   clearButton: false,
	//   dataSource: "selectDataSource",
	//   dataTextField: "Name",
	//   dataValueField: "Id",
	//   delay: 100,
	//   enable: true,
	//   enforceMinLength: true,
	//   filter: "contains",
	//   fixedGroupTemplate: "Fixed header: #: data #",
	//   footerTemplate: "Total <strong>#: instance.dataSource.total() #</strong> items found",
	//   groupTemplate: "Group template: #: data #",
	//   height: 500,
	//   highlightFirst: false,
	//   ignoreCase: false,
	//   minLength: 3,
	//   maxSelectedItems: 3,
	//   noDataTemplate: "No Data!",
	//   placeholder: "Select...",
	//   popup: {
	//     appendTo: $('#container'),
	//     origin: "top left",
	//     position: "bottom right"
	//   },
	//   headerTemplate: "<div><h2>Fruits</h2></div>",
	//   itemTemplate: "<span><img src='/img/#: id #.png' alt='#: name #' />#: name #</span>",
	//   tagTemplate: kendo.template($('#tagTemplate').html()),
	//   tagMode: "multiple",
	//   value: [
	//     {"productName": "Item 1","productId": "1"},
	//     {"productName": "Item 2","productId": "2"},
	//     {"productName": "Item 3","productId": "3"}
	//   ],
	//   valuePrimitive: true,
	//   virtual: {
	//     itemHeight: 26,
	//     mapValueTo: "index",
	//     valueMapper: function(options) {
	//       $.ajax({
	//         url: 'https://demos.telerik.com/kendo-ui/service/Orders/ValueMapper',
	//         type: 'GET',
	//         dataType: 'jsonp',
	//         data: convertValues(options.value),
	//         success: function (data) {
	//           //the **data** is either index or array of indices.
	//           //Example:
	//           // 10258 -> 10 (index in the Orders collection)
	//           // [10258, 10261] -> [10, 14] (indices in the Orders collection)
	//           options.success(data);
	//         }
	//       })
	//     }
	//   }
	// });
}

func ExampleKendoDataSource_ToJavaScript_Virtual_Boolean_FALSE() {
	javaScript := KendoUiMultiSelect{
		Html: HtmlElementFormSelect{
			Global: HtmlGlobalAttributes{
				Id: "multiselect",
			},
		},
		Animation: KendoAnimation{
			Open: KendoOpen{
				Effects:  EFFECT_ZOOM_IN,
				Duration: 300,
			},
			Close: KendoClose{
				Effects:  EFFECT_ZOOM_OUT,
				Duration: 300,
			},
		},
		AutoBind:           FALSE,
		AutoClose:          FALSE,
		AutoWidth:          TRUE,
		ClearButton:        FALSE,
		DataSource:         "selectDataSource",
		DataTextField:      "Name",
		DataValueField:     "Id",
		Delay:              100,
		Enable:             TRUE,
		EnforceMinLength:   TRUE,
		Filter:             FILTER_CONTAINS,
		FixedGroupTemplate: "Fixed header: #: data #",
		FooterTemplate:     "Total <strong>#: instance.dataSource.total() #</strong> items found",
		GroupTemplate:      "Group template: #: data #",
		Height:             500,
		HighlightFirst:     FALSE,
		IgnoreCase:         FALSE,
		MinLength:          3,
		MaxSelectedItems:   3,
		NoDataTemplate:     "No Data!",
		Placeholder:        "Select...",
		Popup: KendoPopup{
			AppendTo: JavaScript{
				Code: "$('#container')",
			},
			Origin:   ORIGIN_TOP_LEFT,
			Position: POSITION_BOTTOM_RIGHT,
		},
		HeaderTemplate: "<div><h2>Fruits</h2></div>",
		ItemTemplate:   "<span><img src='/img/#: id #.png' alt='#: name #' />#: name #</span>",
		TagTemplate: JavaScript{
			Code: "kendo.template($('#tagTemplate').html())",
		},
		TagMode: TAG_MODE_MULTIPLE,
		Value: []map[string]interface{}{
			{"productName": "Item 1", "productId": "1"},
			{"productName": "Item 2", "productId": "2"},
			{"productName": "Item 3", "productId": "3"},
		},
		ValuePrimitive: TRUE,
		Virtual:        FALSE,
	}

	fmt.Printf("%s", javaScript.ToJavaScript())

	// Output:
	// $("multiselect")kendoMultiSelect({
	//   animation: {
	//     close: {
	//       effects: "zoom:out",
	//       duration: 300
	//     },
	//     open: {
	//       effects: "zoom:in",
	//       duration: 300
	//     }
	//   },
	//   autoBind: false,
	//   autoClose: false,
	//   autoWidth: true,
	//   clearButton: false,
	//   dataSource: "selectDataSource",
	//   dataTextField: "Name",
	//   dataValueField: "Id",
	//   delay: 100,
	//   enable: true,
	//   enforceMinLength: true,
	//   filter: "contains",
	//   fixedGroupTemplate: "Fixed header: #: data #",
	//   footerTemplate: "Total <strong>#: instance.dataSource.total() #</strong> items found",
	//   groupTemplate: "Group template: #: data #",
	//   height: 500,
	//   highlightFirst: false,
	//   ignoreCase: false,
	//   minLength: 3,
	//   maxSelectedItems: 3,
	//   noDataTemplate: "No Data!",
	//   placeholder: "Select...",
	//   popup: {
	//     appendTo: $('#container'),
	//     origin: "top left",
	//     position: "bottom right"
	//   },
	//   headerTemplate: "<div><h2>Fruits</h2></div>",
	//   itemTemplate: "<span><img src='/img/#: id #.png' alt='#: name #' />#: name #</span>",
	//   tagTemplate: kendo.template($('#tagTemplate').html()),
	//   tagMode: "multiple",
	//   value: [
	//     {"productName": "Item 1","productId": "1"},
	//     {"productName": "Item 2","productId": "2"},
	//     {"productName": "Item 3","productId": "3"}
	//   ],
	//   valuePrimitive: true,
	//   virtual: false
	// });
}
