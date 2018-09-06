package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoOffset struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/draggable/configuration/cursoroffset#cursorOffset

	  If set, specifies the offset of the hint relative to the mouse cursor/finger. By default, the hint is initially
	  positioned on top of the draggable source offset. The option accepts an object with two keys: top and left.

	  Example
	  <div id="draggable"></div>

	  <script>
	    $("#draggable").kendoDraggable({
	      hint: function(element) {
	        return element.clone();
	      },
	      cursorOffset: { top: 30, left: 100 }
	    });
	  </script>
	  <style>
	    #draggable {
	      width: 50px;
	      height: 50px;
	      background-color: orange;
	      border: 2px solid green;
	    }
	  </style>
	*/
	Top int `jsObject:"top"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/draggable/configuration/cursoroffset#cursorOffset

	  If set, specifies the offset of the hint relative to the mouse cursor/finger. By default, the hint is initially
	  positioned on top of the draggable source offset. The option accepts an object with two keys: top and left.

	  Example
	  <div id="draggable"></div>

	  <script>
	    $("#draggable").kendoDraggable({
	      hint: function(element) {
	        return element.clone();
	      },
	      cursorOffset: { top: 30, left: 100 }
	    });
	  </script>
	  <style>
	    #draggable {
	      width: 50px;
	      height: 50px;
	      background-color: orange;
	      border: 2px solid green;
	    }
	  </style>
	*/
	Left int `jsObject:"left"`

	*ToJavaScriptConverter
}

func (el *KendoOffset) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoOffset.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
