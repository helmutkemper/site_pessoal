package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUiMultiSelect struct {
	Html HtmlElementFormSelect `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog

	  This kendo dialog has been placed here to create a dialog window when you need to add more data to the data source
	*/
	Dialog KendoUiDialog `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/animation

	  Configures the opening and closing animations of the suggestion popup. Setting the <b><u>animation</u></b> option to <b><u>false</u></b> will disable the opening and closing animations. As a result the suggestion popup will open and close instantly.
	  <b><u>animation:true</u></b> is not a valid configuration.

	  Example - disable open and close animations
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     animation: false
	   });
	   </script>
	*/
	Animation interface{} `jsObject:"animation" jsType:"*KendoAnimation,Boolean"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/autobind

	  Controls whether to bind the widget to the data source on initialization. (default: true)

	  Example
	  <select id="multiselect" multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	  </select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	      autoBind: false
	  });
	  </script>
	*/
	AutoBind Boolean `jsObject:"autoBind"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/autoclose

	  Controls whether to close the widget suggestion list on item selection. (default: true)

	  Example
	  <select id="multiselect" multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	  </select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	      autoClose: false
	  });
	  </script>
	*/
	AutoClose Boolean `jsObject:"autoClose"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/autowidth

	  If set to true, the widget automatically adjusts the width of the popup element and does not wrap up the item label.
	  Note: Virtualized list doesn't support the auto-width functionality.

	  Example - enable autoWidth
	  <select id="multiselect" style="width: 100px;"></select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    autoWidth: true,
	    dataSource: {
	      data: ["Short item", "An item with really, really long text"]
	    }
	  });
	  </script>
	*/
	AutoWidth Boolean `jsObject:"autoWidth"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/clearbutton

	  Unless this options is set to false, a button will appear when hovering the widget. Clicking that button will reset the widget's value and will trigger the change event. (default: true)

	  Example - disable the clear button
	  <select id="multiselect" multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	  </select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	      clearButton: false
	  });
	  </script>
	*/
	ClearButton Boolean `jsObject:"clearButton"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/datasource

	  The data source of the widget which is used to display suggestions for the current value. Can be a JavaScript object which represents a valid data source configuration, a JavaScript array or an existing <a href="/kendo-ui/api/javascript/data/datasource">kendo.data.DataSource</a> instance.
	  If the <b><u>dataSource</u></b> option is set to a JavaScript object or array the widget will initialize a new <a href="/kendo-ui/api/javascript/data/datasource">kendo.data.DataSource</a> instance using that value as data source configuration.
	  If the <b><u>dataSource</u></b> option is an existing <a href="/kendo-ui/api/javascript/data/datasource">kendo.data.DataSource</a> instance the widget will use that instance and will <strong>not</strong> initialize a new one.

	  Example - set dataSource as a JavaScript object
	  <select id="multiselect" multiple="multiple"></select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    dataSource: {
	      data: ["One", "Two"]
	    }
	  });
	  </script>

	  Example - set dataSource as a JavaScript array
	  <select id="multiselect" multiple="multiple"></select>
	  <script>
	  var data = ["One", "Two"];
	  $("#multiselect").kendoMultiSelect({
	    dataSource: data
	  });
	  </script>

	  Example - set dataSource as an existing kendo.data.DataSource instance
	  <select id="multiselect" multiple="multiple"></select>
	  <script>
	  var dataSource = new kendo.data.DataSource({
	    transport: {
	      read: {
	        url: "https://demos.telerik.com/kendo-ui/service/products",
	        dataType: "jsonp"
	      }
	    }
	  });
	  $("#multiselect").kendoMultiSelect({
	    dataSource: dataSource,
	    dataTextField: "ProductName",
	    dataValueField: "ProductID"
	  });
	  </script>
	*/
	DataSource interface{} `jsObject:"dataSource" jsType:"*KendoDataSource,string,*map[string]interface {},[]string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/datatextfield

	  The field of the data item that provides the text content of the list items. The widget will filter the data source based on this field. (default: "")
	  Important: When <b>dataTextField</b> is defined, the <b>dataValueField</b> option also should be set.

	  Example - set the dataTextField
	  <select id="multiselect" multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	  </select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	      dataSource: [
	          { Name: "Parent1", Id: 1 },
	          { Name: "Parent2", Id: 2 }
	      ],
	      dataTextField: "Name",
	      dataValueField: "Id"
	  });
	  </script>
	*/
	DataTextField string `jsObject:"dataTextField"`

	/*
	  @sse https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/datavaluefield

	  The field of the data item that provides the value of the widget. (default: "")
	  Important When <b>dataValueField</b> is defined, the <b>dataTextField<b> option also should be set.

	  Example - set the dataValueField
	  <select id="multiselect" multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	  </select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	      dataSource: [
	          { Name: "Parent1", Id: 1 },
	          { Name: "Parent2", Id: 2 }
	      ],
	      dataTextField: "Name",
	      dataValueField: "Id"
	  });
	  </script>
	*/
	DataValueField string `jsObject:"dataValueField"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/delay

	  Specifies the delay in milliseconds after which the MultiSelect will start filtering dataSource. (default: 200)

	  Example - set the delay
	  <select id="multiselect" multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	  </select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	      delay: 1000 // wait 1 second before clearing the user input
	  });
	  </script>
	*/
	Delay int `jsObject:"delay"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/enable

	  If set to <b>false</b> the widget will be disabled and will not allow user input. The widget is enabled by default and allows user input. (default: true)

	  Example - disable the widget
	  <select id="multiselect" multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	  </select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    enable: false
	  });
	  </script>
	*/
	Enable Boolean `jsObject:"enable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/enforceminlength

	  If set to <b>true</b> the widget will not show all items when the text of the search input cleared. By default the widget shows all items when the text of the search input is cleared. Works in conjunction with minLength.

	  Example - enforce minLength
	  <select id="multiselect"></select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	      placeholder: "Select products...",
	      dataTextField: "ProductName",
	      dataValueField: "ProductID",
	      autoBind: false,
	      minLength: 3,
	      enforceMinLength: true,
	      dataSource: {
	          type: "odata",
	          serverFiltering: true,
	          transport: {
	              read: {
	                  url: "//demos.telerik.com/kendo-ui/service/Northwind.svc/Products",
	              }
	          }
	      },
	      value: [
	          { ProductName: "Chang", ProductID: 2 },
	          { ProductName: "Uncle Bob's Organic Dried Pears", ProductID: 7 }
	      ]
	  });
	  </script>
	*/
	EnforceMinLength Boolean `jsObject:"enforceMinLength"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/filter

	  The filtering method used to determine the suggestions for the current value. Filtration is turned of by default, and can be performed over <b>string</b> values only (either the widget's data has to be an array of strings, or over the field, configured in the dataTextField option). The supported filter values are <b>startswith</b>, <b>endswith</b> and <b>contains</b>. (default: "startswith")

	  Example - set the filter
	  <select id="multiselect" multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	  </select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    filter: "contains"
	  });
	  </script>
	*/
	Filter KendoFilter `jsObject:"filter"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/fixedgrouptemplate

	  The template used to render the fixed header group. By default the widget displays only the value of the current group.

	  Example
	  <select id="customers" style="width: 400px;"></select>
	  <script>
	      $(document).ready(function() {
	          $("#customers").kendoMultiSelect({
	              placeholder: "Select customers...",
	              dataTextField: "ContactName",
	              dataValueField: "CustomerID",
	              fixedGroupTemplate: "Fixed header: #: data #",
	              height: 400,
	              dataSource: {
	                  type: "odata",
	                  transport: {
	                      read: "https://demos.telerik.com/kendo-ui/service/Northwind.svc/Customers"
	                  },
	                  group: { field: "Country" }
	              }
	          });
	      });
	  </script>
	*/
	FixedGroupTemplate interface{} `jsObject:"fixedGroupTemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/footertemplate

	  The template used to render the footer template. The footer template receives the widget itself as a part of the data argument. Use the widget fields directly in the template.

	  Parameters:
	  instance Object - The widget instance.

	  Example - specify footerTemplate as a string
	  <select id="customers" style="width: 400px;"></select>
	  <script>
	  $("#customers").kendoMultiSelect({
	    dataSource: [
	      { id: 1, name: "Apples" },
	      { id: 2, name: "Oranges" }
	    ],
	    dataTextField: "name",
	    dataValueField: "id",
	    footerTemplate: 'Total <strong>#: instance.dataSource.total() #</strong> items found'
	  });
	  </script>
	*/
	FooterTemplate interface{} `jsObject:"footerTemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/grouptemplate

	  The template used to render the groups. By default the widget displays only the value of the group.

	  Example
	  <select id="customers" style="width: 400px;"></select>
	  <script>
	      $(document).ready(function() {
	          $("#customers").kendoMultiSelect({
	              placeholder: "Select customers...",
	              dataTextField: "ContactName",
	              dataValueField: "CustomerID",
	              groupTemplate: "Group template: #: data #",
	              height: 400,
	              dataSource: {
	                  type: "odata",
	                  transport: {
	                      read: "https://demos.telerik.com/kendo-ui/service/Northwind.svc/Customers"
	                  },
	                  group: { field: "Country" }
	              }
	          });
	      });
	  </script>
	*/
	GroupTemplate interface{} `jsObject:"groupTemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/height

	  The height of the suggestion popup in pixels. The default value is 200 pixels. (default: 200)

	  Example - set the height
	  <select id="multiselect" multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	  </select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    height: 500
	  });
	  </script>
	*/
	Height int `jsObject:"height"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/highlightfirst

	  If set to <b>true</b> the first suggestion will be automatically highlighted. (default: true)

	  Example - set highlightFirst
	  <select id="multiselect" multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	  </select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    highlightFirst: false
	  });
	  </script>
	*/
	HighlightFirst Boolean `jsObject:"highlightFirst"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/ignorecase

	  If set to <b>false</b> case-sensitive search will be performed to find suggestions. The widget performs case-insensitive searching by default. (default: true)

	  Example - disable case-insensitive suggestions
	  <select id="multiselect" multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	  </select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    ignoreCase: false
	  });
	  </script>
	*/
	IgnoreCase Boolean `jsObject:"ignoreCase"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/minlength

	  The minimum number of characters the user must type before a search is performed. Set to a higher value if the search could match a lot of items. A zero value means that a request will be made as soon as the user focuses the widget. (default: 1)
	  Widget will initiate a request when input value is cleared. If you would like to prevent this behavior please check the filtering event for more details.

	  Example - set minLength
	  <select id="multiselect" multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	  </select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    minLength: 3
	  });
	  </script>
	*/
	MinLength int `jsObject:"minLength"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/maxselecteditems

	  Defines the limit of the selected items. If set to null widget will not limit number of the selected items.

	  Example
	  <select id="multiselect" multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	    <option>Item3</option>
	    <option>Item4</option>
	  </select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	      maxSelectedItems: 3 //only three or less items could be selected
	  });
	  </script>
	*/
	MaxSelectedItems int `jsObject:"maxSelectedItems"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/nodatatemplate

	  The template used to render the "no data" template, which will be displayed if no results are found or the underlying data source is empty. The noData template receives the widget itself as a part of the data argument. The template will be evaluated on every widget data bound.  (default: "NO DATA FOUND.")
	  Important The popup will open when 'noDataTemplate' is defined

	  Example - specify headerTemplate as a string
	  <select id="multiselect"></select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    dataSource: [
	      { id: 1, name: "Apples" },
	      { id: 2, name: "Oranges" }
	    ],
	    dataTextField: "name",
	    dataValueField: "id",
	    noDataTemplate: 'No Data!'
	  });
	  </script>
	*/
	NoDataTemplate interface{} `jsObject:"noDataTemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/placeholder

	  The hint displayed by the widget when it is empty. Not set by default. (default: "")

	  Example - specify placeholder option
	  <select id="multiselect" multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	  </select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    placeholder: "Select..."
	  });
	  </script>

	  Example - specify placeholder as HTML attribute
	  <select id="multiselect" data-placeholder="Select..." multiple="multiple">
	    <option>Item1</option>
	    <option>Item2</option>
	  </select>

	  <script>
	  $("#multiselect").kendoMultiSelect();
	  </script>
	*/
	Placeholder string `jsObject:"placeholder"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/popup

	  The options that will be used for the popup initialization. For more details about the available options refer to Popup documentation.

	  Example - append the popup to a specific element
	  <div id="container">
	    <select id="multiselect"></select>
	  </div>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    dataSource: [
	      { id: 1, name: "Apples" },
	      { id: 2, name: "Oranges" }
	    ],
	    dataTextField: "name",
	    dataValueField: "id",
	    popup: {
	      appendTo: $("#container")
	    }
	  });
	  </script>
	*/
	Popup KendoPopup `jsObject:"popup"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/headertemplate

	  Specifies a static HTML content, which will be rendered as a header of the popup element.

	  Important: The header content should be wrapped with a HTML tag if it contains more than one element. This is applicable also when header content is just a string/text.
	  Important: Widget does not pass a model data to the header template. Use this option only with static HTML.

	  Example - specify headerTemplate as a string
	  <select id="multiselect"></select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    dataSource: [
	      { id: 1, name: "Apples" },
	      { id: 2, name: "Oranges" }
	    ],
	    dataTextField: "name",
	    dataValueField: "id",
	    headerTemplate: '<div><h2>Fruits</h2></div>'
	  });
	  </script>
	*/
	HeaderTemplate interface{} `jsObject:"headerTemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/itemtemplate

	  The template used to render the items in the popup list.

	  Example - specify template as a function
	  <select id="multiselect" multiple="multiple"></select>
	  <script id="itemTemplate" type="text/x-kendo-template">
	    <span>
	      <img src="/img/#: id #.png" alt="#: name #" />
	      #: name #
	    </span>
	  </script>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    dataSource: [
	      { id: 1, name: "Apples" },
	      { id: 2, name: "Oranges" }
	    ],
	    dataTextField: "name",
	    dataValueField: "id",
	    itemTemplate: kendo.template($("#itemTemplate").html())
	  });
	  </script>

	  Example - specify template as a string
	  <select id="multiselect" multiple="multiple"></select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    dataSource: [
	      { id: 1, name: "Apples" },
	      { id: 2, name: "Oranges" }
	    ],
	    dataTextField: "name",
	    dataValueField: "id",
	    itemTemplate: '<span><img src="/img/#: id #.png" alt="#: name #" />#: name #</span>'
	  });
	  </script>
	*/
	ItemTemplate interface{} `jsObject:"itemTemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/tagtemplate

	  The template used to render the tags.

	  Template Data for the 'multiple' tagMode
	  data Object
	  The dataitem that corresponds to the selected value.

	  Template Data for the 'single' tagMode
	  data.values Array
	  A list of the selected values.

	  data.dataItems Array
	  A list of the selected data items.

	  data.currentTotal Array
	  The current dataSource total value. If it is server filtered, it will show the current length of the view.

	  data.maxTotal Array
	  The maximum total value of the dataSource. Unlike the currentTotal, this value will keep the maximum reached total value. Usable when the tag shows the total of the available items.

	  Example - specify template as a function
	  <select id="multiselect" multiple="multiple"></select>
	  <script id="tagTemplate" type="text/x-kendo-template">
	    <span>
	      <img src="/img/#: id #.png" alt="#: name #" />
	      #: name #
	    </span>
	  </script>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    dataSource: [
	      { id: 1, name: "Apples" },
	      { id: 2, name: "Oranges" }
	    ],
	    dataTextField: "name",
	    dataValueField: "id",
	    tagTemplate: kendo.template($("#tagTemplate").html())
	  });
	  </script>

	  Example - specify template as a string
	  <select id="multiselect" multiple="multiple"></select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    dataSource: [
	      { id: 1, name: "Apples" },
	      { id: 2, name: "Oranges" }
	    ],
	    dataTextField: "name",
	    dataValueField: "id",
	    tagTemplate: '<span><img src="/img/#: id #.png" alt="#: name #" />#: name #</span>'
	  });
	  </script>

	  Example - specify a template displaying the number of the selected values
	  <select id="multiselect" multiple="multiple"></select>
	  <script id="tagTemplate" type="text/x-kendo-template">
	      # if (values.length < 3) { #
	          # for (var idx = 0; idx < values.length; idx++) { #
	              #:values[idx]#
	              # if (idx < values.length - 1) {#, # } #
	          # } #
	      # } else { #
	         #:values.length# out of #:maxTotal#
	      # } #
	  </script>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    dataSource: [
	      { id: 1, name: "Apples" },
	      { id: 2, name: "Oranges" },
	      { id: 2, name: "Bananas" }
	    ],
	    dataTextField: "name",
	    dataValueField: "id",
	    tagTemplate: kendo.template($("#tagTemplate").html()),
	    tagMode: "single"
	  });
	  </script>
	*/
	TagTemplate interface{} `jsObject:"tagTemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/tagmode

	  The mode used to render the selected tags. The available modes are: (default: "multiple")

	  <b>multiple</b> - renders a tag for every selected value
	  <b>single</b> - renders only one tag that shows the number of the selected values

	  Every tagMode has a specific tagTemplate value. If you would like to control the content of the rendered tags, set a custom a tagTemplate value.
	*/
	TagMode KendoTagMode `jsObject:"tagMode"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/value

	  Define the value of the widget (default: [])

	  Example
	  <select id="multiselect" multiple="multiple"></select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	       dataSource: ["Item1", "Item2", "Item3", "Item4"],
	       value: ["Item2", "Item3"]
	  });
	  </script>

	  Important: Define a list of data items if widget is not initially bound

	  Example
	  <select id="multiselect" multiple="multiple"></select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	      autoBind: false,
	      dataTextField: "productName",
	      dataValueField: "productId",
	      value: [{ productName: "Item 1", productId: "1" }, { productName: "Item 2", productId: "2" }]
	  });
	  </script>
	*/
	Value []map[string]interface{} `jsObject:"value"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/valueprimitive

	  Specifies the value binding behavior for the widget. If set to true, the View-Model field will be updated with the selected item value field. If set to false, the View-Model field will be updated with the selected item. (default: false)

	  Example - specify that the View-Model field should be updated with the selected item value
	  <select id="multiselect" multiple="multiple" data-bind="value: selectedProductId, source: products"></select>
	  <script>
	  $("#multiselect").kendoMultiSelect({
	    valuePrimitive: true,
	    dataTextField: "name",
	    dataValueField: "id"
	  });
	  var viewModel = kendo.observable({
	    selectedProductId: [],
	    products: [
	      { id: 1, name: "Coffee" },
	      { id: 2, name: "Tea" },
	      { id: 3, name: "Juice" }
	    ]
	  });

	  kendo.bind($("#multiselect"), viewModel);
	  </script>
	*/
	ValuePrimitive Boolean `jsObject:"valuePrimitive"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/multiselect/configuration/virtual

	  Enables the virtualization feature of the widget. The configuration can be set on an object, which contains two properties - <b>itemHeight</b> and <b>valueMapper</b>.

	  For detailed information, refer to the article on virtualization.

	  Example - MultiSelect with a virtualized list
	  <select id="orders" style="width: 400px;"></select>
	  <script>
	      $(document).ready(function() {
	          $("#orders").kendoMultiSelect({
	              placeholder: "Select addresses...",
	              dataTextField: "ShipName",
	              dataValueField: "OrderID",
	              height: 520,
	              virtual: {
	                  itemHeight: 26,
	                  valueMapper: function(options) {
	                      $.ajax({
	                          url: "https://demos.telerik.com/kendo-ui/service/Orders/ValueMapper",
	                          type: "GET",
	                          dataType: "jsonp",
	                          data: convertValues(options.value),
	                          success: function (data) {
	                              options.success(data);
	                          }
	                      })
	                  }
	              },
	              dataSource: {
	                  type: "odata",
	                  transport: {
	                      read: "https://demos.telerik.com/kendo-ui/service/Northwind.svc/Orders"
	                  },
	                  schema: {
	                      model: {
	                          fields: {
	                              OrderID: { type: "number" },
	                              Freight: { type: "number" },
	                              ShipName: { type: "string" },
	                              OrderDate: { type: "date" },
	                              ShipCity: { type: "string" }
	                          }
	                      }
	                  },
	                  pageSize: 80,
	                  serverPaging: true,
	                  serverFiltering: true
	              }
	          });
	      });

	      function convertValues(value) {
	          var data = {};

	          value = $.isArray(value) ? value : [value];

	          for (var idx = 0; idx < value.length; idx++) {
	              data["values[" + idx + "]"] = value[idx];
	          }

	          return data;
	      }
	  </script>

	  Example - MultiSelect widget with a declarative virtualization configuration
	  <div class="demo-section k-header">
	    <h4>Search for shipping name</h4>
	    <select id="orders" style="width: 400px"
	           data-role="multiselect"
	           data-bind="value: order, source: source"
	           data-text-field="ShipName"
	           data-value-field="OrderID"
	           data-filter="contains"
	           data-virtual="{itemHeight:26,valueMapper:orderValueMapper}"
	           data-height="520"
	            ></select>
	  </div>

	  <script>
	      $(document).ready(function() {
	          var model = kendo.observable({
	                  order: [10548],
	            source: new kendo.data.DataSource({
	              type: "odata",
	              transport: {
	                read: "https://demos.telerik.com/kendo-ui/service/Northwind.svc/Orders"
	              },
	              schema: {
	                model: {
	                  fields: {
	                    OrderID: { type: "number" },
	                    Freight: { type: "number" },
	                    ShipName: { type: "string" },
	                    OrderDate: { type: "date" },
	                    ShipCity: { type: "string" }
	                  }
	                }
	              },
	              pageSize: 80,
	              serverPaging: true,
	              serverFiltering: true
	            })
	          });


	          kendo.bind($(document.body), model);
	      });

	      function orderValueMapper(options) {
	          $.ajax({
	            url: "https://demos.telerik.com/kendo-ui/service/Orders/ValueMapper",
	            type: "GET",
	            dataType: "jsonp",
	            data: convertValues(options.value),
	            success: function (data) {
	              options.success(data);
	            }
	          })
	      }

	      function convertValues(value) {
	          var data = {};

	          value = $.isArray(value) ? value : [value];

	          for (var idx = 0; idx < value.length; idx++) {
	              data["values[" + idx + "]"] = value[idx];
	          }

	          return data;
	      }
	  </script>
	*/
	Virtual interface{} `jsObject:"virtual" jsType:"*KendoVirtual,Boolean"`

	*ToJavaScriptConverter
}

func (el *KendoUiMultiSelect) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	if reflect.DeepEqual(el.DataSource, KendoDataSource{}) == false && reflect.DeepEqual(el.Dialog, KendoUiDialog{}) == false {

		for k := range el.Dialog.Actions {

			if el.Dialog.Actions[k].ButtonType == BUTTON_TYPE_ADD_AND_CLOSE {

				el.Dialog.Actions[k].Action = JavaScript{
					Code: string(el.GetId()) + "AddAndCloseButton",
				}

			} else if el.Dialog.Actions[k].ButtonType == BUTTON_TYPE_ADD {

				el.Dialog.Actions[k].Action = JavaScript{
					Code: string(el.GetId()) + "AddButton",
				}

			}

		}

		/*if el.NoDataTemplate != nil {
		  switch el.NoDataTemplate.(type) {
		  case HtmlElementScript:

		    for k := range el.NoDataTemplate.(HtmlElementScript).Content {

		      switch el.NoDataTemplate.(HtmlElementScript).Content[ k ].(type) {
		      case *HtmlElementFormButton:

		        if el.NoDataTemplate.(HtmlElementScript).Content[ k ].(*HtmlElementFormButton).ButtonType == BUTTON_TYPE_ADD_IN_TEMPLATE {
		          //el.NoDataTemplate.(HtmlElementScript).Content[ k ].(HtmlElementFormButton).Global.OnClick = "addNewItemToKendoDataSource('id:#: instance.element[0].id #')"

		          //reflect.ValueOf(&el.NoDataTemplate.(HtmlElementScript).Content[ k ].(*HtmlElementFormButton).Global).FieldByName("OnClick").SetString("esta vivo")
		          p := reflect.ValueOf(&el.NoDataTemplate.(HtmlElementScript).Content[ k ].(*HtmlElementFormButton).Global.OnClick)
		          p.Elem().SetString("esta vivo!!!!!!!!")
		          //fmt.Printf( "------%v\n\n\n\n\n\n\n\n", p.Elem().CanSet() )
		          fmt.Printf( "------%v--------\n\n\n\n\n\n\n\n", el.NoDataTemplate.(HtmlElementScript).Content[ k ].(*HtmlElementFormButton).Global.OnClick )
		          //el.NoDataTemplate.(HtmlElementScript).Content[ k ].(HtmlElementFormButton).Global.OnClick = "esta vivo"
		        }

		      }

		    }

		  }
		}*/

	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("kendoMultiSelect.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoMultiSelect({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiMultiSelect) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiMultiSelect) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiMultiSelect) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
