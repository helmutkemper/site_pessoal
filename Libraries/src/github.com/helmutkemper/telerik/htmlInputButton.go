package telerik

import (
	"bytes"
	"reflect"
)

// <input> elements of type "button"  are rendered as simple push buttons, which can be programmed to control custom
// functionality anywhere on a webpage as required when assigned an event handler function (typically for the click
// event).
//
// <input type="button" value="Click Me">
type HtmlInputButton struct {
	/*
	  The name of the control, which is submitted with the form data.
	  @see typeNamesForAutocomplete.go
	  Ex.: const NAMES_FOR_AUTOCOMPLETE_NAME
	*/
	Name string `htmlAttr:"name"`

	/*
	  The initial value of the control. This attribute is optional except when the value of the type attribute is radio or
	  checkbox.
	  Note that when reloading the page, Gecko and IE will ignore the value specified in the HTML source, if the value was
	  changed before the reload.
	*/
	Value string `htmlAttr:"value"`

	/*
	  The form element that the input element is associated with (its form owner). The value of the attribute must be an id
	  of a <form> element in the same document. If this attribute is not specified, this <input> element must be a
	  descendant of a <form> element. This attribute enables you to place <input> elements anywhere within a document, not
	  just as descendants of their form elements. An input can only be associated with one form.
	*/
	Form string `htmlAttr:"from"`

	/*
	  This Boolean attribute indicates that the form control is not available for interaction. In particular, the click
	  event will not be dispatched on disabled controls. Also, a disabled control's value isn't submitted with the form.
	  Unlike other browsers, Firefox will by default persist the dynamic disabled state of an <input> across page loads. Use
	  the autocomplete attribute to control this feature.
	*/
	Disabled Boolean `htmlAttrSet:"disabled"`

	Global HtmlGlobalAttributes

	*ToJavaScriptConverter
}

func (el *HtmlInputButton) SetOmitHtml(value Boolean) {
	el.Global.DoNotUseThisFieldOmitHtml = value
}
func (el *HtmlInputButton) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Global.DoNotUseThisFieldOmitHtml == TRUE {
		return []byte{}
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	//return `<input ` + el.Global.String() + ` type="button" ` + el.Name.ToAttr("name") + el.Value.ToAttr("value") + el.Form.ToAttr("form") + el.Disabled.ToAttrSet("disabled") + `>`

	buffer.Write([]byte(`<input`))
	buffer.Write(el.Global.ToHtml())
	buffer.Write([]byte(` type="button" `))
	buffer.Write(data)
	buffer.Write([]byte(`>`))

	return buffer.Bytes()
}
