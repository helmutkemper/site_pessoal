package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoUiButton struct {
	Html HtmlElementFormButton `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/button#configuration-enable

	  Indicates whether the <strong>Button</strong> should be enabled or disabled. By default, it is enabled, unless a <b><u>disabled="disabled"</u></b> attribute is detected.

	  Example
	   <button id="button" type="button">Foo</button>
	   <script>
	   $("#button").kendoButton({
	       enable: false
	   });
	   </script>
	*/
	Enable Boolean `jsObject:"enable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/button#configuration-icon

	  Defines a name of an existing icon in the Kendo UI theme sprite. The icon will be applied as background image of a <b><u>span</u></b> element inside the <strong>Button</strong>. The <b><u>span</u></b> element can be added automatically by the widget, or an existing element can be used, if it has a <b><u>k-icon</u></b> CSS class applied. For a list of available icon names, please refer to the <a href="http://demos.telerik.com/kendo-ui/web/styling/icons.html">Icons demo</a>.

	  Example
	   <button id="button" type="button">Cancel</button>
	   <script>
	   $("#button").kendoButton({
	       icon: "cancel"
	   });
	   </script>
	*/
	Icon string `jsObject:"icon"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/button#configuration-imageUrl

	  Defines a URL, which will be used for an <b><u>img</u></b> element inside the Button. The URL can be relative or absolute. In case it is relative, it will be evaluated with relation to the web page URL.
	  The <b><u>img</u></b> element can be added automatically by the widget, or an existing element can be used, if it has a <b><u>k-image</u></b> CSS class applied.

	  Example
	   <button id="button" type="button">Edit</button>
	   <script>
	   $("#button").kendoButton({
	       imageUrl: "/images/edit-icon.gif"
	   });
	   </script>
	*/
	ImageUrl string `jsObject:"imageUrl"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/button#configuration-spriteCssClass

	  Defines a CSS class (or multiple classes separated by spaces), which will be used for applying a background image to a <b><u>span</u></b> element inside the <strong>Button</strong>. In case you want to use an icon from the Kendo UI theme sprite background image, it is easier to use the <a href="https://docs.telerik.com/kendo-ui/api/javascript/ui/button#configuration-icon"><b><u>icon</u></b> property</a>.
	  The <b><u>span</u></b> element can be added automatically by the widget, or an existing element can be used, if it has a <b><u>k-sprite</u></b> CSS class applied.

	  Example
	   <button id="button" type="button">Edit</button>
	   <script>
	   $("#button").kendoButton({
	       spriteCssClass: "myEditIcon"
	   });
	   </script>
	*/

	SpriteCssClass string `jsObject:"spriteCssClass"`

	*ToJavaScriptConverter
}

func (el *KendoUiButton) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("kendoButton.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoButton({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiButton) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiButton) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiButton) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
