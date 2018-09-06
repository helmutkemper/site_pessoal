package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridColumnMenuMessages struct {
	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columnmenu.messages#columnMenu.messages.columns
	//
	// The text message displayed for the column selection menu item. (default: "Columns")
	//
	//    Example - set the column selection message
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
	//            columns: "Choose columns"
	//          }
	//        },
	//        dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	Columns string `jsObject:"columns"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columnmenu.messages#columnMenu.messages.filter
	//
	// The text message displayed for the filter menu item. (default: "Filter")
	//
	//    Example - set the filter message
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
	//            filter: "Apply filter",
	//          }
	//        },
	//        filterable: true,
	//        dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	Filter string `jsObject:"filter"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columnmenu.messages#columnMenu.messages.sortAscending
	//
	// The text message displayed for the menu item which performs ascending sort. (default: "Sort Ascending")
	//
	//    Example - set the sort ascending message
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
	//            sortAscending: "Sort (asc)",
	//          }
	//        },
	//        sortable: true,
	//        dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	SortAscending string `jsObject:"sortAscending"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columnmenu.messages#columnMenu.messages.sortDescending
	//
	// The text message displayed for the menu item which performs descending sort. (default: "Sort Descending")
	//
	//    Example - set the sort descending message
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
	//            sortDescending: "Sort (desc)",
	//          }
	//        },
	//        sortable: true,
	//        dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	SortDescending string `jsObject:"sortDescending"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columnmenu.messages#columnMenu.messages.settings
	//
	// The text message displayed in the menu header (available in mobile mode only). (default: "Column Settings")
	//
	//    Example - mobile column menu header
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
	//            settings: "Column Options",
	//          }
	//        },
	//        mobile: "phone",
	//        sortable: true,
	//        dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	Settings string `jsObject:"settings"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columnmenu.messages#columnMenu.messages.done
	//
	// The text message displayed in the menu header button (available in mobile mode only). (default: "Done")
	//
	//    Example - mobile column menu header button text
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
	//            done: "Ok",
	//          }
	//        },
	//        mobile: "phone",
	//        sortable: true,
	//        dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	Done string `jsObject:"done"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columnmenu.messages#columnMenu.messages.lock
	//
	// The text message displayed in the column menu for locking a column. (default: "Lock")
	//
	//    Example - column menu lock button text
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { locked: true, field: "id", width:200 },
	//          { field: "name", width:400 },
	//          { field: "age", width:400 }
	//        ],
	//        columnMenu: {
	//          messages: {
	//            lock: "Pin this column"
	//          }
	//        },
	//        sortable: true,
	//        dataSource: [
	//          { id: 1, name: "Jane Doe", age: 30 },
	//          { id: 2, name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	Lock string `jsObject:"lock"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columnmenu.messages#columnMenu.messages.unlock
	//
	// The text message displayed in the column menu for unlocking a column. (default: "Unlock")
	//
	//    Example - column menu unlock button text
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { locked: true, field: "id", width:200 },
	//          { field: "name", width:400 },
	//          { field: "age", width:400 }
	//        ],
	//        columnMenu: {
	//          messages: {
	//            unlock: "Unpin this column"
	//          }
	//        },
	//        sortable: true,
	//        dataSource: [
	//          { id: 1, name: "Jane Doe", age: 30 },
	//          { id: 2, name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	Unlock string `jsObject:"unlock"`

	*ToJavaScriptConverter
}

func (el *KendoGridColumnMenuMessages) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridColumnMenuMessages.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
