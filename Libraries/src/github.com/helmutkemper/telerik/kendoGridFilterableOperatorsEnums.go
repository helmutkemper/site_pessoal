package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridFilterableOperatorsEnums struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.enums#filterable.operators.enums.eq

	  The text of the "equal" filter operator. (default: "Is equal to")
	*/
	//
	//    Example - set the enum "equal" operator
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "productName" },
	//        { field: "category", values: [
	//            {text: "Beverages", value: 1 },
	//            {text: "Food", value: 2 }
	//          ]
	//        }
	//      ],
	//      dataSource: [
	//        { productName: "Tea", category: 1 },
	//        { productName: "Ham", category: 2 }
	//      ],
	//      filterable: {
	//        operators: {
	//          enums: {
	//            eq: "Equal to"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Eq string `jsObject:"eq"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.enums#filterable.operators.enums.isnull

	  The text of the "isnull" filter operator. (default: "Is null")
	*/
	//
	//    Example - set the enum "isnull" operator
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "productName" },
	//        { field: "category", values: [
	//            {text: "Beverages", value: 1 },
	//            {text: "Food", value: 2 }
	//          ]
	//        }
	//      ],
	//      dataSource: [
	//        { productName: "Tea", category: 1 },
	//        { productName: "Ham", category: 2 }
	//      ],
	//      filterable: {
	//        operators: {
	//          enums: {
	//            isnull: "Null"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	IsNull string `jsObject:"isnull"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.enums#filterable.operators.enums.isnotnull

	  The text of the "isnotnull" filter operator. (default: "Is not null")
	*/
	//
	//    Example - set the enum "isnotnull" operator
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "productName" },
	//        { field: "category", values: [
	//            {text: "Beverages", value: 1 },
	//            {text: "Food", value: 2 }
	//          ]
	//        }
	//      ],
	//      dataSource: [
	//        { productName: "Tea", category: 1 },
	//        { productName: "Ham", category: 2 }
	//      ],
	//      filterable: {
	//        operators: {
	//          enums: {
	//            isnotnull: "Not null"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	IsNotNull string `jsObject:"isnotnull"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.enums#filterable.operators.enums.neq

	  The text of the "not equal" filter operator. (default: "Is not equal to")
	*/
	//
	//    Example - set the enum "not equal" operator
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "productName" },
	//        { field: "category", values: [
	//            {text: "Beverages", value: 1 },
	//            {text: "Food", value: 2 }
	//          ]
	//        }
	//      ],
	//      dataSource: [
	//        { productName: "Tea", category: 1 },
	//        { productName: "Ham", category: 2 }
	//      ],
	//      filterable: {
	//        operators: {
	//          enums: {
	//            neq: "Not equal to"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Neq string `jsObject:"neq"`

	*ToJavaScriptConverter
}

func (el *KendoGridFilterableOperatorsEnums) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridFilterableOperatorsEnums.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
