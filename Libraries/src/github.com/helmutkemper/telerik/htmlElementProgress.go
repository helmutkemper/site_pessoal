package telerik

// Introduced in HTML5
//
// The HTML <progress> element displays an indicator showing the completion progress of a task, typically displayed as
// a progress bar.
type HtmlElementFormProgress struct {
	/*
	  The initial value of the control. This attribute is optional except when the value of the type attribute is radio or
	  checkbox.
	  Note that when reloading the page, Gecko and IE will ignore the value specified in the HTML source, if the value was
	  changed before the reload.
	*/
	Value string

	/*
	  The maximum (numeric or date-time) value for this item, which must not be less than its minimum (min attribute) value.
	*/
	Max Int

	/*
	  Content inside html tag
	*/
	Content string

	Global HtmlGlobalAttributes
} /*
func(el *HtmlElementFormProgress)String() string {
  return `<progress ` + el.Global.String() + el.Value.ToAttr("value") + el.Max.ToAttr("max") + `>` + el.Content.String() + `</progress>`
}*/
