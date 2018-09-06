package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUiDropDownList struct {
	Html HtmlInputText `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation

	  Configures the opening and closing animations of the suggestion popup. Setting the <b><u>animation</u></b> option to <b><u>false</u></b> will disable the opening and closing animations. As a result the suggestion popup will open and close instantly.
	  <b><u>animation:true</u></b> is not a valid configuration.

	  Example - disable open and close animations
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	     dataSource: {
	       data: ["One", "Two"]
	     },
	     animation: false
	   });
	   </script>
	*/
	Animation interface{} `jsObject:"animation" jsType:"*KendoAnimation,Boolean"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-autoBind

	  Controls whether to bind the widget to the data source on initialization.

	  Example
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	       autoBind: false
	   });
	   </script>
	*/
	AutoBind Boolean `jsObject:"autoBind"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-autoWidth

	  If set to <b><u>true</u></b>, the widget automatically adjusts the width of the popup element and does not wrap up the item label.
	  Note: Virtualized list doesn't support the auto-width functionality.

	  Example - enable autoWidth
	   <input id="dropdownlist" style="width: 100px;" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	     autoWidth: true,
	     dataSource: {
	       data: ["Short item", "An item with really, really long text"]
	     }
	   });
	   </script>
	*/
	AutoWidth Boolean `jsObject:"autoWidth"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-cascadeFrom

	  Use it to set the Id of the parent DropDownList widget. <a href="/kendo-ui/web/dropdownlist/cascading">Help topic showing how cascading functionality works</a>

	  Example
	   <input id="parent" />
	   <input id="child" />
	   <script>
	   $("#parent").kendoDropDownList({
	       dataTextField: "parentName",
	       dataValueField: "parentId",
	       dataSource: [
	           { parentName: "Parent1", parentId: 1 },
	           { parentName: "Parent2", parentId: 2 }
	       ]
	   });

	   $("#child").kendoDropDownList({
	       cascadeFrom: "parent",
	       dataTextField: "childName",
	       dataValueField: "childId",
	       dataSource: [
	           { childName: "Child1", childId: 1, parentId: 1 },
	           { childName: "Child2", childId: 2, parentId: 2 },
	           { childName: "Child3", childId: 3, parentId: 1 },
	           { childName: "Child4", childId: 4, parentId: 2 }
	       ]
	   });
	   </script>
	*/
	CascadeFrom string `jsObject:"cascadeFrom"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-cascadeFromField

	  Defines the field to be used to filter the data source. If not defined the <a href="/kendo-ui/api/javascript/ui/dropdownlist#configuration-dataValueField">parent's dataValueField option will be used</a>. <a href="/kendo-ui/web/dropdownlist/cascading">Help topic showing how cascading functionality works</a>

	  Example
	   <input id="parent" />
	   <input id="child" />
	   <script>
	   $("#parent").kendoDropDownList({
	       dataTextField: "name",
	       dataValueField: "id",
	       dataSource: [
	           { name: "Parent1", id: 1 },
	           { name: "Parent2", id: 2 }
	       ]
	   });

	   $("#child").kendoDropDownList({
	       cascadeFrom: "parent",
	       cascadeFromField: "parentId",
	       dataTextField: "name",
	       dataValueField: "id",
	       dataSource: [
	           { name: "Child1", id: 1, parentId: 1 },
	           { name: "Child2", id: 2, parentId: 2 },
	           { name: "Child3", id: 3, parentId: 1 },
	           { name: "Child4", id: 4, parentId: 2 }
	       ]
	   });
	   </script>
	*/
	CascadeFromField string `jsObject:"cascadeFromField"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-dataSource

	  The data source of the widget which is used to display a list of values. Can be a JavaScript object which represents a valid data source configuration, a JavaScript array or an existing <a href="/kendo-ui/api/javascript/data/datasource">kendo.data.DataSource</a> instance.
	  If the <b><u>dataSource</u></b> option is set to a JavaScript object or array the widget will initialize a new <a href="/kendo-ui/api/javascript/data/datasource">kendo.data.DataSource</a> instance using that value as data source configuration.
	  If the <b><u>dataSource</u></b> option is an existing <a href="/kendo-ui/api/javascript/data/datasource">kendo.data.DataSource</a> instance the widget will use that instance and will <strong>not</strong> initialize a new one.

	  Example - set dataSource as a JavaScript object
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	     dataSource: {
	       data: ["One", "Two"]
	     }
	   });
	   </script>
	*/
	DataSource interface{} `jsObject:"dataSource" jsType:"*KendoDataSource,string,*map[string]interface {},[]string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-dataTextField

	  The field of the data item that provides the text content of the list items. The widget will filter the data source based on this field.
	  When <code>dataTextField</code> is defined, the <code>dataValueField</code> option also should be set.

	  Example - set the dataTextField
	   <input id="dropdownlist" />
	   <script>
	     $("#dropdownlist").kendoDropDownList({
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
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-dataValueField

	  The field of the data item that provides the value of the widget.
	  When <code>dataValueField</code> is defined, the <code>dataTextField</code> option also should be set.

	  Example - set the dataValueField
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
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
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-delay

	  Specifies the delay in milliseconds before the search-text typed by the end user is cleared.

	  Example - set the delay
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	       delay: 1000 // wait 1 second before clearing the user input
	   });
	   </script>
	*/
	Delay int `jsObject:"delay"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-enable

	  If set to <b><u>false</u></b> the widget will be disabled and will not allow user input. The widget is enabled by default and allows user input.

	  Example - disable the widget
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	     enable: false
	   });
	   </script>
	*/
	Enable Boolean `jsObject:"enable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-enforceMinLength

	  If set to <b><u>true</u></b> the widget will not show all items when the text of the search input cleared. By default the widget shows all items when the text of the search input is cleared. Works in conjunction with <a href="https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-minLength">minLength</a>.

	  Example - enforce minLength
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	       filter: "startswith",
	       autoBind: false,
	       minLength: 3,
	       enforceMinLength: true,
	       dataTextField: "ProductName",
	       dataValueField: "ProductID",
	       dataSource: {
	           type: "odata",
	           serverFiltering: true,
	           transport: {
	               read: {
	                   url: "//demos.telerik.com/kendo-ui/service/Northwind.svc/Products",
	               }
	           }
	       }
	   });
	   </script>
	*/
	EnforceMinLength Boolean `jsObject:"enforceMinLength"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-filter

	  The filtering method used to determine the suggestions for the current value. Filtration is turned off by default, and can be performed over <b><u>string</u></b> values only (either the widget's data has to be an array of strings, or over the field, configured in the <a href="https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-dataTextField"><b><u>dataTextField</u></b></a> option). The supported filter values are <b><u>startswith</u></b>, <b><u>endswith</u></b> and <b><u>contains</u></b>.

	  Example - set the filter
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	     dataSource: ["Chai", "Chang", "Tofu"],
	     filter: "contains"
	   });
	   </script>
	*/
	Filter string `jsObject:"filter"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-fixedGroupTemplate

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">template</a> used to render the fixed header group. By default the widget displays only the value of the current group.
	*/
	FixedGroupTemplate interface{} `jsObject:"fixedGroupTemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-footerTemplate

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">template</a> used to render the footer template. The footer template receives the widget itself as a part of the data argument. Use the widget fields directly in the template.
	  The widget instance.

	  Parameters
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
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
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-groupTemplate

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">template</a> used to render the groups. By default the widget displays only the value of the group.
	*/
	GroupTemplate interface{} `jsObject:"groupTemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-height

	  The height of the suggestion popup in pixels. The default value is 200 pixels.

	  Example - set the height
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	     dataTextField: "ProductName",
	     dataValueField: "ProductID",
	     height: 500,
	     dataSource: {
	       transport: {
	         read: {
	           dataType: "jsonp",
	           url: "//demos.telerik.com/kendo-ui/service/Products",
	         }
	       }
	     }
	   });
	   </script>
	*/
	Height int `jsObject:"height"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-ignoreCase

	  If set to <b><u>false</u></b> case-sensitive search will be performed to find suggestions. The widget performs case-insensitive searching by default.

	  Example - disable case-insensitive suggestions
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	     ignoreCase: false
	   });
	   </script>
	*/
	IgnoreCase Boolean `jsObject:"ignoreCase"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-index

	  The index of the initially selected item. The index is <b><u>0</u></b> based.

	  Example - select second item
	   <input id="dropdownlist" />
	   <script>
	   var items = [{ text: "Item 1", value: "1" }, { text: "Item 2", value: "2" }];
	   $("#dropdownlist").kendoDropDownList({
	       dataTextField: "text",
	       dataValueField: "value",
	       dataSource: items,
	       index: 1
	   });
	   </script>
	*/
	Index int `jsObject:"index"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-minLength

	  The minimum number of characters the user must type before a filter is performed. Set to higher value than <b><u>1</u></b> if the search could match a lot of items.
	  Widget will initiate a request when input value is cleared. If you would like to prevent this behavior please check the <a href="#events-filtering">filtering</a> event for more details.

	  Example - set minLength
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	     dataSource: ["Chai", "Chang", "Tofu"],
	     filter: "contains",
	     minLength: 3
	   });
	   </script>
	*/
	MinLength int `jsObject:"minLength"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-noDataTemplate

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">template</a> used to render the "no data" template, which will be displayed if no results are found or the underlying data source is empty. The noData template receives the widget itself as a part of the data argument. The template will be evaluated on every widget data bound.
	  <strong>Important</strong> The popup will open when 'noDataTemplate' is defined

	  Example - specify headerTemplate as a string
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	     dataSource: [],
	     dataTextField: "name",
	     dataValueField: "id",
	     noDataTemplate: 'No Data!'
	   });
	   </script>
	*/
	NoDataTemplate interface{} `jsObject:"noDataTemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-popup

	  The options that will be used for the popup initialization. For more details about the available options refer to <a href="/kendo-ui/api/javascript/ui/popup">Popup</a> documentation.

	  Example - append the popup to a specific element
	   <div id="container">
	       <input id="dropdownlist" />
	   </div>
	   <script>
	   $("#dropdownlist").kendoDropDownList({
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
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-optionLabel

	  Define the text of the default empty item. If the value is an object, then the widget will use it as a valid data item. Note that the optionLabel will not be available if the widget is empty.

	  Example - specify optionLabel as a string
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	       dataSource: ["Apples", "Oranges"],
	       optionLabel: "Select a fruit..."
	   });
	   </script>
	*/
	OptionLabel string `jsObject:"optionLabel"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-optionLabelTemplate

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">template</a> used to render the option label.
	  Define the <a href="/kendo-ui/api/javascript/kendo#configuration-optionLabel">optionLabel</a> as <strong>object</strong> if complex template structure is used
	*/
	OptionLabelTemplate interface{} `jsObject:"optionLabelTemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-headerTemplate

	  Specifies a static HTML content, which will be rendered as a header of the popup element.

	  Example - specify headerTemplate as a string
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
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
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-template

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">template</a> used to render the items. By default the widget displays only the text of the data item (configured via <b><u>dataTextField</u></b>).

	  Example - specify template as a function
	   <input id="dropdownlist" />
	   <script id="template" type="text/x-kendo-template">
	     <span>
	       <img src="/img/#: id #.png" alt="#: name #" />
	       #: name #
	     </span>
	   </script>
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	     dataSource: [
	       { id: 1, name: "Apples" },
	       { id: 2, name: "Oranges" }
	     ],
	     dataTextField: "name",
	     dataValueField: "id",
	     template: kendo.template($("#template").html())
	   });
	   </script>
	*/
	Template interface{} `jsObject:"template" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-valueTemplate

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">valueTemplate</a> used to render the selected value. By default the widget displays only the text of the data item (configured via <b><u>dataTextField</u></b>).

	  Example - specify valueTemplate as a function
	   <input id="dropdownlist" />
	   <script id="valueTemplate" type="text/x-kendo-template">
	       <img src="/img/#: id #.png" alt="#: name #" />
	       #: name #
	   </script>
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	     dataSource: [
	       { id: 1, name: "Apples" },
	       { id: 2, name: "Oranges" }
	     ],
	     dataTextField: "name",
	     dataValueField: "id",
	     valueTemplate: kendo.template($("#valueTemplate").html())
	   });
	   </script>
	*/
	ValueTemplate interface{} `jsObject:"valueTemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-text

	  The text of the widget used when the <b><u>autoBind</u></b> is set to <b><u>false</u></b>.

	  Example - specify text of the widget
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	        autoBind: false,
	        text: "Chai"
	   });
	   </script>
	*/
	Text string `jsObject:"text"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-value

	  The value of the widget.

	  Example - specify value of the widget
	   <input id="dropdownlist" />
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	        dataSource: ["Item1", "Item2"],
	        value: "Item1"
	   });
	   </script>
	*/
	Value string `jsObject:"value"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-valuePrimitive

	  Specifies the <a href="/kendo-ui/framework/mvvm/bindings/value">value binding</a> behavior for the widget when the initial model value is null. If set to true, the View-Model field will be updated with the selected item value field. If set to false, the View-Model field will be updated with the selected item.

	  Example - specify that the View-Model field should be updated with the selected item value
	   <select id="dropdown" data-bind="value: selectedProductId, source: products" >
	   </select>

	   <script>
	   $("#dropdown").kendoDropDownList({
	     valuePrimitive: true,
	     dataTextField: "name",
	     dataValueField: "id",
	     optionLabel: "Select product..."
	   });
	   var viewModel = kendo.observable({
	     selectedProductId: null,
	     products: [
	       { id: 1, name: "Coffee" },
	       { id: 2, name: "Tea" },
	       { id: 3, name: "Juice" }
	     ]
	   });

	   kendo.bind($("#dropdown"), viewModel);
	   </script>
	*/
	ValuePrimitive Boolean `jsObject:"valuePrimitive"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-virtual

	  Enables the virtualization feature of the widget. The configuration can be set on an object, which contains two properties - <b><u>itemHeight</u></b> and <b><u>valueMapper</u></b>.
	  For detailed information, refer to the <a href="/kendo-ui/controls/editors/combobox/virtualization">article on virtualization</a>.
	*/
	Virtual KendoVirtual `jsObject:"virtual"`

	*ToJavaScriptConverter
}

func (el *KendoUiDropDownList) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("kendoDropDownList.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoDropDownList({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiDropDownList) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiDropDownList) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiDropDownList) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
