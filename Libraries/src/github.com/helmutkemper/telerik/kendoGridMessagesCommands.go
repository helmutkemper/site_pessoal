package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridMessagesCommands struct {
	// this field don't exist in documentation, but, exits in example
	Select string `jsObject:"select"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/excel

	  Defines the text of the "Export to Excel" button of the grid toolbar.
	*/
	//
	//    Example
	//
	//    <div id="grid"></div>
	//    <script>
	//    $("#grid").kendoGrid({
	//      toolbar: [ "excel" ],
	//      columns: [
	//        { field: "name" },
	//        { field: "age" }
	//      ],
	//      dataSource: [
	//        { name: "Jane Doe", age: 30 },
	//        { name: "John Doe", age: 33 }
	//      ],
	//      messages: {
	//        commands: {
	//          excel: "Excel export"
	//        }
	//      }
	//    });
	//    </script>
	//
	Excel string `jsObject:"excel"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/save

	  Defines the text of the "Save Changes" button located in the toolbar of the widget.
	*/
	//
	//    Example
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
	//          save: "Save changes"
	//        }
	//      }
	//    });
	//    </script>
	//
	Save string `jsObject:"save"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/destroy

	  Defines the text of the "Delete" button rendered in `inline´ or `popup´ editing mode.
	*/
	//
	//    Example
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
	//          destroy: "Delete"
	//        }
	//      }
	//    });
	//    </script>
	//
	Destroy string `jsObject:"destroy"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/update

	  Defines the text of the "Update" button that is rendered in `inline´ or `popup´ editing mode.
	*/
	//
	//    Example
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
	//          update: "Update"
	//        }
	//      }
	//    });
	//    </script>
	//
	Update string `jsObject:"update"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/cancel

	  Defines the text of the "Cancel Changes" button located in the toolbar of the widget.
	*/
	//
	//    Example
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
	//          cancel: "Cancel changes"
	//        }
	//      }
	//    });
	//    </script>
	//
	Cancel string `jsObject:"cancel"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/edit

	  Defines the text of the "Edit" button that is rendered in `inline´ or `popup´ editing mode.
	*/
	//
	//    Example
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
	//          edit: "Edit"
	//        }
	//      }
	//    });
	//    </script>
	//
	Edit string `jsObject:"edit"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/create

	  Defines the text of the "Add new record" button located in the toolbar of the widget.
	*/
	//
	//    Example
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
	//          create: "Add new record"
	//        }
	//      }
	//    });
	//    </script>
	//
	Create string `jsObject:"create"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/canceledit

	  Defines the text of the "Cancel" button that is rendered in `inline´ or `popup´ editing mode.
	*/
	//
	//    Example
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
	//          canceledit: "Cancel"
	//        }
	//      }
	//    });
	//    </script>
	//
	CancelEdit string `jsObject:"canceledit"`

	*ToJavaScriptConverter
}

func (el *KendoGridMessagesCommands) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()

	/*
	  Select
	  Excel
	  Save
	  Destroy
	  Update
	  Cancel
	  Edit
	  Create
	  CancelEdit
	*/

	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridMessagesCommands.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
