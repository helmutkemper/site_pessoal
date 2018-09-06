package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoSort struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/sort#sort.dir

	  The sort order (direction).

	  The supported values are:
	  "asc" (ascending order)
	  "desc" (descending order)
	  Example - specify the sort order (direction)
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Jane Doe", age: 30 },
	      { name: "John Doe", age: 33 }
	    ],
	    // order by "age" in descending order
	    sort: { field: "age", dir: "desc" }
	  });
	  dataSource.fetch(function(){
	    var data = dataSource.view();
	    console.log(data[0].age); // displays "33"
	  });
	  </script>
	*/
	Dir KendoDirection `jsObject:"dir"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/sort#sort.field

	  The field by which the data items are sorted.

	  Example - specify the sort field
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    data: [
	      { name: "Jane Doe", age: 30 },
	      { name: "John Doe", age: 33 }
	    ],
	    // order by "age" in descending order
	    sort: { field: "age", dir: "desc" }
	  });
	  dataSource.fetch(function(){
	    var data = dataSource.view();
	    console.log(data[0].age); // displays "33"
	  });
	  </script>
	*/
	Field string `jsObject:"field"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/sort#sort.compare

	  Function which can be used for custom comparing of the DataSource items.

	  Example - use a custom compare function to compare items in the DataSource
	  <div id="grid"></div>
	  <script>
	    var numbers = {
	      "one"  : 1,
	      "two"  : 2,
	      "three": 3,
	      "four" : 4,
	    };

	    var dataSource = new kendo.data.DataSource({
	      data: [
	        { id: 1, item: "two" },
	        { id: 2, item: "one" },
	        { id: 3, item: "three" },
	        { id: 4, item: "four" }
	      ],
	      sort: { field: "item", dir: "asc", compare: function(a, b) {
	        return numbers[a.item] - numbers[b.item];
	      }
	            }
	    });

	    $("#grid").kendoGrid({
	      dataSource: dataSource,
	      sortable: true,
	      columns: [{
	        field: "item",
	        sortable: {
	          compare: function(a, b) {
	            return numbers[a.item] - numbers[b.item];
	          }
	        }
	      }]
	    });
	  </script>
	*/
	Compare JavaScript `jsObject:"compare"`

	*ToJavaScriptConverter
}

func (el *KendoSort) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoSort.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
