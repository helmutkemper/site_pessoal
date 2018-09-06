package telerik

import (
	"bytes"
	"reflect"
)

// The HTML <ul> element represents an unordered list of items, typically rendered as a bulleted list.
//
// @see https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul
type HtmlElementUl struct {
	Name string `htmlAttr:"-"`
	/*
	  Content inside html tag
	*/
	Content Content `htmlAttr:"-"`

	Global HtmlGlobalAttributes `htmlAttr:"-"`

	*ToJavaScriptConverter `htmlAttr:"-"`
}

func (el *HtmlElementUl) SetOmitHtml(value Boolean) {
	el.Global.DoNotUseThisFieldOmitHtml = value
}
func (el *HtmlElementUl) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Global.DoNotUseThisFieldOmitHtml == TRUE {
		return []byte{}
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	buffer.Write([]byte(`<ul`))
	buffer.Write(el.Global.ToHtml())
	buffer.Write(data)
	buffer.Write([]byte(`>`))
	buffer.Write(el.Content.ToHtml())
	buffer.Write([]byte(`</ul>`))

	return buffer.Bytes()
}
func (el *HtmlElementUl) ToJavaScript() []byte {
	var buffer bytes.Buffer

	buffer.Write(el.Content.ToJavaScript())

	return buffer.Bytes()
}

func (el *HtmlElementUl) ToHtmlSupport() []byte {
	return el.Content.ToHtmlSupport()
}
