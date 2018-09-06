package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoDataModel struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/model/fields/id#id

	  The value of the Model's ID. This field is available only if the id is defined in the Model configuration.
	*/
	Id string `jsObject:"id"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/model/fields/idfield#idField

	  The name of the Model's ID field. This field is available only if the id is defined in the Model configuration.
	  <script>
	  var Person = kendo.data.Model.define({
	      id: "personId",
	      fields: {
	          "name": {
	              type: "string"
	          },
	          "age": {
	              type: "number"
	          }
	      }
	  });

	  var person = new Person({
	      personId: 1,
	      name: "John Doe",
	      age: 42
	  });

	  console.log(person.id); // outputs 1
	  console.log(person.idField); // outputs "personId"
	  </script>
	*/
	IdField string `jsObject:"idField"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/model/fields/uid
	  The unique identifier of the Model. Inherited from ObservableObject. More info can be found in the uid section of the ObservableObject API reference.

	  The main benefit of uid's is to represent a link between data items (that may not have an ID of their own) and the corresponding rendered DOM elements (list items, table rows, etc). The uid's are generated randomly and they are not persisted across data or web page reloads.

	  https://docs.telerik.com/kendo-ui/api/javascript/data/observableobject/fields/uid
	*/
	UId string `jsObject:"uid"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/data/model/fields/dirty#dirty

	  Indicates whether the model is modified.

	  Example - using the dirty field
	  <script>
	  var model = new kendo.data.Model({
	      name: "John Doe"
	  });

	  console.log(model.dirty); // outputs "false"
	  model.set("name", "Jane Doe");
	  console.log(model.dirty); // outputs "true"
	  </script>
	*/
	Dirty Boolean `jsObject:"dirty"`

	// fixme:
	Fields map[string]KendoField `jsObject:"fields"`

	*ToJavaScriptConverter
}

func (el *KendoDataModel) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoDataModel.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
