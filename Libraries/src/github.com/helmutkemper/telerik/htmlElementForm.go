package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
)

// The HTML <form> element represents a document section that contains interactive controls to submit information to a
// web server.
//
// It is possible to use the :valid and :invalid CSS pseudo-classes to style a <form> element.
type HtmlElementForm struct {
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
	  This Boolean attribute indicates that the form control is not available for interaction. In particular, the click
	  event will not be dispatched on disabled controls. Also, a disabled control's value isn't submitted with the form.
	  Unlike other browsers, Firefox will by default persist the dynamic disabled state of an <input> across page loads. Use
	  the autocomplete attribute to control this feature.
	*/
	Disabled Boolean `htmlAttrSet:"disabled"`

	/*
	  If the value of the type attribute is file, then this attribute will indicate the types of files that the server
	  accepts, otherwise it will be ignored. The value must be a comma-separated list of unique content type specifiers:
	  A file extension starting with the STOP character (U+002E). (e.g. .jpg, .png, .doc).
	  A valid MIME type with no extensions.
	  audio/* representing sound files.
	  video/* representing video files.
	  image/* representing image files.
	*/
	Accept []string `htmlAttr:"accept"`

	/*
	  A space- or comma-delimited list of character encodings that the server accepts. The browser uses them in the order in
	  which they are listed. The default value, the reserved string "UNKNOWN", indicates the same encoding as that of the
	  document containing the form element.
	  In previous versions of HTML, the different character encodings could be delimited by spaces or commas. In HTML5, only
	  spaces are allowed as delimiters.
	*/
	AcceptCharset string `htmlAttr:"acceptcharset"`

	/*
	  The URI of a program that processes the form information. This value can be overridden by a formaction attribute on a
	  <button> or <input> element.
	*/
	Action string `htmlAttr:"action"`

	/*
	  This is a nonstandard attribute used by iOS Safari Mobile which controls whether and how the text value for textual
	  form control descendants should be automatically capitalized as it is entered/edited by the user. If the
	  autocapitalize attribute is specified on an individual form control descendant, it trumps the form-wide autocapitalize
	  setting. The non-deprecated values are available in iOS 5 and later. The default value is sentences. Possible values
	  are:
	  > none: Completely disables automatic capitalization
	  > sentences: Automatically capitalize the first letter of sentences.
	  > words: Automatically capitalize the first letter of words.
	  > characters: Automatically capitalize all characters.
	*/
	AutoCapitalize AutoCapitalize `htmlAttr:"autocapitalize"`

	/*
	  This attribute indicates whether the value of the control can be automatically completed by the browser.
	  Possible values are:
	  off: The user must explicitly enter a value into this field for every use, or the document provides its own
	  auto-completion method. The browser does not automatically complete the entry.
	  on: The browser is allowed to automatically complete the value based on values that the user has entered during
	  previous uses, however on does not provide any further information about what kind of data the user might be expected
	  to enter.
	  @see typeNamesForAutocomplete.go
	*/
	AutoComplete Boolean `htmlAttrOnOff:"autocomplete"`

	/*
	  When the value of the method attribute is post, enctype is the MIME type of content that is used to submit the form to
	  the server. Possible values are:
	  > application/x-www-form-urlencoded: The default value if the attribute is not specified.
	  > multipart/form-data: The value used for an <input> element with the type attribute set to "file".
	  > text/plain (HTML5)
	  This value can be overridden by a formenctype attribute on a <button> or <input> element.
	*/
	EncType string `htmlAttr:"enctype"`

	/*
	  The HTTP method that the browser uses to submit the form. Possible values are:
	  > post: Corresponds to the HTTP POST method ; form data are included in the body of the form and sent to the server.
	  > get: Corresponds to the HTTP GET method; form data are appended to the action attribute URI with a '?' as separator,
	  and the resulting URI is sent to the server. Use this method when the form has no side-effects and contains only ASCII
	  characters.
	  This value can be overridden by a formmethod attribute on a <button> or <input> element.
	*/
	Method string `htmlAttr:"method"`

	/*
	  This Boolean attribute indicates that the form is not to be validated when submitted. If this attribute is not
	  specified (and therefore the form is validated), this default setting can be overridden by a formnovalidate attribute
	  on a <button> or <input> element belonging to the form.
	*/
	NoValidate Boolean `htmlAttrSet:"novalidate"`

	/*
	  A name or keyword indicating where to display the response that is received after submitting the form. In HTML 4, this
	  is the name/keyword for a frame. In HTML5, it is a name/keyword for a browsing context (for example, tab, window, or
	  inline frame). The following keywords have special meanings:
	  > _self: Load the response into the same HTML 4 frame (or HTML5 browsing context) as the current one. This value is the
	  default if the attribute is not specified.
	  > _blank: Load the response into a new unnamed HTML 4 window or HTML5 browsing context.
	  > _parent: Load the response into the HTML 4 frameset parent of the current frame, or HTML5 parent browsing context of
	  the current one. If there is no parent, this option behaves the same way as _self.
	  > _top: HTML 4: Load the response into the full original window, and cancel all other frames. HTML5: Load the response
	  into the top-level browsing context (i.e., the browsing context that is an ancestor of the current one, and has no
	  parent). If there is no parent, this option behaves the same way as _self.
	  > iframename: The response is displayed in a named <iframe>.
	  HTML5: This value can be overridden by a formtarget attribute on a <button> or <input> element.
	*/
	Target string `htmlAttr:"target"`

	Global HtmlGlobalAttributes `htmlAttr:"-"`

	*ToJavaScriptConverter `htmlAttr:"-"`
}

func (el *HtmlElementForm) SetOmitHtml(value Boolean) {
	el.Global.DoNotUseThisFieldOmitHtml = value
}
func (el *HtmlElementForm) ToJavaScript() []byte {
	return []byte{}

	var ret bytes.Buffer
	if el.Global.Id == "" {
		el.Global.Id = GetAutoId()
	}

	el.Content.FilterFormElements()

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("HtmlElementForm.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Global.Id + `").kendoMultiSelect({`))
	ret.Write(data)
	ret.Write([]byte(`});`))

	return ret.Bytes()
}
func (el *HtmlElementForm) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Global.DoNotUseThisFieldOmitHtml == TRUE {
		return []byte{}
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	buffer.Write([]byte(`<form`))
	buffer.Write(el.Global.ToHtml())
	buffer.Write(data)
	buffer.Write([]byte(`>`))
	buffer.Write(el.Content.ToHtml())
	buffer.Write([]byte(`</form>`))

	return buffer.Bytes()
}
func (el *HtmlElementForm) ToHtmlSupport() []byte {
	return el.Content.ToHtmlSupport()
}
