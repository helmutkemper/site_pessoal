package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUiDropTarget struct {
	Html HtmlInputText `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/droptarget#configuration-group

	  Used to group sets of draggable and drop targets. A draggable with the same group value as a drop target will be accepted by the drop target.

	  Example
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

	*ToJavaScriptConverter
}

func (el *KendoUiDropTarget) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("kendoDropTarget.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoDropTarget({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiDropTarget) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiDropTarget) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiDropTarget) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
