package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoVirtual struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-virtual.itemHeight  Specifies the height of the virtual item. All items in the virtualized list <strong>must</strong> have the same height. If the developer does not specify one, the framework will automatically set <b><u>itemHeight</u></b> based on the current theme and font size.
	  Example - DropDownList widget with a virtualized list
	   <input id="orders" style="width: 400px" />
	   <script>
	       $(document).ready(function() {
	           $("#orders").kendoDropDownList({
	               template: '<span class="order-id">#= OrderID #</span> #= ShipName #, #= ShipCountry #',
	               dataTextField: "ShipName",
	               dataValueField: "OrderID",
	               filter: "contains",
	               virtual: {
	                   itemHeight: 26,
	                   valueMapper: function(options) {
	                       $.ajax({
	                           url: "https://demos.telerik.com/kendo-ui/service/Orders/ValueMapper",
	                           type: "GET",
	                           dataType: "jsonp",
	                           data: convertValues(options.value),
	                           success: function (data) {
	                               //the **data** is either index or array of indices.
	                               //Example:
	                               // 10258 -> 10 (index in the Orders collection)
	                               // [10258, 10261] -> [10, 14] (indices in the Orders collection)

	                               options.success(data);
	                           }
	                       })
	                   }
	               },
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
	ItemHeight int `jsObject:"itemHeight"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-virtual.mapValueTo  The changes introduced with the Kendo UI R3 2016 release enable you to determine if the <b><u>valueMapper</u></b> must resolve a <em>value to an <b><u>index</u></b></em> or a <em>value to a <b><u>dataItem</u></b></em>. This is configured through the <b><u>mapValueTo</u></b> option that accepts two possible values - <b><u>"index"</u></b> or <b><u>"dataItem"</u></b>. By default, the <b><u>mapValueTo</u></b> is set to <b><u>"index"</u></b>, which does not affect the current behavior of the virtualization process.
	  For more information, refer to the <a href="/kendo-ui/controls/editors/combobox/virtualization#value-mapping">article on virtualization</a>.
	  Example - DropDownList widget with declarative virtualization config
	   <div class="demo-section k-header">
	       <h4>Search for shipping name</h4>
	       <input id="orders" style="width: 400px"
	              data-role="dropdownlist"
	              data-bind="value: order, source: source"
	              data-text-field="ShipName"
	              data-value-field="OrderID"
	              data-filter="contains"
	              data-virtual="{itemHeight:26,valueMapper:orderValueMapper}"
	              data-height="520"
	              />
	   </div>

	   <script>
	       $(document).ready(function() {
	           var model = kendo.observable({
	                   order: "10548",
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
	MapValueTo KendoMapValueTo `jsObject:"mapValueTo"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-virtual.valueMapper  The widget calls the <b><u>valueMapper</u></b> function when the widget receives a value, that is not fetched from the remote server yet. The widget will pass the selected value(s) in the <b><u>valueMapper</u></b> function. In turn, the valueMapper implementation should return the <strong>respective data item(s) index/indices</strong>.
	  <strong>Important</strong>
	  As of the Kendo UI R3 2016 release, the implementation of the <b><u>valueMapper</u></b> function is optional. It is required only if the widget contains an initial value or if the <b><u>value</u></b> method is used.

	*/
	ValueMapper JavaScript `jsObject:"valueMapper"`

	*ToJavaScriptConverter
}

func (el *KendoVirtual) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoVirtual.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
