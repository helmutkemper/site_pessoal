package telerik

import (
	"bytes"
	"reflect"
	"sort"
)

// The HTML <optgroup> element creates a grouping of options within a <select> element.
type HtmlOptGroup struct {
	/*
	  The name of the group of options, which the browser can use when labeling the options in the user interface. This
	  attribute is mandatory if this element is used.
	*/
	Label string `htmlAttrSet:"label"`

	Options []HtmlOptions `htmlAttr:"-"`

	/*
	  This Boolean attribute indicates that the form control is not available for interaction. In particular, the click
	  event will not be dispatched on disabled controls. Also, a disabled control's value isn't submitted with the form.
	  Unlike other browsers, Firefox will by default persist the dynamic disabled state of an <input> across page loads. Use
	  the autocomplete attribute to control this feature.
	*/
	Disabled Boolean `htmlAttrSet:"disabled"`

	Global HtmlGlobalAttributes `htmlAttr:"-"`

	*ToJavaScriptConverter `htmlAttr:"-"`
}

func (el *HtmlOptGroup) SetOmitHtml(value Boolean) {
	el.Global.DoNotUseThisFieldOmitHtml = value
}
func (el *HtmlOptGroup) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Global.DoNotUseThisFieldOmitHtml == TRUE {
		return []byte{}
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	buffer.Write([]byte(`<optgroup`))
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
		buffer.Write([]byte(` value="`))
		buffer.WriteString(el.Options[k].Key)
		buffer.Write([]byte(`">`))
		buffer.WriteString(el.Options[k].Label)
		buffer.Write([]byte(`</option>`))
	}

	buffer.Write([]byte(`</optgroup>`))

	return buffer.Bytes()
}
