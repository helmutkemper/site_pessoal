package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridColumnMenu struct {
	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columnmenu.columns
	//
	// If set to true the column menu would allow the user to select (show and hide) grid columns. By default the column
	// menu allows column selection. (default: true)
	//
	//    Example - disable column selection
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" }
	//        ],
	//        columnMenu: {
	//          columns: false
	//        },
	//        sortable: true,
	//        dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	Columns Boolean `jsObject:"columns"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columnmenu.filterable
	//
	// If set to true the column menu would allow the user to filter the grid. By default the column menu allows the user
	// to filter if filtering is enabled via the filterable. (default: true)
	//
	//    Example - disable column menu filtering
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" }
	//        ],
	//        columnMenu: {
	//          filterable: false
	//        },
	//        filterable: true,
	//        dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	Filterable Boolean `jsObject:"filterable"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columnmenu.sortable
	//
	// If set to true the column menu would allow the user to sort the grid by the column field. By default the column
	// menu allows the user to sort if sorting is enabled via the sortable option. (default: true)
	//
	// Important:
	//
	// If this option is set to false the user could still sort by clicking the column header cell.
	//
	//    Example - disable column menu sorting
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" }
	//        ],
	//        columnMenu: {
	//          sortable: false
	//        },
	//        sortable: true,
	//        dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	Sortable Boolean `jsObject:"sortable"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columnmenu.messages
	//
	// The text messages displayed in the column menu. Use it to customize or localize the column menu messages.
	//
	//    Example - customize column menu messages
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" }
	//        ],
	//        columnMenu: {
	//          messages: {
	//            columns: "Choose columns",
	//            filter: "Apply filter",
	//            sortAscending: "Sort (asc)",
	//            sortDescending: "Sort (desc)"
	//          }
	//        },
	//        sortable: true,
	//        filterable: true,
	//        dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	Messages KendoGridColumnMenuMessages `jsObject:"messages"`

	*ToJavaScriptConverter
}

func (el *KendoGridColumnMenu) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridColumnMenu.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
