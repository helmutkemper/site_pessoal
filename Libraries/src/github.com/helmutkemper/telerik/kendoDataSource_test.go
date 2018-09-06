package telerik

import "fmt"

func ExampleKendoDataSource_ToJavaScript_OfflineStorage_1() {
	javaScript := KendoDataSource{
		//VarName: "dataSource",
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
	}

	fmt.Printf("%s", javaScript.ToJavaScript())

	// Output:
	// var dataSource = new kendo.data.DataSource({aggregate: [{aggregate: "sum",field: "age",},{aggregate: "min",field: "age",},{aggregate: "max",field: "age",},],autoSync: true,batch: true,data: [{"name": "Jane Doe","age": 30,},{"name": "Jane Doe","age": 33,},],inPlaceSort: false,offlineStorage: "products-offline",});
}

func ExampleKendoDataSource_ToJavaScript_OfflineStorage_2() {
	javaScript := KendoDataSource{
		//VarName: "dataSource",
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
		InPlaceSort: FALSE,
		OfflineStorage: KendoOfflineStorage{
			GetItem: "products-key",
			SetItem: "products-key",
		},
	}

	fmt.Printf("%s", javaScript.ToJavaScript())

	// Output:
	// var dataSource = new kendo.data.DataSource({aggregate: [{aggregate: "sum",field: "age",},{aggregate: "min",field: "age",},{aggregate: "max",field: "age",},],autoSync: true,batch: true,data: [{"age": 30,"name": "Jane Doe",},{"name": "Jane Doe","age": 33,},],inPlaceSort: false,offlineStorage: { getItem: function() { return JSON.parse(sessionStorage.getItem("products-key")); }, setItem: function(item) { sessionStorage.setItem("products-key", JSON.stringify(item)); } },});
}

func ExampleKendoDataSource_ToJavaScript_OfflineStorage_3() {
	javaScript := KendoDataSource{
		//VarName: "dataSource",
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
		InPlaceSort: FALSE,
		OfflineStorage: JavaScript{
			Code: `{ getItem: function() { return JSON.parse(sessionStorage.getItem("products-key")); }, setItem: function(item) { sessionStorage.setItem("products-key", JSON.stringify(item)); } }`,
		},
	}

	fmt.Printf("%s", javaScript.ToJavaScript())

	// Output:
	// var dataSource = new kendo.data.DataSource({aggregate: [{aggregate: "sum",field: "age",},{aggregate: "min",field: "age",},{aggregate: "max",field: "age",},],autoSync: true,batch: true,data: [{"name": "Jane Doe","age": 30,},{"name": "Jane Doe","age": 33,},],inPlaceSort: false,offlineStorage: { getItem: function() { return JSON.parse(sessionStorage.getItem("products-key")); }, setItem: function(item) { sessionStorage.setItem("products-key", JSON.stringify(item)); } },});
}

func ExampleKendoDataSource_ToJavaScript_Filter_Array() {
	javaScript := KendoDataSource{
		//VarName: "dataSource",
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
		Filter: []KendoComplexFilter{
			{
				Field:    "category",
				Operator: OPERATOR_EQUAL_TO,
				Value:    "Beverages",
			},
			{
				Field:    "name",
				Operator: OPERATOR_NOT_EQUAL_TO,
				Value:    "Coffee",
			},
		},
		Group: []KendoGroup{
			{
				Field: "category",
			},
			{
				Field: "subcategory",
			},
		},
		InPlaceSort: FALSE,
		OfflineStorage: JavaScript{
			Code: `{ getItem: function() { return JSON.parse(sessionStorage.getItem("products-key")); }, setItem: function(item) { sessionStorage.setItem("products-key", JSON.stringify(item)); } }`,
		},
		Page:     1,
		PageSize: 25,
		Schema: KendoSchema{
			Aggregates: "aggregates",
			Data:       "statuses",
			Errors:     "errors",
			Groups:     "groups",
			Model: KendoDataModel{
				Id: "personId",
			},
			Parser: JavaScript{
				Code: "function(response) { var products = []; for (var i = 0; i < response.length; i++) { var product = { id: response[i].ProductID, name: response[i].ProductName }; products.push(product); } return products; }",
			},
			Total: "total",
			Type:  KENDO_TYPE_DATA_JSON,
		},
		ServerAggregates: FALSE,
		ServerFiltering:  FALSE,
		ServerGrouping:   FALSE,
		ServerPaging:     FALSE,
		ServerSorting:    FALSE,
		Sort: KendoSort{
			Dir:   DIRECTION_ASC,
			Field: "age",
			Compare: JavaScript{
				Code: "function(a, b) { return numbers[a.item] - numbers[b.item]; }",
			},
		},
		Transport: KendoTransport{
			Create: KendoCreate{
				Cache:       FALSE,
				ContentType: "application/json",
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
				DataType: KENDO_TYPE_DATA_JSON_JSON,
				Type:     HTML_METHOD_GET,
				Url:      "https://demos.telerik.com/kendo-ui/service/products/create",
			},
			Destroy: KendoDestroy{
				Cache:       FALSE,
				ContentType: "application/json",
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
				DataType: KENDO_TYPE_DATA_JSON_JSON,
				Type:     HTML_METHOD_DELETE,
				Url:      "https://demos.telerik.com/kendo-ui/service/products",
			},
			ParameterMap: JavaScript{
				Code: "function(data, type) { if (type == 'read') { /* send take as '$top' and skip as '$skip' */ return { $top: data.take, $skip: data.skip } } }",
			},
			Push: JavaScript{
				Code: "function(callbacks) { hub.on('create', function(result) { console.log('push create'); callbacks.pushCreate(result); }); hub.on('update', function(result) { console.log('push update'); callbacks.pushUpdate(result); }); hub.on('destroy', function(result) { console.log('push destroy'); callbacks.pushDestroy(result); }); }",
			},
			Read: KendoRead{
				Cache:       FALSE,
				ContentType: "application/json",
				Data: map[string]interface{}{
					"skip": 0,
					"take": 2,
				},
				DataType: KENDO_TYPE_DATA_JSON_JSON,
				Type:     HTML_METHOD_GET,
				Url:      "https://demos.telerik.com/kendo-ui/service/products",
			},
			Submit: JavaScript{
				Code: "function(e) { var data = e.data; console.log(data); /* send batch update to desired URL, then notify success/error */ e.success(data.updated,'update'); e.success(data.created,'create'); e.success(data.destroyed,'destroy'); e.error(null, 'customerror', 'custom error'); }",
			},
			Update: KendoUpdate{
				Cache:       FALSE,
				ContentType: "application/json",
				Data: map[string]interface{}{
					"skip": 0,
					"take": 2,
				},
				DataType: KENDO_TYPE_DATA_JSON_JSON,
				Type:     HTML_METHOD_PUT,
				Url:      "https://demos.telerik.com/kendo-ui/service/products",
			},
		},
		Type: KENDO_TYPE_ODATA,
	}

	fmt.Printf("%s", javaScript.ToJavaScript())

	// Output:
	// var dataSource = new kendo.data.DataSource({aggregate: [{aggregate: "sum",field: "age",},{aggregate: "min",field: "age",},{aggregate: "max",field: "age",},],autoSync: true,batch: true,data: [{"name": "Jane Doe","age": 30,},{"name": "Jane Doe","age": 33,},],filter: [{field: "category",operator: "eq",value: "Beverages",},{field: "name",operator: "neq",value: "Coffee",},],group: [{field: "category",},{field: "subcategory",},],inPlaceSort: false,offlineStorage: { getItem: function() { return JSON.parse(sessionStorage.getItem("products-key")); }, setItem: function(item) { sessionStorage.setItem("products-key", JSON.stringify(item)); } },page: 1,pageSize: 25,schema: {aggregates: "aggregates",data: "statuses",errors: "errors",groups: "groups",model: {id: "personId",},groups: function(response) { var products = []; for (var i = 0; i < response.length; i++) { var product = { id: response[i].ProductID, name: response[i].ProductName }; products.push(product); } return products; },total: "total",type: "json",},serverAggregates: false,serverFiltering: false,serverGrouping: false,serverPaging: false,serverSorting: false,sort: { dir: "asc",field: "age",compare: function(a, b) { return numbers[a.item] - numbers[b.item]; },},transport: { destroy: { cache: false,contentType: "application/json",contentType: [{"name": "Jane Doe","age": 30,},{"name": "Jane Doe","age": 33,},],dataType: "json",type: "DELETE",url: "https://demos.telerik.com/kendo-ui/service/products",},parameterMap: function(data, type) { if (type == 'read') { /* send take as '$top' and skip as '$skip' */ return { $top: data.take, $skip: data.skip } } },create: { cache: false,contentType: "application/json",data: [{"name": "Jane Doe","age": 30,},{"name": "Jane Doe","age": 33,},],dataType: "json",type: "GET",url: "https://demos.telerik.com/kendo-ui/service/products/create",},push: function(callbacks) { hub.on('create', function(result) { console.log('push create'); callbacks.pushCreate(result); }); hub.on('update', function(result) { console.log('push update'); callbacks.pushUpdate(result); }); hub.on('destroy', function(result) { console.log('push destroy'); callbacks.pushDestroy(result); }); },read: { cache: false,contentType: "application/json",data: {skip: 0,take: 2,},dataType: "json",type: "GET",url: "https://demos.telerik.com/kendo-ui/service/products",},submit: function(e) { var data = e.data; console.log(data); /* send batch update to desired URL, then notify success/error */ e.success(data.updated,'update'); e.success(data.created,'create'); e.success(data.destroyed,'destroy'); e.error(null, 'customerror', 'custom error'); },update: { cache: false,contentType: "application/json",data: {skip: 0,take: 2,},dataType: "json",type: "PUT",url: "https://demos.telerik.com/kendo-ui/service/products",},},type: "odata",});
}

func ExampleKendoDataSource_ToJavaScript_Filter_Object() {
	javaScript := KendoDataSource{
		//VarName: "dataSource",
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
		Filter: KendoComplexFilter{
			Field:    "category",
			Operator: OPERATOR_EQUAL_TO,
			Value:    "Beverages",
		},
		Group: []KendoGroup{
			{
				Field: "category",
			},
			{
				Field: "subcategory",
			},
		},
		InPlaceSort: FALSE,
		OfflineStorage: JavaScript{
			Code: `{ getItem: function() { return JSON.parse(sessionStorage.getItem("products-key")); }, setItem: function(item) { sessionStorage.setItem("products-key", JSON.stringify(item)); } }`,
		},
		Page:     1,
		PageSize: 25,
		Schema: KendoSchema{
			Aggregates: "aggregates",
			Data:       "statuses",
			Errors:     "errors",
			Groups:     "groups",
			Model: KendoDataModel{
				Id: "personId",
			},
			Parser: JavaScript{
				Code: "function(response) { var products = []; for (var i = 0; i < response.length; i++) { var product = { id: response[i].ProductID, name: response[i].ProductName }; products.push(product); } return products; }",
			},
			Total: "total",
			Type:  KENDO_TYPE_DATA_JSON,
		},
		ServerAggregates: FALSE,
		ServerFiltering:  FALSE,
		ServerGrouping:   FALSE,
		ServerPaging:     FALSE,
		ServerSorting:    FALSE,
		Sort: KendoSort{
			Dir:   DIRECTION_ASC,
			Field: "age",
			Compare: JavaScript{
				Code: "function(a, b) { return numbers[a.item] - numbers[b.item]; }",
			},
		},
		Transport: KendoTransport{
			Create: KendoCreate{
				Cache:       FALSE,
				ContentType: "application/json",
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
				DataType: KENDO_TYPE_DATA_JSON_JSON,
				Type:     HTML_METHOD_GET,
				Url:      "https://demos.telerik.com/kendo-ui/service/products/create",
			},
			Destroy: KendoDestroy{
				Cache:       FALSE,
				ContentType: "application/json",
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
				DataType: KENDO_TYPE_DATA_JSON_JSON,
				Type:     HTML_METHOD_DELETE,
				Url:      "https://demos.telerik.com/kendo-ui/service/products",
			},
			ParameterMap: JavaScript{
				Code: "function(data, type) { if (type == 'read') { /* send take as '$top' and skip as '$skip' */ return { $top: data.take, $skip: data.skip } } }",
			},
			Push: JavaScript{
				Code: "function(callbacks) { hub.on('create', function(result) { console.log('push create'); callbacks.pushCreate(result); }); hub.on('update', function(result) { console.log('push update'); callbacks.pushUpdate(result); }); hub.on('destroy', function(result) { console.log('push destroy'); callbacks.pushDestroy(result); }); }",
			},
			Read: KendoRead{
				Cache:       FALSE,
				ContentType: "application/json",
				Data: map[string]interface{}{
					"skip": 0,
					"take": 2,
				},
				DataType: KENDO_TYPE_DATA_JSON_JSON,
				Type:     HTML_METHOD_GET,
				Url:      "https://demos.telerik.com/kendo-ui/service/products",
			},
			Submit: JavaScript{
				Code: "function(e) { var data = e.data; console.log(data); /* send batch update to desired URL, then notify success/error */ e.success(data.updated,'update'); e.success(data.created,'create'); e.success(data.destroyed,'destroy'); e.error(null, 'customerror', 'custom error'); }",
			},
			Update: KendoUpdate{
				Cache:       FALSE,
				ContentType: "application/json",
				Data: map[string]interface{}{
					"skip": 0,
					"take": 2,
				},
				DataType: KENDO_TYPE_DATA_JSON_JSON,
				Type:     HTML_METHOD_PUT,
				Url:      "https://demos.telerik.com/kendo-ui/service/products",
			},
		},
		Type: KENDO_TYPE_ODATA,
	}

	fmt.Printf("%s", javaScript.ToJavaScript())

	// Output:
	// var dataSource = new kendo.data.DataSource({aggregate: [{aggregate: "sum",field: "age",},{aggregate: "min",field: "age",},{aggregate: "max",field: "age",},],autoSync: true,batch: true,data: [{"name": "Jane Doe","age": 30,},{"name": "Jane Doe","age": 33,},],filter: { field: "category",operator: "eq",value: "Beverages",},group: [{field: "category",},{field: "subcategory",},],inPlaceSort: false,offlineStorage: { getItem: function() { return JSON.parse(sessionStorage.getItem("products-key")); }, setItem: function(item) { sessionStorage.setItem("products-key", JSON.stringify(item)); } },page: 1,pageSize: 25,schema: {aggregates: "aggregates",data: "statuses",errors: "errors",groups: "groups",model: {id: "personId",},groups: function(response) { var products = []; for (var i = 0; i < response.length; i++) { var product = { id: response[i].ProductID, name: response[i].ProductName }; products.push(product); } return products; },total: "total",type: "json",},serverAggregates: false,serverFiltering: false,serverGrouping: false,serverPaging: false,serverSorting: false,sort: { dir: "asc",field: "age",compare: function(a, b) { return numbers[a.item] - numbers[b.item]; },},transport: { destroy: { cache: false,contentType: "application/json",contentType: [{"name": "Jane Doe","age": 30,},{"age": 33,"name": "Jane Doe",},],dataType: "json",type: "DELETE",url: "https://demos.telerik.com/kendo-ui/service/products",},parameterMap: function(data, type) { if (type == 'read') { /* send take as '$top' and skip as '$skip' */ return { $top: data.take, $skip: data.skip } } },create: { cache: false,contentType: "application/json",data: [{"name": "Jane Doe","age": 30,},{"name": "Jane Doe","age": 33,},],dataType: "json",type: "GET",url: "https://demos.telerik.com/kendo-ui/service/products/create",},push: function(callbacks) { hub.on('create', function(result) { console.log('push create'); callbacks.pushCreate(result); }); hub.on('update', function(result) { console.log('push update'); callbacks.pushUpdate(result); }); hub.on('destroy', function(result) { console.log('push destroy'); callbacks.pushDestroy(result); }); },read: { cache: false,contentType: "application/json",data: {take: 2,skip: 0,},dataType: "json",type: "GET",url: "https://demos.telerik.com/kendo-ui/service/products",},submit: function(e) { var data = e.data; console.log(data); /* send batch update to desired URL, then notify success/error */ e.success(data.updated,'update'); e.success(data.created,'create'); e.success(data.destroyed,'destroy'); e.error(null, 'customerror', 'custom error'); },update: { cache: false,contentType: "application/json",data: {skip: 0,take: 2,},dataType: "json",type: "PUT",url: "https://demos.telerik.com/kendo-ui/service/products",},},type: "odata",});
}
