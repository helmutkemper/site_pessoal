package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridSortable struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/sortable.initialdirection

	  Determines the inital (from un-sorted to sorted state) sort direction. The supported values are `asc´ and `desc´. (default: asc)
	*/
	//
	//    Example - disable sorting
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        {
	//          field: "id",
	//          sortable: {
	//            initialDirection: "desc"
	//          }
	//        },
	//        { field: "name" }
	//      ],
	//      sortable: true,
	//      dataSource: [ { id: 1, name: "Jane Doe" }, { id: 2, name: "John Doe" } ]
	//    });
	//    </script>
	//
	InitialDirection KendoDirection `jsObject:"initialDirection"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/sortable.showindexes

	  If set to `true´ the user will see sort sequence indicators for sorted columns. (default: true)
	*/
	//
	//    Example - do not allow unsorting
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      sortable: {
	//        showIndexes: true,
	//        mode: "multiple"
	//      }
	//    });
	//    </script>
	//
	ShowIndexes Boolean `jsObject:"showIndexes"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/sortable.allowunsort

	  If set to `true´ the user can get the grid in unsorted state by clicking the sorted column header. If set to `false´ the user will not be able to unsort the column once sorted. (default: true)
	*/
	//
	//    Example - disable unsorting
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { sortable: { allowUnsort: false }, field: "id" },
	//        { field: "name" }
	//      ],
	//      sortable: true,
	//      dataSource: [ { id: 1, name: "Jane Doe" }, { id: 2, name: "John Doe" } ]
	//    });
	//    </script>
	//
	AllowUnSort Boolean `jsObject:"allowUnsort"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/sortable.mode

	  The sorting mode. If set to "single" the user can sort by one column. If set to "multiple" the user can sort by multiple columns. (default: "single")
	*/
	//
	//    Example - allow multiple column sorting
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      sortable: {
	//        mode: "multiple"
	//      }
	//    });
	//    </script>
	//
	Mode KendoTagMode `jsObject:"mode"`

	*ToJavaScriptConverter
}

func (el *KendoGridSortable) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridSortable.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
