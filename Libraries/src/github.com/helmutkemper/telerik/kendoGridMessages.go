package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridMessages struct {
	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/messages.commands
	//
	// Defines the text of the command buttons that are shown within the Grid. Used primarily for localization.
	//
	//    Example
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" },
	//          { command: ["edit", "destroy"] }
	//        ],
	//        dataSource: [
	//          { name: "Jane Doe", age: 30 },
	//          { name: "John Doe", age: 33 }
	//        ],
	//        toolbar: ["create", "save", "cancel"],
	//        messages: {
	//          commands: {
	//            cancel: "Cancel changes",
	//            canceledit: "Cancel",
	//            create: "Add new record",
	//            destroy: "Delete",
	//            edit: "Edit",
	//            save: "Save changes",
	//            select: "Select",
	//            update: "Update"
	//          }
	//        }
	//      });
	//    </script>
	Commands KendoGridMessagesCommands `jsObject:"commands"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/messages.norecords
	//
	// Defines the text of the "noRecords" option that is rendered when no records are available in current view. The
	// "noRecords" options should be set to true.
	//
	//    Example
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" },
	//          { command: ["edit", "destroy"] }
	//        ],
	//        dataSource: {
	//          data: [
	//            { name: "Jane Doe", age: 30 },
	//            { name: "John Doe", age: 33 }
	//          ],
	//          page: 2,
	//          pageSize: 10
	//        },
	//        noRecords: true,
	//        height: 200,
	//        messages: {
	//          noRecords: "There is no data on current page"
	//        }
	//      });
	//    </script>
	NoRecords string `jsObject:"noRecords"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/messages.expandcollapsecolumnheader
	//
	// Allows the customization of the text in the column header for the expand or collapse columns. Sets the value to
	// make the widget compliant with the web accessibility standards. (default: "")
	//
	//    Example
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { name: "Jane Doe", age: 30, city: "London" },
	//            { name: "John Doe", age: 33, city: "Berlin" }
	//          ]
	//        },
	//        detailInit: function (e) {
	//          e.detailCell.text("City: " + e.data.city);
	//        },
	//        height: 200,
	//        messages: {
	//          expandCollapseColumnHeader: "E/C"
	//        }
	//      });
	//    </script>
	ExpandCollapseColumnHeader string `jsObject:"expandCollapseColumnHeader"`

	*ToJavaScriptConverter
}

func (el *KendoGridMessages) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridMessages.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
