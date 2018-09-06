package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridFilterableOperatorsNumber struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.number#filterable.operators.number.neq

	  The text of the "not equal" filter operator. (default: "Is not equal to")
	*/
	//
	//    Example - set the number "not equal" operator
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
	//            neq: "Not equal to"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Neq string `jsObject:"neq"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.number#filterable.operators.number.lt

	  The text of the "less than" filter operator. (default: "Is less than")
	*/
	//
	//    Example - set the number "less than" operator
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
	//            lt: "Less than"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Lt string `jsObject:"lt"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.number#filterable.operators.number.eq

	  The text of the "equal" filter operator. (default: "Is equal to")
	*/
	//
	//    Example - set the number "equal" operator
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
	//            eq: "Equal to"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Eq string `jsObject:"eq"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.number#filterable.operators.number.isnull

	  The text of the "isnull" filter operator. (default: "Is null")
	*/
	//
	//    Example - set the number "isnull" operator
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
	//            isnull: "Null"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	IsNull string `jsObject:"isnull"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.number#filterable.operators.number.lte

	  The text of the "less than or equal" filter operator. (default: "Is less than or equal to")
	*/
	//
	//    Example - set the number "less than or equal to" operator
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
	//            lte: "Less than or equal to"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Lte string `jsObject:"lte"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.number#filterable.operators.number.gt

	  The text of the "greater than" filter operator. (default: "Is greater than")
	*/
	//
	//    Example - set the number "greater than" operator
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
	//            gt: "Greater than"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Gt string `jsObject:"gt"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.number#filterable.operators.number.gte

	  The text of the "greater than or equal" filter operator. (default: "Is greater than or equal to")
	*/
	//
	//    Example - set the number "greater than or equal to" operator
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
	//            gte: "Greater than or equal to"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Gte string `jsObject:"gte"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.number#filterable.operators.number.isnotnull

	  The text of the "isnotnull" filter operator. (default: "Is not null")
	*/
	//
	//    Example - set the number "isnotnull" operator
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

func (el *KendoGridFilterableOperatorsNumber) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridFilterableOperatorsNumber.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
