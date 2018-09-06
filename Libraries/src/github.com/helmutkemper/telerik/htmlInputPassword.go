package telerik

import (
	"bytes"
	"reflect"
)

// <input> elements of type "password" provide a way for the user to securely enter a password. The element is presented
// as a one-line plain text editor control in which the text is obscured so that it cannot be read, usually by replacing
// each character with a symbol such as the asterisk ("*") or a dot ("•"). This character will vary depending on the
// user agent and OS.
//
// Specifics of how the entry process works may vary from browser to browser; mobile devices, for example, often display
// the typed character for a moment before obscuring it, to allow the user to be sure they pressed the key they meant to
// press; this is helpful given the small size of keys and the ease with which the wrong one can be pressed, especially
// on virtual keyboards.
//
// Any forms involving sensitive information like passwords (e.g. login forms) should be served over HTTPS; Many
// browsers now implement mechanisms to warn against insecure login forms; see Insecure passwords.
//
// <input id="userPassword" type="password">
type HtmlInputPassword struct {
	/*
	  The name of the control, which is submitted with the form data.
	  @see typeNamesForAutocomplete.go
	  Ex.: const NAMES_FOR_AUTOCOMPLETE_NAME
	*/
	Name string `htmlAttr:"name"`

	/*
	  The initial value of the control. This attribute is optional except when the value of the type attribute is radio or
	  checkbox.
	  Note that when reloading the page, Gecko and IE will ignore the value specified in the HTML source, if the value was
	  changed before the reload.
	*/
	Value string `htmlAttr:"value"`

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
	  If the value of the type attribute is text, email, search, password, tel, or url, this attribute specifies the maximum
	  number of characters (in UTF-16 code units) that the user can enter. For other control types, it is ignored. It can
	  exceed the value of the size attribute. If it is not specified, the user can enter an unlimited number of characters.
	  Specifying a negative number results in the default behavior (i.e. the user can enter an unlimited number of
	  characters). The constraint is evaluated only when the value of the attribute has been changed.
	*/
	MaxLength int `htmlAttr:"maxlength"`

	/*
	  If the value of the type attribute is text, email, search, password, tel, or url, this attribute specifies the minimum
	  number of characters (in Unicode code points) that the user can enter. For other control types, it is ignored.
	*/
	MinLength int `htmlAttr:"minlength"`

	/*
	  A regular expression that the control's value is checked against. The pattern must match the entire value, not just
	  some subset. Use the title attribute to describe the pattern to help the user. This attribute applies when the value
	  of the type attribute is text, search, tel, url, email, or password, otherwise it is ignored. The regular expression
	  language is the same as JavaScript RegExp algorithm, with the 'u' parameter that makes it treat the pattern as a
	  sequence of unicode code points. The pattern is not surrounded by forward slashes.
	*/
	Pattern string `htmlAttr:"pattern"`

	/*
	  A hint to the user of what can be entered in the control . The placeholder text must not contain carriage returns or
	  line-feeds.
	*/
	PlaceHolder string `htmlAttr:"placeholder"`

	/*
	  This attribute indicates that the user cannot modify the value of the control. The value of the attribute is
	  irrelevant. If you need read-write access to the input value, do not add the "readonly" attribute. It is ignored if
	  the value of the type attribute is hidden, range, color, checkbox, radio, file, or a button type (such as button or
	  submit).
	*/
	Readonly Boolean `htmlAttrSet:"readonly"`

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

	Global HtmlGlobalAttributes `htmlAttr:"-"`

	*ToJavaScriptConverter `htmlAttr:"-"`
}

func (el *HtmlInputPassword) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Global.DoNotUseThisFieldOmitHtml == TRUE {
		return []byte{}
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	buffer.Write([]byte(`<input type="password"`))
	buffer.Write(el.Global.ToHtml())
	buffer.Write(data)
	buffer.Write([]byte(`>`))

	return buffer.Bytes()
}
func (el *HtmlInputPassword) GetId() []byte {
	if el.Global.Id == "" {
		el.Global.Id = GetAutoId()
	}
	return []byte(el.Global.Id)
}
func (el *HtmlInputPassword) GetName() []byte {
	if el.Name == "" {
		el.Name = GetAutoId()
	}
	return []byte(el.Name)
}
func (el *HtmlInputPassword) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Global.Id == "" {
		el.Global.Id = GetAutoId()
	}

	ret.Write([]byte(`$("#` + el.Global.Id + `").addClass('k-textbox');`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
