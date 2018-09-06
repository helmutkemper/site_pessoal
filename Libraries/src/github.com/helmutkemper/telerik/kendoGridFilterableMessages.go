package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridFilterableMessages struct {
	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.and
	//
	// The text of the option which represents the "and" logical operation. (default: "And")
	//
	//    Example - set the "and" message
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
	//            and: "and"
	//          }
	//        }
	//      });
	//    </script>
	And string `jsObject:"and"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.clear
	//
	// The text of the button which clears the filter. (default: "Clear")
	//
	//    Example - set the "clear" message
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
	//            clear: "Clear filter"
	//          }
	//        }
	//      });
	//    </script>
	Clear string `jsObject:"clear"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.filter
	//
	// The text of the button which applies the filter. (default: "Filter")
	//
	//    Example - set the "filter" message
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
	//            filter: "Apply filter"
	//          }
	//        }
	//      });
	//    </script>
	Filter string `jsObject:"filter"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.info
	//
	// The text of the information message on the top of the filter menu. (default: "Show items with value that: ")
	//
	//    Example - set the "info" message
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
	//            info: "Filter by: "
	//          }
	//        }
	//      });
	//    </script>
	Info string `jsObject:"info"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.title
	//
	// The text rendered for the title attribute of the filter menu form. (default: "Show items with value that: ")
	Title string `jsObject:"title"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.isFalse
	//
	// The text of the radio button for false values. Displayed when filtering Boolean fields. (default: "is false")
	//
	//    Example - set the "isFalse" message
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "active" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { active: true },
	//            { active: false }
	//          ],
	//          schema: {
	//            model: {
	//              fields: {
	//                active: { type: "boolean" }
	//              }
	//            }
	//          }
	//        },
	//        filterable: {
	//          messages: {
	//            isFalse: "False"
	//          }
	//        }
	//      });
	//    </script>
	IsFalse string `jsObject:"isFalse"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.isTrue
	//
	// The text of the radio button for true values. Displayed when filtering Boolean fields. (default: "is true")
	//
	//    Example - set the "isTrue" message
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "active" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { active: true },
	//            { active: false }
	//          ],
	//          schema: {
	//            model: {
	//              fields: {
	//                active: { type: "boolean" }
	//              }
	//            }
	//          }
	//        },
	//        filterable: {
	//          messages: {
	//            isTrue: "True"
	//          }
	//        }
	//      });
	//    </script>
	IsTrue string `jsObject:"isTrue"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.or
	//
	// The text of the option which represents the "or" logical operation. (default: "Or")
	//
	//    Example - set the "or" message
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
	//            or: "or"
	//          }
	//        }
	//      });
	//    </script>
	Or string `jsObject:"or"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.search
	//
	// The placeholder of the search input for columns with the search option set to true. (default: "Search")
	//
	//    Example - set the "search" message
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          {
	//            field: "category",
	//            filterable: {
	//              multi: true,
	//              search: true
	//            }
	//          }
	//        ],
	//        dataSource: [
	//          { category: "Foo" },
	//          { category: "Boo" }
	//        ],
	//        filterable: {
	//          messages: {
	//            search: "Search category"
	//          }
	//        }
	//      });
	//    </script>
	Search string `jsObject:"search"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.selectValue
	//
	// The text of the DropDownList displayed in the filter menu for columns whose values option is set.
	// (default: "-Select value-")
	//
	//    Example - set the "selectValue" message
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "productName" },
	//          { field: "category", values: [
	//              { text: "Beverages", value: 1 },
	//              { text: "Food", value: 2 },
	//            ]
	//          }
	//        ],
	//        dataSource: [
	//          { productName: "Tea", category: 1 },
	//          { productName: "Ham", category: 2 }
	//        ],
	//        filterable: {
	//          messages: {
	//            selectValue: "Select category"
	//          }
	//        }
	//      });
	//    </script>
	SelectValue string `jsObject:"selectValue"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.cancel
	//
	// The text of the cancel button in the filter menu header (available in mobile mode only). (default: "Cancel")
	//
	//    Example - set the cancel button text
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "productName" },
	//          { field: "category", values: [
	//              { text: "Beverages", value: 1 },
	//              { text: "Food", value: 2 },
	//            ]
	//          }
	//        ],
	//        dataSource: [
	//          { productName: "Tea", category: 1 },
	//          { productName: "Ham", category: 2 }
	//        ],
	//        mobile: "phone",
	//        filterable: {
	//          messages: {
	//            cancel: "Reject"
	//          }
	//        }
	//      });
	//    </script>
	Cancel string `jsObject:"cancel"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.selectedItemsFormat
	//
	// The format string for selected items count in filter menu when search option set to true.
	// (default: "{0} items selected")
	//
	//    Example - set the "selectedItemsFormat" text
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          {
	//            field: "category",
	//            filterable: {
	//              multi: true,
	//              search: true
	//            }
	//          }
	//        ],
	//        dataSource: [
	//          { category: "Foo" },
	//          { category: "Boo" }
	//        ],
	//        filterable: {
	//          messages: {
	//            selectedItemsFormat: "There are {0} selected items"
	//          }
	//        }
	//      });
	//    </script>
	SelectedItemsFormat string `jsObject:"selectedItemsFormat"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.operator
	//
	// The text of the operator item in filter menu (available in mobile mode only). (default: "Operator")
	//
	//    Example - set the text of operator item
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "productName" },
	//          { field: "category", values: [
	//              { text: "Beverages", value: 1 },
	//              { text: "Food", value: 2 },
	//            ]
	//          }
	//        ],
	//        dataSource: [
	//          { productName: "Tea", category: 1 },
	//          { productName: "Ham", category: 2 }
	//        ],
	//        mobile: "phone",
	//          filterable: {
	//            messages: {
	//              operator: "Choose operator"
	//            }
	//          }
	//      });
	//    </script>
	Operator string `jsObject:"operator"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.value
	//
	// The text of the value item in filter menu (available in mobile mode only). (default: "Value")
	//
	//    Example - set the text of value item
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "productName" },
	//          { field: "category", values: [
	//              { text: "Beverages", value: 1 },
	//              { text: "Food", value: 2 },
	//            ]
	//          }
	//        ],
	//        dataSource: [
	//          { productName: "Tea", category: 1 },
	//          { productName: "Ham", category: 2 }
	//        ],
	//        mobile: "phone",
	//        filterable: {
	//          messages: {
	//            value: "Choose value"
	//          }
	//        }
	//      });
	//    </script>
	Value string `jsObject:"value"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.messages#filterable.messages.checkAll
	//
	// The label used for the check-all checkbox.
	//
	//    Example - change the checkAll default message.
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [{
	//          field: "country",
	//          filterable: {
	//            multi:true,
	//            messages: {
	//              checkAll: "Do select all"
	//            },
	//            itemTemplate: function(e) {
	//              return "<span><label><span>#= data.country|| data.all #</span><input type='checkbox' name='" + e.field + "' value='#= data.country#'/></label></span>"
	//            }
	//          }
	//        }],
	//        filterable: true,
	//        dataSource: [ { country: "BG" }, { country: "USA" } ]
	//      });
	//    </script>
	CheckAll string `jsObject:"checkAll"`

	*ToJavaScriptConverter
}

func (el *KendoGridFilterableMessages) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridFilterableMessages.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
