package telerik

import (
	"bytes"
	"reflect"
)

// fixme incompleto
// The HTML <script> element is used to embed or reference executable code; this is typically used to embed or refer to
// JavaScript code. The <script> element can also be used with other languages, such as WebGL's GLSL shader programming
// language.
//
// @see https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
type HtmlElementScript struct {

	/*
	  This attribute indicates the type of script represented. The value of this attribute will be in one of the following
	  categories:

	    Omitted or a JavaScript MIME type: For HTML5-compliant browsers this indicates the script is JavaScript. HTML5
	    specification urges authors to omit the attribute rather than provide a redundant MIME type. In earlier browsers,
	    this identified the scripting language of the embedded or imported (via the src attribute) code. JavaScript MIME
	    types are listed in the specification.

	    module: HTML5 For HTML5-compliant browsers the code is treated as a JavaScript module. The processing of the script
	    contents is not affected by the charset and defer attributes. For information on using module, see ES6 in Depth:
	    Modules.

	    Any other value: The embedded content is treated as a data block which won't be processed by the browser. Developers
	    must use a valid MIME type that is not a JavaScript MIME type to denote data blocks. The src attribute will be
	    ignored.

	  Note:  in Firefox you could specify the version of JavaScript contained in a <script> element by including a
	  non-standard version parameter inside the type attribute — for example  type="application/javascript;version=1.8".
	  This has been removed in Firefox 59 (see bug 1428745).

	  Deprecated attributes:

	    language: Like the type attribute, this attribute identifies the scripting language in use. Unlike the type
	    attribute, however, this attribute’s possible values were never standardized. The type attribute should be used
	    instead.
	*/
	Type HtmlScriptType `htmlAttr:"type" required:"true" requiredMessage:"Please, set script type for all elements 'telerik.HtmlElementScript{}'"`
	// fixme: fazer

	/*
	  Content inside html tag
	*/
	Content Content `htmlAttr:"-"`

	Global HtmlGlobalAttributes `htmlAttr:"-"`

	*ToJavaScriptConverter `htmlAttr:"-"`
}

func (el *HtmlElementScript) SetOmitHtml(value Boolean) {
	el.Global.DoNotUseThisFieldOmitHtml = value
}
func (el *HtmlElementScript) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Global.DoNotUseThisFieldOmitHtml == TRUE {
		return []byte{}
	}

	if el.Type == SCRIPT_TYPE_JAVASCRIPT {
		return []byte{}
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	buffer.Write([]byte(`<script`))
	buffer.Write(el.Global.ToHtml())
	buffer.Write(data)
	buffer.Write([]byte(`>`))
	buffer.Write(el.Content.ToHtml())
	buffer.Write([]byte(`</script>`))

	return buffer.Bytes()
}
func (el *HtmlElementScript) ToElementScriptTag() []byte {
	var buffer bytes.Buffer

	if el.Global.Id == "" {
		el.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	buffer.Write([]byte(`<script`))
	buffer.Write(el.Global.ToHtml())
	buffer.Write(data)
	buffer.Write([]byte(`>`))
	buffer.Write(el.Content.ToHtml())
	buffer.Write([]byte(`</script>`))

	return buffer.Bytes()
}
func (el *HtmlElementScript) ToJavaScript() []byte {
	var buffer bytes.Buffer

	if el.Type == SCRIPT_TYPE_KENDO_TEMPLATE {
		return []byte{}
	}

	buffer.Write([]byte(`<script>`))
	buffer.Write(el.Content.ToJavaScript())
	buffer.Write([]byte(`</script>`))

	return buffer.Bytes()
}
func (el *HtmlElementScript) ToKendoTemplate() []byte {
	var buffer bytes.Buffer

	if el.Global.Id == "" {
		el.Global.Id = GetAutoId()
	}

	buffer.Write([]byte(`kendo.template($('#`))
	buffer.Write([]byte(el.Global.Id))
	buffer.Write([]byte(`').html())`))

	return buffer.Bytes()
}
func (el *HtmlElementScript) GetId() []byte {
	if el.Global.Id == "" {
		el.Global.Id = GetAutoId()
	}
	return []byte(el.Global.Id)
}

//containerCreateTemplateExposedPortsAddNewPort
