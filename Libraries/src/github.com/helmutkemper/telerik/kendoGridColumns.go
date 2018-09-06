package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridColumns struct {
	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.aggregates
	//
	// The aggregate(s) which are calculated when the grid is grouped by the columns field. The supported aggregates are
	// "average", "count", "max", "min" and "sum".
	//
	//    Example - set column aggregates
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "firstName", groupable: false },
	//          { field: "lastName" }, /* group by this column to see the footer template */
	//          { field: "age", groupable: false,
	//             aggregates: [ "count", "min", "max" ],
	//             groupFooterTemplate: "age total: #: count #, min: #: min #, max: #: max #"
	//          }
	//        ],
	//        groupable: true,
	//        dataSource: {
	//          data: [
	//            { firstName: "Jane", lastName: "Doe", age: 30 },
	//            { firstName: "John", lastName: "Doe", age: 33 }
	//          ]
	//        }
	//      });
	//    </script>
	Aggregates []KendoAggregate `jsObject:"aggregates"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.attributes
	//
	// HTML attributes of the table cell (<td>) rendered for the column.
	//
	// Important:
	//
	// HTML attributes which are JavaScript keywords (e.g. class) must be quoted.
	//
	//    Example - specify column HTML attributes
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [ {
	//          field: "name",
	//          title: "Name",
	//          attributes: {
	//            "class": "table-cell",
	//            "style": "text-align: right; font-size: 14px"
	//          }
	//        } ],
	//        dataSource: [ { name: "Jane Doe" }, { name: "John Doe" }]
	//      });
	//    </script>
	//
	//  Observation:
	//
	// The table cells would look like this: <td class="table-cell" style="text-align: right; font-size: 14px">...</td>.
	Attributes map[string]interface{} `jsObject:"attributes"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.columns
	//
	// The columns which should be rendered as child columns under this group column header.
	//
	// Important:
	//
	// Note that group column cannot be data bound and supports limited number of bound column settings - such as title,
	// headerTemplate, locked
	//
	//    Example - set column group column for displaying multicolumn headers
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          {
	//            title: "Personal Info",
	//            columns: [
	//              { field: "name" },
	//              { field: "birthdate" }
	//            ]
	//          },
	//          {
	//            title: "Location",
	//            columns: [
	//              { field: "city" },
	//              { field: "country" }
	//            ]
	//          },
	//          {
	//            field: "phone"
	//          }
	//        ],
	//        editable: true,
	//        dataSource: [ { name: "Jane Doe", birthdate: new Date("1995/05/04"), city: "London", country: "UK", phone: "555-444-333" } ]
	//      });
	//    </script>
	Columns []KendoColumnsFields `jsObject:"columns"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.command#columns.command
	//
	// The configuration of the column command(s). If set the column would display a button for every command. Commands
	// can be custom or built-in ("edit" or "destroy").
	//
	// The "edit" built-in command switches the current table row in edit mode.
	//
	// The "destroy" built-in command removes the data item to which the current table row is bound.
	//
	// Custom commands are supported by specifying the click option.
	//
	// Important:
	//
	// The built-in "edit" and "destroy" commands work only if editing is enabled via the editable option. The "edit"
	// command supports "inline" and "popup" editing modes.
	//
	//    Example - set command as a string
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { command: "destroy" } // displays the built-in "destroy" command
	//        ],
	//        editable: true,
	//        dataSource: [ { name: "Jane Doe" } ]
	//      });
	//    </script>
	//
	//
	//    Example - set command as array of strings
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { command: ["edit", "destroy"] } // displays the built-in "edit" and "destroy" commands
	//        ],
	//        editable: "inline",
	//        dataSource: [ { name: "Jane Doe" } ]
	//      });
	//    </script>
	Command interface{} `jsObject:"command" jsType:"[]TypeKendoGridColumnsCommand,[]KendoGridColumnsCommand"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.editable#columns.editable
	//
	// The JavaScript function executed when the cell/row is about to be opened for edit. The result returned will
	// determine whether an editor for the column will be created.
	//
	//    Example - conditionally edit a cell
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          {
	//            field: "salary",
	//            editable: function (dataItem) {
	//              return dataItem.name === "Jane";
	//            }
	//          }
	//        ],
	//        editable: true,
	//        dataSource: [ { name: "Jane", salary: 2000 }, { name: "Bill", salary: 2000 } ]
	//      });
	//    </script>
	Editable JavaScript `jsObject:"Editable"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.editor#columns.editor
	//
	// Provides a way to specify a custom editing UI for the column. Use the container parameter to create the editing UI.
	//
	// Important:
	//
	// The editing UI should contain an element whose name HTML attribute is set as the column field.
	//
	// Validation settings defined in the model.fields configuration will not be applied automatically. In order the
	// validation to work, the developer is responsible for attaching the corresponding validation attributes to the
	// editor input the data-bind attribute is whitespace sensitive. In case the custom editor is a widget, the developer
	// should customize the validation warning tooltip position in order to avoid visual issues.
	//
	// Parameters:
	//
	// container jQuery
	// The jQuery object representing the container element.
	//
	// options Object
	// options.field String
	// The name of the field to which the column is bound.
	//
	// options.format String
	// The format string of the column specified via the format option.
	//
	// options.model kendo.data.Model
	// The model instance to which the current table row is bound.
	//
	// options.values Array
	// Array of values specified via the values option.
	//
	//    Example - create a custom column editor using the Kendo UI AutoComplete
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [ {
	//          field: "name",
	//          editor: function(container, options) {
	//            // create an input element
	//            var input = $("<input/>");
	//            // set its name to the field to which the column is bound ('name' in this case)
	//            input.attr("name", options.field);
	//            // append it to the container
	//            input.appendTo(container);
	//            // initialize a Kendo UI AutoComplete
	//            input.kendoAutoComplete({
	//              dataTextField: "name",
	//              dataSource: [
	//                { name: "Jane Doe" },
	//                { name: "John Doe" }
	//              ]
	//            });
	//          }
	//        } ],
	//        editable: true,
	//        dataSource: [ { name: "Jane Doe" }, { name: "John Doe" } ]
	//      });
	//    </script>
	//
	//
	//    Example - create a custom column editor with validation
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [ {
	//          field: "num",
	//          editor: function(container, options) {
	//            //create input element and add the validation attribute
	//            var input = $('<input name="' + options.field + '" required="required" />');
	//            //append the editor
	//            input.appendTo(container);
	//            //enhance the input into NumericTextBox
	//            input.kendoNumericTextBox();
	//
	//            //create tooltipElement element, NOTE: data-for attribute should match editor's name attribute
	//            var tooltipElement = $('<span class="k-invalid-msg" data-for="' + options.field + '"></span>');
	//            //append the tooltip element
	//            tooltipElement.appendTo(container);
	//          }
	//        } ],
	//        editable: true,
	//        scrollable: false,
	//        dataSource: {
	//          data: [ { num: 1 }, { num: 2 } ],
	//          schema: {
	//            model: {
	//              fields: {
	//                num: { type: "number", validation: { required: true } }
	//              }
	//            }
	//          }
	//        }
	//      });
	//    </script>
	Editor JavaScript `jsObject:"editor"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.encoded#columns.encoded
	//
	// If set to true the column value will be HTML-encoded before it is displayed. If set to false the column value will
	// be displayed as is. By default the column value is HTML-encoded.
	//
	//    Example - prevent HTML encoding
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name", encoded: false }
	//        ],
	//        dataSource: [ { name: "<strong>Jane Doe</strong>" } ]
	//      });
	//    </script>
	Encoded Boolean `jsObject:"encoded"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.field
	//
	// The field to which the column is bound. The value of this field is displayed in the column's cells during data
	// binding. Only columns that are bound to a field can be sortable or filterable. The field name should be a valid
	// Javascript identifier and should contain only alphanumeric characters (or "$" or "_"), and may not start with a
	// digit.
	//
	//    Example - specify the column field
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          // create a column bound to the "name" field
	//          { field: "name" },
	//          // create a column bound to the "age" field
	//          { field: "age" }
	//        ],
	//        dataSource: [ { name: "Jane", age: 30 }, { name: "John", age: 33 }]
	//      });
	//    </script>
	Field string `jsObject:"field"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.filterable
	//
	// If set to true a filter menu will be displayed for this column when filtering is enabled. If set to false the
	// filter menu will not be displayed. By default a filter menu is displayed for all columns when filtering is enabled
	// via the filterable option. (default: true)
	//
	// Can be set to a JavaScript object which represents the filter menu configuration.
	//
	//    Example - disable filtering
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name", filterable: false },
	//          { field: "age" }
	//        ],
	//        filterable: true,
	//        dataSource: [ { name: "Jane", age: 30 }, { name: "John", age: 33 }]
	//      });
	//    </script>
	Filterable KendoGridColumnsFilterable `jsObject:"filterable"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.footerattributes
	//
	// HTML attributes of the column footer. The footerAttributes option can be used to set the HTML attributes of that cell.
	//
	// Important:
	//
	// HTML attributes which are JavaScript keywords (e.g. class) must be quoted.
	//
	//    Example - set the column footer HTML attributes
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age",
	//            footerTemplate: "Min: #: min # Max: #: max #",
	//            footerAttributes: {
	//              "class": "table-footer-cell",
	//              style: "text-align: right; font-size: 14px"
	//            }
	//          }
	//        ],
	//        dataSource: {
	//          data: [
	//            { name: "Jane Doe", age: 30 },
	//            { name: "John Doe", age: 33 }
	//          ],
	//          aggregate: [
	//            { field: "age", aggregate: "min" },
	//            { field: "age", aggregate: "max" }
	//          ]
	//        }
	//      });
	//    </script>
	//
	// Observation:
	//
	// The table footer cell will look like this:
	//
	// <td class="table-footer-cell" style="text-align: right; font-size: 14px">Min: 30 Max: 33</td>.
	FooterAttributes map[string]interface{} `jsObject:"footerattributes"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.footertemplate
	//
	// The template which renders the footer table cell for the column.
	//
	// The fields which can be used in the template are:
	//
	// average - the value of the "average" aggregate (if specified)
	// count - the value of the "count" aggregate (if specified)
	// max - the value of the "max" aggregate (if specified)
	// min - the value of the "min" aggregate (if specified)
	// sum - the value of the "sum" aggregate (if specified)
	// data - provides access to all available aggregates, e.g. data.fieldName1.sum or data.fieldName2.average
	//
	// Important:
	//
	// If the grid is bound using source binding, it will initially be assigned with an empty dataSource without any
	// aggregates. In order to avoid a JavaScript error for an undefined aggregate when the footer is rendered with the
	// empty dataSource, you should check if the field is defined in the template data before accessing the value. If no
	// groups are specified for the actual dataSource, then you will also need to use the field name to access the
	// aggregate value.
	//
	//    Example - specify column footer template
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age",
	//            footerTemplate: "Min: #: min # Max: #: max #"
	//          }
	//        ],
	//        dataSource: {
	//        data: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ],
	//        aggregate: [
	//          { field: "age", aggregate: "min" },
	//          { field: "age", aggregate: "max" }
	//        ]
	//      }
	//    });
	//    </script>
	//
	//
	//    Example - specify footer template when using source binding
	//
	//    <div data-role="grid" data-bind="source:dataSource"
	//     data-columns='["category", "name", {"field": "price", "footerTemplate": "Total: #: data.price ? sum: 0 #"}]'></div>
	//    <script>
	//      $(function() {
	//        var viewModel = kendo.observable({
	//          dataSource: new kendo.data.DataSource({
	//          data: [
	//            { category: "Beverages", name: "Chai", price: 18 },
	//            { category: "Beverages", name: "Chang", price: 19 },
	//            { category: "Seafood", name: "Konbu", price: 6 }
	//          ],
	//          group: [{field: "category"}],
	//          aggregate: [
	//            { field: "price", aggregate: "sum" }
	//          ]
	//        })
	//      });
	//      kendo.bind($("body"), viewModel);
	//    });
	//</script>
	//
	//
	//    Example - specify footer template when using source binding and there are no groups
	//
	//    <div data-role="grid" data-bind="source:dataSource"
	//     data-columns='["category", "name", {"field": "price", "footerTemplate": "Total: #: data.price ? data.price.sum: 0 #"}]'></div>
	//    <script>
	//      $(function() {
	//        var viewModel = kendo.observable({
	//          dataSource: new kendo.data.DataSource({
	//            data: [
	//              { category: "Beverages", name: "Chai", price: 18 },
	//              { category: "Beverages", name: "Chang", price: 19 },
	//              { category: "Seafood", name: "Konbu", price: 6 }
	//            ],
	//            aggregate: [
	//              { field: "price", aggregate: "sum" }
	//            ]
	//          })
	//        });
	//        kendo.bind($("body"), viewModel);
	//      });
	//    </script>
	FooterTemplate interface{} `jsObject:"footertemplate" jsType:"*JavaScript,string"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.format
	//
	// The format that is applied to the value before it is displayed. Takes the form "{0:format}" where "format" is a
	// standard number format, custom number format, standard date format or a custom date format.
	//
	// @see https://docs.telerik.com/kendo-ui/api/javascript/kendo
	//
	// Important:
	//
	// The kendo.format function is used to format the value.
	//
	//    Example - specify the column format string
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [ {
	//            field: "date",
	//            format: "{0: yyyy-MM-dd HH:mm:ss}"
	//          },
	//          {
	//            field: "number",
	//            format: "{0:c}"
	//          }
	//        ],
	//        filterable: true,
	//        dataSource: [ { date: new Date(), number: 3.1415 } ]
	//      });
	//    </script>
	Format string `jsObject:"format"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.groupable
	//
	// If set to false the column will not be groupable (requires Grid groupable property to be enabled). By default all
	// columns are groupable. (default: true)
	//
	//    Example - disable grouping for individual column
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        groupable: true,
	//        columns: [
	//          { field: "name", groupable: false },
	//          { field: "age"}
	//        ],
	//        dataSource: {
	//          data: [
	//            { name: "Jane Doe", age: 30 },
	//            { name: "John Doe", age: 30 }
	//          ]
	//        }
	//      });
	//    </script>
	Groupable Boolean `jsObject:"groupable"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.groupheadertemplate
	//
	// The template which renders the group header when the grid is grouped by the column field. By default the name of
	// the field and the current group value is displayed.
	//
	// The fields which can be used in the template are:
	//
	// value - the current group value
	// field - the current group field
	// average - the value of the "average" aggregate (if specified)
	// count - the value of the "count" aggregate (if specified)
	// max - the value of the "max" aggregate (if specified)
	// min - the value of the "min" aggregate (if specified)
	// sum - the value of the "sum" aggregate (if specified)
	// aggregates - provides access to all available aggregates, e.g. aggregates.fieldName1.sum or aggregates.fieldName2.average
	// items - the data items for current group. Returns groups if the data items are grouped (in case there are child groups)
	//
	//    Example - set the group header template
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age",
	//            groupHeaderTemplate: "Age: #= value # total: #= count #"
	//          }
	//        ],
	//        dataSource: {
	//          data: [
	//            { name: "Jane Doe", age: 30 },
	//            { name: "John Doe", age: 30 }
	//          ],
	//          group: { field: "age", aggregates: [ { field: "age", aggregate: "count" }] }
	//        }
	//      });
	//    </script>
	//
	//
	//    Example - use items field inside the group header template
	//
	//    <div id="grid"></div>
	//    <script>
	//      var filterAdmins = function(item) {
	//        return item.role === "admin";
	//      };
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age",
	//            groupHeaderTemplate: "Admin count: #=items.filter(filterAdmins).length#"
	//          },
	//          {field: "role" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { name: "Jane Doe", age: 30, role: "admin" },
	//            { name: "John Doe", age: 30, role: "guest" },
	//            { name: "Peter", age: 30, role: "admin" }
	//          ],
	//          group: { field: "age", aggregates: [ { field: "age", aggregate: "count" }] }
	//        }
	//      });
	//    </script>
	GroupHeaderTemplate interface{} `jsObject:"groupheadertemplate" jsType:"*JavaScript,string"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.groupfootertemplate
	//
	// The template which renders the group footer when the grid is grouped by the column field. By default the group
	// footer is not displayed.
	//
	// The fields which can be used in the template are:
	//
	// average - the value of the "average" aggregate (if specified)
	// count - the value of the "count" aggregate (if specified)
	// max - the value of the "max" aggregate (if specified)
	// min - the value of the "min" aggregate (if specified)
	// sum - the value of the "sum" aggregate (if specified)
	// data - provides access to all available aggregates, e.g. data.fieldName1.sum or data.fieldName2.average
	// group - provides information for the current group. An object with three fields - field, value and items. items field contains the data items for current group. Returns groups if the data items are grouped (in case there are child groups)
	//
	// Important:
	//
	// If the template is declared as a function the group field is accessible only through the data field, e.g.
	// data.fieldName1.group.value.
	//
	//    Example - set the group header template
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age",
	//            groupFooterTemplate: "Total: #= count #"
	//          }
	//        ],
	//        dataSource: {
	//          data: [
	//            { name: "Jane Doe", age: 30 },
	//            { name: "John Doe", age: 30 }
	//          ],
	//          group: { field: "age", aggregates: [ { field: "age", aggregate: "count" }] }
	//        }
	//      });
	//    </script>
	//
	//
	//    Example - set the group header template as function
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age",
	//            groupFooterTemplate: function(e) {
	//              return "Total: " + e.age.count;
	//            }
	//          }
	//        ],
	//        dataSource: {
	//          data: [
	//            { name: "Jane Doe", age: 30 },
	//            { name: "John Doe", age: 30 }
	//          ],
	//          group: { field: "age", aggregates: [ { field: "age", aggregate: "count" }] }
	//        }
	//      });
	//    </script>
	GroupFooterTemplate interface{} `jsObject:"groupfootertemplate" jsType:"*JavaScript,string"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.headerattributes
	//
	// HTML attributes of the column header. The grid renders a table header cell (<th>) for every column. The
	// headerAttributes option can be used to set the HTML attributes of that th.
	//
	// Important:
	//
	// HTML attributes which are JavaScript keywords (e.g. class) must be quoted.
	//
	//    Example - set the column header HTML attributes
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [{
	//          field: "name",
	//          headerAttributes: {
	//            "class": "table-header-cell",
	//            style: "text-align: right; font-size: 14px"
	//          }
	//        }],
	//        dataSource: [ { name: "Jane Doe" }, { name: "John Doe" } ]
	//      });
	//    </script>
	//
	// Observation:
	//
	// The table header cell will look like this:
	// <th class="table-header-cell" style="text-align: right; font-size: 14px">name</th>.
	HeaderAttributes map[string]interface{} `jsObject:"headerattributes"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.headertemplate
	//
	// The template which renders the column header content. By default the value of the title column option is displayed
	// in the column header cell.
	//
	// Important:
	//
	// If sorting is enabled, the column header content will be wrapped in a <a> element. As a result the template must
	// contain only inline elements.
	//
	//    Example - column header template as a string
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [{
	//          field: "name",
	//          headerTemplate: '<input type="checkbox" id="check-all" /><label for="check-all">Check All</label>'
	//        }],
	//        dataSource: [ { name: "Jane Doe" }, { name: "John Doe" } ]
	//      });
	//    </script>
	//
	//
	//
	//    Example - column header template as a Kendo UI template function with conditional logic
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [ {
	//          field: "name",
	//          headerTemplate: kendo.template('# if (true) { # <input type="checkbox" id="check-all" /><label for="check-all">Check All</label> # } else { # this will never be displayed # } #')
	//        }],
	//        dataSource: [ { name: "Jane Doe" }, { name: "John Doe" } ]
	//      });
	//    </script>
	HeaderTemplate interface{} `jsObject:"headertemplate" jsType:"*JavaScript,string"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.hidden
	//
	// If set to true the column will not be displayed in the grid. By default all columns are displayed. (default: false)
	//
	//    Example - hide columns
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { hidden: true, field: "id" },
	//          { field: "name" }
	//        ],
	//        dataSource: [ { id: 1, name: "Jane Doe" }, { id: 2, name: "John Doe" } ]
	//      });
	//    </script>
	Hidden Boolean `jsObject:"hidden"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.locked
	//
	// If set to true the column will be displayed as locked (frozen) in the grid. Also see the information about Locked
	// Columns in the Grid Appearance article.
	//
	// Important:
	//
	// Row template and detail features are not supported in combination with column locking. If multi-column
	// headers are used, it is possible to lock (freeze) a column at the topmost level only.
	//
	//    Example - locked columns
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { locked: true, field: "id", width:200 },
	//          { field: "name", width:800 }
	//        ],
	//        dataSource: [ { id: 1, name: "Jane Doe" }, { id: 2, name: "John Doe" } ]
	//      });
	//    </script>
	Locked Boolean `jsObject:"locked"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.lockable
	//
	// If set to false the column will remain in the side of the grid into which its own locked configuration placed it.
	//
	// Important:
	//
	// This option is meaningful when the grid has columns which are configured with a locked value. Setting it explicitly
	// to false will prevent the user from locking or unlocking this column using the user interface.
	//
	//    Example - lockable columns
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { locked: true, field: "id", lockable: false, width:250 },
	//          { locked: true, field: "age", width:250 },
	//          { field: "name", width:250 },
	//          { field: "city", lockable: false, width:250 }
	//        ],
	//        dataSource: [
	//          { id: 1, name: "Jane Doe", age: 31, city: "Boston" },
	//          { id: 2, name: "John Doe", age: 55, city: "New York" }
	//        ]
	//      });
	//    </script>
	Lockable Boolean `jsObject:"lockable"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.minresizablewidth
	//
	// The pixel screen width below which the user will not be able to resize the column via the UI.
	//
	// Important:
	//
	// This option is meaningful when the grid is set as resizable.
	//
	//    Example - set the column width as a number
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        resizable: true,
	//        columns: [
	//          { field: "name", minResizableWidth: 80 },
	//          { field: "age" }
	//        ],
	//        dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	MinResizableWidth int `jsObject:"minresizablewidth"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.minscreenwidth
	//
	// The pixel screen width below which the column will be hidden. The setting takes precedence over the hidden setting,
	// so the two should not be used at the same time.
	//
	//    Example - lockable columns
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "id", width: 250, minScreenWidth: 500 }, //column will become hidden if screen size is less than 500px
	//          { field: "name", width: 250 }, //column will always be visible
	//          { field: "age", width: 250, minScreenWidth: 750 } //column will become hidden if screen size is less than 750px
	//        ],
	//        dataSource: [
	//          { id: 1, name: "Jane Doe", age: 31, city: "Boston" },
	//          { id: 2, name: "John Doe", age: 55, city: "New York" }
	//        ]
	//      });
	//    </script>
	MinScreenWidth int `jsObject:"minscreenwidth"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.selectable
	//
	// If set to true the grid will render a select column with checkboxes in each cell, thus enabling multi-row
	// selection. The header checkbox allows users to select/deselect all the rows on the current page. (default: false)
	//
	// Important:
	//
	// The checkbox column selection functionality has two limitations:
	//
	// It is not integrated with the keyboard navigation
	// It is not integrated with the currently existing select functionality of the grid which is enabled via the
	// selectable option. They are mutually exclusive and we recommend using only one of the approaches for enabling
	// selection.
	//
	//    Example - disable sorting
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { selectable: true },
	//          { field: "name" }
	//        ],
	//        dataSource: [ { name: "Jane Doe" }, { name: "John Doe" } ]
	//      });
	//    </script>
	Selectable Boolean `jsObject:"selectable"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.sortable
	//
	// If set to true the user can click the column header and sort the grid by the column field when sorting is enabled.
	// If set to false sorting will be disabled for this column. By default all columns are sortable if sorting is enabled
	// via the sortable option.
	//
	//    Example - disable sorting
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { sortable: false, field: "id" },
	//          { field: "name" }
	//        ],
	//        sortable: true,
	//        dataSource: [ { id: 1, name: "Jane Doe" }, { id: 2, name: "John Doe" } ]
	//      });
	//    </script>
	Sortable interface{} `jsObject:"sortable" jsType:"*KendoColumnSortable,Boolean"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.template
	//
	// The template which renders the column content. The grid renders table rows (<tr>) which represent the data source
	// items. Each table row consists of table cells (<td>) which represent the grid columns. By default the HTML-encoded
	// value of the field is displayed in the column.
	//
	// Important:
	//
	// Use the template to customize the way the column displays its value.
	//
	//    Example - set the template as a string (wrap the column value in HTML)
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [ {
	//          field: "name",
	//          template: "<strong>#: name # </strong>"
	//        }],
	//        dataSource: [ { name: "Jane Doe" }, { name: "John Doe" } ]
	//      });
	//    </script>
	//
	//
	//
	//    Example - set the template as a function returned by kendo.template
	//
	//    <div id="grid"></div>
	//    <script id="name-template" type="text/x-kendo-template">
	//      <strong>#: name #</strong>
	//    </script>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [ {
	//          field: "name",
	//          template: kendo.template($("#name-template").html())
	//        }],
	//        dataSource: [ { name: "Jane Doe" }, { name: "John Doe" } ]
	//      });
	//    </script>
	//
	//
	//
	//    Example - set the template as a function which returns a string
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [ {
	//          field: "name",
	//          template: function(dataItem) {
	//            return "<strong>" + kendo.htmlEncode(dataItem.name) + "</strong>";
	//          }
	//        }],
	//        dataSource: [ { name: "Jane Doe" }, { name: "John Doe" } ]
	//      });
	//    </script>
	Template interface{} `jsObject:"template" jsType:"*JavaScript,string"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.title
	//
	// The text that is displayed in the column header cell. If not set the field is used.
	//
	// Important:
	//
	// Column titles should not contain HTML entities or tags. If such exist, they should be encoded.
	//
	//    Example - set the title of the column
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [ { field: "name", title: "Name" } ],
	//        dataSource: [ { name: "Jane Doe" }, { name: "John Doe" } ]
	//      });
	//    </script>
	Title string `jsObject:"title"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.width
	//
	// The width of the column. Numeric values are treated as pixels. For more important information, please refer to
	// Column Widths.
	//
	//    Example - set the column width as a string
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name", width: "200px" },
	//          { field: "age" }
	//        ],
	//        dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	//
	//
	//
	//    Example - set the column width as a number
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name", width: 200 },
	//          { field: "age" }
	//        ],
	//        dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	Width int `jsObject:"width"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.values
	//
	// An array of values that will be displayed instead of the bound value. Each item in the array must have a text and value fields.
	//
	// Important:
	//
	// Use the values option to display user-friendly text instead of database values.
	//
	//    Example - specify column values
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "productName" },
	//          { field: "category", values: [
	//            { text: "Beverages", value: 1 },
	//            { text: "Food", value: 2 }
	//          ] }
	//        ],
	//        dataSource: [
	//          { productName: "Tea", category: 1 },
	//          { productName: "Ham", category: 2 }
	//        ]
	//      });
	//    </script>
	//
	// This example displays "Beverages" and "Food" in the "category" column instead of "1" and "2".
	// @see https://demos.telerik.com/kendo-ui/grid/foreignkeycolumn
	Values []map[string]interface{} `jsObject:"values"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.menu
	//
	// If set to true the column will be visible in the grid column menu. By default the column menu includes all
	// data-bound columns (ones that have their field set).
	//
	//    Example - hide a column from the column menu
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "id", menu: false },
	//          { field: "name" },
	//          { field: "age" }
	//        ],
	//        columnMenu: true,
	//        dataSource: [
	//          { id: 1, name: "Jane Doe", age: 30 },
	//          { id: 2, name: "John Doe", age: 33 }
	//        ]
	//      });
	//    </script>
	Menu Boolean `jsObject:"menu"`

	*ToJavaScriptConverter
}

func (el *KendoGridColumns) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridColumns.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
