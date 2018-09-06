package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUiGrid struct {
	Html HtmlElementDiv `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/allowcopy

	  If set to `true´ and selection of the Grid is enabled the user could copy the selection into the clipboard and paste
	  it into Excel or other similar programs that understand TSV/CSV formats. By default allowCopy is disabled and the
	  default format is TSV.

	  Can be set to a JavaScript object which represents the allowCopy configuration. (default: false)

	*/
	//
	//    Example - enable allowCopy
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//        selectable: "multiple cell",
	//        allowCopy: true,
	//        columns: [
	//            { field: "productName" },
	//            { field: "category" }
	//        ],
	//        dataSource: [
	//            { productName: "Tea", category: "Beverages" },
	//            { productName: "Coffee", category: "Beverages" },
	//            { productName: "Ham", category: "Food" },
	//            { productName: "Bread", category: "Food" }
	//        ]
	//    });
	//    </script>
	//
	AllowCopy interface{} `jsObject:"allowCopy" jsType:"Boolean,KendoGridAllowCopy"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/altrowtemplate

	  The template [ https://docs.telerik.com/kendo-ui/api/javascript/kendo/methods/template ] which renders the alternating
	  table rows. Be default the grid renders a table row (`<tr>´) for every data source item.

	  Important:

	  The outermost HTML element in the template must be a table row (`<tr>´). That table row must have the `uid´ data
	  attribute set to `#= uid #´. The grid uses the `uid´ data attribute to determine the data to which a table row is
	  bound to.

	  Set the `class´ of the table row to `k-alt´ to get the default "alternating" look and feel.

	*/
	//
	//    Example - specify alternating row template as a function
	//
	//    <div id="grid"></div>
	//    <script id="alt-template" type="text/x-kendo-template">
	//        <tr data-uid="#= uid #">
	//            <td colspan="2">
	//                <strong>#: name #</strong>
	//                <strong>#: age #</strong>
	//            </td>
	//        </tr>
	//    </script>
	//    <script>
	//    $("#grid").kendoGrid({
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      altRowTemplate: kendo.template($("#alt-template").html())
	//    });
	//    </script>
	//
	//
	//    Example - specify alternating row template as a string
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      dataSource: [ { name: "Jane Doe", age: 30 }, { name: "John Doe", age: 33 } ],
	//      altRowTemplate: '<tr data-uid="#= uid #"><td colspan="2"><strong>#: name #</strong><strong>#: age #</strong></td></tr>'
	//    });
	//    </script>
	//
	AltRowTemplate interface{} `jsObject:"altRowTemplate" jsType:"JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/autobind

	  If set to `false´, the Grid will not bind to the data source during initialization, i.e. it will not call the fetch
	  [ https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/methods/fetch ] method of the dataSource
	  [ https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/fields/datasource ] instance. In such scenarios data
	  binding will occur when the change [ https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/events/change ]
	  event of the dataSource instance is fired. By default, `autoBind´ is set to `true´ and the widget will bind to the
	  data source specified in the configuration. (default: true)

	  Important:

	  Setting `autoBind´ to `false´ is useful when multiple widgets are bound to the same data source. Disabling automatic
	  binding ensures that the shared data source doesn't make more than one request to the remote service.

	*/
	//
	//    Example - disable automatic binding
	//
	//    <div id="grid"></div>
	//    <script>
	//    var dataSource = new kendo.data.DataSource({
	//      data: [ { name: "Jane Doe" }, { name: "John Doe" }]
	//    });
	//    $("#grid").kendoGrid({
	//      autoBind: false,
	//      dataSource: dataSource
	//    });
	//    dataSource.read(); // "read()" will fire the "change" event of the dataSource and the widget will be bound
	//    </script>
	//
	AutoBind Boolean `jsObject:"autoBind"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columnmenu

	  If set to `true´ the grid will display the column menu when the user clicks the chevron icon in the column headers.
	  The column menu allows the user to show and hide columns, filter and sort (if filtering and sorting are enabled).

	  By default the column menu is not enabled.

	  Can be set to a JavaScript object which represents the column menu configuration. (default: false)

	*/
	//
	//    Example - enable the column menu
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      columnMenu: true,
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ]
	//    });
	//    </script>
	//
	ColumnMenu interface{} `jsObject:"columnMenu" jsType:"Boolean,KendoGridColumnMenu"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columnresizehandlewidth

	  Defines the width of the column resize handle in pixels. Apply a larger value for easier grasping. (default: 3)

	*/
	//
	//    <div id="grid"></div>
	//    <script>
	//      var dataSource = new kendo.data.DataSource({
	//        data: [ { name: "Jane Doe", age: 11 }, { name: "John Doe", age: 12 }]
	//      });
	//      $("#grid").kendoGrid({
	//        columnResizeHandleWidth:20,
	//        dataSource: dataSource,
	//        resizable:true
	//      });
	//    </script>
	//
	ColumnResizeHandleWidth int `jsObject:"columnResizeHandleWidth"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns

	  The configuration of the grid columns. An array of JavaScript objects or strings. JavaScript objects are interpreted
	  as column configurations. Strings are interpreted as the columns.field to which the column is bound. The grid will
	  create a column for every item of the array.

	  Important:

	  If this setting is not specified the grid will create a column for every field of the data item.

	*/
	//
	//    Example - specify grid columns as array of strings
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: ["name", "age"], // two columns bound to the "name" and "age" fields
	//      dataSource: [ { name: "Jane", age: 31 }, { name: "John", age: 33 }]
	//    });
	//    </script>
	//
	//
	//    Example - specify grid columns as array of objects
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [{
	//        field: "name",// create a column bound to the "name" field
	//        title: "Name" // set its title to "Name"
	//      }, {
	//        field: "age",// create a column bound to the "age" field
	//        title: "Age" // set its title to "Age"
	//      }],
	//      dataSource: [ { name: "Jane", age: 30 }, { name: "John", age: 33 }]
	//    });
	//    </script>
	//
	Columns interface{} `jsObject:"columns" jsType:"[]KendoGridColumns,[]string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/datasource

	  The data source of the widget which is used render table rows. Can be a JavaScript object which represents a valid
	  data source configuration, a JavaScript array or an existing kendo.data.DataSource
	  [ https://docs.telerik.com/kendo-ui/api/javascript/data/datasource ] instance.

	  If the `dataSource´ option is set to a JavaScript object or array the widget will initialize a
	  new kendo.data.DataSource [ https://docs.telerik.com/kendo-ui/api/javascript/data/datasource ] instance using that
	  value as data source configuration.

	  If the `dataSource´ option is an existing kendo.data.DataSource
	  [ https://docs.telerik.com/kendo-ui/api/javascript/data/datasource ] instance the widget will use that instance and
	  will not initialize a new one.

	*/
	//
	//    Example - set dataSource as a JavaScript object
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
	//        ]
	//      }
	//    });
	//    </script>
	//
	//
	//    Example - set dataSource as a JavaScript array
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ]
	//    });
	//    </script>
	//
	//
	//    Example - set dataSource as an existing kendo.data.DataSource instance
	//
	//    <div id="grid"></div>
	//    <script>
	//    var dataSource = new kendo.data.DataSource({
	//      transport: {
	//        read: {
	//          url: "https://demos.telerik.com/kendo-ui/service/products",
	//          dataType: "jsonp"
	//        }
	//      },
	//      pageSize: 10
	//    });
	//    $("#grid").kendoGrid({
	//      dataSource: dataSource,
	//      pageable: true
	//    });
	//    </script>
	//
	DataSource interface{} `jsObject:"dataSource" jsType:"*KendoDataSource,string,*map[string]interface {},[]string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/detailtemplate

	  The template [ https://docs.telerik.com/kendo-ui/api/javascript/kendo/methods/template ] which renders the detail
	  rows.

	  Check [ http://demos.telerik.com/kendo-ui/grid/detailtemplate ] Detail Template for a live demo.

	  Important:

	  The detail template content cannot be wider than the total width of all master columns, unless the detail template is
	  scrollable.

	*/
	//
	//    Example - specify detail template as a function
	//
	//    <script id="detail-template" type="text/x-kendo-template">
	//      <div>
	//        Name: #: name #
	//      </div>
	//      <div>
	//        Age: #: age #
	//      </div>
	//    </script>
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      detailTemplate: kendo.template($("#detail-template").html())
	//    });
	//    </script>
	//
	//
	//    Example - specify detail template as a string
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      detailTemplate: "<div>Name: #: name #</div><div>Age: #: age #</div>"
	//    });
	//    </script>
	//
	DetailTemplate interface{} `jsObject:"detailTemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/editable

	  If set to `true´ the user would be able to edit the data to which the grid is bound. By default editing is disabled.

	  Can be set to a string ("inline", "incell" or "popup") to specify the editing mode. The default editing mode is
	  "incell".

	  Can be set to a JavaScript object which represents the editing configuration. (default: false)

	  Important:

	  The "inline" and "popup" editing modes are triggered by the "edit" column command. Thus it is required to have a
	  column with an "edit" command.

	  The "incell" editing mode combined with DataSource `autoSync: true´ setting is not supported when using server-side
	  grouping in the Grid. To be able to save edited values on each change, you can disable server-side grouping or trigger
	  a DataSource `sync()´ manually inside the cellClose event
	  [ https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/events/cellclose ].

	*/
	//
	//    Example - enable editing
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        toolbar: ["save"],
	//        columns: [
	//          { field: "name" },
	//          { field: "age" }
	//        ],
	//        dataSource: {
	//         data: [
	//          { id: 1, name: "Jane Doe", age: 30 },
	//          { id: 2, name: "John Doe", age: 33 }
	//         ],
	//         schema:{
	//          model: {
	//           id: "id",
	//           fields: {
	//             age: { type: "number"}
	//           }
	//          }
	//         }
	//        },
	//        editable: true
	//      });
	//    </script>
	//
	//
	//    Example - enable popup editing
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" },
	//        { command: "edit" }
	//      ],
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
	//      },
	//      editable: "popup"
	//    });
	//    </script>
	//
	Editable interface{} `jsObject:"editable" jsType:"Boolean,KendoGridEditorMode,KendoGridEditable,JavaScript"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/excel

	  Configures the Kendo UI Grid Excel export settings.
	*/
	Excel KendoGridExcel `jsObject:"excel"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/filterable

	  If set to `true´ the user can filter the data source using the grid filter menu. Filtering is disabled by default.

	  Can be set to a JavaScript object which represents the filter menu configuration. (default: false)
	*/
	//
	//    Example - enable filtering
	//
	//        <div id="grid"></div>
	//        <script>
	//            $("#grid").kendoGrid({
	//              columns: [
	//                  { field: "name" },
	//                  { field: "age" }
	//              ],
	//              filterable: true,
	//              dataSource: {
	//               data: [
	//                { id: 1, name: "Jane Doe", age: 30 },
	//                { id: 2, name: "John Doe", age: 33 }
	//               ],
	//               schema:{
	//                model: {
	//                 id: "id",
	//                 fields: {
	//                   age: { type: "number"}
	//                 }
	//                }
	//               }
	//              }
	//          });
	//        </script>
	//
	Filterable interface{} `jsObject:"filterable" jsType:"Boolean,KendoGridFilterable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/groupable

	  If set to `true´ the user could group the grid by dragging the column header cells. By default grouping is disabled.

	  Can be set to a JavaScript object which represents the grouping configuration. (default: false)

	*/
	//
	//    Example - enable grouping
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "productName" },
	//        { field: "category" }
	//      ],
	//      dataSource: [
	//        { productName: "Tea", category: "Beverages" },
	//        { productName: "Coffee", category: "Beverages" },
	//        { productName: "Ham", category: "Food" },
	//        { productName: "Bread", category: "Food" }
	//      ],
	//      groupable: true
	//    });
	//    </script>
	//
	Groupable interface{} `jsObject:"groupable" jsType:"Boolean,*KendoGridGroupable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/height

	  The height of the grid. Numeric values are treated as pixels.

	*/
	//
	//    Example - set the height as a number
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "productName" },
	//        { field: "category" }
	//      ],
	//      dataSource: [
	//        { productName: "Tea", category: "Beverages" },
	//        { productName: "Coffee", category: "Beverages" },
	//        { productName: "Ham", category: "Food" },
	//        { productName: "Bread", category: "Food" }
	//      ],
	//      height: 100
	//    });
	//    </script>
	//
	//
	//    Example - set the height as a string
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "productName" },
	//        { field: "category" }
	//      ],
	//      dataSource: [
	//        { productName: "Tea", category: "Beverages" },
	//        { productName: "Coffee", category: "Beverages" },
	//        { productName: "Ham", category: "Food" },
	//        { productName: "Bread", category: "Food" }
	//      ],
	//      height: "10em"
	//    });
	//    </script>
	//
	Height int `jsObject:"height"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/messages

	  Defines the text of the command buttons that are shown within the Grid. Used primarily for localization.

	*/
	//
	//    Example - change the messages
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" },
	//        { command: ["edit", "destroy"] }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      toolbar: ["create", "save", "cancel"],
	//      messages: {
	//        commands: {
	//          cancel: "Cancel changes",
	//          canceledit: "Cancel",
	//          create: "Add new record",
	//          destroy: "Delete",
	//          edit: "Edit",
	//          save: "Save changes",
	//          select: "Select",
	//          update: "Update"
	//        }
	//      }
	//    });
	//    </script>
	//
	Messages KendoGridMessages `jsObject:"messages"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/mobile

	  If set to `true´ and the grid is viewed on mobile browser it will use adaptive rendering.

	  Can be set to a string `phone´ or `tablet´ which will force the widget to use adaptive rendering regardless of browser
	  type.

	  The grid uses same layout for both `phone´ and `tablet´. (default: false)

	*/
	//
	//    Example - enable adaptive rendering auto detect
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//       columns: [
	//         { field: "name" },
	//         { field: "age" },
	//         { command: "destroy" }
	//       ],
	//       dataSource: [
	//         { name: "Jane Doe", age: 30 },
	//         { name: "John Doe", age: 33 }
	//       ],
	//       filterable: true,
	//       columnMenu: true,
	//       mobile: true
	//    });
	//    </script>
	//
	//
	//    Example - force adaptive rendering
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//       columns: [
	//         { field: "name" },
	//         { field: "age" },
	//         { command: "destroy" }
	//       ],
	//       dataSource: [
	//         { name: "Jane Doe", age: 30 },
	//         { name: "John Doe", age: 33 }
	//       ],
	//       filterable: true,
	//       columnMenu: true,
	//       mobile: "phone"
	//    });
	//    </script>
	//
	Mobile interface{} `jsObject:"mobile" jsType:"Boolean,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/navigatable

	  If set to `true´ the use could navigate the widget using the keyboard navigation. By default keyboard navigation is
	  disabled. (default: false)

	*/
	//
	//    Example - enable keyboard navigation
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      navigatable: true,
	//      selectable: true
	//    });
	//    </script>
	//
	Navigatable Boolean `jsObject:"navigatable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/norecords

	  If set to `true´ and current view contains no records, message similar to "No records available" will be displayed.
	  By default this option is disabled. (default: false)

	*/
	//
	//    Example - enable noRecords message
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      noRecords: true,
	//      dataSource: []
	//    });
	//    </script>
	//
	NoRecords interface{} `jsObject:"noRecords" jsType:"Boolean,KendoGridNoRecords"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pageable

	  If set to `true´ the grid will display a pager. By default paging is disabled.

	  Can be set to a JavaScript object which represents the pager configuration. (default: false)

	  Important:

	  Don't forget to set a pageSize
	  [ https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/pagesize ], no matter if paging is
	  performed client-side or server-side. A `pageSize´ can be defined in the `pageable´ settings, or in the dataSource
	  [ https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/datasource ] settings. If an already existing
	  datasource instance is passed to the grid, then the pagesize
	  [ https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/pagesize ] option should be set in
	  the dataSource's settings and not in the `pageable´ settings.

	*/
	//
	//    Example - enable paging
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "productName" },
	//        { field: "category" }
	//      ],
	//      dataSource: [
	//        { productName: "Tea", category: "Beverages" },
	//        { productName: "Coffee", category: "Beverages" },
	//        { productName: "Ham", category: "Food" },
	//        { productName: "Bread", category: "Food" }
	//      ],
	//      pageable: {
	//        pageSize: 2
	//      }
	//    });
	//    </script>
	//
	Pageable interface{} `jsObject:"pageable" jsType:"Boolean,KendoGridPageable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/pdf

	  Configures the Kendo UI Grid PDF export settings.

	*/
	Pdf KendoGridPdf `jsObject:"pdf"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/persistselection

	  Sets a value indicating whether the selection will be persisted when sorting, paging, filtering and etc are performed.
	  (default: false)

	  Important:

	  Selection persistence works only for row selection.

	  In order for selection persistence to work correctly, you need to define an ID field in schema.model
	  [ https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/schema.model ].

	  Selection persistence does not work for new items when the Grid DataSource is in offline mode. In offline mode, newly
	  added items do not have IDs, which are required for selection persistence to work.

	*/
	//
	//    Example - enables selection persistence
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
	//          { id: 1, name: "Jane Doe", age: 30 },
	//          { id: 2, name: "John Doe", age: 33},
	//          { id: 3, name: "Jim Doe", age: 30 },
	//          { id: 4, name: "Jack Doe", age: 33}
	//        ],
	//        schema: {
	//          model: { id: "id" }
	//        }
	//      },
	//      pageable: {
	//        pageSize: 2
	//      },
	//      selectable: "multiple, row",
	//      persistSelection: true
	//    });
	//    </script>
	//
	PersistSelection Boolean `jsObject:"persistSelection"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/reorderable

	  If set to `true´ the user could reorder the columns by dragging their header cells. By default reordering is disabled.

	  Multi-level headers allow reordering only in same level. (default: false)

	*/
	//
	//    Example - enable column reordering
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      reorderable: true
	//    });
	//    </script>
	//
	Reorderable Boolean `jsObject:"reorderable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/resizable

	  If set to `true´, users can resize columns by dragging the edges (resize handles) of their header cells. As of Kendo
	  UI Q1 2015, users can also auto-fit a column by double-clicking its resize handle. In this case the column will assume
	  the smallest possible width, which allows the column content to fit without wrapping.

	  By default, column resizing is disabled. (default: false)

	*/
	//
	//    Example - enable column resizing
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      resizable: true
	//    });
	//    </script>
	//
	Resizable Boolean `jsObject:"resizable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/rowtemplate

	  The template [ https://docs.telerik.com/kendo-ui/api/javascript/kendo/methods/template ] which renders rows. Be
	  default renders a table row (`<tr>´) for every data source item.

	  Important:

	  The outermost HTML element in the template must be a table row (`<tr>´). That table row must have the `uid´ data
	  attribute set to `#= uid #´. The grid uses the `uid´ data attribute to determine the data to which a table row is
	  bound to.

	*/
	//
	//    Example - specify row template as a function
	//
	//    <div id="grid"></div>
	//    <script id="template" type="text/x-kendo-template">
	//        <tr data-uid="#= uid #">
	//            <td colspan="2">
	//                <strong>#: name #</strong>
	//                <strong>#: age #</strong>
	//            </td>
	//        </tr>
	//    </script>
	//    <script>
	//    $("#grid").kendoGrid({
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      rowTemplate: kendo.template($("#template").html())
	//    });
	//    </script>
	//
	//
	//    Example - specify row template as a string
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      dataSource: [ { name: "Jane Doe", age: 30 }, { name: "John Doe", age: 33 } ],
	//      rowTemplate: '<tr data-uid="#= uid #"><td colspan="2"><strong>#: name #</strong><strong>#: age #</strong></td></tr>'
	//    });
	//    </script>
	//
	RowTemplate interface{} `jsObject:"rowTemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/scrollable

	  If set to `true´ the grid will display a scrollbar when the total row height (or width) exceeds the grid height (or
	  width). By default scrolling is enabled.

	  Can be set to a JavaScript object which represents the scrolling configuration. (default: true)

	*/
	//
	//    Example - disable scrolling
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      scrollable: false
	//    });
	//    </script>
	//
	Scrollable interface{} `jsObject:"scrollable" jsType:"Boolean,*KendoGridScrollable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/selectable

	  If set to `true´ the user would be able to select grid rows. By default selection is disabled. (default: false)

	  Can also be set to the following string values:

	  > "row" - the user can select a single row.</li>
	  > "cell" - the user can select a single cell.</li>
	  > "multiple, row" - the user can select multiple rows.</li>
	  > "multiple, cell" - the user can select multiple cells.</li>

	  Important:

	  When the selectable property is set to "multiple, row" or "multiple, cell" the Grid cannot be scrollable on mobile
	  devices as both are listening on the same event.

	*/
	//
	//    Example - set selectable as a boolean
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      selectable: true
	//    });
	//    </script>
	//
	//
	//    Example - set selectable as a string
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      selectable: "multiple, row"
	//    });
	//    </script>
	//
	Selectable interface{} `jsObject:"selectable" jsType:"Boolean,KendoGridSelectable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/sortable

	  If set to `true´ the user could sort the grid by clicking the column header cells. By default sorting is disabled.

	  Can be set to a JavaScript object which represents the sorting configuration. (default: false)
	*/
	//
	//    Example - enable sorting
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      sortable: true
	//    });
	//    </script>
	//
	Sortable interface{} `jsObject:"sortable" jsType:"Boolean,*KendoGridSortable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/toolbar

	  If a `String´ value is assigned to the `toolbar´ configuration option, it will be treated as a single string template
	  for the whole grid Toolbar, and the string value will be passed as an argument to a kendo.template()
	  [ https://docs.telerik.com/kendo-ui/api/javascript/kendo/methods/template ] function.

	  If a `Function´ value is assigned (it may be a kendo.template() function call or a generic function reference), then
	  the return value of the function will be used to render the Grid Toolbar contents.

	  Important:

	  If the grid is instantiated with MVVM, The template passed will not be bound to the grid view model context. You may
	  bind the toolbar element manually afterwards, using `kendo.bind(gridWidgetInstance.element.find("k-grid-toolbar"))´

	  Observation:

	  If an `Array´ value is assigned, it will be treated as the list of commands displayed in the grid's Toolbar. Commands
	  can be custom or built-in ("cancel", "create", "save", "excel", "pdf").

	  The "cancel" built-in command reverts any data changes done by the end user.

	  The "create" command adds an empty data item to the grid.

	  The "save" command persists any data changes done by the end user.

	  The "excel" command exports the grid data in MS Excel format.

	  The "pdf" command exports the grid data in PDF format.

	*/
	//
	//    Example - configure the Grid Toolbar as a string template
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      toolbar: "<p>My string template in a paragraph.</p>",
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: {
	//        data: [
	//          { id: 1, name: "Jane Doe", age: 30 },
	//          { id: 2, name: "John Doe", age: 33},
	//        ],
	//        schema: {
	//          model: { id: "id" }
	//        }
	//      },
	//      editable: true
	//    });
	//    </script>
	//
	//
	//    Example - configure the Grid Toolbar template with a function
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      toolbar: kendo.template("<p>My function template.</p>"),
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: {
	//        data: [
	//          { id: 1, name: "Jane Doe", age: 30 },
	//          { id: 2, name: "John Doe", age: 33},
	//        ],
	//        schema: {
	//          model: { id: "id" }
	//        }
	//      },
	//      editable: true
	//    });
	//    </script>
	//
	//
	//    Example - configure the Grid Toolbar as an array of commands
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      toolbar: [
	//        { name: "create" },
	//        { name: "save" },
	//        { name: "cancel" }
	//      ],
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: {
	//        data: [
	//          { id: 1, name: "Jane Doe", age: 30 },
	//          { id: 2, name: "John Doe", age: 33},
	//        ],
	//        schema: {
	//          model: { id: "id" }
	//        }
	//      },
	//      editable: true
	//    });
	//    </script>
	//
	Toolbar interface{} `jsObject:"toolbar"  jsType:"*JavaScript,[]string,string,[]KendoGridToolbar"`

	*ToJavaScriptConverter
}

func (el *KendoUiGrid) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoUiGrid.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoGrid({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiGrid) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiGrid) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiGrid) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
