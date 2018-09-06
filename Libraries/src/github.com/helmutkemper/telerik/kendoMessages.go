package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoMessages struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-messages.close  The title of the close button.
	  Example
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     messages:{
	       close: "Close Me!"
	     }
	   });
	   </script>
	*/
	Close string `jsObject:"close"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-messages.promptInput  The title of the prompt input.
	  Example
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     messages:{
	       promptInput: "Input!"
	     }
	   });
	   </script>
	*/
	PromptInput string `jsObject:"promptinput"`

	*ToJavaScriptConverter
}

func (el *KendoMessages) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoMessages.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
