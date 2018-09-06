package telerik

type HtmlMouseEvent int

var htmlMouseEvents = [...]string{
	"",
	"abort",
	"beforeinput",
	"blur",
	"click",
	"compositionstart",
	"compositionupdate",
	"compositionend",
	"dblclick",
	"error",
	"focus",
	"focusin",
	"focusout",
	"input",
	"keydown",
	"keypress",
	"keyup",
	"load",
	"mousedown",
	"mouseenter",
	"mouseleave",
	"mousemove",
	"mouseout",
	"mouseover",
	"mouseup",
	"resize",
	"scroll",
	"select",
	"unload",
	"wheel",
}

func (el HtmlMouseEvent) String() string {
	return htmlMouseEvents[el]
}

const (
	MOUSE_EVENT_ABORT HtmlMouseEvent = iota + 1
	MOUSE_EVENT_BEFOREINPUT
	MOUSE_EVENT_BLUR
	MOUSE_EVENT_CLICK
	MOUSE_EVENT_COMPOSITIONSTART
	MOUSE_EVENT_COMPOSITIONUPDATE
	MOUSE_EVENT_COMPOSITIONEND
	MOUSE_EVENT_DBLCLICK
	MOUSE_EVENT_ERROR
	MOUSE_EVENT_FOCUS
	MOUSE_EVENT_FOCUSIN
	MOUSE_EVENT_FOCUSOUT
	MOUSE_EVENT_INPUT
	MOUSE_EVENT_KEYDOWN
	MOUSE_EVENT_KEYPRESS
	MOUSE_EVENT_KEYUP
	MOUSE_EVENT_LOAD
	MOUSE_EVENT_MOUSEDOWN
	MOUSE_EVENT_MOUSEENTER
	MOUSE_EVENT_MOUSELEAVE
	MOUSE_EVENT_MOUSEMOVE
	MOUSE_EVENT_MOUSEOUT
	MOUSE_EVENT_MOUSEOVER
	MOUSE_EVENT_MOUSEUP
	MOUSE_EVENT_RESIZE
	MOUSE_EVENT_SCROLL
	MOUSE_EVENT_SELECT
	MOUSE_EVENT_UNLOAD
	MOUSE_EVENT_WHEEL
)
