package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoColorMessages struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker/configuration/messages#messages.apply

	  Allows customization of the "Apply" button text. (default: "Apply")

	  Example
	  <input id="colorpicker" type="color" />
	  <script>
	  $("#colorpicker").kendoColorPicker({
	    messages: {
	      apply: "Update"
	    }
	  })
	  </script>
	*/
	Apply string `jsObject:"apply"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker/configuration/messages#messages.cancel

	  Allows customization of the "Cancel" button text. (default: "Cancel")

	  Example
	  <input id="colorpicker" type="color" />
	  <script>
	  $("#colorpicker").kendoColorPicker({
	    messages: {
	      cancel: "Discard"
	    }
	  })
	  </script>
	*/
	Cancel string `jsObject:"cancel"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker/configuration/messages#messages.previewInput

	  Allows customization of the "Color Hexadecimal Code" preview input title. (default: "Color Hexadecimal Code")

	  Example
	  <input id="colorpicker" type="color" />
	  <script>
	  $("#colorpicker").kendoColorPicker({
	    messages: {
	      previewInput: "Edit Color"
	    }
	  })
	  </script>
	*/
	PreviewInput string `jsObject:"previewInput"`

	*ToJavaScriptConverter
}

func (el *KendoColorMessages) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoColorMessages.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
