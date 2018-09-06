package telerik

import (
	"bytes"
	"reflect"
)

// The HTML Content Division element (<div>) is the generic container for flow content. It has no effect on the content
// or layout until styled using CSS. As a "pure" container, the <div> element does not inherently represent anything.
// Instead, it's used to group content so it can be easily styled using the class or id attributes, marking a section of
// a document as being written in a different language (using the lang attribute), and so on.
//
// @see https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
type HtmlElementDiv struct {
	Name string `htmlAttr:"-"`
	/*
	  Content inside html tag
	*/
	Content Content `htmlAttr:"-"`

	Global HtmlGlobalAttributes `htmlAttr:"-"`

	*ToJavaScriptConverter `htmlAttr:"-"`
}

func (el *HtmlElementDiv) SetOmitHtml(value Boolean) {
	el.Global.DoNotUseThisFieldOmitHtml = value
}
func (el *HtmlElementDiv) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Global.DoNotUseThisFieldOmitHtml == TRUE {
		return []byte{}
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	buffer.Write([]byte(`<div`))
	buffer.Write(el.Global.ToHtml())
	buffer.Write(data)
	buffer.Write([]byte(`>`))
	buffer.Write(el.Content.ToHtml())
	buffer.Write([]byte(`</div>`))

	return buffer.Bytes()
}
func (el *HtmlElementDiv) ToJavaScript() []byte {
	var buffer bytes.Buffer

	buffer.Write(el.Content.ToJavaScript())

	return buffer.Bytes()
}

func (el *HtmlElementDiv) ToHtmlSupport() []byte {
	return el.Content.ToHtmlSupport()
}
