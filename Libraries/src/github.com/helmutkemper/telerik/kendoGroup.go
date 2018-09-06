package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGroup struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/group#group.aggregates

	  The aggregates which are calculated during grouping.
	  The supported aggregates are:
	  "average"
	  "count"
	  "max"
	  "min"
	  "sum"

	  Example - set group aggregates
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Tea", category: "Beverages", price: 1 },
	      { name: "Coffee", category: "Beverages", price: 2 },
	      { name: "Ham", category: "Food", price: 3 },
	    ],
	    group: {
	      field: "category",
	      aggregates: [
	        { field: "price", aggregate: "max" },
	        { field: "price", aggregate: "min" }
	      ]
	    }
	  });
	  dataSource.fetch(function(){
	    var view = dataSource.view();
	    var beverages = view[0];
	    console.log(beverages.aggregates.price.max); // displays "2"
	    console.log(beverages.aggregates.price.min); // displays "1"
	    var food = view[1];
	    console.log(food.aggregates.price.max); // displays "3"
	    console.log(food.aggregates.price.min); // displays "3"
	  });
	  </script>
	*/
	Aggregates []KendoAggregates `jsObject:"aggregates"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/group#group.dir

	  The sort order of the group.
	  The supported values are:
	  "asc" (ascending order)
	  "desc" (descending order)
	  The default sort order is ascending.

	  Example - sort the groups in descending order
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Tea", category: "Beverages"},
	      { name: "Ham", category: "Food"},
	    ],
	    // group by "category" in descending order
	    group: { field: "category", dir: "desc" }
	  });
	  dataSource.fetch(function(){
	    var view = dataSource.view();
	    var food = view[0];
	    console.log(food.value); // displays "Food"
	    var beverages = view[1];
	    console.log(beverages.value); // displays "Beverages"
	  });
	  </script>
	*/
	Direction KendoDirection `jsObject:"dir"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/group#group.field

	  The data item field to group by.

	  Example - set the field
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
	    var beverages = view[0];
	    console.log(beverages.items[0].name); // displays "Tea"
	    console.log(beverages.items[1].name); // displays "Coffee"
	    var food = view[1];
	    console.log(food.items[0].name); // displays "Ham"
	  });
	  </script>
	*/
	Field string `jsObject:"field"`

	*ToJavaScriptConverter
}

func (el *KendoGroup) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGroup.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
