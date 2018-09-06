package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridFilterable struct {
	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.extra
	//
	// If set to true the filter menu allows the user to input a second criterion. (default: true)
	//
	//    Example - disable the extra filtering criteria
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" }
	//        ],
	//        filterable: {
	//          extra: false
	//        },
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema:{
	//            model: {
	//              id: "id",
	//              fields: {
	//                age: { type: "number"}
	//              }
	//            }
	//          }
	//        }
	//      });
	//    </script>
	Extra Boolean `jsObject:"extra"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages
	//
	// The text messages displayed in the filter menu. Use it to customize or localize the filter menu messages.
	//
	//    Example - customize filter menu messages
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema:{
	//            model: {
	//              id: "id",
	//              fields: {
	//                age: { type: "number"}
	//              }
	//            }
	//          }
	//        },
	//        filterable: {
	//          messages: {
	//            and: "and",
	//            or: "or",
	//            filter: "Apply filter",
	//            clear: "Clear filter"
	//          }
	//        }
	//      });
	//    </script>
	Messages KendoGridFilterableMessages `jsObject:"messages"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.mode
	//
	// If set to row the user would be able to filter via extra row added below the headers. By default filtering is using
	// the menu mode.
	//
	// Can also be set to the following string values:
	//
	// > "row" - the user can filter via extra row within the header.
	// > "menu" - the user can filter via the menu after clicking the filter icon.
	// > "menu, row" - the user can filter with both modes above enabled.
	//
	// Important:
	//
	// When the filterable.mode property is set to "row" or "menu, row", the Grid dataSource instance is copied and
	// applied to the Kendo UI AutoComplete widgets used for string filtering. This will cause one additional read request
	// per string column. The AutoComplete dataSources do not perform paging and will use a collection of the unique
	// column values only.
	//
	//    Example - set mode option to use both "menu" and "row" modes simultaneously
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" }
	//        ],
	//        filterable: {
	//          mode: "menu, row"
	//        },
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema:{
	//            model: {
	//              id: "id",
	//              fields: {
	//                age: { type: "number"}
	//              }
	//            }
	//          }
	//        }
	//      });
	//    </script>
	Mode KendoGridFilterableMode `jsObject:"mode"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators
	//
	// The text of the filter operators displayed in the filter menu.
	//
	// Important:
	//
	// If operators are defined manually, the default messages will be overridden too. To control the operators and still
	// use the default messages, retrieve them from the FilterCell prototype -
	// kendo.ui.FilterCell.fn.options.operators.{type}, where the type can be "string", "date", "number", and "enums".
	//
	// If the same options are specific to a column, it is possible to use the column filterable configuration of the
	// Grid.
	//
	// In multiple Grids, it is possible to override the filterable options of the Kendo UI FilterMenu before the Grids
	// are initialized. Then the new filter options will be available for all Grids without further configurations.
	//
	//    Example - override the filterable options in multiple Grids
	//
	//    <h4>Grid One</h4>
	//    <div id="gridOne"></div>
	//    <h4>Grid Two</h4>
	//    <div id="gridTwo"></div>
	//
	//    <script>
	//      kendo.ui.FilterMenu.fn.options.operators.string = {
	//        eq: "Equal to",
	//        neq: "Not equal to"
	//      };
	//
	//      $("#gridOne").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema: {
	//            model: {
	//              id: "id",
	//              fields: {
	//                name: { type: "string" },
	//                age: { type: "number" }
	//              }
	//            }
	//          }
	//        },
	//        filterable: {
	//          extra: false
	//        }
	//      });
	//
	//      $("#gridTwo").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema: {
	//            model: {
	//              id: "id",
	//              fields: {
	//                name: { type: "string" },
	//                age: { type: "number" }
	//              }
	//            }
	//          }
	//        },
	//        filterable: {
	//          extra: false
	//        }
	//      });
	//    </script>
	Operators KendoGridFilterableOperators `jsObject:"operators"`

	*ToJavaScriptConverter
}

func (el *KendoGridFilterable) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridFilterable.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
