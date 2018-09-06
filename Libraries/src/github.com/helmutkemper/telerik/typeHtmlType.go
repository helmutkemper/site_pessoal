package telerik

type HtmlType int

var htmlTypes = [...]string{
	"",
	"button",
	"checkbox",
	"color",
	"date",
	"datetime-local",
	"email",
	"file",
	"hidden",
	"image",
	"month",
	"number",
	"password",
	"Note",
	"radio",
	"range",
	"reset",
	"search",
	"submit",
	"tel",
	"text",
	"time",
	"url",
	"week",
}

func (el HtmlType) String() string {
	return htmlTypes[el]
}

const (
	// A push button with no default behavior.
	INPUT_TYPE_BUTTON HtmlType = iota + 1

	// A check box allowing single values to be selected/deselected.
	INPUT_TYPE_CHECKBOX

	// HTML5 A control for specifying a color. A color picker's UI has no required features other than accepting simple
	// colors as text (more info).
	INPUT_TYPE_COLOR

	// HTML5 A control for entering a date (year, month, and day, with no time).
	INPUT_TYPE_DATE

	// HTML5 A control for entering a date and time, with no time zone.
	INPUT_TYPE_DATETIME_LOCAL

	// HTML5 A field for editing an e-mail address.
	INPUT_TYPE_EMAIL

	// A control that lets the user select a file. Use the accept attribute to define the types of files that the control
	// can select.
	INPUT_TYPE_FILE

	// A control that is not displayed but whose value is submitted to the server.
	INPUT_TYPE_HIDDEN

	// A graphical submit button. You must use the src attribute to define the source of the image and the alt attribute
	// to define alternative text. You can use the height and width attributes to define the size of the image in pixels.
	INPUT_TYPE_IMAGE

	// HTML5 A control for entering a month and year, with no time zone.
	INPUT_TYPE_MONTH

	// HTML5 A control for entering a number.
	INPUT_TYPE_NUMBER

	// A single-line text field whose value is obscured. Use the maxlength and minlength attributes to specify the maximum
	// length of the value that can be entered.
	//
	// Any forms involving sensitive information like passwords (e.g. login forms) should be served over HTTPS; Firefox
	// now implements multiple mechanisms to warn against insecure login forms — see Insecure passwords. Other browsers
	// are also implementing similar mechanisms.
	INPUT_TYPE_PASSWORD

	// A radio button, allowing a single value to be selected out of multiple choices.
	INPUT_TYPE_RADIO

	// HTML5 A control for entering a number whose exact value is not important.
	INPUT_TYPE_RANGE

	// A button that resets the contents of the form to default values.
	INPUT_TYPE_RESET

	// HTML5 A single-line text field for entering search strings. Line-breaks are automatically removed from the input
	// value.
	INPUT_TYPE_SEARCH

	// A button that submits the form.
	INPUT_TYPE_SUBMIT

	// HTML5 A control for entering a telephone number.
	INPUT_TYPE_TEL

	// A single-line text field. Line-breaks are automatically removed from the input value.
	INPUT_TYPE_TEXT

	// HTML5 A control for entering a time value with no time zone.
	INPUT_TYPE_TIME

	// HTML5 A field for entering a URL.
	INPUT_TYPE_URL

	// HTML5 A control for entering a date consisting of a week-year number and a week number with no time zone.
	INPUT_TYPE_WEEK
)
