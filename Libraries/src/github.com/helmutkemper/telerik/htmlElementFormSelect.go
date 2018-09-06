package telerik

import (
	"bytes"
	"reflect"
	"sort"
)

// The HTML <select> element represents a control that provides a menu
// Fixme: rever!
// @see https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
type HtmlElementFormSelect struct {
	/*
	  The name of the control, which is submitted with the form data.
	*/
	Name string `htmlAttr:"name"`

	/*
	  The form element that the input element is associated with (its form owner). The value of the attribute must be an id
	  of a <form> element in the same document. If this attribute is not specified, this <input> element must be a
	  descendant of a <form> element. This attribute enables you to place <input> elements anywhere within a document, not
	  just as descendants of their form elements. An input can only be associated with one form.
	*/
	Form string `htmlAttr:"form"`

	/*
	  This Boolean attribute indicates that the form control is not available for interaction. In particular, the click
	  event will not be dispatched on disabled controls. Also, a disabled control's value isn't submitted with the form.
	  Unlike other browsers, Firefox will by default persist the dynamic disabled state of an <input> across page loads. Use
	  the autocomplete attribute to control this feature.
	*/
	Disabled Boolean `htmlAttrSet:"disabled"`

	/*
	  This Boolean attribute indicates whether the user can enter more than one value. This attribute applies when the type
	  attribute is set to email or file, otherwise it is ignored.
	*/
	Multiple string `htmlAttr:"multiple"`

	/*
	  This attribute specifies that the user must fill in a value before submitting a form. It cannot be used when the type
	  attribute is hidden, image, or a button type (submit, reset, or button). The :optional and :required CSS
	  pseudo-classes will be applied to the field as appropriate.
	*/
	Required Boolean `htmlAttrSet:"required"`

	/*
	  The initial size of the control. This value is in pixels unless the value of the type attribute is text or password,
	  in which case it is an integer number of characters. Starting in HTML5, this attribute applies only when the type
	  attribute is set to text, search, tel, url, email, or password, otherwise it is ignored. In addition, the size must be
	  greater than zero. If you do not specify a size, a default value of 20 is used. HTML5 simply states "the user agent
	  should ensure that at least that many characters are visible", but different characters can have different widths in
	  certain fonts. In some browsers, a certain string with x characters will not be entirely visible even if size is
	  defined to at least x.
	*/
	Size int `htmlAttr:"size"`

	/*
	  The initial value of the control. This attribute is optional except when the value of the type attribute is radio or
	  checkbox.
	  Note that when reloading the page, Gecko and IE will ignore the value specified in the HTML source, if the value was
	  changed before the reload.
	*/
	Value string `htmlAttr:"-"`

	Options []HtmlOptions `htmlAttr:"-"`
	//OptionsGroup                           `htmlAttr:"options"`

	Global HtmlGlobalAttributes `htmlAttr:"-"`

	*ToJavaScriptConverter `htmlAttr:"-"`
}

func (el *HtmlElementFormSelect) SetOmitHtml(value Boolean) {
	el.Global.DoNotUseThisFieldOmitHtml = value
}
func (el *HtmlElementFormSelect) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Global.DoNotUseThisFieldOmitHtml == TRUE {
		return []byte{}
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	buffer.Write([]byte(`<select`))
	buffer.Write(el.Global.ToHtml())
	buffer.Write(data)
	buffer.Write([]byte(`>`))

	// ordena as chaves
	keys := make([]int, 0)
	for k := range el.Options {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for k := range keys {
		buffer.Write([]byte(`<option`))

		if el.Value != "" && el.Value == el.Options[k].Key {
			buffer.Write([]byte(` selected`))
		}

		buffer.Write([]byte(` value="`))
		buffer.WriteString(el.Options[k].Key)

		if el.Options[k].Label == "" {
			buffer.Write([]byte(`"/>`))
		} else {
			buffer.Write([]byte(`">`))
			buffer.WriteString(el.Options[k].Label)
			buffer.Write([]byte(`</option>`))
		}
	}
	buffer.WriteString(`</select>`)

	return buffer.Bytes()
}
func (el *HtmlElementFormSelect) ToJavaScript() []byte {
	return []byte{}
}
func (el *HtmlElementFormSelect) GetId() []byte {
	if el.Global.Id == "" {
		el.Global.Id = GetAutoId()
	}
	return []byte(el.Global.Id)
}
func (el *HtmlElementFormSelect) GetName() []byte {
	if el.Name == "" {
		el.Name = GetAutoId()
	}
	return []byte(el.Name)
}
