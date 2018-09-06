package telerik

import (
	"bytes"
	"reflect"
)

// The HTML <button> element represents a clickable button, which can be used in forms, or anywhere in a document that
// needs simple, standard button functionality. By default, HTML buttons are typically presented in a style similar to
// that of the host platform the user agent is running on, but you can change the appearance of the button using CSS.
type HtmlElementFormButton struct {
	ButtonType SupportCustomButtonType `htmlAttr:"-"`

	/*
	  The name of the control, which is submitted with the form data.
	*/
	Name string `htmlAttr:"name"`

	/*
	  Content inside html tag
	*/
	Content Content `htmlAttr:"-"`

	/*
	  The form element that the input element is associated with (its form owner). The value of the attribute must be an id
	  of a <form> element in the same document. If this attribute is not specified, this <input> element must be a
	  descendant of a <form> element. This attribute enables you to place <input> elements anywhere within a document, not
	  just as descendants of their form elements. An input can only be associated with one form.
	*/
	Form string `htmlAttr:"form"`

	/*
	  The URI of a program that processes the information submitted by the button. If specified, it overrides the action
	  attribute of the button's form owner.
	*/
	FormAction string `htmlAttr:"action"`

	/*
	  If the button is a submit button, this attribute specifies the type of content that is used to submit the form to the
	  server. Possible values are:
	  > application/x-www-form-urlencoded: The default value if the attribute is not specified.
	  > multipart/form-data: Use this value if you are using an <input> element with the type attribute set to file.
	  > text/plain
	*/
	FormEncType string `htmlAttr:"enctype"`

	/*
	  If the button is a submit button, this attribute specifies the HTTP method that the browser uses to submit the form.
	  Possible values are:
	  > post: The data from the form are included in the body of the form and sent to the server.
	  > get: The data from the form are appended to the form attribute URI, with a '?' as a separator, and the resulting URI
	  is sent to the server. Use this method when the form has no side-effects and contains only ASCII characters.
	*/
	FormMethod string `htmlAttr:"method"`

	/*
	  If the button is a submit button, this Boolean attribute specifies that the form is not to be validated when it is
	  submitted. If this attribute is specified, it overrides the novalidate attribute of the button's form owner.
	*/
	FormNoValidate Boolean `htmlAttr:"novalidate"`

	/*
	  If the button is a submit button, this attribute is a name or keyword indicating where to display the response that is
	  received after submitting the form. This is a name of, or keyword for, a browsing context (for example, tab, window,
	  or inline frame). If this attribute is specified, it overrides the target attribute of the button's form owner. The
	  following keywords have special meanings:
	  > _self: Load the response into the same browsing context as the current one. This value is the default if the
	  attribute is not specified.
	  > _blank: Load the response into a new unnamed browsing context.
	  > _parent: Load the response into the parent browsing context of the current one. If there is no parent, this option
	  behaves the same way as _self.
	  > _top: Load the response into the top-level browsing context (that is, the browsing context that is an ancestor of
	  the current one, and has no parent). If there is no parent, this option behaves the same way as _self.
	*/
	FormTarget string `htmlAttr:"target"`

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

func (el *HtmlElementFormButton) SetOmitHtml(value Boolean) {
	el.Global.DoNotUseThisFieldOmitHtml = value
}
func (el *HtmlElementFormButton) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Global.DoNotUseThisFieldOmitHtml == TRUE {
		return []byte{}
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	buffer.Write([]byte(`<button`))
	buffer.Write(el.Global.ToHtml())
	buffer.Write(data)
	buffer.Write([]byte(`>`))
	buffer.Write(el.Content.ToHtml())
	buffer.Write([]byte(`</button>`))

	return buffer.Bytes()
}
func (el *HtmlElementFormButton) ToJavaScript() []byte {
	var buffer bytes.Buffer

	//buffer.Write( el.Content.ToJavaScript() )

	return buffer.Bytes()
}
