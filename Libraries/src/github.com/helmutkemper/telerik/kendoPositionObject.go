package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoPositionObject struct {
	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/position#position.top
	//
	// Specifies the initial top position of the window. Numeric values are treated as pixels. String values can specify
	// pixels, percentages, ems or other valid values.
	Top interface{} `jsObject:"top"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/position#position.left
	//
	// Specifies the initial left position of the window. Numeric values are treated as pixels. String values can specify
	// pixels or percentages, ems or other valid values.
	Left interface{} `jsObject:"left"`

	*ToJavaScriptConverter
}

func (el *KendoPositionObject) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoPositionObject.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
