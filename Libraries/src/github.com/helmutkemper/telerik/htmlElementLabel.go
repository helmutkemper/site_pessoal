package telerik

import (
	"bytes"
	"reflect"
)

// The HTML <label> element represents a caption for an item in a user interface.
type HtmlElementFormLabel struct {
	/*
	  The form element that the input element is associated with (its form owner). The value of the attribute must be an id
	  of a <form> element in the same document. If this attribute is not specified, this <input> element must be a
	  descendant of a <form> element. This attribute enables you to place <input> elements anywhere within a document, not
	  just as descendants of their form elements. An input can only be associated with one form.
	*/
	Form string `htmlAttr:"form"`

	/*
	  The id of a labelable form-related element in the same document as the label element. The first such element in the
	  document with an ID matching the value of the for attribute is the labeled control for this label element.
	  A label element can have both a for attribute and a contained control element, as long as the for attribute points to
	  the contained control element.
	*/
	For string `htmlAttr:"for"`

	/*
	  Content inside html tag
	*/
	Content Content `htmlAttr:"-"`

	Global HtmlGlobalAttributes `htmlAttr:"-"`

	*ToJavaScriptConverter `htmlAttr:"-"`
}

func (el *HtmlElementFormLabel) SetOmitHtml(value Boolean) {
	el.Global.DoNotUseThisFieldOmitHtml = value
}
func (el *HtmlElementFormLabel) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Global.DoNotUseThisFieldOmitHtml == TRUE {
		return []byte{}
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	buffer.Write([]byte(`<label`))
	buffer.Write(el.Global.ToHtml())
	buffer.Write(data)
	buffer.Write([]byte(`>`))
	buffer.Write(el.Content.ToHtml())
	buffer.Write([]byte(`</label>`))

	return buffer.Bytes()
}
func (el *HtmlElementFormLabel) ToJavaScript() []byte {
	return []byte{}
}
