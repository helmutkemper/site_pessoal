package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

// fixme: https://docs.telerik.com/kendo-ui/api/javascript/data/datasource/configuration/schema#schema.model
// fixme: falta fazer o resto
/*
Example - set the model as a JavaScript object
EditPreviewOpen In Dojo
<script>
var dataSource = new kendo.data.DataSource({
  schema: {
    model: {
      id: "ProductID",
      fields: {
        ProductID: {
          //this field will not be editable (default value is true)
          editable: false,
          // a defaultValue will not be assigned (default value is false)
          nullable: true
        },
        ProductName: {
          //set validation rules
          validation: { required: true }
        },
        UnitPrice: {
          //data type of the field {Number|String|Boolean|Date} default is String
          type: "number",
          // used when new model is created
          defaultValue: 42,
          validation: { required: true, min: 1 }
        }
      }
    }
  }
});
</script>
Example - set the model as an existing kendo.data.Model instance
EditPreviewOpen In Dojo
<script>
var Product = kendo.data.Model.define({
  id: "ProductID",
  fields: {
    ProductID: {
      //this field will not be editable (default value is true)
      editable: false,
      // a defaultValue will not be assigned (default value is false)
      nullable: true
    },
    ProductName: {
      //set validation rules
      validation: { required: true }
    },
    UnitPrice: {
      //data type of the field {Number|String|Boolean|Date} default is String
      type: "number",
      // used when new model is created
      defaultValue: 42,
      validation: { required: true, min: 1 }
    }
  }
});
var dataSource = new kendo.data.DataSource({
  schema: {
    model: Product
  }
});
</script>
*/
type KendoField struct {
	Type         JavaScriptType         `jsObject:"type"`
	DefaultValue interface{}            `jsObject:"defaultValue"`
	Validation   map[string]interface{} `jsObject:"validation"`

	*ToJavaScriptConverter
}

func (el *KendoField) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoField.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
func (el *KendoField) String() string {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoField.Error: %v", err.Error())
		return ""
	}

	return string(ret)
}
