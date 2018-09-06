package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridColumnsFilterable struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.cell

	  Specifies options for the filter header cell when filter mode is set to 'row'.

	  Can be set to a JavaScript object which represents the filter cell configuration.

	*/
	//
	//    Example - cell filtering options
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        {
	//            field: "name",
	//            filterable: {
	//                cell: {
	//                    enabled: true,
	//                    delay: 1500
	//                }
	//            },
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
	Cell KendoGridColumnsCell `jsObject:"cell"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.extra

	  If set to `true´ the filter menu allows the user to input a second criterion. (default: true)

	*/
	//
	//    Example - disable the extra filtering criteria
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      filterable: {
	//        extra: false
	//      },
	//      dataSource: {
	//       data: [
	//        { id: 1, name: "Jane Doe", age: 30 },
	//        { id: 2, name: "John Doe", age: 33 }
	//       ],
	//       schema:{
	//        model: {
	//         id: "id",
	//         fields: {
	//           age: { type: "number"}
	//         }
	//        }
	//       }
	//      }
	//    });
	//    </script>
	//
	Extra Boolean `jsObject:"extra"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.multi

	  Use this option to enable the MultiCheck filtering support for that column. (default: false)

	  Important:

	  If you have enabled the columns.multi option and your Grid uses serverPaging (or ServerOperations(true) when using the
	  MVC wrappers) you will need to provide columns.filterable.dataSource. Otherwise, a negative impact on the performance
	  could be observed.

	*/
	//
	//    Example - enable checkbox filtering support.
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [ {
	//        field: "country",
	//        filterable: {
	//            multi:true
	//        }
	//      } ],
	//    filterable: true,
	//      dataSource: [ { country: "BG" }, { country: "USA" } ]
	//    });
	//    </script>
	//
	Multi Boolean `jsObject:"multi"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.datasource

	  The dataSource configuration for the items that will be used when columns.filterable.multi
	  [ https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.multi ] is enabled.

	*/
	//
	//    Example - provide custom DataSource for the checkbox filtering.
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [ {
	//        field: "country",
	//        filterable: {
	//            multi:true,
	//            dataSource: [{country: "BG"},{country: "GRM"}, {country: "USA"}]
	//        }
	//      } ],
	//    filterable: true,
	//      dataSource: [ { country: "BG" }, { country: "USA" } ]
	//    });
	//    </script>
	//
	DataSource interface{} `jsObject:"dataSource" jsType:"KendoDataSource,string,map[string]interface {},[]string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.checkall

	  Controls whether to show or not the checkAll checkbox before the other checkboxes when using checkbox filtering.
	  (default: true)

	*/
	//
	//    Example - provide custom DataSource for the FilterMultiCheck filtering.
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [ {
	//        field: "country",
	//        filterable: {
	//        multi:true,
	//          checkAll: false
	//        }
	//      }],
	//    filterable: true,
	//      dataSource: [ { country: "BG" }, { country: "USA" } ]
	//    });
	//    </script>
	//
	CheckAll Boolean `jsObject:"checkAll"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.itemtemplate

	  Allows customization on the logic that renders the checkboxes when using checkbox filtering.

	*/
	//
	//    Example - provide custom DataSource for the FilterMultiCheck filtering.
	//
	//    <div id="grid"></div>
	//    <script>
	//        $("#grid").kendoGrid({
	//            columns: [ {
	//                field: "country",
	//                filterable: {
	//                    multi:true,
	//                    itemTemplate: function(e) {
	//                        return "<span><label><span>#= data.country|| data.all #</span><input type='checkbox' name='" + e.field + "' value='#= data.country#'/></label></span>"
	//                    }
	//                }
	//            }],
	//            filterable: true,
	//            dataSource: [ { country: "BG" }, { country: "USA" } ]
	//        });
	//    </script>
	//
	ItemTemplate interface{} `jsObject:"itemTemplate" jsType:"JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.operators

	  The property is identical to filterable.operators
	  [ https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.operators ], but is used for a
	  specific column.

	*/
	Operators KendoGridFilterableEnums `jsObject:"operators"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.search

	  Controls whether to show a search box when checkbox filtering is enabled
	  [ https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.multi ].
	  (default: false)

	*/
	//
	//    Example - Enable checkbox filter search
	//
	//    <div id="grid"></div>
	//    <script>
	//        $("#grid").kendoGrid({
	//            columns: [{
	//                field: "country",
	//                filterable: {
	//                    multi: true,
	//                    search: true
	//                }
	//            }],
	//            filterable: true,
	//            dataSource: [{ country: "BG" }, { country: "USA" }]
	//        });
	//    </script>
	//
	Search Boolean `jsObject:"search"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.ignorecase

	  Toggles between case-insensitive (default) and case-sensitive searching.
	  [ https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.search ]

	  (default: true)

	*/
	//
	//    Example - Enable checkbox filter search
	//
	//    <div id="grid"></div>
	//    <script>
	//        $("#grid").kendoGrid({
	//            columns: [{
	//                field: "country",
	//                filterable: {
	//                    multi: true,
	//                    search: true,
	//                    ignoreCase: true
	//                }
	//            }],
	//            filterable: true,
	//            dataSource: [{ country: "BG" }, { country: "USA" }]
	//        });
	//    </script>
	//
	IgnoreCase Boolean `jsObject:"ignoreCase"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.ui

	  The role data attribute of the widget used in the filter menu or a JavaScript function which initializes that widget.

	  Important:

	  This feature is not supported for columns which have their <a href="columns.values">values</a> option set and Boolean
	  columns.

	  If filterable.mode [ https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable.mode ] is set
	  to 'row', columns.filterable.cell.template
	  [ https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable.cell.template ] should be
	  used to customize the input.

	*/
	//
	//    Example - specify the filter UI as a string
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [ {
	//        field: "date",
	//        filterable: {
	//          ui: "datetimepicker" // use Kendo UI DateTimePicker
	//        }
	//      } ],
	//      filterable: true,
	//      dataSource: [ { date: new Date() }, { date: new Date() } ]
	//    });
	//    </script>
	//
	//
	//    Example - initialize the filter UI
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [ {
	//        field: "date",
	//        filterable: {
	//          ui: function(element) {
	//            element.kendoDateTimePicker(); // initialize a Kendo UI DateTimePicker
	//          }
	//        }
	//      } ],
	//        filterable: true,
	//        dataSource: [ { date: new Date() }, { date: new Date() } ]
	//    });
	//    </script>
	//
	Ui JavaScript `jsObject:"ui"`

	*ToJavaScriptConverter
}

func (el *KendoGridColumnsFilterable) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridColumnsFilterable.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
