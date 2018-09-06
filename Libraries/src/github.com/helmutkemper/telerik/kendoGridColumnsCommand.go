package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridColumnsCommand struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.command#columns.command.classname

	  The CSS class applied to the command button.
	*/
	//
	//    Example - set the CSS class of the command
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { command: [{ className: "btn-destroy", name: "destroy", text: "Remove" }] }
	//      ],
	//      editable: true,
	//      dataSource: [ { name: "Jane Doe" } ]
	//    });
	//    </script>
	//    <style>
	//    .btn-destroy {
	//        color: red;
	//    }
	//    </style>
	//
	ClassName string `jsObject:"className"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.command#columns.command.click

	  The JavaScript function executed when the user clicks the command button. The function receives a jQuery Event
	  [ http://api.jquery.com/category/events/event-object/ ] as an argument.

	  The function context (available via the `this´ keyword) will be set to the grid instance.

	  Important:

	  Grid custom commands are rendered as anchors (`<a>´) with no `href´ value. Prevent the click event in the click
	  function in order to avoid shifting of the page scroll position.

	*/
	//
	//    Example - handle the click event of the custom command button
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { command: [{
	//            name: "details",
	//            click: function(e) {
	//                // prevent page scroll position change
	//                e.preventDefault();
	//                // e.target is the DOM element representing the button
	//                var tr = $(e.target).closest("tr"); // get the current table row (tr)
	//                // get the data bound to the current table row
	//                var data = this.dataItem(tr);
	//                console.log("Details for: " + data.name);
	//            }
	//          }]
	//       }
	//      ],
	//      dataSource: [ { name: "Jane Doe" } ]
	//    });
	//    </script>
	//
	Click JavaScript `jsObject:"click"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.command#columns.command.iconclass

	  The class for the web font icon [ http://docs.telerik.com/kendo-ui/styles-and-layout/icons-web ] of the button.
	  When it is defined as an object it allows to customize the web font icon for the "edit", "update" and "cancel" command
	  buttons.

	  Important:

	  Grid commands are rendered as anchors (`<a>´) with a `span´ inside. The icon for the button depends on the iconClass
	  which is rendered as a class for the inner span.

	  Default commands have a predefined iconClass value.

	*/
	//
	//    Example - provide an iconClass for the grid command column
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { command: [{
	//            name: "copy",
	//            iconClass: "k-icon k-i-copy"
	//            }]
	//       }
	//      ],
	//      dataSource: [ { name: "Jane Doe" } ]
	//    });
	//    </script>
	//
	//
	//    Example - provide an custom iconClass for the update and cancel command buttons
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { command: [{
	//            name: "edit",
	//            iconClass: {
	//                edit: "k-icon k-i-edit",
	//                update: "k-icon k-i-copy",
	//                cancel: "k-icon k-i-arrow-60-up"
	//              }
	//            }]
	//       }
	//      ],
	//      dataSource: [ { name: "Jane Doe" } ],
	//      editable: "inline"
	//    });
	//    </script>
	//
	IconClass interface{} `jsObject:"iconClass" jsType:"string,KendoGridColumnsIconClass,CssClassIcon"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.command#columns.command.template

	  The template of the command column.

	  Important:

	  Add the `k-grid-[command.name]´ to any element in the template which requires the click
	  [ https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.command#columns.command.click ]
	  handler to be called.

	*/
	//
	//    Example - customize the template of the command column
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { command: [
	//            {
	//              // for click to work when there is template, add class "k-grid-[command.name]" to some element, otherwise the click handler will not be triggered
	//              name: "settings",
	//              template: "Some text in the command column <a class='k-button k-grid-settings'><span class='k-icon k-i-settings'></span>Settings</a>",
	//              click(e){
	//                kendo.alert("Settings clicked!")
	//              }
	//            }
	//          ]
	//          }
	//        ],
	//        dataSource: [{ name: "Jane Doe" }]
	//      });
	//    </script>
	//
	Template interface{} `jsObject:"template" jsType:"JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.command#columns.command.name

	  The name of the command. The built-in commands are "edit" and "destroy". Can be set to a custom value.
	*/
	//
	//    Example - set the command name
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { command: [{ name: "edit" }] }
	//      ],
	//      editable: "popup",
	//      dataSource: [ { name: "Jane Doe" } ]
	//    });
	//    </script>
	//
	Name interface{} `jsObject:"name" jsType:"TypeKendoGridColumnsCommand,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.command#columns.command.text

	  The text displayed by the command button and the "cancel", "edit" and "update" texts of the edit command. If not set
	  the <a href="columns.command.name">name</a> option is used as the button text.

	*/
	//
	//    Example - customize the text of the command
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { command: [{ name: "destroy", text: "Remove" }] }
	//      ],
	//      editable: true,
	//      dataSource: [ { name: "Jane Doe" } ]
	//    });
	//    </script>
	//
	//
	//    Example - customize the "edit", "cancel" and "update" text of the edit command
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { field: "age" },
	//        { command: [{ name: "edit",
	//                      text: { edit: "Custom edit", cancel: "Custom cancel", update: "Custom update" } }] }
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
	//      editable: {
	//        mode: "inline"
	//      }
	//    });
	//    </script>
	//
	Text interface{} `jsObject:"text"  jsType:"KendoGridMessagesCommands,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/columns.command#columns.command.visible

	  The JavaScript function executed on initialization of the row which will determine whether the command button will be
	  visible. The function receives a the data item object for the row as an argument.

	*/
	//
	//    Example - set the command name
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      columns: [
	//        { field: "name" },
	//        { command: [{ name: "edit", visible: function(dataItem) { return dataItem.name==="Jane" } }] }
	//      ],
	//      editable: "popup",
	//      dataSource: [ { name: "Jane" }, { name: "Bill" } ]
	//    });
	//    </script>
	//
	Visible JavaScript `jsObject:"visible"`

	*ToJavaScriptConverter
}

func (el *KendoGridColumnsCommand) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridColumnsCommand.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
