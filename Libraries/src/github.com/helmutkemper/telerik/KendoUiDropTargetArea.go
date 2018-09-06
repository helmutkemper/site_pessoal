package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUiDropTargetArea struct {
	Html HtmlInputText `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/droptargetarea#configuration-group

	  Used to group sets of draggable and drop targets. A draggable with the same group value as a drop target will be accepted by the drop target.

	  Example
	   <p>Area accepts only draggable elements from orange group</p>
	   <div id="area">
	     <div id="leftArea"></div>
	     <div id="rightArea"></div>
	   </div>
	   <div class="orange"></div>
	   <div class="orange"></div>
	   <div class="purple"></div>

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

	     $("#area").kendoDropTargetArea({
	         group: "orangeGroup",
	         filter: "#leftArea, #rightArea",
	         drop: onDrop
	     });

	     function onDrop(e) {
	       e.draggable.destroy();
	       e.draggable.element.css("opacity", 0.3);
	     }
	   </script>
	   <style>
	     .orange, .purple{
	       width: 50px;
	       height: 50px;
	       border: 2px solid green;
	       margin: 5px;
	       display: inline-block;
	     }
	     .orange { background-color: orange; }
	     .purple { background-color: purple; }
	     #area {
	         width: 300px;
	         height: 300px;
	         line-height: 300px;
	         background-color: gray;
	     }
	     #leftArea, #rightArea {
	       width: 140px;
	       height: 100px;
	       border: 2px solid green;
	       margin: 2px;
	       background-color: orange;
	       display: inline-block;
	       vertical-align: middle;
	     }
	   </style>
	*/
	Group string `jsObject:"group"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/droptargetarea#configuration-filter

	  Selector to filter the drop targets in the area. Every matched element acts as a drop target and fires events on the DropTargetArea. <strong>Specifying the filter is mandatory.</strong>

	  Example
	   <p>Area accepts only draggable elements from orange group</p>
	   <div id="area">
	     <div id="droptarget" class="orange"></div>
	   </div>
	   <div id="draggable" class="purple"></div>

	   <script>
	     $("#draggable").kendoDraggable({
	       hint: function(element) {
	         return element.clone();
	       }
	     });

	     $("#area").kendoDropTargetArea({
	         filter: "#droptarget",
	         drop: onDrop
	     });

	     function onDrop(e) {
	       e.dropTarget.toggleClass("orange").toggleClass("purple");
	       e.draggable.element.toggleClass("orange").toggleClass("purple");
	     }
	   </script>
	   <style>
	     #draggable {
	       width: 50px;
	       height: 50px;
	       border: 2px solid green;
	       margin: 5px;
	       display: inline-block;
	     }
	     .orange { background-color: orange; }
	     .purple { background-color: purple; }
	     #area {
	         width: 300px;
	         height: 300px;
	         line-height: 300px;
	         background-color: gray;
	     }
	     #droptarget {
	       width: 100px;
	       height: 100px;
	       border: 2px solid green;
	       margin: 0 96px;
	       display: inline-block;
	       vertical-align: middle;
	     }
	   </style>
	*/
	Filter string `jsObject:"filter"`

	*ToJavaScriptConverter
}

func (el *KendoUiDropTargetArea) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("kendoDropTargetArea.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoDropTargetArea({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiDropTargetArea) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiDropTargetArea) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiDropTargetArea) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
