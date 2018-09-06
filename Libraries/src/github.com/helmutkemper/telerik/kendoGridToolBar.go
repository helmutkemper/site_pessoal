package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridToolbar struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/toolbar.text

	  The text displayed by the command button. If not set the <a href="toolbar.name">name</a>` option would be used as the button text instead.
	*/
	//
	//    Example - set the text of the toolbar button
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      toolbar: [
	//        { name: "create", text: "Add new" }
	//      ],
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//        { id: 1, name: "Jane Doe", age: 30 },
	//        { id: 2, name: "John Doe", age: 33 }
	//      ],
	//      editable: true
	//    });
	//    </script>
	//
	Text string `jsObject:"text"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/toolbar.template

	  The template [ https://docs.telerik.com/kendo-ui/api/javascript/kendo/methods/template ] which renders the command. By default renders a button.
	*/
	//
	//    Example - set the template as a function
	//
	//    <div id="grid"></div>
	//    <script id="template" type="text/x-kendo-template">
	//    <a class="k-button" href="\#" onclick="return toolbar_click()">Command</a>
	//    </script>
	//    <script>
	//    function toolbar_click() {
	//      console.log("Toolbar command is clicked!");
	//      return false;
	//    }
	//    $("#grid").kendoGrid({
	//      toolbar: [
	//        { template: kendo.template($("#template").html()) }
	//      ],
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//      ]
	//    });
	//    </script>
	//
	//
	//    Example - set the template as a string
	//
	//    <div id="grid"></div>
	//    <script>
	//    function toolbar_click() {
	//      console.log("Toolbar command is clicked!");
	//      return false;
	//    }
	//    $("#grid").kendoGrid({
	//      toolbar: [
	//        {
	//          template: '<a class="k-button" href="\\#" onclick="return toolbar_click()">Command</a>'
	//        }
	//      ],
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//      ]
	//    });
	//    </script>
	//
	Template interface{} `jsObject:"template" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/toolbar.iconclass

	  The class for the <a href="http://docs.telerik.com/kendo-ui/styles-and-layout/icons-web">web font icon</a> of the button that will be rendered in the toolbar.

	  Important:
	  Grid commands are rendered as anchors (`<a>´) with a `span´ inside. The icon for the button depends on the iconClass which is rendered as a class for the inner span.
	  Built-in commands have a predefined iconClass value.

	*/
	//
	//    Example - provide an iconClass for a toolbar command
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      toolbar: [
	//        { name: "copy", iconClass: "k-icon k-i-copy" },
	//        { name: "save" },
	//        { name: "custom" }
	//      ],
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: {
	//        data: [
	//          { id: 1, name: "Jane Doe", age: 30 },
	//          { id: 2, name: "John Doe", age: 33 }
	//        ],
	//        schema: {
	//          model: { id: "id" }
	//        }
	//      },
	//      editable: true
	//    });
	//    </script>
	//
	IconClass string `jsObject:"iconClass"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/toolbar.name

	  The name of the toolbar command. Either a built-in ("cancel", "create", "save", "excel", "pdf") or custom. The `name´ is reflected in one of the CSS classes, which is applied to the button - `k-grid-name´.
	  This class can be used to obtain reference to the button after Grid initialization and attach click handlers.
	*/
	//
	//    Example - specify the name of the command
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      toolbar: [
	//        { name: "create" },
	//        { name: "save" },
	//        { name: "custom" }
	//      ],
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: {
	//        data: [
	//          { id: 1, name: "Jane Doe", age: 30 },
	//          { id: 2, name: "John Doe", age: 33 }
	//        ],
	//        schema: {
	//          model: { id: "id" }
	//        }
	//      },
	//      editable: true
	//    });
	//
	//    $(".k-grid-custom").click(function(e){
	//        // handler body
	//    });
	//    </script>
	//
	Name KendoGridToolBarName `jsObject:"name"`

	*ToJavaScriptConverter
}

func (el *KendoGridToolbar) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridToolbar.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
