package telerik

// <input> elements of type "reset"  are rendered as buttons, with a default click event handler that resets all of the
// inputs in the form to their initial values.
//
// <input type="reset" value="Reset the form">
type HtmlInputReset struct {
	/*
	  The name of the control, which is submitted with the form data.
	  @see typeNamesForAutocomplete.go
	  Ex.: const NAMES_FOR_AUTOCOMPLETE_NAME
	*/
	Name string

	/*
	  The initial value of the control. This attribute is optional except when the value of the type attribute is radio or
	  checkbox.
	  Note that when reloading the page, Gecko and IE will ignore the value specified in the HTML source, if the value was
	  changed before the reload.
	*/
	Value string

	/*
	  The form element that the input element is associated with (its form owner). The value of the attribute must be an id
	  of a <form> element in the same document. If this attribute is not specified, this <input> element must be a
	  descendant of a <form> element. This attribute enables you to place <input> elements anywhere within a document, not
	  just as descendants of their form elements. An input can only be associated with one form.
	*/
	Form string

	/*
	  This Boolean attribute indicates that the form control is not available for interaction. In particular, the click
	  event will not be dispatched on disabled controls. Also, a disabled control's value isn't submitted with the form.
	  Unlike other browsers, Firefox will by default persist the dynamic disabled state of an <input> across page loads. Use
	  the autocomplete attribute to control this feature.
	*/
	Disabled Boolean

	Global HtmlGlobalAttributes
} /*
func(el *HtmlInputReset)String() string {
  return `<input ` + el.Global.String() + ` type="reset" ` + el.Name.ToAttr("name") + el.Value.ToAttr("value") + el.Form.ToAttr("form") + el.Disabled.ToAttrSet("disabled") + `>`
}*/
