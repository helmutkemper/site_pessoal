package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUiContextMenu struct {
	Html HtmlElementDiv `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/contextmenu#configuration-alignToAnchor

	  Specifies that ContextMenu should be shown aligned to the target or the filter element if specified.

	  Example
	   <div id="target">Target</div>
	   <ul id="context-menu">
	       <li>Item 1
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	       <li>Item 2
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	   </ul>
	   <script>
	       $("#context-menu").kendoContextMenu({
	           target: "#target",
	           alignToAnchor: true
	       });
	   </script>
	*/
	AlignToAnchor Boolean `jsObject:"alignToAnchor"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/contextmenu#configuration-animation

	  A collection of <strong>Animation</strong> objects, used to change default animations. A value of <b><u>false</u></b> will disable all animations in the widget.
	  <b><u>animation:true</u></b> is not a valid configuration.
	  Available animations for the <strong>ContextMenu</strong> are listed below. Each animation has a reverse options which is used for the <strong>close</strong> effect by default, but can be over-ridden by setting the <strong>close</strong> animation. Each animation also has a direction which can be set off the animation (i.e. <strong>slideIn:Down</strong>).
	  <b>slideIn</b> - ContextMenu content slides in from the top.
	  <b>fadeIn</b> - ContextMenu content fades in.
	  <b>expand</b> - ContextMenu content expands from the top down. Similar to slideIn.

	  slideIn
	   <div id="target">Target</div>
	   <ul id="context-menu">
	       <li>Item 1
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	       <li>Item 2
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	   </ul>
	   <script>
	       $("#context-menu").kendoContextMenu({
	           target: "#target",
	           animation: {
	               open: {
	                   effects: "fadeIn"
	               }
	           }
	       });
	   </script>
	*/
	Animation interface{} `jsObject:"animation" jsType:"*KendoAnimation,Boolean"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/contextmenu#configuration-appendTo

	  The DOM element to which the ContextMenu will be appended. The element needs to be relatively positioned.

	  Example
	   <div id="container">Container</div>
	   <ul id="context-menu">
	       <li>Item 1
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	       <li>Item 2
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	   </ul>
	   <script>
	       $("#context-menu").kendoContextMenu({
	           appendTo: "#container"
	       });
	   </script>
	*/
	AppendTo interface{} `jsObject:"appendTo" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/contextmenu#configuration-closeOnClick

	  Specifies that sub menus should close after item selection (provided they won't navigate).

	  Example
	   <div id="target">Target</div>
	   <ul id="context-menu">
	       <li>Item 1
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	       <li>Item 2
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	   </ul>
	   <script>
	       $("#context-menu").kendoContextMenu({
	           target: "#target",
	           closeOnClick: false
	       });
	   </script>
	*/
	CloseOnClick Boolean `jsObject:"closeOnClick"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/contextmenu#configuration-copyAnchorStyles

	  Copies and uses the styles from the anchor.

	  Example
	   <span class="k-icon k-i-filter" id="span-target"></span>
	   <ul id="context-menu-span">
	     <li>Item 1
	       <ul>
	         <li>Sub Item 1</li>
	         <li>Sub Item 2</li>
	         <li>Sub Item 3</li>
	       </ul>
	     </li>
	     <li>Item 2
	       <ul>
	         <li>Sub Item 1</li>
	         <li>Sub Item 2</li>
	         <li>Sub Item 3</li>
	       </ul>
	     </li>
	   </ul>
	   <script>

	     $("#context-menu-span").kendoContextMenu({
	       target: "#span-target",
	       alignToAnchor: true,
	       copyAnchorStyles: false
	     });
	   </script>
	*/
	CopyAnchorStyles Boolean `jsObject:"copyAnchorStyles"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/contextmenu#configuration-dataSource

	  The data source of the widget which is used to render its items. Can be a JSON object/Array that contains an item or an Array of items to be rendered. Refer to the example below for a list of the supported properties.

	  Example
	   <div id="target">Target</div>
	   <ul id="context-menu">
	       <li>Item 1
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	       <li>Item 2
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	   </ul>
	   <script>
	       $(document).ready(function() {
	           $("#context-menu").kendoContextMenu({
	               target: "#target",
	               dataSource:
	                   [{
	                       text: "Item 1",
	                       cssClass: "myClass",                         // Add custom CSS class to the item, optional, added 2012 Q3 SP1.
	                       url: "http://www.kendoui.com"                // Link URL if navigation is needed, optional.
	                   },
	                   {
	                       text: "<b>Item 2</b>",
	                       encoded: false,                              // Allows use of HTML for item text
	                       content: "text"                              // content within an item
	                   },
	                   {
	                       text: "Item 3",
	                       imageUrl: "http://www.kendoui.com/test.jpg", // Item image URL, optional.
	                       items: [{                                    // Sub item collection
	                            text: "Sub Item 1"
	                       },
	                       {
	                            text: "Sub Item 2"
	                       }]
	                   },
	                   {
	                       text: "Item 4",
	                       spriteCssClass: "imageClass3"                // Item image sprite CSS class, optional.
	                   }]
	           })
	       });
	   </script>
	*/
	DataSource interface{} `jsObject:"dataSource" jsType:"*KendoDataSource,string,*map[string]interface {},[]string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/contextmenu#configuration-direction

	  Specifies ContextMenu's sub menu opening direction. Can be "top", "bottom", "left", "right". The example below will initialize the sub menus to open to the left.

	  Example
	   <div id="target">Target</div>
	   <ul id="context-menu">
	       <li>Item 1
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	       <li>Item 2
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	   </ul>
	   <script>
	       $("#context-menu").kendoContextMenu({
	           target: "#target",
	           direction: "left"
	       });
	   </script>
	*/
	Direction KendoContextMenuDirection `jsObject:"direction"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/contextmenu#configuration-filter

	  Specifies ContextMenu filter selector - the ContextMenu will only be shown on items that satisfy the provided selector.

	  Show the ContextMenu on some elements inside the target
	   <div id="target">Target
	       <div class="box"></div>
	       <div></div>
	       <div class="box"></div>
	   </div>
	   <ul id="context-menu">
	       <li>Item 1
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	       <li>Item 2
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	   </ul>
	   <script>
	       $("#context-menu").kendoContextMenu({
	           target: "#target",
	           filter: ".box"
	       });
	   </script>
	*/
	Filter string `jsObject:"filter"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/contextmenu#configuration-hoverDelay

	  Specifies the delay in ms before the sub menus are opened/closed - used to avoid accidental closure on leaving.

	  Example
	   <div id="target">Target</div>
	   <ul id="context-menu">
	       <li>Item 1
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	       <li>Item 2
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	   </ul>
	   <script>
	       $("#context-menu").kendoContextMenu({
	           target: "#target",
	           hoverDelay: 200
	       });
	   </script>
	*/
	HoverDelay int `jsObject:"hoverDelay"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/contextmenu#configuration-orientation

	  Root menu orientation. Could be horizontal or vertical.

	  Example
	   <div id="target">Target</div>
	   <ul id="context-menu">
	       <li>Item 1
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	       <li>Item 2
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	   </ul>
	   <script>
	       $("#context-menu").kendoContextMenu({
	           target: "#target",
	           orientation: "horizontal"
	       });
	   </script>
	*/          //
	Orientation KendoOrientation `jsObject:"orientation"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/contextmenu#configuration-popupCollision

	  Specifies how ContextMenu should adjust to screen boundaries. By default the strategy is <strong>"fit"</strong> for a sub menu with a horizontal parent or the root menu, meaning it will move to fit in screen boundaries in all directions, and <strong>"fit flip"</strong> for a sub menu with vertical parent, meaning it will fit vertically and flip over its parent horizontally. You can also switch off the screen boundary detection completely if you set the <strong>popupCollision</strong> to false.

	  Example
	   <div id="target">Target</div>
	   <ul id="context-menu">
	       <li>Item 1
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	       <li>Item 2
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	   </ul>
	   <script>
	       $("#context-menu").kendoContextMenu({
	           target: "#target",
	           popupCollision: false
	       });
	   </script>
	*/
	PopupCollision Boolean `jsObject:"popupCollision"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/contextmenu#configuration-showOn

	  Specifies the event or events on which ContextMenu should open. By default ContextMenu will show on <em>contextmenu</em> event on desktop and <em>hold</em> event on touch devices. Could be any pointer/mouse/touch event, also several, separated by spaces.

	  Show the ContextMenu on left click
	   <div id="target">Target</div>
	   <ul id="context-menu">
	       <li>Item 1
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	       <li>Item 2
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	   </ul>
	   <script>
	       $("#context-menu").kendoContextMenu({
	           target: "#target",
	           showOn: "click"
	       });
	   </script>
	*/
	ShowOn HtmlMouseEvent `jsObject:"showOn"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/contextmenu#configuration-target

	  Specifies the element on which ContextMenu should open. The default element is the document body.

	  Show the ContextMenu on element with ID target
	   <div id="target">Target</div>
	   <ul id="context-menu">
	       <li>Item 1
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	       <li>Item 2
	           <ul>
	               <li>Sub Item 1</li>
	               <li>Sub Item 2</li>
	               <li>Sub Item 3</li>
	           </ul>
	       </li>
	   </ul>
	   <script>
	       $("#context-menu").kendoContextMenu({
	           target: "#target"
	       });
	   </script>
	*/
	Target interface{} `jsObject:"target" jsType:"*JavaScript,string"`

	*ToJavaScriptConverter
}

func (el *KendoUiContextMenu) ToJavaScript() []byte {
	var ret bytes.Buffer

	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoUiContextMenu.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoContextMenu({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiContextMenu) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiContextMenu) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiContextMenu) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
