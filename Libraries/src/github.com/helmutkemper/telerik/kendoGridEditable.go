package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridEditable struct {
	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/editable.confirmation
	//
	// If set to true the grid will display a confirmation dialog when the user clicks the "destroy" command button.
	// (default: true)
	//
	// Can be set to a string which will be used as the confirmation text.
	//
	// Can be set to a function which will be called, passing the model instance, to return the confirmation text.
	//
	//    Example - disable delete confirmation
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" },
	//          { command: "destroy" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema:{
	//            model: {
	//              id: "id",
	//              fields: {
	//                age: { type: "number"}
	//              }
	//            }
	//          }
	//        },
	//        editable: {
	//          confirmation: false
	//        }
	//      });
	//    </script>
	//
	//
	//
	//    Example - set delete confirmation text
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" },
	//          { command: "destroy" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema:{
	//            model: {
	//              id: "id",
	//              fields: {
	//                age: { type: "number"}
	//              }
	//            }
	//          }
	//        },
	//        editable: {
	//          confirmation: "Are you sure that you want to delete this record?"
	//        }
	//      });
	//    </script>
	//
	//
	//
	//    Example - set delete confirmation as function
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" },
	//          { command: "destroy" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema:{
	//            model: {
	//              id: "id",
	//              fields: {
	//                age: { type: "number"}
	//              }
	//            }
	//          }
	//        },
	//        editable: {
	//          confirmation: function(e) {
	//            return  "Are you sure that you want to delete record for " + e.name + "?";
	//          }
	//        }
	//      });
	//    </script>
	Confirmation interface{} `jsObject:"confirmation" jsType:"*JavaScript,string,Boolean"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/editable.canceldelete
	//
	// If confirmation is enabled the grid will display a confirmation dialog when the user clicks the "destroy" command
	// button. If the grid is in mobile mode this text will be used for the cancel button.
	//
	//    Example - change the cancel delete button text
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" },
	//          { command: "destroy" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema:{
	//            model: {
	//              id: "id",
	//              fields: {
	//                age: { type: "number"}
	//              }
	//            }
	//          }
	//        },
	//        mobile: "phone",
	//        editable: {
	//          confirmation: true,
	//          cancelDelete: "No"
	//        }
	//      });
	//    </script>
	CancelDelete string `jsObject:"canceldelete"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/editable.confirmdelete
	//
	// If confirmation is enabled the grid will display a confirmation dialog when the user clicks the "destroy" command
	// button. If the grid is in mobile mode this text will be used for the confirm button. (default: "Delete")
	//
	//    Example - change the confirm delete button text
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" },
	//          { command: "destroy" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema:{
	//            model: {
	//              id: "id",
	//              fields: {
	//                age: { type: "number"}
	//              }
	//            }
	//          }
	//        },
	//        mobile: "phone",
	//        editable: {
	//          confirmation: true,
	//          confirmDelete: "Yes"
	//        }
	//      });
	//    </script>
	ConfirmDelete string `jsObject:"confirmdelete"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/editable.createat
	//
	// The position at which new data items are inserted in the grid. Must be set to either "top" or "bottom". By default
	// new data items are inserted at the top. (default: "top")
	//
	//    Example - insert new data items at the bottom
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
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema:{
	//            model: {
	//              id: "id",
	//              fields: {
	//                age: { type: "number"}
	//              }
	//            }
	//          }
	//        },
	//        editable: {
	//          createAt: "bottom"
	//        },
	//        toolbar: ["create"]
	//      });
	//    </script>
	Createat string `jsObject:"createat"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/editable.destroy
	//
	// If set to true the user can delete data items from the grid by clicking the "destroy" command button. Deleting is
	// enabled by default. (default: true)
	//
	//    Example - disable deleting
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" },
	//          { command: "destroy" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema:{
	//            model: {
	//              id: "id",
	//              fields: {
	//                age: { type: "number"}
	//              }
	//            }
	//          }
	//        },
	//        editable: {
	//          destroy: false
	//        },
	//        toolbar: ["create"]
	//      });
	//    </script>
	Destroy Boolean `jsObject:"destroy"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/editable.mode
	//
	// The editing mode to use. The supported editing modes are "incell", "inline" and "popup". (default: "incell")
	//
	// Important:
	//
	// The "inline" and "popup" editing modes are triggered by the "edit" column command. Thus it is required to have a
	// column with an "edit" command.
	//
	//    Example - specify inline editing mode
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" },
	//          { command: "edit" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema:{
	//            model: {
	//              id: "id",
	//              fields: {
	//                age: { type: "number"}
	//              }
	//            }
	//          }
	//        },
	//        editable: {
	//          mode: "inline"
	//        }
	//      });
	//    </script>
	Mode KendoGridEditorMode `jsObject:"mode"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/editable.template
	//
	// The template which renders popup editor.
	//
	// The template should contain elements whose name HTML attributes are set as the editable fields. This is how the
	// grid will know which field to update. The other option is to use MVVM bindings in order to bind HTML elements to
	// data item fields.
	//
	// Use the role data attribute to initialize Kendo UI widgets in the template. Check data attribute initialization for
	// more info. The validation that is set in schema.model(/api/javascript/data/datasource/configuration/schema.model)
	// is not mapped automatically. As a result, when you use the editable.template option, you have to add the validation
	// for every element manually.
	//
	//    Example - customize the popup editor
	//
	//    <script id="popup-editor" type="text/x-kendo-template">
	//      <h3>Edit Person</h3>
	//      <p>
	//        <label>Name:<input name="name" /></label>
	//      </p>
	//      <p>
	//        <label>Age: <input data-role="numerictextbox" name="age" /></label>
	//      </p>
	//    </script>
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" },
	//          { command: "edit" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema:{
	//            model: {
	//              id: "id",
	//              fields: {
	//                age: { type: "number"}
	//              }
	//            }
	//          }
	//        },
	//        editable: {
	//          mode: "popup",
	//          template: kendo.template($("#popup-editor").html())
	//        }
	//      });
	//    </script>
	//
	//
	//
	//    Example - using MVVM in the popup editor template
	//
	//    <script id="popup-editor" type="text/x-kendo-template">
	//      <h3>Edit Person</h3>
	//      <p>
	//        <label>Name:<input data-bind="value:name" /></label>
	//      </p>
	//      <p>
	//        <label>Age:<input data-role="numerictextbox" data-bind="value:age" /></label>
	//      </p>
	//    </script>
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" },
	//          { command: "edit" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema:{
	//            model: {
	//              id: "id",
	//              fields: {
	//                age: { type: "number"}
	//              }
	//            }
	//          }
	//        },
	//        editable: {
	//          mode: "popup",
	//          template: kendo.template($("#popup-editor").html())
	//        }
	//      });
	//    </script>
	Template interface{} `jsObject:"template" jsType:"*JavaScript,string"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/editable.update
	//
	// If set to true the user can edit data items when editing is enabled. (default: true)
	//
	//    Example - enable only deleting
	//
	//    <div id="grid"></div>
	//    <script>
	//      $("#grid").kendoGrid({
	//        columns: [
	//          { field: "name" },
	//          { field: "age" },
	//          { command: "destroy" }
	//        ],
	//        dataSource: {
	//          data: [
	//            { id: 1, name: "Jane Doe", age: 30 },
	//            { id: 2, name: "John Doe", age: 33 }
	//          ],
	//          schema:{
	//            model: {
	//              id: "id",
	//              fields: {
	//                age: { type: "number"}
	//              }
	//            }
	//          }
	//        },
	//        editable: {
	//          mode: "incell",
	//          update: false
	//        }
	//      });
	//    </script>
	Update Boolean `jsObject:"update"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/editable.window
	//
	// Configures the Kendo UI Window instance, which is used when the Grid edit mode is "popup". The configuration is
	// optional.
	//
	// For more information, please refer to the Window configuration API.
	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window
	//
	//    Example - Grid popup Window configuration
	//
	//    <div id="grid"></div>
	//    <script>
	//      function myOpenEventHandler(e) {
	//        var confirm =   window.confirm('Do you want to edit this record?');
	//        if(!confirm){
	//          e.preventDefault()
	//        }
	//      }
	//
	//      var dataSource = new kendo.data.DataSource({
	//        data: [
	//          { id: 1, name: "Jane Doe", age: 30 },
	//          { id: 2, name: "John Doe", age: 33 }
	//        ],
	//        schema:{
	//          model: {
	//            id: "id",
	//            fields: {
	//              age: { type: "number"}
	//            }
	//          }
	//        }
	//      });
	//
	//      $("#grid").kendoGrid({
	//        columns:['name','age', {command:'edit'}],
	//        dataSource:dataSource,
	//        editable: {
	//          mode: "popup",
	//          window: {
	//            title: "My Custom Title",
	//            animation: false,
	//            open: myOpenEventHandler
	//          }
	//        }
	//      });
	//    </script>
	Window KendoUiWindow `jsObject:"window"`

	*ToJavaScriptConverter
}

func (el *KendoGridEditable) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridEditable.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
