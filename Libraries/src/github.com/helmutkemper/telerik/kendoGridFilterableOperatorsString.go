package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridFilterableOperatorsString struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.string#filterable.operators.string.isnotempty

	  The text of the "isnotempty" filter operator. (default: "Is not empty")
	*/
	//
	//    Example - set the string "isnotempty" operator
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
	//            isnotempty: "Not empty"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	IsNotEmpty string `jsObject:"isnotempty"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.string#filterable.operators.string.isnull

	  The text of the "isnull" filter operator. (default: "Is null")
	*/
	//
	//    Example - set the string "isnull" operator
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
	//            isnull: "Null"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	IsNull string `jsObject:"isnull"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.string#filterable.operators.string.isempty

	  The text of the "isempty" filter operator. (default: "Is empty")
	*/
	//
	//    Example - set the string "isempty" operator
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
	//            isempty: "Empty"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	IsEmpty string `jsObject:"isempty"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.string#filterable.operators.string.doesnotcontain

	  The text of the "does not contain" filter operator. (default: "Does not contain")
	*/
	//
	//    Example - set the string "does not contain" operator
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
	//            doesnotcontain: "Doesn't contain"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	DoesNotContain string `jsObject:"doesnotcontain"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.string#filterable.operators.string.contains

	  The text of the "contains" filter operator. (default: "Contains")
	*/
	//
	//    Example - set the string "contains" operator
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
	//            contains: "Contains"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Contains string `jsObject:"contains"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.string#filterable.operators.string.endswith

	  The text of the "ends with" filter operator. (default: "Ends with")
	*/
	//
	//    Example - set the string "ends with" operator
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
	//            endswith: "Ends"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	EndsWith string `jsObject:"endswith"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.string#filterable.operators.string.startswith

	  The text of the "starts with" filter operator. (default: "Starts with")
	*/
	//
	//    Example - set the string "starts with" operator
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
	//            startswith: "Starts"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	StartsWith string `jsObject:"startswith"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.string#filterable.operators.string.eq

	  The text of the "equal" filter operator. (default: "Is equal to")
	*/
	//
	//    Example - set the string "equal" operator
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
	//            eq: "Equal to"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Eq string `jsObject:"eq"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.string#filterable.operators.string.neq

	  The text of the "not equal" filter operator. (default: "Is not equal to")
	*/
	//
	//    Example - set the string "not equal" operator
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
	//            neq: "Not equal to"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Neq string `jsObject:"neq"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.string#filterable.operators.string.isnotnull

	  The text of the "isnotnull" filter operator. (default: "Is not null")
	*/
	//
	//    Example - set the string "isnotnull" operator
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
	//            isnotnull: "Not null"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	IsNotNull string `jsObject:"isnotnull"`

	*ToJavaScriptConverter
}

func (el *KendoGridFilterableOperatorsString) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridFilterableOperatorsString.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
