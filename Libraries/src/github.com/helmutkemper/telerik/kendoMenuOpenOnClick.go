package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type MenuOpenOnClick struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/menu/configuration/openonclick#openOnClick.rootMenuItems

	  Specifies that the root menus will be opened only on item click. (default: false)

	*/
	//    Example
	//
	//    <ul id="menu">
	//      <li>Item 1
	//        <ul>
	//          <li>Sub Item 1</li>
	//          <li>Sub Item 2</li>
	//          <li>Sub Item 3</li>
	//        </ul>
	//      </li>
	//      <li>Item 2
	//        <ul>
	//          <li>Sub Item 1</li>
	//          <li>Sub Item 2</li>
	//          <li>Sub Item 3</li>
	//        </ul>
	//      </li>
	//    </ul>
	//    <script>
	//      $("#menu").kendoMenu({
	//        openOnClick: {
	//          rootMenuItems: true
	//        }
	//      });
	//    </script>
	RootMenuItems Boolean `jsObject:"rootMenuItems"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/menu/configuration/openonclick#openOnClick.subMenuItems

	  Specifies that the sub menus will be opened only on item click. (default: false)

	*/
	//    Example
	//
	//    <ul id="menu">
	//      <li>Item 1
	//        <ul>
	//          <li>Sub Item 1</li>
	//          <li>Sub Item 2
	//            <ul>
	//              <li>Sub Item 2.1</li>
	//              <li>Sub Item 2.2</li>
	//              <li>Sub Item 2.3</li>
	//            </ul>
	//          </li>
	//          <li>Sub Item 3</li>
	//        </ul>
	//      </li>
	//      <li>Item 2
	//        <ul>
	//          <li>Sub Item 1</li>
	//          <li>Sub Item 2
	//            <ul>
	//              <li>Sub Item 2.1</li>
	//              <li>Sub Item 2.2</li>
	//              <li>Sub Item 2.3</li>
	//            </ul>
	//          </li>
	//          <li>Sub Item 3</li>
	//        </ul>
	//      </li>
	//    </ul>
	//    <script>
	//      $("#menu").kendoMenu({
	//        openOnClick: {
	//          subMenuItems: true
	//        }
	//      });
	//    </script>
	SubMenuItems Boolean `jsObject:"subMenuItems"`

	*ToJavaScriptConverter
}

func (el *MenuOpenOnClick) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoClose.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
