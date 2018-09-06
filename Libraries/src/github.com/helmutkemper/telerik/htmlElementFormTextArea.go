package telerik

import (
	"bytes"
	"reflect"
)

// The HTML <textarea> element represents a multi-line plain-text editing control.
type HtmlElementFormTextArea struct {
	/*
	  The name of the control, which is submitted with the form data.
	  @see typeNamesForAutocomplete.go
	  Ex.: const NAMES_FOR_AUTOCOMPLETE_NAME
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
	AutoComplete Boolean `htmlAttrOnOff:"AutoComplete"`

	/*
	  The visible width of the text control, in average character widths. If it is specified, it must be a positive integer.
	  If it is not specified, the default value is 20 (HTML5).
	*/
	Cols int `htmlAttr:"cols"`

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
	MinLength int `htmlAttr:"maxlength"`

	/*
	  A hint to the user of what can be entered in the control . The placeholder text must not contain carriage returns or
	  line-feeds.
	*/
	PlaceHolder string `htmlAttr:"placeholder"`

	/*
	  This attribute specifies that the user must fill in a value before submitting a form. It cannot be used when the type
	  attribute is hidden, image, or a button type (submit, reset, or button). The :optional and :required CSS
	  pseudo-classes will be applied to the field as appropriate.
	*/
	Required Boolean `htmlAttrSet:"Required"`

	/*
	  The number of visible text lines for the control.
	*/
	Rows int `htmlAttr:"rows"`

	/*
	  Setting the value of this attribute to true indicates that the element needs to have its spelling and grammar checked.
	  The value default indicates that the element is to act according to a default behavior, possibly based on the parent
	  element's own spellcheck value. The value false indicates that the element should not be checked.
	*/
	SpellCheck Boolean `htmlAttr:"spellcheck"`

	/*
	  Indicates how the control wraps text. Possible values are:
	  > hard: The browser automatically inserts line breaks (CR+LF) so that each line has no more than the width of the
	  control; the cols attribute must be specified.
	  > soft: The browser ensures that all line breaks in the value consist of a CR+LF pair, but does not insert any
	  additional line breaks.
	  > off:  Like soft but changes appearance to white-space: pre so line segments exceeding cols are not wrapped and area
	  becomes horizontally scrollable.
	  If this attribute is not specified, soft is its default value.
	*/
	Wrap Warp `htmlAttr:"wrap"`

	Global HtmlGlobalAttributes `htmlAttr:"-"`

	*ToJavaScriptConverter `htmlAttr:"-"`
}

func (el *HtmlElementFormTextArea) SetOmitHtml(value Boolean) {
	el.Global.DoNotUseThisFieldOmitHtml = value
}
func (el *HtmlElementFormTextArea) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Global.DoNotUseThisFieldOmitHtml == TRUE {
		return []byte{}
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)

	buffer.Write([]byte(`<textarea`))
	buffer.Write(el.Global.ToHtml())
	buffer.Write(data)
	buffer.Write([]byte(`>`))
	buffer.Write(el.Content.ToHtml())
	buffer.Write([]byte(`</textarea>`))

	return buffer.Bytes()
}
func (el *HtmlElementFormTextArea) GetId() []byte {
	if el.Global.Id == "" {
		el.Global.Id = GetAutoId()
	}
	return []byte(el.Global.Id)
}
func (el *HtmlElementFormTextArea) GetName() []byte {
	if el.Name == "" {
		el.Name = GetAutoId()
	}
	return []byte(el.Name)
}
