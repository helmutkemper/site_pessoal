package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridFilterableOperatorsDate struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.date#filterable.operators.date.gt

	  The text of the "greater than" filter operator. (default: "Is after")
	*/
	//
	//    Example - set the date "greater than" operator
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
	//            gt: "After"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Gt string `jsObject:"gt"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.date#filterable.operators.date.isnotnull

	  The text of the "isnotnull" filter operator. (default: "Is not null")
	*/
	//
	//    Example - set the date "isnotnull" operator
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
	//            isnotnull: "Null"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	IsNotNull string `jsObject:"isnotnull"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.date#filterable.operators.date.lte

	  The text of the "less than or equal" filter operator. (default: "Is before or equal to")
	*/
	//
	//    Example - set the date "less than or equal" operator
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
	//            lte: "Before or equal to"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Lte string `jsObject:"lte"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.date#filterable.operators.date.neq

	  The text of the "not equal" filter operator. (default: "Is not equal to")
	*/
	//
	//    Example - set the date "not equal" operator
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
	//            neq: "Not equal"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Neq string `jsObject:"neq"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.date#filterable.operators.date.isnull

	  The text of the "isnull" filter operator. (default: "Is null")
	*/
	//
	//    Example - set the date "isnull" operator
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
	//            isnull: "Null"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	IsNull string `jsObject:"isnull"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.date#filterable.operators.date.lt

	  The text of the "less than" filter operator. (default: "Is before")
	*/
	//
	//    Example - set the date "less than" operator
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
	//            lt: "Before"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Lt string `jsObject:"lt"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.date#filterable.operators.date.gte

	  The text of the "greater than or equal" filter operator. (default: "Is after or equal to")
	*/
	//
	//    Example - set the date "greater than or equal" operator
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
	//            gte: "After or equal to"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Gte string `jsObject:"gte"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators.date#filterable.operators.date.eq

	  The text of the "equal" filter operator. (default: "Is equal to")
	*/
	//
	//    Example - set the date "equal" operator
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
	//            eq: "Equal"
	//          }
	//        }
	//      }
	//    });
	//    </script>
	//
	Eq string `jsObject:"eq"`

	*ToJavaScriptConverter
}

func (el *KendoGridFilterableOperatorsDate) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridFilterableOperatorsDate.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
