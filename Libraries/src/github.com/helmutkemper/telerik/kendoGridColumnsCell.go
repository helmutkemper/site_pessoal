package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridColumnsCell struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.cell#columns.filterable.cell.delay

	  Specifies the delay of the AutoComplete widget which will provide the suggest functionality (when using String type
	  column). (default: 200)

	*/
	//
	//    Example - Specifying delay option for the AutoComplete widget used to make suggestions while filtering.
	//
	//    <div id="grid"></div>
	//    <script>
	//          $("#grid").kendoGrid({
	//              columns: [
	//                {
	//                    field: "name",
	//                    filterable: {
	//                        cell: {
	//                            delay: 1500
	//                        }
	//                    }
	//                },
	//                { field: "age" } ],
	//            filterable: { mode: "row" },
	//            dataSource: {
	//                data: [ { name: "Jane", age: 30 }, { name: "John", age: 33 }],
	//              schema:{
	//                model:{
	//                    fields: {
	//                        age: { type: "number" }
	//                      }
	//                  }
	//              }
	//            }
	//          });
	//    </script>
	//
	Delay int `jsObject:"delay"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.cell#columns.filterable.cell.enabled

	  When set to false the Grid will not render the cell filtering widget for that specific column.
	  (default: true)

	*/
	//
	//    Example - Disable the cell filtering for a specific column.
	//
	//    <div id="grid"></div>
	//    <script>
	//          $("#grid").kendoGrid({
	//              columns: [
	//                {
	//                    field: "name",
	//                    filterable: {
	//                        cell: {
	//                            enabled: false
	//                        }
	//                    }
	//                },
	//                { field: "age" } ],
	//            filterable: { mode: "row" },
	//            dataSource: {
	//                data: [ { name: "Jane", age: 30 }, { name: "John", age: 33 }],
	//              schema:{
	//                model:{
	//                    fields: {
	//                        age: { type: "number" }
	//                      }
	//                  }
	//              }
	//            }
	//          });
	//    </script>
	//
	Enabled Boolean `jsObject:"enabled"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.cell#columns.filterable.cell.minlength

	  Specifies the minLength option of the AutoComplete widget when column is of type string.
	  (default: 1)

	*/
	//
	//    Example - Specifying minLength of the AutoComplete widget when using filter cell.
	//
	//    <div id="grid"></div>
	//    <script>
	//          $("#grid").kendoGrid({
	//              columns: [
	//                {
	//                    field: "name",
	//                    filterable: {
	//                        cell: {
	//                            minLength: 3
	//                        }
	//                    }
	//                },
	//                { field: "age" } ],
	//            filterable: { mode: "row" },
	//            dataSource: {
	//                data: [ { name: "Jane", age: 30 }, { name: "John", age: 33 }],
	//              schema:{
	//                model:{
	//                    fields: {
	//                        age: { type: "number" }
	//                      }
	//                  }
	//              }
	//            }
	//            });
	//    </script>
	//
	MinLength int `jsObject:"minLength"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.cell#columns.filterable.cell.datasource

	  Specifies a custom dataSource for the AutoComplete when the type of the column is `string´. Can be a JavaScript object
	  which represents a valid data source configuration, a JavaScript array, or an existing kendo.data.DataSource
	  [ https://docs.telerik.com/kendo-ui/api/javascript/data/datasource ] instance.

	  It is not recommended that you use the same `dataSource´ instance for the Grid and the AutoComplete because it causes
	  negative side effects.

	  If the `dataSource´ options is missing, a new cloned instance of the Grid's dataSource will be used.

	  If the `dataSource´ option is an existing kendo.data.DataSource
	  [ https://docs.telerik.com/kendo-ui/api/javascript/data/datasource ] instance, the widget will use that instance and
	  will not initialize a new one.

	*/
	//
	//    Example - custom cell filter autocomplete dataSource
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        {
	//            field: "name",
	//            filterable: {
	//                cell: {
	//                    dataSource: new kendo.data.DataSource({
	//                        data: [
	//                            { someField: "Jane" },
	//                            { someField: "Jake" },
	//                            { someField: "John" }
	//                        ]
	//                    }),
	//                    dataTextField: "someField"
	//                }
	//            }
	//        },
	//        { field: "age" }
	//      ],
	//      filterable: {
	//          mode: "row"
	//      },
	//      dataSource: {
	//        data: [ { name: "Jane", age: 30 }, { name: "John", age: 33 }],
	//        schema:{
	//            model:{
	//                fields: {
	//                    age: { type: "number" }
	//                }
	//            }
	//        }
	//      }
	//    });
	//    </script>
	//
	DataSource interface{} `jsObject:"dataSource" jsType:"*KendoDataSource,string,*map[string]interface {},[]string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.cell#columns.filterable.cell.inputwidth

	  Specifies the width of the input before it is initialized or turned into a widget. Provides convenient way to set the
	  width according to the column width.

	*/
	//
	//    Example - Specifying inputWidth option for the filter cell of a column
	//
	//    <div id="grid"></div>
	//    <script>
	//          $("#grid").kendoGrid({
	//              columns: [
	//                {
	//                    field: "name",
	//                    filterable: {
	//                        cell: {
	//                            inputWidth: 333
	//                        }
	//                    }
	//                },
	//                { field: "age" } ],
	//            filterable: { mode: "row" },
	//            dataSource: {
	//                data: [ { name: "Jane", age: 30 }, { name: "John", age: 33 }],
	//              schema:{
	//                model:{
	//                    fields: {
	//                        age: { type: "number" }
	//                      }
	//                  }
	//              }
	//            }
	//          });
	//    </script>
	//
	InputWidth int `jsObject:"inputWidth"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.cell#columns.filterable.cell.datatextfield

	  Specifies the name of the field which will provide the text representation for the AutoComplete suggestion (when using
	  String type column) when CustomDataSource is provided. By default the name of the field bound to the column will be
	  used.

	*/
	//
	//    Example - Using custom dataSource and providing dataTextField option
	//
	//    <div id="grid"></div>
	//    <script>
	//          $("#grid").kendoGrid({
	//            columns: [
	//              {
	//                  field: "name",
	//                  filterable: {
	//                      cell: {
	//                          dataSource: new kendo.data.DataSource({ data: [
	//                              { someField: "Jane" },
	//                              { someField: "Jake" },
	//                              { someField: "John" }
	//                          ] }),
	//                          dataTextField: "someField"
	//                      }
	//                  }
	//              },
	//              { field: "age" }
	//            ],
	//            filterable: { mode: "row" },
	//            dataSource: {
	//                data: [ { name: "Jane", age: 30 }, { name: "John", age: 33 }],
	//              schema:{
	//                model:{
	//                    fields: {
	//                        age: { type: "number" }
	//                      }
	//                  }
	//              }
	//            }
	//         });
	//    </script>
	//
	DataTextField string `jsObject:"dataTextField"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.cell#columns.filterable.cell.template

	  JavaScript function which will customize how the input for the filter value is rendered.
	  The function receives an object argument with two fields:

	  `element´ - the default input inside the filter cell;

	  `dataSource´ - a Kendo UI DataSource instance, which has the same settings as the Grid dataSource, but will only
	  contain data items with unique values for the current column.

	  This instance is also used by the default AutoComplete widget, which is used inside the filter cell if no template is
	  set. Keep in mind that the passed dataSource instance may still not be populated at the time the template function is
	  called, if the Grid uses remote binding.

	*/
	//
	//    Example - Using template for the filter cell
	//
	//    <div id="grid"></div>
	//    <script>
	//          $("#grid").kendoGrid({
	//              columns: [
	//                {
	//                    field: "color",
	//                    filterable: {
	//                        cell: {
	//                            template: function (args) {
	//                                // create a DropDownList of unique values (colors)
	//                                args.element.kendoDropDownList({
	//                                    dataSource: args.dataSource,
	//                                    dataTextField: "color",
	//                                    dataValueField: "color",
	//                                    valuePrimitive: true
	//                                });
	//
	//                                // or
	//
	//                                // create a ColorPicker
	//                                // args.element.kendoColorPicker();
	//                            },
	//                            showOperators: false
	//                        }
	//                    }
	//                },
	//                { field: "size" } ],
	//            filterable: { mode: "row" },
	//            dataSource: [ { color: "#ff0000", size: 30 }, { color: "#000000", size: 33 }] });
	//    </script>
	//
	Template interface{} `jsObject:"template" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.cell#columns.filterable.cell.showoperators

	  Specifies whether to show or hide the DropDownList with the operators. (default: true)

	*/
	//
	//    Example - Hide the operators dropdownlist for cell filtering.
	//
	//    <div id="grid"></div>
	//    <script>
	//          $("#grid").kendoGrid({
	//              columns: [
	//                {
	//                    field: "name",
	//                    filterable: {
	//                        cell: {
	//                            showOperators: false,
	//                            operator: "contains"
	//                        }
	//                    }
	//                },
	//                { field: "age" } ],
	//            filterable: { mode: "row" },
	//            dataSource: {
	//                data: [ { name: "Jane", age: 30 }, { name: "John", age: 33 }],
	//              schema:{
	//                model:{
	//                    fields: {
	//                        age: { type: "number" }
	//                      }
	//                  }
	//              }
	//            }
	//          });
	//    </script>
	//
	ShowOperators Boolean `jsObject:"showOperators"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.cell#columns.filterable.cell.suggestionoperator

	  Specifies the AutoComplete `filter´ option. The possible values are the same as the ones for the AutoComplete
	  `filter´ option - `"startswith"´, `"endswith"´, `"contains"´. The `"contains"´ operator performs a case-insensitive
	  search. To perform a case-sensitive filtering, set a custom filtering function through the dataSource.filter.operator
	  [ https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/filter.operator ] option.
	  (default: "startswith")

	  Important:

	  This operator is completely independent from the operator used for the filtering on this column. For more information,
	  check operator
	  [ https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.cell.operator ]

	*/
	//
	//    Example - Specifying suggestionOperator option for the filter cell of a column
	//
	//    <div id="grid"></div>
	//    <script>
	//          $("#grid").kendoGrid({
	//              columns: [
	//                {
	//                    field: "name",
	//                    filterable: {
	//                        cell: {
	//                            suggestionOperator: "contains"
	//                        }
	//                    }
	//                },
	//                { field: "age" } ],
	//            filterable: { mode: "row" },
	//            dataSource: {
	//                data: [ { name: "Jane", age: 30 }, { name: "John", age: 33 }],
	//              schema:{
	//                model:{
	//                    fields: {
	//                        age: { type: "number" }
	//                      }
	//                  }
	//              }
	//            }
	//          });
	//    </script>
	//
	SuggestionOperator KendoFilter `jsObject:"suggestionOperator"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.cell#columns.filterable.cell.operator

	  Specifies the default operator that will be used for the cell filtering. (default: "eq")

	  Important:

	  If you want to change how the AutoComplete suggestions are filtered use
	<a href="columns.filterable.cell.suggestionoperator">suggestionOperator</a>.

	*/
	//
	//    Example - Specifying default operator for cell filtering.
	//
	//    <div id="grid"></div>
	//    <script>
	//          $("#grid").kendoGrid({
	//              columns: [
	//                {
	//                    field: "name",
	//                    filterable: {
	//                        cell: {
	//                            operator: "neq"
	//                        }
	//                    }
	//                },
	//                { field: "age" } ],
	//            filterable: { mode: "row" },
	//            dataSource: {
	//                data: [ { name: "Jane", age: 30 }, { name: "John", age: 33 }],
	//              schema:{
	//                model:{
	//                    fields: {
	//                        age: { type: "number" },
	//                        name: {type: "string"}
	//                      }
	//                  }
	//              }
	//            }
	//         });
	//    </script>
	//
	Operator KendoFilter `jsObject:"operator"`

	*ToJavaScriptConverter
}

func (el *KendoGridColumnsCell) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoCreate.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
