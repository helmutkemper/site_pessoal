package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUiWindow struct {
	Html HtmlElementDiv `jsObject:"-"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/actions
	//
	// The buttons for interacting with the window. Predefined array values are "Close", "Refresh", "Minimize", and
	// "Maximize".
	//
	//    Example
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        actions: [ "Minimize", "Maximize" ]
	//      });
	//    </script>
	Actions []KendoWindowAction `jsObject:"actions"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/animation
	//
	// A collection of {Animation} objects, used to change default animations. A value of false will disable all animations in the widget.
	//
	// Important:
	//
	// animation:true is not a valid configuration.
	//
	//    Example - disable animation
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        animation: false
	//      });
	//    </script>
	Animation interface{} `jsObject:"animation" jsType:"*KendoAnimation,Boolean"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/appendto
	//
	// The element that the Window will be appended to. Beneficial if the Window is used together with a form. Note that
	// this does not constrain the window dragging within the given element. (default: document.body)
	//
	// Important:
	//
	// Appending the Window to an element with an overflow:hidden, overflow:auto or overflow:scroll style may result in
	// undesired behavior, because the Window will not be able to be displayed outside the element's boundaries. Unwanted
	// scrollbars may appear as well.
	//
	//    Example - set the window container to be the form with id="mainForm"
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        appendTo: "form#mainForm"
	//      });
	//    </script>
	AppendTo string `jsObject:"appendto"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/autofocus
	//
	// Determines whether the Window will be focused automatically when opened. The property also influences the focus
	// behavior when the Window is clicked when already opened. (default: true)
	//
	//    Example - set the autoFocus property
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        autoFocus: false
	//      });
	//    </script>
	AutoFocus Boolean `jsObject:"autofocus"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/content
	//
	// Specifies a URL or request options that the window should load its content from.
	//
	// Important:
	//
	// For URLs starting with a protocol (e.g. http://), a container iframe element is automatically created. This
	// behavior may change in future versions, so it is advisable to always use the iframe configuration option.
	//
	//    Example - fetch content from the server
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        content: "/details"
	//      });
	//    </script>
	Content interface{} `jsObject:"content" jsType:"*JavaScript,JavaScript,string,Content"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/draggable
	//
	// Enables (true) or disables (false) the ability for users to move/drag the widget. (default: true)
	//
	//    Example - disable window dragging
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        draggable: false
	//      });
	//    </script>
	Draggable Boolean `jsObject:"draggable"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/iframe
	//
	// Explicitly states whether a content iframe should be created. For more information, please read Using iframes.
	//
	//    Example - load full page
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        content: "http://www.telerik.com/",
	//        iframe: true
	//      });
	//    </script>
	IFrame Boolean `jsObject:"iframe"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/height
	//
	// Specifies height of the window.
	//
	//    Example
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        height: 400
	//      });
	//    </script>
	Height int `jsObject:"height"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/maxheight
	//
	// The maximum height (in pixels) that may be achieved by resizing the window.
	//
	//    Example
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        maxHeight: 300
	//      });
	//    </script>
	MaxHeight int `jsObject:"maxheight"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/maxwidth
	//
	// The maximum width (in pixels) that may be achieved by resizing the window.
	//
	//    Example
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        maxWidth: 300
	//      });
	//    </script>
	MaxWidth int `jsObject:"maxheight"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/minheight
	//
	// The minimum height (in pixels) that may be achieved by resizing the window. (default: 50)
	//
	//    Example
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        minHeight: 100
	//      });
	//    </script>
	MinHeight int `jsObject:"minheight"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/minwidth
	//
	// The minimum width (in pixels) that may be achieved by resizing the window.
	//
	//    Example
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        minWidth: 100
	//      });
	//    </script>
	MinWidth int `jsObject:"minWidth"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/modal
	//
	// Specifies whether the window should show a modal overlay over the page. (default: false)
	//
	//    Example
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        modal: true
	//      });
	//    </script>
	Modal Boolean `jsObject:"modal"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/pinned
	//
	// Specifies whether the window should be pinned, i.e. it will not move together with the page content during
	// scrolling. (default: false)
	//
	//    Example
	//
	//    <div style="height: 5000px;"></div>
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        pinned: true,
	//        position: { top: 100 }
	//      });
	//    </script>
	Pinned Boolean `jsObject:"pinned"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/position
	//
	// A collection of one or two members, which define the initial Window's top and/or left position on the page.
	//
	//    Example
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        position: {
	//          top: 100,
	//          left: "20%"
	//        }
	//      });
	//    </script>
	Position KendoPositionObject `jsObject:"position"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/resizable
	//
	// Enables (true) or disables (false) the ability for users to resize a Window. (default: true)
	//
	//    Example
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        resizable: false
	//      });
	//    </script>
	Resizable Boolean `jsObject:"resizable"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/scrollable
	//
	// Enables (true) or disables (false) the ability for users to scroll the window contents. (default: true)
	//
	//    Example
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        scrollable: false
	//      });
	//    </script>
	Scrollable Boolean `jsObject:"scrollable"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/title
	//
	// The text in the window title bar. If false, the window will be displayed without a title bar. Note that this will
	// prevent the window from being dragged, and the window titlebar buttons will not be shown.
	//
	//    Example
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        title: "Customer details"
	//      });
	//    </script>
	//
	//
	//
	//    Example - create a window without a title
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        title: false
	//      });
	//    </script>
	Title interface{} `jsObject:"title" jsType:"Boolean,string"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/visible
	//
	// Specifies whether the window will be initially visible. (default: true)
	//
	//    Example - show a dialog after one second delay
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        visible: false
	//      });
	//      setTimeout(function() {
	//        $("#dialog").data("kendoWindow").open();
	//      }, 1000);
	//    </script>
	Visible Boolean `jsObject:"visible"`

	// @see https://docs.telerik.com/kendo-ui/api/javascript/ui/window/configuration/width
	//
	// Specifies width of the window.
	//
	//    Example
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        width: 400
	//      });
	//    </script>
	//
	//
	//
	//    Example - specify window width in percent
	//
	//    <div id="dialog"></div>
	//    <script>
	//      $("#dialog").kendoWindow({
	//        width: "50%"
	//      });
	//    </script>
	Width interface{} `jsObject:"width" jsType:"int,string"`

	*ToJavaScriptConverter
}

func (el *KendoUiWindow) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoUiWindow.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoWindow({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiWindow) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiWindow) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiWindow) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
