package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoComplexFilter struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/filter#filter.field

	  The data item field to which the filter operator is applied.

	  Example - set the filter field
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
	*/
	Field string `jsObject:"field"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/filter#filter.filters

	  The nested filter expressions. Supports the same options as filter. Filters can be nested indefinitely.

	  Example - nested filters
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
	Filters []KendoFilters `jsObject:"field"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/filter#filter.logic

	  The logical operation to use when the filter.filters option is set.
	  The supported values are: "and" or "or"

	  Example - set the filter logic
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
	Logic KendoLogic `jsObject:"logic"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/filter#filter.operator

	  The filter operator (comparison).
	  The supported operators are:
	  "eq" (equal to)
	  "neq" (not equal to)
	  "isnull" (is equal to null)
	  "isnotnull" (is not equal to null)
	  "lt" (less than)
	  "lte" (less than or equal to)
	  "gt" (greater than)
	  "gte" (greater than or equal to)
	  "startswith"
	  "endswith"
	  "contains"
	  "doesnotcontain"
	  "isempty"
	  "isnotempty"
	  The last five are supported only for string fields.

	  Example - set the filter operator
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
	*/
	Operator KendoOperator `jsObject:"operator"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/filter#filter.value

	  The value to which the field is compared. The value has to be of the same type as the field.
	  By design, the "\n" is removed from the filter before the filtering is performed. That is why an "\n" identifier from the filter will not match data items whose corresponding fields contain new lines.

	  Example - specify the filter value
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Jane Doe", birthday: new Date(1983, 1, 1) },
	      { name: "John Doe", birthday: new Date(1980, 1, 1)}
	    ],
	    filter: { field: "birthday", operator: "gt", value: new Date(1980, 1, 1) }
	  });
	  dataSource.fetch(function(){
	    var view = dataSource.view();
	    console.log(view.length); // displays "1"
	    console.log(view[0].name); // displays "Jane Doe"
	  });
	  </script>
	*/
	Value interface{} `jsObject:"value" jsType:"int,int64,float32,float64,string,Boolean"`

	*ToJavaScriptConverter
}

func (el *KendoComplexFilter) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoComplexFilter.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
