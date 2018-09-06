package telerik

// The HTML Output element (<output>) is a container element into which a site or app can inject the results of a
// calculation or the outcome of a user action.
type HtmlElementFormOutput struct {
	/*
	  The name of the control, which is submitted with the form data.
	*/
	Name string

	/*
	  Content inside html tag
	*/
	Content Content

	/*
	  The form element that the input element is associated with (its form owner). The value of the attribute must be an id
	  of a <form> element in the same document. If this attribute is not specified, this <input> element must be a
	  descendant of a <form> element. This attribute enables you to place <input> elements anywhere within a document, not
	  just as descendants of their form elements. An input can only be associated with one form.
	*/
	Form string

	/*
	  The id of a labelable form-related element in the same document as the label element. The first such element in the
	  document with an ID matching the value of the for attribute is the labeled control for this label element.
	  A label element can have both a for attribute and a contained control element, as long as the for attribute points to
	  the contained control element.
	*/
	For string

	Global HtmlGlobalAttributes
} /*
func(el *HtmlElementFormOutput)String() string {
  return `<output ` + el.Global.String() + el.Name.ToAttr("name") + el.Form.ToAttr("form") + el.For.ToAttr("for") + `>` + el.Content.String() + `</output>`
}*/
