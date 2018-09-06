package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoClose struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.close.effects  The effect(s) to use when playing the close animation. Multiple effects should be separated with a space.
	  <a href="/kendo-ui/api/javascript/effects/common">Complete list of available animations</a>

	*/
	Effects KendoEffect `jsObject:"effects"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-animation.close.duration  The duration of the close animation in milliseconds.

	*/
	Duration int `jsObject:"duration"`

	*ToJavaScriptConverter
}

func (el *KendoClose) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoClose.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
