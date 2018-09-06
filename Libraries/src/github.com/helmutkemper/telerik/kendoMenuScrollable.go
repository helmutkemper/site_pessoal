package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type MenuScrollable struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/menu/configuration/scrollable#scrollable.distance

	  Sets the scroll amount (in pixels) that the Menu scrolls when the scroll buttons are hovered. Each such distance is
	  animated and then another animation starts with the same distance. If clicking a scroll button, the Menu scrolls with
	  2x the distance. (default: 50)

	*/
	//    <ul id="menu" style="width:150px;">
	//      <li>Item 1</li>
	//      <li>Item 2</li>
	//      <li>Item 3</li>
	//    </ul>
	//
	//    <script>
	//      $("#menu").kendoMenu({
	//        scrollable: {
	//          distance: 20
	//        }
	//      });
	//    </script>
	Distance int `jsObject:"distance"`

	*ToJavaScriptConverter
}

func (el *MenuScrollable) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoClose.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
