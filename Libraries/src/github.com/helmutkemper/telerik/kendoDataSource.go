package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoDataSource struct {
	//VarName                                 string              `jsObject:"-"`

	/*
	  @sse https://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate

	  The aggregates which are calculated when the data source populates with data.

	  The supported aggregates are:

	  > "average" - Only for Number.
	  > "count" - string, Number and Date.
	  > "max" - Number and Date.
	  > "min" - Number and Date.
	  > "sum" - Only for Number.

	  The data source calculates aggregates client-side unless the serverAggregates option is set to true.

	  EXAMPLE - SPECIFY AGGREGATES
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Jane Doe", age: 30 },
	      { name: "John Doe", age: 33 }
	    ],
	    aggregate: [
	      { field: "age", aggregate: "sum" },
	      { field: "age", aggregate: "min" },
	      { field: "age", aggregate: "max" }
	    ]
	  });
	  dataSource.fetch(function(){
	    var results = dataSource.aggregates().age;
	    console.log(results.sum, results.min, results.max); // displays "63 30 33"
	  });
	  </script>
	*/

	Aggregate []KendoAggregates `jsObject:"aggregate"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/autosync#autoSync

	  If set to true the data source would automatically save any changed data items by calling the sync method. By default, changes are not automatically saved. (default: false)

	  Example - enable auto sync
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    autoSync: true,
	    transport: {
	      read:  {
	        url: "https://demos.telerik.com/kendo-ui/service/products",
	        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      },
	      update: {
	        url: "https://demos.telerik.com/kendo-ui/service/products/update",
	        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      }
	    },
	    schema: {
	      model: { id: "ProductID" }
	    }
	  });
	  dataSource.fetch(function() {
	    var product = dataSource.at(0);
	    product.set("UnitPrice", 20); // auto-syncs and makes request to https://demos.telerik.com/kendo-ui/service/products/update
	  });
	  </script>
	*/
	AutoSync Boolean `jsObject:"autoSync"`

	/*
	  @sse https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/batch#batch

	  If set to <b>true</b>, the data source will batch CRUD operation requests. For example, updating two data items would cause one HTTP request instead of two. By default, the data source makes a HTTP request for every CRUD operation. (default: false)
	  The changed data items are sent as <b>models</b> by default. This can be changed via the parameterMap option.

	  Example - enable the batch mode
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    batch: true,
	    transport: {
	      read:  {
	        url: "https://demos.telerik.com/kendo-ui/service/products",
	        dataType: "jsonp" //"jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      },
	      update: {
	        url: "https://demos.telerik.com/kendo-ui/service/products/update",
	        dataType: "jsonp" //"jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      }
	    },
	    schema: {
	      model: { id: "ProductID" }
	    }
	  });
	  dataSource.fetch(function() {
	    var product = dataSource.at(0);
	    product.set("UnitPrice", 20);
	    var anotherProduct = dataSource.at(1);
	    anotherProduct.set("UnitPrice", 20);
	    dataSource.sync(); // causes only one request to "https://demos.telerik.com/kendo-ui/service/products/update"
	  });
	  </script>
	*/
	Batch Boolean `jsObject:"batch"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/data#data

	  The array of data items which the data source contains. The data source will wrap those items as kendo.data.ObservableObject or kendo.data.Model (if schema.model is set).

	  Example - set the data items of a data source
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Jane Doe", age: 30 },
	      { name: "John Doe", age: 33 }
	    ]
	  });
	  dataSource.fetch(function(){
	    var janeDoe = dataSource.at(0);
	    console.log(janeDoe.name); // displays "Jane Doe"
	  });
	  </script>
	*/
	Data interface{} `jsObject:"data" jsType:"string,*map[string]interface {},[]string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/filter#filter

	  The filters which are applied over the data items. By default, no filter is applied.
	  The data source filters the data items client-side unless the serverFiltering option is set to <b>true</b>.

	  Example - set a single filter
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Jane Doe" },
	      { name: "John Doe" }
	    ],
	    filter: { field: "name", operator: "startswith", value: "Jane" }
	  });
	  dataSource.fetch(function(){
	    var view = dataSource.view();
	    console.log(view.length); // displays "1"
	    console.log(view[0].name); // displays "Jane Doe"
	  });
	  </script>

	  Example - set filter as conjunction (and)
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Tea", category: "Beverages" },
	      { name: "Coffee", category: "Beverages" },
	      { name: "Ham", category: "Food" }
	    ],
	    filter: [
	      // leave data items which are "Beverage" and not "Coffee"
	      { field: "category", operator: "eq", value: "Beverages" },
	      { field: "name", operator: "neq", value: "Coffee" }
	    ]
	  });
	  dataSource.fetch(function(){
	    var view = dataSource.view();
	    console.log(view.length); // displays "1"
	    console.log(view[0].name); // displays "Tea"
	  });
	  </script>

	  Example - set filter as disjunction (or)
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Tea", category: "Beverages" },
	      { name: "Coffee", category: "Beverages" },
	      { name: "Ham", category: "Food" }
	    ],
	    filter: {
	      // leave data items which are "Food" or "Tea"
	      logic: "or",
	      filters: [
	        { field: "category", operator: "eq", value: "Food" },
	        { field: "name", operator: "eq", value: "Tea" }
	      ]
	    }
	  });
	  dataSource.fetch(function(){
	    var view = dataSource.view();
	    console.log(view.length); // displays "2"
	    console.log(view[0].name); // displays "Tea"
	    console.log(view[1].name); // displays "Ham"
	  });
	  </script>
	*/
	Filter interface{} `jsObject:"filter" jsType:"*KendoComplexFilter,*[]KendoComplexFilter"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/group#group

	  The grouping configuration of the data source. If set, the data items will be grouped when the data source is populated. By default, grouping is not applied.
	  The data source groups the data items client-side unless the serverGrouping option is set to true.

	  Example - set a group as an object
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Tea", category: "Beverages" },
	      { name: "Coffee", category: "Beverages" },
	      { name: "Ham", category: "Food" }
	    ],
	    // group by the "category" field
	    group: { field: "category" }
	  });
	  dataSource.fetch(function(){
	    var view = dataSource.view();
	    console.log(view.length); // displays "2"
	    var beverages = view[0];
	    console.log(beverages.value); // displays "Beverages"
	    console.log(beverages.items[0].name); // displays "Tea"
	    console.log(beverages.items[1].name); // displays "Coffee"
	    var food = view[1];
	    console.log(food.value); // displays "Food"
	    console.log(food.items[0].name); // displays "Ham"
	  });
	  </script>

	  Example - set a group as an array (subgroups)
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Pork", category: "Food", subcategory: "Meat" },
	      { name: "Pepper", category: "Food", subcategory: "Vegetables" },
	      { name: "Beef", category: "Food", subcategory: "Meat" }
	    ],
	    group: [
	      // group by "category" and then by "subcategory"
	      { field: "category" },
	      { field: "subcategory" },
	    ]
	  });
	  dataSource.fetch(function(){
	    var view = dataSource.view();
	    console.log(view.length); // displays "1"
	    var food = view[0];
	    console.log(food.value); // displays "Food"
	    var meat = food.items[0];
	    console.log(meat.value); // displays "Meat"
	    console.log(meat.items.length); // displays "2"
	    console.log(meat.items[0].name); // displays "Pork"
	    console.log(meat.items[1].name); // displays "Beef"
	    var vegetables = food.items[1];
	    console.log(vegetables.value); // displays "Vegetables"
	    console.log(vegetables.items.length); // displays "1"
	    console.log(vegetables.items[0].name); // displays "Pepper"
	  });
	  </script>
	*/
	Group interface{} `jsObject:"group" jsType:"*KendoGroup,*[]KendoGroup"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/inplacesort#inPlaceSort

	  If set to <b>true</b> the original <b>Array</b> used as data will be sorted when sorting operation is performed. This setting supported only with local data, bound to a JavaScript array via the data option. (default: false)
	*/
	InPlaceSort Boolean `jsObject:"inPlaceSort"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/offlinestorage#offlineStorage

	  Example - set offline storage key
	  <script>
	  var dataSource = new kendo.data.DataSource({
	      offlineStorage: "products-offline",
	      transport: {
	          read: {
	              url: "https://demos.telerik.com/kendo-ui/service/products",
	              type: "jsonp"
	          }
	      }
	  });
	  </script>
	*/
	OfflineStorage interface{} `jsObject:"offlineStorage" jsType:"*OfflineStorage,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/page#page

	  The page of data which the data source will return when the view method is invoked or request from the remote service.
	  The data source will page the data items client-side unless the serverPaging option is set to <b>true</b>.

	  Example - set the current page
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Tea", category: "Beverages" },
	      { name: "Coffee", category: "Beverages" },
	      { name: "Ham", category: "Food" }
	    ],
	    // set the second page as the current page
	    page: 2,
	    pageSize: 2
	  });
	  dataSource.fetch(function(){
	    var view = dataSource.view();
	    console.log(view.length); // displays "1"
	    console.log(view[0].name); // displays "Ham"
	  });
	  </script>
	*/
	Page int `jsObject:"page"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/pagesize#pageSize

	  The number of data items per page. The property has no default value. <b>That is why to use paging, make sure some pageSize value is set</b>.
	  The data source will page the data items client-side unless the serverPaging option is set to <b>true</b>.

	  Example - set the page size
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Tea", category: "Beverages" },
	      { name: "Coffee", category: "Beverages" },
	      { name: "Ham", category: "Food" }
	    ],
	    page: 1,
	    // a page of data contains two data items
	    pageSize: 2
	  });
	  dataSource.fetch(function(){
	    var view = dataSource.view();
	    console.log(view.length); // displays "2"
	    console.log(view[0].name); // displays "Tea"
	    console.log(view[1].name); // displays "Coffee"
	  });
	  </script>
	*/
	PageSize int `jsObject:"pageSize"`

	/*
	  https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/schema#schema

	  The configuration used to parse the remote service response.

	  Example - specify the schema of the remote service
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      read: {
	        url: "https://demos.telerik.com/kendo-ui/service/twitter/search",
	        dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	        data: { q: "html5" } // search for tweets that contain "html5"
	      }
	    },
	    schema: {
	      data: function(response) {
	        return response.statuses; // twitter's response is { "statuses": [ / * results * / ] }
	      }
	    }
	  });
	  dataSource.fetch(function(){
	    var data = this.data();
	    console.log(data.length);
	  });
	  </script>
	*/
	Schema KendoSchema `jsObject:"schema"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/serveraggregates#serverAggregates

	  If set to true, the data source will leave the aggregate calculation to the remote service. By default, the data source calculates aggregates client-side.
	  <b>Configure schema.aggregates if you set serverAggregates to true</b>.
	  For more information and tips about client and server data operations, refer to the introductory article on the DataSource.
	  https://docs.telerik.com/kendo-ui/framework/datasource/overview#mixed-data-operations-mode

	  Example - enable server aggregates
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      // transport configuration
	    },
	    serverAggregates: true,
	    aggregate: [
	      { field: "age", aggregate: "sum" }
	    ],
	    schema: {
	      aggregates: "aggregates" // aggregate results are returned in the "aggregates" field of the response
	    }
	  });
	  </script>
	*/
	ServerAggregates Boolean `jsObject:"serverAggregates"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/serverfiltering#serverFiltering

	  If set to true, the data source will leave the filtering implementation to the remote service. By default, the data source performs filtering client-side.
	  By default, the filter is sent to the server following jQuery's conventions.

	  For example, the filter { logic: "and", filters: [ { field: "name", operator: "startswith", value: "Jane" } ] } is sent as:

	  filter[logic]: and
	  filter[filters][0][field]: name
	  filter[filters][0][operator]: startswith
	  filter[filters][0][value]: Jane
	  Use the parameterMap option to send the filter option in a different format.

	  For more information and tips about client and server data operations, refer to the introductory article on the DataSource.
	  https://docs.telerik.com/kendo-ui/framework/datasource/overview#mixed-data-operations-mode

	  Example - enable server filtering
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      // transport configuration
	    },
	    serverFiltering: true,
	    filter: { logic: "and", filters: [ { field: "name", operator: "startswith", value: "Jane" } ] }
	  });
	  </script>
	*/
	ServerFiltering Boolean `jsObject:"serverFiltering"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/servergrouping#serverGrouping

	  If set to true, the data source will leave the grouping implementation to the remote service. By default, the data source performs grouping client-side. (default: false)
	  By default, the group is sent to the server following jQuery's conventions.

	  For example, the group { field: "category", dir: "desc" } is sent as:

	  group[0][field]: category
	  group[0][dir]: desc
	  Use the parameterMap option to send the group option in a different format.

	  For more information and tips about client and server data operations, refer to the introductory article on the DataSource.
	  https://docs.telerik.com/kendo-ui/framework/datasource/overview#mixed-data-operations-mode

	  Example - enable server grouping
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      // transport configuration
	    },
	    serverGrouping: true,
	    group: { field: "category", dir: "desc" }
	  });
	  </script>
	*/
	ServerGrouping Boolean `jsObject:"serverGrouping"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/serverpaging#serverPaging

	  If set to true, the data source will leave the data item paging implementation to the remote service. By default, the data source performs paging client-side. (default: false)
	  Configure schema.total if you set serverPaging to true. In addition, pageSize should be set no matter if paging is performed client-side or server-side.

	  The following options are sent to the server when server paging is enabled:

	  page - the page of data item to return (1 means the first page).
	  pageSize - the number of items to return.
	  skip - how many data items to skip.
	  take - the number of data items to return (the same as pageSize).
	  Use the parameterMap option to send the paging options in a different format.

	  For more information and tips about client and server data operations, refer to the introductory article on the DataSource.
	  https://docs.telerik.com/kendo-ui/framework/datasource/overview#mixed-data-operations-mode

	  Example - enable server paging
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      // transport configuration
	    },
	    serverPaging: true,
	    schema: {
	      total: "total" // total is returned in the "total" field of the response
	    }
	  });
	  </script>
	*/
	ServerPaging Boolean `jsObject:"serverPaging"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/serversorting#serverSorting

	  If set to true, the data source will leave the data item sorting implementation to the remote service. By default, the data source performs sorting client-side. (default: false)

	  By default, the sort is sent to the server following jQuery's conventions.

	  For example, the sort { field: "age", dir: "desc" } is sent as:

	  sort[0][field]: age
	  sort[0][dir]: desc
	  Use the parameterMap option to send the paging options in a different format.

	  For more information and tips about client and server data operations, refer to the introductory article on the DataSource.
	  https://docs.telerik.com/kendo-ui/framework/datasource/overview#mixed-data-operations-mode

	  Example - enable server sorting
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      / * transport configuration * /
	    },
	    serverSorting: true,
	    sort: { field: "age", dir: "desc" }
	  });
	  </script>
	*/
	ServerSorting Boolean `jsObject:"serverSorting"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/sort#sort

	  The sort order which will be applied over the data items. By default the data items are not sorted.
	  The data source sorts the data items client-side unless the serverSorting option is set to true.

	  Example - sort the data items
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Jane Doe", age: 30 },
	      { name: "John Doe", age: 33 }
	    ],
	    sort: { field: "age", dir: "desc" }
	  });
	  dataSource.fetch(function(){
	    var data = dataSource.view();
	    console.log(data[0].age); // displays "33"
	  });
	  </script>

	  Example - sort the data items by multiple fields
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Tea", category: "Beverages" },
	      { name: "Coffee", category: "Beverages" },
	      { name: "Ham", category: "Food" }
	    ],
	    sort: [
	      // sort by "category" in descending order and then by "name" in ascending order
	      { field: "category", dir: "desc" },
	      { field: "name", dir: "asc" }
	    ]
	  });
	  dataSource.fetch(function(){
	    var data = dataSource.view();
	    console.log(data[1].name); // displays "Coffee"
	  });
	  </script>
	*/
	Sort KendoSort `jsObject:"sort"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/transport#transport

	  The configuration used to load and save the data items. A data source is remote or local based on the way of it retrieves data items.

	  Remote data sources load and save data items from and to a remote end-point (also known as remote service or server). The transport option describes the remote service configuration - URL, HTTP verb, HTTP headers, and others. The transport option can also be used to implement custom data loading and saving.

	  Local data sources are bound to a JavaScript array via the data option.

	  Example - specify the remote service configuration
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      read: {
	        url: "https://demos.telerik.com/kendo-ui/service/products",
	        dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
	      }
	    }
	  });
	  dataSource.fetch(function() {
	    var products = dataSource.data();
	    console.log(products[0].ProductName); // displays "Chai"
	  });
	  </script>
	*/
	Transport KendoTransport `jsObject:"transport"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/type#type

	  If set, the data source will use a predefined transport and/or schema.

	  The supported values are:

	  "odata" which supports the OData v.2 protocol - http://www.odata.org/
	  "odata-v4" which partially supports odata version 4 - https://github.com/telerik/ui-for-aspnet-mvc-examples/tree/master/grid/odata-v4-web-api-binding
	  "signalr"

	  Example - enable OData support
	  <script>
	  var dataSource= new kendo.data.DataSource({
	    type: "odata",
	    transport: {
	      read: "https://demos.telerik.com/kendo-ui/service/Northwind.svc/Orders"
	    },
	    pageSize: 20,
	    serverPaging: true
	  });
	  dataSource.fetch(function() {
	    console.log(dataSource.view().length); // displays "20"
	  });
	  </script>
	*/
	Type KendoType `jsObject:"type"`

	*ToJavaScriptConverter
}

func (el *KendoDataSource) ToJavaScript() []byte {
	var ret bytes.Buffer

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoDataSource.Error: %v", err.Error())
		return []byte{}
	}

	//if el.VarName == "" {
	ret.Write([]byte(` new kendo.data.DataSource({`))
	ret.Write(data)
	ret.Write([]byte(`}),`))
	//} else {
	//  ret.Write( []byte(`` + el.VarName + ` = new kendo.data.DataSource({`) )
	//  ret.Write( data )
	//  ret.Write( []byte(`});`) )
	//}

	return ret.Bytes()
}
