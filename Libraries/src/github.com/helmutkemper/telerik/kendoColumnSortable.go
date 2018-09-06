package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoColumnSortable struct {
	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.sortable#columns.sortable.allowUnsort
	//
	// If set to true the user can get the grid in unsorted state by clicking the sorted column header. If set to false
	// the user will not be able to unsort the column once sorted. (default: true)
	//
	//    Example - disable unsorting
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { sortable: { allowUnsort: false }, field: "id" },
	//          { field: "name" }
	//        ],
	//        sortable: true,
	//        dataSource: [ { id: 1, name: "Jane Doe" }, { id: 2, name: "John Doe" } ]
	//      });
	//    </script>
	AllowUnSort Boolean `jsObject:"allowUnsort"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.sortable#columns.sortable.compare
	//
	// A JavaScript function which is used to compare the values. It has the same signature as the compare function
	// accepted by Array.sort.
	//
	// The basic function implementation is as follows (pseudo-code):
	//
	//    function compare(a, b, descending) {
	//      if (a is less than b by some ordering criterion) {
	//        return -1;
	//      }
	//
	//      if (a is greater than b by the ordering criterion) {
	//        return 1;
	//      }
	//
	//      // a must be equal to b
	//      return 0;
	//    }
	//
	// One notable exception is that we also supply a third parameter that indicates the sort direction (true for
	// descending). See How-to: Stable Sort in Chrome for more details on how this can be useful.
	//
	//    Example - define custom compare function
	//
	//    <div id="grid"></div>
	//    <script>
	//      var numbers = {
	//        "one"  : 1,
	//        "two"  : 2,
	//        "three": 3
	//      };
	//
	//      var dataSource = new kendo.data.DataSource({
	//        data: [
	//          { id: 1, item: "two" },
	//          { id: 2, item: "one" },
	//          { id: 3, item: "three" }
	//        ]
	//      });
	//
	//      $("#grid").kendoGrid({
	//        dataSource: dataSource,
	//        sortable: true,
	//        columns: [{
	//          field: "item",
	//          sortable: {
	//            compare: function(a, b) {
	//              return numbers[a.item] - numbers[b.item];
	//            }
	//          }
	//        }]
	//      });
	//    </script>
	Compare *JavaScript `jsObject:"compare"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.sortable#columns.sortable.initialDirection
	//
	// Determines the inital (from un-sorted to sorted state) sort direction. The supported values are asc and desc.
	// (default: asc)
	//
	//    Example - disable sorting
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          {
	//            field: "id",
	//            sortable: {
	//              initialDirection: "desc"
	//            }
	//          },
	//          { field: "name" }
	//        ],
	//        sortable: true,
	//        dataSource: [ { id: 1, name: "Jane Doe" }, { id: 2, name: "John Doe" } ]
	//      });
	//    </script>
	InitialDirection KendoDirection `jsObject:"initialDirection"`

	*ToJavaScriptConverter
}

func (el *KendoColumnSortable) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoColumnSortable.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
