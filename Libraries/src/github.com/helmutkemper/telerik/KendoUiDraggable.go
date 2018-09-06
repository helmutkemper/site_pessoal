package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUiDraggable struct {
	Html HtmlElementDiv `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/draggable#configuration-axis

	  Constrains the hint movement to either the horizontal (x) or vertical (y) axis. Can be set to either "x" or "y".

	  Example - initialize horizontally draggable element
	   <div id="draggable"></div>

	   <script>
	     $("#draggable").kendoDraggable({
	       hint: function(element) {
	         return element.clone();
	       },
	       axis: "x"
	     });
	   </script>
	   <style>
	     #draggable {
	       width: 50px;
	       height: 50px;
	       background-color: orange;
	       border: 2px solid green;
	     }
	   </style>
	*/
	Axis KendoAxis `jsObject:"axis"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/draggable#configuration-autoScroll

	  If set to <b><u>true</u></b> the widget will auto-scroll the container when the mouse/finger is close to the top/bottom of it.

	  Example - use autoScroll in a scrollable container
	   <div style="width: 200px; height: 200px; overflow: auto">
	       <div style="width: 1000px; height: 1000px;">
	           <div id="draggable"></div>
	       </div>
	   </div>

	   <script>
	     $("#draggable").kendoDraggable({ hint: function(element) { return element.clone(); }, autoScroll: true });
	   </script>

	   <style>
	     #draggable {
	       width: 50px;
	       height: 50px;
	       background-color: orange;
	       border: 2px solid green;
	     }
	   </style>
	*/
	AutoScroll Boolean `jsObject:"autoScroll"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/draggable#configuration-container

	  If set, the hint movement is constrained to the container boundaries.

	  Example
	   <div id="container">
	     <div id="draggable"></div>
	   </div>

	   <script>
	     $("#draggable").kendoDraggable({
	       hint: function(element) {
	         return element.clone();
	       },
	       container: $("#container")
	     });
	   </script>
	   <style>
	     #container {
	       width: 200px;
	       height: 200px;
	       border: 1px dashed red;
	     }
	     #draggable {
	       width: 50px;
	       height: 50px;
	       background-color: orange;
	       border: 2px solid green;
	     }
	   </style>
	*/
	Container JavaScript `jsObject:"container"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/draggable#configuration-cursorOffset

	  If set, specifies the offset of the hint relative to the mouse cursor/finger. By default, the hint is initially positioned on top of the draggable source offset. The option accepts an object with two keys: <b><u>top</u></b> and <b><u>left</u></b>.

	  Example
	   <div id="draggable"></div>

	   <script>
	     $("#draggable").kendoDraggable({
	       hint: function(element) {
	         return element.clone();
	       },
	       cursorOffset: { top: 30, left: 100 }
	     });
	   </script>
	   <style>
	     #draggable {
	       width: 50px;
	       height: 50px;
	       background-color: orange;
	       border: 2px solid green;
	     }
	   </style>
	*/
	CursorOffset KendoOffset `jsObject:"cursorOffset"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/draggable#configuration-distance

	  The required distance that the mouse should travel in order to initiate a drag.

	  Example
	   <div id="draggable"></div>

	   <script>
	     $("#draggable").kendoDraggable({
	       hint: function(element) {
	         return element.clone();
	       },
	       distance: 50
	     });
	   </script>
	   <style>
	     #draggable {
	       width: 50px;
	       height: 50px;
	       background-color: orange;
	       border: 2px solid green;
	     }
	   </style>
	*/
	Distance int `jsObject:"distance"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/draggable#configuration-filter

	  Selects child elements that are draggable if a widget is attached to a container.

	  Example
	   <div id="container">
	     <div class="draggable"></div>
	     <div class="static"></div>
	     <div class="static"></div>
	     <div class="static"></div>
	     <div class="draggable"></div>
	   </div>
	   <script>
	     $("#container").kendoDraggable({
	       filter: ".draggable",
	       hint: function(element) {
	         return element.clone();
	       }
	     });
	   </script>
	   <style>
	     .draggable, .static {
	       width: 50px;
	       height: 50px;
	       border: 2px solid green;
	       margin: 5px;
	     }
	     .draggable { background-color: orange; }
	     .static{ background-color: purple; }
	   </style>
	*/
	Filter string `jsObject:"filter"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/draggable#configuration-group

	  Used to group sets of draggable and drop targets. A draggable with the same group value as a drop target will be accepted by the drop target.

	  Example - grouping draggable elements
	   <div class="orange"></div>
	   <div class="orange"></div>
	   <div class="purple"></div>
	   <div class="purple"></div>
	   <div id="orangeArea"></div>
	   <div id="purpleArea"></div>

	   <script>
	     $(".orange").kendoDraggable({
	       group: "orangeGroup",
	       hint: function(element) {
	         return element.clone();
	       }
	     });

	     $(".purple").kendoDraggable({
	       group: "purpleGroup",
	       hint: function(element) {
	         return element.clone();
	       }
	     });

	     $("#orangeArea").kendoDropTarget({ group: "orangeGroup", drop: onDrop });
	     $("#purpleArea").kendoDropTarget({ group: "purpleGroup", drop: onDrop });

	     function onDrop(e) {
	       e.draggable.destroy();
	       e.draggable.element.remove();
	     }
	   </script>
	   <style>
	     .orange, .purple{
	       width: 50px;
	       height: 50px;
	       border: 2px solid green;
	       margin: 5px;
	     }
	     #orangeArea, #purpleArea {
	       width: 200px;
	       height: 200px;
	       border: 2px solid green;
	       margin: 5px;
	     }
	     .orange, #orangeArea { background-color: orange; }
	     .purple, #purpleArea { background-color: purple; }
	   </style>
	*/
	Group string `jsObject:"group"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/draggable#configuration-hint

	  Provides a way for customization of the drag indicator. If a function is supplied, it receives one argument - the draggable element's jQuery object.

	  Example - customizing draggable tooltip
	   <div id="draggable"></div>

	   <script>
	     $("#draggable").kendoDraggable({
	       hint: function(element) {
	         var hintElement = $("<div id='hint'></div>");
	         hintElement.css({
	           "background-image": "url('http://www.telerik.com/image/kendo-logo.png')",
	           "width": "230px",
	           "height": "80px"
	         });
	         return hintElement;
	       }
	     });
	   </script>
	   <style>
	     #draggable {
	       width: 50px;
	       height: 50px;
	       background-color: orange;
	       border: 2px solid green;
	     }
	   </style>
	*/
	Hint interface{} `jsObject:"hint" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/draggable#configuration-holdToDrag

	  Suitable for touch oriented user interface, in order to avoid collision with the touch scrolling gesture. When set to <b><u>true</u></b>, the widget will be activated after the user taps and holds the finger on the element for a short amount of time.
	  The <em>draggable</em> will also be activated by pressing, holding and lifting the finger without any movement. Dragging it afterwards will initiate the drag immediately. The activated mode can be canceled by calling <a href="https://docs.telerik.com/kendo-ui/api/javascript/ui/draggable#methods-cancelHold"><b><u>cancelHold</u></b></a>.

	  Example - hold to drag
	   <div id="draggable"></div>

	   <p id="alert" style="display:none">dragToHold activated...</p>

	   <script>
	     $("#draggable").kendoDraggable({
	       holdToDrag: true,
	       hold: function(e) {
	           $("#draggable").css("background", "#f00");
	           $("#alert").show();
	       },
	       hint: function(element) {
	         var hintElement = $("<div id='hint'></div>");
	         hintElement.css({
	           background: "#3c0",
	           width: 200,
	           height: 200
	         });
	         return hintElement;
	       }
	     });
	   </script>

	   <style>
	     #draggable {
	       width: 200px;
	       height: 200px;
	       background: #f90;
	       border: 2px solid #c60;
	     }
	   </style>
	*/
	HoldToDrag Boolean `jsObject:"holdToDrag"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/draggable#configuration-ignore

	  Specifies child elements for which the drag will not be initialized. Useful if the draggable contains input elements.

	  Example
	   <div id="container">
	       <input type="text" />
	       <div>Foo</div>
	   </div>

	   <script>
	     $("#container").kendoDraggable({
	       ignore: "input",
	       hint: function(element) {
	         return element.clone();
	       }
	     });
	   </script>
	   <style>
	       #container {
	           width: 50px;
	           height: 50px;
	           border: 2px solid green;
	           margin: 5px;
	       }
	       #container input
	       {
	           width: 90%;
	       }
	   </style>
	*/
	Ignore string `jsObject:"ignore"`

	*ToJavaScriptConverter
}

func (el *KendoUiDraggable) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("kendoDraggable.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoDraggable({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiDraggable) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiDraggable) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiDraggable) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
