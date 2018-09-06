package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUiMenu struct {
	Html HtmlElementUl `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/menu/configuration/animation

	  Configures the opening and closing animations of the suggestion popup. Setting the **animation** option to **false**
	  will disable the opening and closing animations. As a result the suggestion popup will open and close instantly.
	  **animation:true** is not a valid configuration.
	*/
	//    Example - disable open and close animations
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
	//        animation: { open: { effects: "fadeIn" } }
	//      });
	//    </script>
	//
	//
	//
	Animation interface{} `jsObject:"animation" jsType:"*KendoAnimation,Boolean"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/menu/configuration/closeonclick

	  Specifies that sub menus should close after item selection (provided they won't navigate). (default: true)

	*/
	//
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
	//        closeOnClick: false
	//      });
	//    </script>
	CloseOnClick Boolean `jsObject:"closeOnClick"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/menu/configuration/datasource

	  The data source of the widget which is used to display suggestions for the current value. Can be a JavaScript object which represents a valid data source configuration, a JavaScript array or an existing <a href="/kendo-ui/api/javascript/data/datasource">kendo.data.DataSource</a> instance.
	  If the <b><u>dataSource</u></b> option is set to a JavaScript object or array the widget will initialize a new <a href="/kendo-ui/api/javascript/data/datasource">kendo.data.DataSource</a> instance using that value as data source configuration.
	  If the <b><u>dataSource</u></b> option is an existing <a href="/kendo-ui/api/javascript/data/datasource">kendo.data.DataSource</a> instance the widget will use that instance and will <strong>not</strong> initialize a new one.

	*/
	//    Example - set dataSource as a JavaScript object
	//
	//    <ul id="menu"></ul>
	//    <script>
	//      var imgUrl = "https://demos.telerik.com/kendo-ui/content/shared/icons/sports/swimming.png";
	//      $(document).ready(function() {
	//        $("#menu").kendoMenu({
	//          dataSource:
	//          [
	//            {
	//              text: "Item 1",
	//              cssClass: "myClass",                         // Add custom CSS class to the item, optional, added 2012 Q3 SP1.
	//              url: "http://www.kendoui.com",               // Link URL if navigation is needed, optional.
	//              attr: {
	//                custom: 'value',                            // Add attributes with specified values
	//                other: 'value'
	//              }
	//            },
	//            {
	//              text: "<b>Item 2</b>",
	//              encoded: false,                              // Allows use of HTML for item text
	//              content: "text",                              // content within an item
	//              contentAttr: {
	//                style: 'border: 1px solid red; padding: 2px;', // Add attributes to the content container
	//                custom: 'value'
	//              }
	//            },
	//            {
	//              text: "Item 3",
	//              imageAttr: {                                                                   // Add additional image attributes
	//                alt: 'Image',
	//                height: '25px',
	//                width: '25px'
	//              },
	//              imageUrl: imgUrl,                            // Item image URL, optional.
	//              items:
	//              [
	//                {                                    // Sub item collection
	//                  text: "Sub Item 1"
	//                },
	//                {
	//                  text: "Sub Item 2"
	//                }
	//              ]
	//            },
	//            {
	//              text: "Item 4",
	//              spriteCssClass: "imageClass3"                // Item image sprite CSS class, optional.
	//            },
	//            {
	//              text: "Item 5",
	//              select: function(e) {                        // Item select event handler, optional
	//                // e.sender - returns reference to the Kendo Menu widget
	//                // e.target - returns the clicked element. Typically, the span.k-link element.
	//
	//                // handle event
	//            }
	//          }]
	//        })
	//      });
	//    </script>
	DataSource interface{} `jsObject:"dataSource" jsType:"*KendoDataSource,string,*map[string]interface {},[]string,[]MenuObject"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/menu/configuration/direction

	  Specifies Menu opening direction. Can be "top", "bottom", "left", "right". You can also specify different direction
	  for root and sub menu items, separating them with space. The example below will initialize the root menu to open
	  upwards and its sub menus to the left. (default: "default")

	*/
	/*
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
	  //        direction: "top left"
	  //      });
	  //    </script>
	*/
	Direction TypeMenuDirection `jsObject:"direction"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/menu/configuration/hoverdelay

	  Specifies the delay in ms before the menu is opened/closed - used to avoid accidental closure on leaving. (default: 100)

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
	//        hoverDelay: 200
	//      });
	//    </script>
	HoverDelay int `jsObject:"hoverDelay"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/menu/configuration/openonclick

	  Specifies that the root sub menus will be opened on item click. (default: false)

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
	//        openOnClick: true
	//      });
	//    </script>
	OpenOnClick interface{} `jsObject:"openOnClick" jsType:"Boolean,MenuOpenOnClick"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/menu/configuration/orientation

	  Root menu orientation. Could be horizontal or vertical. (default: "horizontal")

	*/
	//
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
	//        orientation: "vertical"
	//      });
	//    </script>
	Orientation TypeMenuOrientation `jsObject:"orientation"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/menu/configuration/popupcollision

	  Specifies how Menu should adjust to screen boundaries. By default the strategy is "fit" for a sub menu with a
	  horizontal parent, meaning it will move to fit in screen boundaries in all directions, and "fit flip" for a sub menu
	  with vertical parent, meaning it will fit vertically and flip over its parent horizontally. You can also switch off
	  the screen boundary detection completely if you set the popupCollision to false.

	*/
	//
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
	//        popupCollision: false
	//      });
	//    </script>
	PopupCollision TypeMenuCollision `jsObject:"popupCollision"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/menu/configuration/scrollable

	  If enabled, the Menu displays buttons that scroll the items when they cannot fit the width or the popups' height of
	  the Menu. By default, scrolling is disabled.

	  The following example demonstrates how to enable the scrolling functionality.

	*/
	//
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
	//        scrollable: true
	//      });
	//    </script>
	Scrollable interface{} `jsObject:"scrollable" jsType:"Boolean,MenuScrollable"`

	*ToJavaScriptConverter
}

func (el *KendoUiMenu) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoUiMenu.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoMenu({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiMenu) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiMenu) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiMenu) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
