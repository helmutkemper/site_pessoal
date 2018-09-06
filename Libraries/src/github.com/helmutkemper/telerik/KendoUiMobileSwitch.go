package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUiMobileSwitch struct {
	Html HtmlInputCheckBox `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/mobile/ui/switch/configuration/checked

	  The checked state of the widget. (default: false)

	  Example
	  <div id="foo" data-role="view">
	    <input type="checkbox" data-role="switch" data-checked="false" />
	    <input type="checkbox" data-role="switch" data-checked="true" />
	  </div>

	  <script>
	  var app = new kendo.mobile.Application();
	  </script>
	*/
	Checked Boolean `jsObject:"checked"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/mobile/ui/switch/configuration/enable

	  If set to false the widget will be disabled and will not allow the user to change its checked state. The widget is
	  enabled by default. (default: true)

	  Example - initialize disabled switch.
	  <div data-role="view">
	    <input type="checkbox" data-role="switch" data-enable="false" />
	  </div>

	  <script>
	    var app = new kendo.mobile.Application();
	  </script>
	*/
	Enable Boolean `jsObject:"enable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/mobile/ui/switch/configuration/offlabel

	  The OFF label. (default: "OFF")

	  Example
	  <div id="foo" data-role="view">
	    <input type="checkbox" data-role="switch" data-off-label="No" data-on-label="Yes" />
	  </div>

	  <script>
	    var app = new kendo.mobile.Application();
	  </script>
	*/
	OffLabel string `jsObject:"offLabel"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/mobile/ui/switch/configuration/onlabel

	  The ON label. (default: "ON")

	  Example
	  <div id="foo" data-role="view">
	    <input type="checkbox" data-role="switch" data-off-label="No" data-on-label="Yes" />
	  </div>

	  <script>
	    var app = new kendo.mobile.Application();
	  </script>
	*/
	OnLabel string `jsObject:"onLabel"`

	*ToJavaScriptConverter
}

func (el *KendoUiMobileSwitch) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("kendoNumericTextBox.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoMobileSwitch({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiMobileSwitch) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiMobileSwitch) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
