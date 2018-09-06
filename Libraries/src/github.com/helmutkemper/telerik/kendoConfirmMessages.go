package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoConfirmMessages struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/confirm/configuration/messages#messages.okText

	  The title of the OK button.

	  Example
	  <div id="confirm"></div>
	  <script>
	  $("#confirm").kendoConfirm({
	    messages:{
	      okText: "OK",
	    }
	  }).data("kendoConfirm").open();
	  </script>
	*/
	OkText string `jsObject:"okText"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/confirm/configuration/messages#messages.cancel

	  The title of the Cancel button.

	  Example
	  <div id="confirm"></div>
	  <script>
	  $("#confirm").kendoConfirm({
	    messages:{
	      cancel: "No"
	    }
	  }).data("kendoConfirm").open();
	  </script>
	*/
	Cancel string `jsObject:"cancel"`

	*ToJavaScriptConverter
}

func (el *KendoConfirmMessages) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoConfirmMessages.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
