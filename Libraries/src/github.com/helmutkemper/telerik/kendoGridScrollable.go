package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoGridScrollable struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/scrollable.virtual

	  If set to true the grid will always display a single page of data. Scrolling would just change the data which is
	  currently displayed. (default: false)

	*/
	Virtual Boolean `jsObject:"virtual"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/grid/configuration/scrollable.endless

	  If set to true the grid will always display a single page of data. Scrolling to the end will load more items untill
	  all items are displayed. (default: false)

	*/
	Endless Boolean `jsObject:"endless"`

	*ToJavaScriptConverter
}

func (el *KendoGridScrollable) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoGridScrollable.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
