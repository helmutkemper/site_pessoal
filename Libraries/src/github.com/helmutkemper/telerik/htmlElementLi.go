package telerik

import (
	"bytes"
	"reflect"
)

// The HTML <li> element is used to represent an item in a list.
// It must be contained in a parent element: an ordered list (<ol>), an unordered list (<ul>), or a menu (<menu>).
// In menus and unordered lists, list items are usually displayed using bullet points. In ordered lists, they are
// usually displayed with an ascending counter on the left, such as a number or letter.
//
// @see https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li
type HtmlElementLi struct {
	Name string `htmlAttr:"-"`
	/*
	  Content inside html tag
	*/
	Content Content `htmlAttr:"-"`

	Global HtmlGlobalAttributes `htmlAttr:"-"`

	*ToJavaScriptConverter `htmlAttr:"-"`
}

func (el *HtmlElementLi) SetOmitHtml(value Boolean) {
	el.Global.DoNotUseThisFieldOmitHtml = value
}
func (el *HtmlElementLi) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Global.DoNotUseThisFieldOmitHtml == TRUE {
		return []byte{}
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	buffer.Write([]byte(`<li`))
	buffer.Write(el.Global.ToHtml())
	buffer.Write(data)
	buffer.Write([]byte(`>`))
	buffer.Write(el.Content.ToHtml())
	buffer.Write([]byte(`</li>`))

	return buffer.Bytes()
}
func (el *HtmlElementLi) ToJavaScript() []byte {
	var buffer bytes.Buffer

	buffer.Write(el.Content.ToJavaScript())

	return buffer.Bytes()
}

func (el *HtmlElementLi) ToHtmlSupport() []byte {
	return el.Content.ToHtmlSupport()
}
