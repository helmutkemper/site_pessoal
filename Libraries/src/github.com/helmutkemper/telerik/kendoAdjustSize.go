package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

/*
@see https://docs.telerik.com/kendo-ui/api/javascript/ui/popup/configuration/adjustsize#adjustSize

Configures the margins, which will be added to the popup size, if its position should end up being next to the viewport edges. By default, the adjustment amount in both dimensions is zero.

The property takes effect only if collision is set to "fit" for the respective dimension (width or height).

Example
<div style="height:500px;">&nbsp;</div>
<p style="text-align:right;"><input id="datepicker" /></p>

<div id="popup">popup that is 100px offset from the bottom-right edge of the page.</div>

<script>
  $("#popup").kendoPopup({
    anchor: $("#datepicker"),
    origin: "bottom right",
    position: "top right",
    collision: "fit",
    adjustSize: {
        width: 100,
        height: 100
    }
  }).data("kendoPopup").open();
</script>
*/
type KendoAdjustSize struct {
	Width int `jsObject:"width"`

	Height int `jsObject:"height"`

	*ToJavaScriptConverter
}

func (el *KendoAdjustSize) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoAdjustSize.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
