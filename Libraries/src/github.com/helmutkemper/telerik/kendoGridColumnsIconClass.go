package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridColumnsIconClass struct {
	Edit   interface{} `jsObject:"edit" jsType:"string,CssClassIcon"`
	Update interface{} `jsObject:"update" jsType:"string,CssClassIcon"`
	Cancel interface{} `jsObject:"cancel" jsType:"string,CssClassIcon"`

	*ToJavaScriptConverter
}

func (el *KendoGridColumnsIconClass) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridColumnsIconClass.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
