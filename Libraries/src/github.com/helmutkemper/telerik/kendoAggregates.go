package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoAggregates struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/aggregate#aggregate

	  The aggregates which are calculated when the data source populates with data.

	  The supported aggregates are:

	  <b>"average"</b> - Only for Number.
	  <b>"count"</b> - String, Number and Date.
	  <b>"max"</b> - Number and Date.
	  <b>"min"</b> - Number and Date.
	  <b>"sum"</b> - Only for Number.
	  The data source calculates aggregates client-side unless the serverAggregates option is set to <b>true</b>.

	  Example - specify aggregates
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

	  Example - specify aggregate function
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Jane Doe", age: 30 },
	      { name: "John Doe", age: 33 }
	    ],
	    aggregate: [
	      { field: "age", aggregate: "sum" }
	    ]
	  });
	  dataSource.fetch(function(){
	    var results = dataSource.aggregates().age;
	    console.log(results.sum); // displays "63"
	  });
	  </script>
	*/
	Aggregate KendoAggregate `jsObject:"aggregate"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/aggregate#aggregate.field

	  The data item field which will be used to calculate the aggregates.

	  Example - specify aggregate field
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Jane Doe", age: 30 },
	      { name: "John Doe", age: 33 }
	    ],
	    aggregate: [
	      { field: "age", aggregate: "sum" }
	    ]
	  });
	  dataSource.fetch(function(){
	    var results = dataSource.aggregates().age;
	    console.log(results.sum); // displays "63"
	  });
	  </script>
	*/
	Field string `jsObject:"field"`

	*ToJavaScriptConverter
}

func (el *KendoAggregates) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoAggregates.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
