package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUiDialog struct {
	GlobalVar Boolean `jsObject:"-"`

	Html HtmlElementDiv `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-actions

	  A collection of objects containing text, action and primary attributes used to specify the dialog buttons.

	  Example
	   <div id="dialog"></div>
	   <script>
	       $("#dialog").kendoDialog({
	         actions: [{
	             text: "OK",
	             action: function(e){
	                 // e.sender is a reference to the dialog widget object
	                 // OK action was clicked
	                 // Returning false will prevent the closing of the dialog
	                 return false;
	             },
	             primary: true
	         },{
	             text: "Cancel"
	         }]
	       });
	   </script>
	*/
	Actions []KendoActions `jsObject:"actions"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-animation

	  A collection of {Animation} objects, used to change default animations. A value of <b><u>false</u></b> will disable all animations in the widget.
	  <b><u>animation:true</u></b> is not a valid configuration.

	  Example - disable animation
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     animation: false
	   });
	   </script>
	*/
	Animation interface{} `jsObject:"animation" jsType:"*KendoAnimation,Boolean"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-buttonLayout

	  Specifies the possible layout of the action buttons in the <strong>Dialog</strong>.
	  Possible values are:

	  Example
	   <div id="dialog"></div>
	   <script>
	       $("#dialog").kendoDialog({
	           buttonLayout: "normal"
	       });
	   </script>
	*/
	ButtonLayout KendoButtonLayout `jsObject:"buttonLayout"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-closable

	  Specifies whether a close button should be rendered at the top corner of the dialog.

	  Example
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     closable: true
	   });
	   </script>
	*/
	Closable Boolean `jsObject:"closable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-content

	  Specifies the content of a <strong>Dialog</strong>.

	  Example - fetch content from the server
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     content: "<em>Dialog content</em>"
	   });
	   </script>
	*/
	Content interface{} `jsObject:"content" jsType:"*JavaScript,JavaScript,string,Content"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-height

	  Specifies height of the dialog.

	  Example
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     height: 400
	   });
	   </script>
	*/
	Height int `jsObject:"height"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-maxHeight

	  The maximum height (in pixels) that may be achieved by resizing the dialog.

	  Example
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     maxHeight: 300
	   });
	   </script>
	*/
	MaxHeight int `jsObject:"maxHeight"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-maxWidth

	  The maximum width (in pixels) that may be achieved by resizing the dialog.

	  Example
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     maxWidth: 300
	   });
	   </script>
	*/
	MaxWidth int `jsObject:"maxWidth"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-messages

	  Defines the text of the labels that are shown within the dialog. Used primarily for localization.

	  Example
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     messages:{
	       close: "Close Me!"
	     }
	   });
	   </script>
	*/
	Messages KendoMessages `jsObject:"messages"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-minHeight

	  The minimum height (in pixels) that may be achieved by resizing the dialog.

	  Example
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     minHeight: 100
	   });
	   </script>
	*/
	MinHeight int `jsObject:"minHeight"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-minWidth

	  The minimum width (in pixels) that may be achieved by resizing the dialog.

	  Example
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     minWidth: 100
	   });
	   </script>
	*/
	MinWidth int `jsObject:"minWidth"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-modal

	  Specifies whether the dialog should show a modal overlay over the page.

	  Example
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     modal: true
	   });
	   </script>
	*/
	Modal Boolean `jsObject:"modal"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-title

	  The text in the dialog title bar. If <b><u>false</u></b>, the dialog will be displayed without a title bar.

	  Example
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     title: "Customer details"
	   });
	   </script>
	*/
	Title string `jsObject:"title"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-visible

	  Specifies whether the dialog will be initially visible.

	  Example - show a dialog after one second delay
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     visible: false
	   });
	   setTimeout(function() {
	     $("#dialog").data("kendoDialog").open();
	   }, 1000);
	   </script>
	*/
	Visible Boolean `jsObject:"visible"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog#configuration-width

	  Specifies width of the dialog.

	  Example
	   <div id="dialog"></div>
	   <script>
	   $("#dialog").kendoDialog({
	     width: 400
	   });
	   </script>
	*/
	Width int `jsObject:"width"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/events/close#close

	  Triggered when a Dialog is closed (by a user or through the close() method).

	  Event Data: e.userTriggered - Boolean
	  Indicates whether the close action has been triggered by the user (by clicking the close button or hitting the escape key). When the close method has been called, this field is false.

	  Example - subscribe to the "close" event during initialization
	  <div id="dialog"></div>
	  <script>
	  $("#dialog").kendoDialog({
	    close: function(e) {
	      // close animation has finished playing
	    }
	  });
	  </script>

	  Example - subscribe to the "close" event after initialization
	  <div id="dialog"></div>
	  <script>
	    function dialog_close(e) {
	      // close animation has finished playing
	    }
	    $("#dialog").kendoDialog();
	    var dialog = $("#dialog").data("kendoDialog");
	    dialog.bind("close", dialog_close);
	  </script>
	*/
	EventClose JavaScript `jsObject:"close"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/events/hide#hide

	  Triggered when a Dialog has finished its closing animation.

	  Example - subscribe to the "hide" event during initialization
	  <div id="dialog"></div>
	  <script>
	  $("#dialog").kendoDialog({
	    hide: function() {
	      // close animation is about to finish
	    }
	  });
	  </script>

	  Example - subscribe to the "hide" event after initialization
	  <div id="dialog"></div>
	  <script>
	    function dialog_hide() {
	      // close animation will start soon
	    }
	    $("#dialog").kendoDialog();
	    var dialog = $("#dialog").data("kendoDialog");
	    dialog.bind("hide", dialog_hide);
	  </script>
	*/
	EventHide JavaScript `jsObject:"hide"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/events/initopen#initOpen

	  Triggered when a Dialog is opened for the first time (i.e. the open() method is called).

	  Example - subscribe to the "initOpen" event during initialization
	  <div id="dialog"></div>
	  <script>
	  $("#dialog").kendoDialog({
	    initOpen: function() {
	      // open animation will start soon
	    }
	  });
	  </script>

	  Example - subscribe to the "initOpen" event after initialization
	  <div id="dialog" style="display: none;"></div>
	  <script>
	    function dialog_initOpen() {
	      // open animation will start soon
	    }
	    $("#dialog").kendoDialog();
	    var dialog = $("#dialog").data("kendoDialog");
	    dialog.bind("initOpen", dialog_initOpen);
	    dialog.open();
	  </script>
	*/
	EventInitOpen JavaScript `jsObject:"initOpen"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/events/open#open

	  Triggered when a Dialog is opened (i.e. the open() method is called).

	  Example - subscribe to the "open" event during initialization
	  <div id="dialog"></div>
	  <script>
	  $("#dialog").kendoDialog({
	    open: function() {
	      // open animation will start soon
	    }
	  });
	  </script>

	  Example - subscribe to the "open" event after initialization
	  <div id="dialog"></div>
	  <script>
	    function dialog_open() {
	      // open animation will start soon
	    }
	    $("#dialog").kendoDialog();
	    var dialog = $("#dialog").data("kendoDialog");
	    dialog.bind("open", dialog_open);
	  </script>
	*/
	EventOpen JavaScript `jsObject:"open"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/events/show#show

	  Triggered when a Dialog has finished its opening animation.

	  Example - subscribe to the "show" event during initialization
	  <div id="dialog"></div>
	  <script>
	  $("#dialog").kendoDialog({
	    show: function() {
	      // open animation has finished playing
	    }
	  });
	  </script>

	  Example - subscribe to the "show" event after initialization
	  <div id="dialog"></div>
	  <script>
	    function dialog_show() {
	      // open animation has finished playing
	    }
	    $("#dialog").kendoDialog();
	    var dialog = $("#dialog").data("kendoDialog");
	    dialog.bind("show", dialog_show);
	  </script>
	*/
	EventShow JavaScript `jsObject:"show"`

	*ToJavaScriptConverter
}

func (el *KendoUiDialog) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("kendoDialog.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoDialog({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}

func (el *KendoUiDialog) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiDialog) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}

/*
@see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/methods/close#close

Closes a Dialog.

Returns: kendo.ui.Dialog Returns the dialog object to support chaining.

Example - close a dialog after one second
<div id="dialog"></div>
<script>
  $("#dialog").kendoDialog();
  var dialog = $("#dialog").data("kendoDialog");
  setTimeout(function() {
    dialog.close();
  }, 1000);
</script>
*/
func (el *KendoUiDialog) MethodClose() []byte {
	if el.GlobalVar == TRUE {
		return []byte(`KendoUiGlobalVarFromId` + el.Html.Global.Id + `.close();`)
	}

	return []byte(`$("#` + el.Html.Global.Id + `").data("kendoDialog").close();`)
}

/*
@see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/methods/content#content

content: Gets or set the content of a dialog. Supports chaining when used as a setter.

Parameters: content String |jQuery (optional)
The content of the Dialog. Can be an HTML string or jQuery object.

Returns: String The current dialog content, if used as a getter. If used as a setter, the method will return the dialog
object to support chaining.

Example - get the dialog content
<div id="dialog">foo</div>
<script>
  $("#dialog").kendoDialog();
  var dialog = $("#dialog").data("kendoDialog");
  console.log(dialog.content()); // logs "foo"
</script>

Example - set the dialog content
<div id="dialog"></div>
<script>
  $("#dialog").kendoDialog();
  var dialog = $("#dialog").data("kendoDialog");
  dialog.content("Kendo UI all the things!");
</script>
*/
func (el *KendoUiDialog) MethodContent(content interface{}) []byte {
	if el.GlobalVar == TRUE {
		switch converted := content.(type) {
		case string:
			return []byte(`KendoUiGlobalVarFromId` + el.Html.Global.Id + `.content("` + converted + `");`)
		case JavaScript:
			return []byte(`KendoUiGlobalVarFromId` + el.Html.Global.Id + `.content(` + converted.String() + `);`)
		case *JavaScript:
			return []byte(`KendoUiGlobalVarFromId` + el.Html.Global.Id + `.content(` + converted.String() + `);`)
		}
	}

	switch converted := content.(type) {
	case string:
		return []byte(`$("#` + el.Html.Global.Id + `").data("kendoDialog").content("` + converted + `");`)
	case JavaScript:
		return []byte(`$("#` + el.Html.Global.Id + `").data("kendoDialog").content(` + converted.String() + `);`)
	case *JavaScript:
		return []byte(`$("#` + el.Html.Global.Id + `").data("kendoDialog").content(` + converted.String() + `);`)
	}
	return []byte(`content must be a string, JavaScript ou *JavaScript`)
}

/*
@see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/methods/destroy#destroy

destroy: Destroys the dialog and its modal overlay, if necessary. Removes the widget HTML elements from the DOM.

Example
<div id="dialog"></div>
<script>
  $("#dialog").kendoDialog();
  var dialog = $("#dialog").data("kendoDialog");
  dialog.destroy();
</script>
*/
func (el *KendoUiDialog) MethodDestroy() []byte {
	if el.GlobalVar == TRUE {
		return []byte(`KendoUiGlobalVarFromId` + el.Html.Global.Id + `.destroy();`)
	}

	return []byte(`$("#` + el.Html.Global.Id + `").data("kendoDialog").destroy();`)
}

/*
@see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/methods/open#open

open: Opens a Dialog and brings it on top of any other open Dialog or Window instances by calling toFront internally.

Returns: kendo.ui.Dialog Returns the dialog object to support chaining.

Example
<div id="dialog"></div>
<script>
  $("#dialog").kendoDialog({
    visible: false
  });
  var dialog = $("#dialog").data("kendoDialog");
  dialog.open();
</script>
*/
func (el *KendoUiDialog) MethodOpen() []byte {
	if el.GlobalVar == TRUE {
		return []byte(`KendoUiGlobalVarFromId` + el.Html.Global.Id + `.open();`)
	}

	return []byte(`$("#` + el.Html.Global.Id + `").data("kendoDialog").open();`)
}

/*
@see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/methods/title#title

title: String (optional) Gets or sets the title of a Dialog. Can be a text string. Supports chaining when used as a
setter. If passed to the method, an HTML string would be escaped.

Parameters: text String (optional)
The title of the Dialog.

Returns: String The current dialog title, if used as a getter. If used as a setter, the method will return the dialog
object to support chaining.

Example - get the title of the dialog
<div id="dialog"></div>
<script>
  $("#dialog").kendoDialog();
  var dialog = $("#dialog").data("kendoDialog");
  var title = dialog.title();
</script>

Example - set the title of a dialog
<div id="dialog"></div>
<script>
  $("#dialog").kendoDialog();
  var dialog = $("#dialog").data("kendoDialog");
  dialog.title("<em>Hello</em>");
</script>
*/
func (el *KendoUiDialog) MethodTitle(title interface{}) []byte {
	if el.GlobalVar == TRUE {
		switch converted := title.(type) {
		case string:
			return []byte(`KendoUiGlobalVarFromId` + el.Html.Global.Id + `.content("` + converted + `");`)
		case JavaScript:
			return []byte(`KendoUiGlobalVarFromId` + el.Html.Global.Id + `.content(` + converted.String() + `);`)
		case *JavaScript:
			return []byte(`KendoUiGlobalVarFromId` + el.Html.Global.Id + `.content(` + converted.String() + `);`)
		}
	}

	switch converted := title.(type) {
	case string:
		return []byte(`$("#` + el.Html.Global.Id + `").data("kendoDialog").content("` + converted + `");`)
	case JavaScript:
		return []byte(`$("#` + el.Html.Global.Id + `").data("kendoDialog").content(` + converted.String() + `);`)
	case *JavaScript:
		return []byte(`$("#` + el.Html.Global.Id + `").data("kendoDialog").content(` + converted.String() + `);`)
	}
	return []byte(`title must be a string, JavaScript ou *JavaScript`)
}

/*
@see https://docs.telerik.com/kendo-ui/api/javascript/ui/dialog/methods/tofront#toFront

toFront
Increases the z-index style of a Dialog wrapper to bring the instance on top of other open Dialogs. This method is executed automatically when the open method is used.

Returns
kendo.ui.Dialog Returns the dialog object to support chaining.

Example
<div id="dialog"></div>
<script>
  $("#dialog").kendoDialog();
  var dialog = $("#dialog").data("kendoDialog");
  dialog.toFront();
</script>
*/
func (el *KendoUiDialog) MethodToFront() []byte {
	if el.GlobalVar == TRUE {
		return []byte(`KendoUiGlobalVarFromId` + el.Html.Global.Id + `.toFront();`)
	}

	return []byte(`$("#` + el.Html.Global.Id + `").data("kendoDialog").toFront();`)
}

func (el *KendoUiDialog) GetGlobalVar() []byte {
	el.GlobalVar = TRUE

	return []byte(`var KendoUiGlobalVarFromId` + el.Html.Global.Id + ` = $("#` + el.Html.Global.Id + `").data("kendoDialog");`)
}
func (el *KendoUiDialog) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
