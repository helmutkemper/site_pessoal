package telerik

import (
	"bytes"
	"reflect"
)

// The HTML <span> element is a generic inline container for phrasing content, which does not inherently represent
// anything. It can be used to group elements for styling purposes (using the class or id attributes), or because they
// share attribute values, such as lang. It should be used only when no other semantic element is appropriate. <span>
// is very much like a <div> element, but <div> is a block-level element whereas a <span> is an inline element.
//
// @see https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span
type HtmlElementSpan struct {

	/*
	  Content inside html tag
	*/
	Content Content `htmlAttr:"-"`

	Global HtmlGlobalAttributes `htmlAttr:"-"`

	*ToJavaScriptConverter `htmlAttr:"-"`
}

func (el *HtmlElementSpan) SetOmitHtml(value Boolean) {
	el.Global.DoNotUseThisFieldOmitHtml = value
}
func (el *HtmlElementSpan) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Global.DoNotUseThisFieldOmitHtml == TRUE {
		return []byte{}
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	buffer.Write([]byte(`<span`))
	buffer.Write(el.Global.ToHtml())
	buffer.Write(data)
	buffer.Write([]byte(`>`))
	buffer.Write(el.Content.ToHtml())
	buffer.Write([]byte(`</span>`))

	return buffer.Bytes()
}
func (el *HtmlElementSpan) ToJavaScript() []byte {
	var buffer bytes.Buffer

	buffer.Write(el.Content.ToJavaScript())

	return buffer.Bytes()
}
