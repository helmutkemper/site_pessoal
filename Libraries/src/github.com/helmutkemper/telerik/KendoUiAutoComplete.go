package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUiAutoComplete struct {
	Html HtmlInputText `jsObject:"-"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation

	  Configures the opening and closing animations of the suggestion popup. Setting the **animation** option to **false**
	  will disable the opening and closing animations. As a result the suggestion popup will open and close instantly.
	  **animation:true** is not a valid configuration.
	*/
	//  Example - disable open and close animations
	//   <input id="autocomplete" />
	//   <script>
	//   $("#autocomplete").kendoAutoComplete({
	//     animation: false
	//   });
	//   </script>
	//
	Animation interface{} `jsObject:"animation" jsType:"*KendoAnimation,Boolean"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-autoWidth

	  If set to <b><u>true</u></b>, the widget automatically adjusts the width of the popup element and does not wrap up the item label.
	  Note: Virtualized list doesn't support the auto-width functionality.

	  Example - enable autoWidth
	   <input id="autocomplete" style="width: 100px;" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     autoWidth: true,
	     dataSource: {
	       data: ["Short item", "An item with really, really long text"]
	     }
	   });
	   </script>
	*/
	AutoWidth Boolean `jsObject:"autoWidth"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-dataSource

	  The data source of the widget which is used to display suggestions for the current value. Can be a JavaScript object which represents a valid data source configuration, a JavaScript array or an existing <a href="/kendo-ui/api/javascript/data/datasource">kendo.data.DataSource</a> instance.
	  If the <b><u>dataSource</u></b> option is set to a JavaScript object or array the widget will initialize a new <a href="/kendo-ui/api/javascript/data/datasource">kendo.data.DataSource</a> instance using that value as data source configuration.
	  If the <b><u>dataSource</u></b> option is an existing <a href="/kendo-ui/api/javascript/data/datasource">kendo.data.DataSource</a> instance the widget will use that instance and will <strong>not</strong> initialize a new one.

	  Example - set dataSource as a JavaScript object
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     dataSource: {
	       data: ["One", "Two"]
	     }
	   });
	   </script>
	*/
	DataSource interface{} `jsObject:"dataSource" jsType:"*KendoDataSource,string,*map[string]interface {},[]string"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-clearButton

	  Unless this options is set to <b><u>false</u></b>, a button will appear when hovering the widget. Clicking that button will reset the widget's value and will trigger the change event.

	  Example - disable the clear button
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	       clearButton: false
	   });
	   </script>
	*/
	ClearButton Boolean `jsObject:"clearButton"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-dataTextField

	  The field of the data item used when searching for suggestions. This is the text that will be displayed in the list of matched results.

	  Example - set the dataTextField
	   <input id="autocomplete" />
	   <script>
	   var data = [
	     { id: 1, name: "Apples" },
	     { id: 2, name: "Oranges" }
	   ];
	   $("#autocomplete").kendoAutoComplete({
	     dataTextField: "name", // The widget is bound to the "name" field
	     dataSource: data
	   });
	   </script>
	*/
	DataTextField string `jsObject:"dataTextField"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-delay

	  The delay in milliseconds between a keystroke and when the widget displays the suggestion popup.

	  Example - set the delay
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     delay: 500
	   });
	   </script>
	*/
	Delay int `jsObject:"delay"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-enable

	  If set to <b><u>false</u></b> the widget will be disabled and will not allow user input. The widget is enabled by default and allows user input.

	  Example - disable the widget
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     enable: false
	   });
	   </script>
	*/
	Enable Boolean `jsObject:"enable"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-enforceMinLength

	  If set to <b><u>true</u></b> the widget will not show all items when the text of the search input cleared. By default the widget shows all items when the text of the search input is cleared. Works in conjunction with <a href="http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-minLength">minLength</a>.

	  Example - enforce minLength
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	       dataTextField: "ProductName",
	       filter: "contains",
	       minLength: 3,
	       enforceMinLength: true,
	       autoBind: false,
	       dataSource: {
	           type: "odata",
	           serverFiltering: true,
	           transport: {
	               read: "//demos.telerik.com/kendo-ui/service/Northwind.svc/Products"
	           }
	       }
	   });
	   </script>
	*/
	EnforceMinLength Boolean `jsObject:"enforceMinLength"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-filter

	  The filtering method used to determine the suggestions for the current value. The default filter is "startswith" - all data items which begin with the current widget value are displayed in the suggestion popup. The supported <b><u>filter</u></b> values are <b><u>startswith</u></b>, <b><u>endswith</u></b> and <b><u>contains</u></b>.

	  Example - set the filter
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     filter: "contains",
	     dataSource: {
	       data: ["One", "Two"]
	     }
	   });
	   </script>
	*/
	Filter KendoFilter `jsObject:"filter"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-fixedGroupTemplate

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">template</a> used to render the fixed header group. By default the widget displays only the value of the current group.

	  Example
	  <input id="customers" style="width: 400px" />
	  <script>
	      $(document).ready(function() {
	          $("#customers").kendoAutoComplete({
	              dataTextField: "ContactName",
	              fixedGroupTemplate: "Fixed group: #:data#",
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
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-footerTemplate

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">template</a> used to render the footer template. The footer template receives the widget itself as a part of the data argument. Use the widget fields directly in the template.
	  The widget instance.

	  Parameters
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     dataSource: [
	       { id: 1, name: "Apples" },
	       { id: 2, name: "Oranges" }
	     ],
	     dataTextField: "name",
	     footerTemplate: 'Total <strong>#: instance.dataSource.total() #</strong> items found'
	   });
	   </script>
	*/
	FooterTemplate interface{} `jsObject:"footerTemplate" jsType:"*JavaScript,string"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-groupTemplate

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">template</a> used to render the groups. By default the widget displays only the value of the group.

	  Example
	  <input id="customers" style="width: 400px" />
	  <script>
	      $(document).ready(function() {
	          $("#customers").kendoAutoComplete({
	              dataTextField: "ContactName",
	              groupTemplate: "Group: #:data#",
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
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-height

	  The height of the suggestion popup in pixels. The default value is 200 pixels.

	  Example - set the height
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     height: 100,
	     dataSource: {
	       data: ["One", "Two", "Three", "Four", "Five", "Six", "Seven"]
	     }
	   });
	   </script>
	*/
	Height int `jsObject:"height"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-highlightFirst

	  If set to <b><u>true</u></b> the first suggestion will be automatically highlighted.

	  Example - set highlightFirst
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     highlightFirst: true,
	     dataSource: {
	       data: ["One", "Two"]
	     }
	   });
	   </script>
	*/
	HighlightFirst Boolean `jsObject:"highlightFirst"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-ignoreCase

	  If set to <b><u>false</u></b> case-sensitive search will be performed to find suggestions. The widget performs case-insensitive searching by default.

	  Example - disable case-insensitive suggestions
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     ignoreCase: false,
	     dataSource: {
	       data: ["One", "Two"]
	     }
	   });
	   </script>
	*/
	IgnoreCase Boolean `jsObject:"ignoreCase"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-minLength

	  The minimum number of characters the user must type before a search is performed. Set to higher value than <b><u>1</u></b> if the search could match a lot of items.
	  Widget will initiate a request when input value is cleared. If you would like to prevent this behavior please check the <a href="#events-filtering">filtering</a> event for more details.

	  Example - set minLength
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     minLength: 3,
	     placeholder: "Type 'one'",
	     dataSource: {
	       data: ["One", "Two"]
	     }
	   });
	   </script>
	*/
	MinLength int `jsObject:"minLength"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-noDataTemplate

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">template</a> used to render the "no data" template, which will be displayed if no results are found or the underlying data source is empty. The noData template receives the widget itself as a part of the data argument. The template will be evaluated on every widget data bound.
	  <strong>Important</strong> The popup will open when 'noDataTemplate' is defined

	  Example - specify headerTemplate as a string
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     dataSource: [
	       { id: 1, name: "Apples" },
	       { id: 2, name: "Oranges" }
	     ],
	     dataTextField: "name",
	     noDataTemplate: 'No Data!'
	   });
	   </script>
	*/
	NoDataTemplate interface{} `jsObject:"noDataTemplate" jsType:"*JavaScript,string"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-placeholder

	  The hint displayed by the widget when it is empty. Not set by default.
	  The Kendo UI AutoComplete widget could also use the value of the <b><u>placeholder</u></b> HTML attribute as hint.

	  Example - specify placeholder
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     placeholder: "Enter value ..."
	   });
	   </script>
	*/
	Placeholder string `jsObject:"placeholder"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-popup

	  The options that will be used for the popup initialization. For more details about the available options refer to <a href="/kendo-ui/api/javascript/ui/popup">Popup</a> documentation.

	  Example - append the popup to a specific element
	   <div id="container">
	       <input id="autocomplete" />
	   </div>
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     dataSource: [
	       { id: 1, name: "Apples" },
	       { id: 2, name: "Oranges" }
	     ],
	     dataTextField: "name",
	     popup: {
	       appendTo: $("#container")
	     }
	   });
	   </script>
	*/
	Popup KendoPopup `jsObject:"popup"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-separator

	  The character used to separate multiple values. Empty by default.
	  As of Q3 2016 the Autocomplete widget supports multiple separators listed in an array. All separators will be replaced with the first array item, which acts as a default separator.
	  Using the separator option will still bind the primitive stringe value of the input. In case you need to bind multiple data items, you can consider the <a href="/kendo-ui/controls/editors/multiselect/overview">MultiSelect widget</a>.

	  Example - set separator to allow multiple values
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     separator: ", ",
	     dataSource: {
	       data: ["One", "Two"]
	     }
	   });
	   </script>
	*/
	Separator string `jsObject:"separator"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-suggest

	  If set to <b><u>true</u></b> the widget will automatically use the first suggestion as its value.

	  Example - enable automatic suggestion
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     suggest: true,
	     dataSource: {
	       data: ["One", "Two"]
	     }
	   });
	   </script>
	*/
	Suggest Boolean `jsObject:"suggest"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-headerTemplate

	  Specifies a static HTML content, which will be rendered as a header of the popup element.
	  <strong>Important</strong> The header content <strong>should be wrapped</strong> with a HTML tag if it contains more than one element. This is applicable also when header content is just a string/text.
	  <strong>Important</strong> Widget does not pass a model data to the header template. Use this option only with static HTML.

	  Example - specify headerTemplate as a string
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     dataSource: [
	       { id: 1, name: "Apples" },
	       { id: 2, name: "Oranges" }
	     ],
	     dataTextField: "name",
	     headerTemplate: '<div><h2>Fruits</h2></div>'
	   });
	   </script>
	*/
	HeaderTemplate interface{} `jsObject:"headerTemplate" jsType:"*JavaScript,string"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-template

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">template</a> used to render the suggestions. By default the widget displays only the text of the suggestion (configured via <b><u>dataTextField</u></b>).

	  Example - specify template as a function
	   <input id="autocomplete" />
	   <script id="template" type="text/x-kendo-template">
	     <span>
	       <img src="/img/#: id #.png" alt="#: name #" />
	       #: name #
	     </span>
	   </script>
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     dataSource: [
	       { id: 1, name: "Apples" },
	       { id: 2, name: "Oranges" }
	     ],
	     dataTextField: "name",
	     template: kendo.template($("#template").html())
	   });
	   </script>
	*/
	Template interface{} `jsObject:"template" jsType:"*JavaScript,string"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-value

	  The value of the widget.

	  Example - specify value of the widget
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     dataSource: [
	       { id: 1, name: "Apples" },
	       { id: 2, name: "Oranges" }
	     ],
	     dataTextField: "name",
	     value: "Oranges"
	   });
	   </script>
	*/
	Value string `jsObject:"value"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-valuePrimitive

	  Specifies the <a href="/kendo-ui/framework/mvvm/bindings/value">value binding</a> behavior for the widget when the initial model value is null. If set to true, the View-Model field will be updated with the selected item text field. If set to false, the View-Model field will be updated with the selected item.

	  Example - specify that the View-Model field should be updated with the selected item text
	   <input id="autocomplete" data-bind="value: productName, source: products" />

	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     valuePrimitive: true,
	     dataTextField: "name"
	   });
	   var viewModel = kendo.observable({
	     productName: null,
	     products: [
	       { id: 1, name: "Coffee" },
	       { id: 2, name: "Tea" },
	       { id: 3, name: "Juice" }
	     ]
	   });

	   kendo.bind($("#autocomplete"), viewModel);
	   </script>
	*/
	ValuePrimitive Boolean `jsObject:"valuePrimitive"`

	/*
	  @see http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-virtual

	  Enables the virtualization feature of the widget. The configuration can be set on an object, which contains two properties - <b><u>itemHeight</u></b> and <b><u>valueMapper</u></b>.
	  For detailed information, refer to the <a href="/kendo-ui/controls/editors/combobox/virtualization">article on virtualization</a>.

	  Example - AutoComplete with a virtualized list
	   <input id="orders" style="width: 400px" />
	   <script>
	       $(document).ready(function() {
	           $("#orders").kendoAutoComplete({
	               template: "#= OrderID # | For: #= ShipName #, #= ShipCountry #",
	               dataTextField: "ShipName",
	               virtual: true,
	               height: 520,
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
	   </script>
	*/
	Virtual KendoVirtual `jsObject:"virtual"`

	*ToJavaScriptConverter
}

func (el *KendoUiAutoComplete) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("kendoAutoComplete.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoAutoComplete({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiAutoComplete) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiAutoComplete) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiAutoComplete) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
