package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridFilterableOperators struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.enums

	  The texts of the filter operators displayed for columns which have their <a href="columns.values">values</a> option set.

	  Important:
	  Omitting an operator will exclude it from the DropDownList with the available operators.

	*/
	//
	//    Example - set enum operators
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
	//            eq: "Equal to",
	//            neq: "Not equal to"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Enums KendoGridFilterableOperatorsEnums `jsObject:"enums"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.date

	  The texts of the filter operators displayed for columns bound to date fields.

	  Important:
	  Omitting an operator will exclude it from the DropDownList with the available operators.

	*/
	//
	//    Example - set date operators
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "date", format: "{0:yyyy-MM-dd}" }
	//      ],
	//      dataSource: {
	//        data: [
	//          { date: kendo.parseDate("2000-10-10") },
	//          { date: new Date() }
	//        ],
	//        schema: {
	//          model: {
	//            fields: {
	//              date: { type: "date" }
	//            }
	//          }
	//        }
	//      },
	//      filterable: {
	//        operators: {
	//          date: {
	//            gt: "After",
	//            lt: "Before"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Date KendoGridFilterableOperatorsDate `jsObject:"date"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.string

	  The texts of the filter operators displayed for columns bound to string fields.

	  Important:
	  Omitting an operator will exclude it from the DropDownList with the available operators.

	*/
	//
	//    Example - set string operators
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe" },
	//        { name: "John Doe" }
	//      ],
	//      filterable: {
	//        operators: {
	//          string: {
	//            eq: "Equal to",
	//            neq: "Not equal to"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	String KendoGridFilterableOperatorsString `jsObject:"string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.number

	  The texts of the filter operators displayed for columns bound to number fields.

	  Important:
	  Omitting an operator will exclude it from the DropDownList with the available operators.

	*/
	//
	//    Example - set number operators
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: {
	//        data: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ],
	//        schema: {
	//          model: {
	//            fields: {
	//              age: { type: "number" }
	//            }
	//          }
	//        }
	//      },
	//      filterable: {
	//        operators: {
	//          number: {
	//            eq: "Equal to",
	//            neq: "Not equal to"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Number KendoGridFilterableOperatorsNumber `jsObject:"number"`

	*ToJavaScriptConverter
}

func (el *KendoGridFilterableOperators) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridFilterableOperators.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
