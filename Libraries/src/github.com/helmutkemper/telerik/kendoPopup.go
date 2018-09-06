package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoPopup struct {

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/popup/configuration/adjustsize#adjustSize

	  Configures the margins, which will be added to the popup size, if its position should end up being next to the viewport edges. By default, the adjustment amount in both dimensions is zero.

	  The property takes effect only if collision is set to "fit" for the respective dimension (width or height).

	  Example
	  <div style="height:500px;">&nbsp;</div>
	  <p style="text-align:right;"><input id="datepicker" /></p>

	  <div id="popup">popup that is 100px offset from the bottom-right edge of the page.</div>

	  <script>
	    $("#popup").kendoPopup({
	      anchor: $("#datepicker"),
	      origin: "bottom right",
	      position: "top right",
	      collision: "fit",
	      adjustSize: {
	          width: 100,
	          height: 100
	      }
	    }).data("kendoPopup").open();
	  </script>
	*/
	AdjustSize KendoAdjustSize `jsObject:"adjustSize"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/popup/configuration/anchor#anchor

	  Specifies the element that will be used as an anchor. The widget will open next to that element.

	  Example - specify an anchor
	  <input id="datepicker" />
	  <div id="popup">CONTENT</div>

	  <script>
	      $("#popup").kendoPopup({
	          anchor: $("#datepicker")
	      }).data("kendoPopup").open();
	  </script>
	*/
	Anchor interface{} `jsObject:"anchor" jsType:"*JavaScript,string"`
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/popup/configuration/animation#animation

	  Configures the opening and closing animations of the suggestion popup. Setting the <b><u>animation</u></b> option to <b><u>false</u></b> will disable the opening and closing animations. As a result the suggestion popup will open and close instantly.
	  <b><u>animation:true</u></b> is not a valid configuration.

	  Example - disable open and close animations
	   <input id="autocomplete" />
	   <script>
	   $("#autocomplete").kendoAutoComplete({
	     animation: false
	   });
	   </script>
	*/
	Animation interface{} `jsObject:"animation" jsType:"*KendoAnimation,Boolean"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-popup.appendTo  Defines a jQuery selector that will be used to find a container element, where the popup will be appended to.
	  Example - append the popup to a specific element
	   <div id="container">
	       <input id="dropdownlist" />
	   </div>
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	     dataSource: [
	       { id: 1, name: "Apples" },
	       { id: 2, name: "Oranges" }
	     ],
	     dataTextField: "name",
	     dataValueField: "id",
	     popup: {
	       appendTo: $("#container")
	     }
	   });
	   </script>
	*/
	AppendTo JavaScript `jsObject:"appendTo"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/popup/configuration/collision#collision

	  Configures how the popup should behave when it cannot be properly displayed and fully visible, if its origin and position settings are obeyed.

	  Valid values are: "fit", "flip", "flip fit" and "fit flip". "Fit" allows the popup to be shifted (moved) until it is fully visible. "Flip" allows the popup to switch its position, according to its anchor. If two words are used, the first one applies to the horizontal dimension and the second one - to the vertical dimension. If one word is used, the setting is applied to both dimensions.

	  <div style="height:500px;">&nbsp;</div>
	  <p style="text-align:right;"><input id="datepicker" /></p>

	  <div id="popup" style="width: 100px; height: 100px;">popup content</div>

	  <script>
	    $("#popup").kendoPopup({
	      anchor: $("#datepicker"),
	      origin: "bottom right",
	      position: "top right",
	      collision: "fit flip"
	    }).data("kendoPopup").open();
	  </script>
	*/
	Collision KendoCollision `jsObject:"collision"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-popup.origin  Specifies how to position the popup element based on anchor point. The value is space separated "y" plus "x" position.
	  The available "y" positions are:
	  The available "x" positions are:
	  Example - append the popup to a specific element
	   <div id="container">
	       <input id="dropdownlist" />
	   </div>
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	     dataSource: [
	       { id: 1, name: "Apples" },
	       { id: 2, name: "Oranges" }
	     ],
	     dataTextField: "name",
	     dataValueField: "id",
	     popup: {
	       origin: "top left"
	     }
	   });
	   </script>
	*/
	Origin KendoOrigin `jsObject:"origin"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dropdownlist#configuration-popup.position  Specifies which point of the popup element to attach to the anchor's origin point. The value is space separated "y" plus "x" position.
	  The available "y" positions are:
	  The available "x" positions are:
	  Example - append the popup to a specific element
	   <div id="container">
	       <input id="dropdownlist" />
	   </div>
	   <script>
	   $("#dropdownlist").kendoDropDownList({
	     dataSource: [
	       { id: 1, name: "Apples" },
	       { id: 2, name: "Oranges" }
	     ],
	     dataTextField: "name",
	     dataValueField: "id",
	     popup: {
	       origin: "top left"
	     }
	   });
	   </script>
	*/
	Position KendoPosition `jsObject:"position"`

	*ToJavaScriptConverter
}

func (el *KendoPopup) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoPopup.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
