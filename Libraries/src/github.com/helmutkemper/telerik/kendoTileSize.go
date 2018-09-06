package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoTileSize struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-tileSize.width  The width of the color cell.
	  Example
	   <input id="colorpicker" type="color" />
	   <script>
	   $("#colorpicker").kendoColorPicker({
	     palette: "basic",
	     tileSize: { width: 40 }
	   });
	   </script>
	*/
	Width int `jsObject:"width"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/colorpicker#configuration-tileSize.height  The height of the color cell.
	  Example
	   <input id="colorpicker" type="color" />
	   <script>
	   $("#colorpicker").kendoColorPicker({
	     palette: "basic",
	     tileSize: { height: 40 }
	   });
	   </script>
	*/
	Height int `jsObject:"height"`

	*ToJavaScriptConverter
}

func (el *KendoTileSize) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoCalendarMessages.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
