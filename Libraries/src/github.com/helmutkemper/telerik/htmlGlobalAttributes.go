package telerik

import (
	"bytes"
	"reflect"
)

// Global attributes are attributes common to all HTML elements; they can be used on all elements, though they may have
// no effect on some elements.
//
// Global attributes may be specified on all HTML elements, even those not specified in the standard. That means that
// any non-standard elements must still permit these attributes, even though using those elements means that the
// document is no longer HTML5-compliant. For example, HTML5-compliant browsers hide content marked as <foo hidden>...
// <foo>, even though <foo> is not a valid HTML element.
type HtmlGlobalAttributes struct {
	/*
	  This field was created for internal use and should not be accessed directly
	*/
	DoNotUseThisFieldOmitHtml Boolean `htmlAttr:"-"`
	/*
	  Provides a hint for generating a keyboard shortcut for the current element. This attribute consists of a
	  space-separated list of characters. The browser should use the first one that exists on the computer keyboard layout.
	*/
	AccessKey string `htmlAttr:"accesskey"`

	/*
	  Is a space-separated list of the classes of the element. Classes allows CSS and JavaScript to select and access
	  specific elements via the class selectors or functions like the method Document.getElementsByClassName().
	*/
	Class string `htmlAttr:"class"`

	/*
	  Is an enumerated attribute indicating if the element should be editable by the user. If so, the browser modifies its
	  widget to allow editing. The attribute must take one of the following values:
	  true or the empty string, which indicates that the element must be editable;
	  false, which indicates that the element must not be editable.
	*/
	ContentEditable Boolean `htmlAttr:"contenteditable"`

	/*
	  Is the id of an <menu> to use as the contextual menu for this element.
	*/
	ContextMenu string `htmlAttr:"contextmenu"`

	/*
	  Forms a class of attributes, called custom data attributes, that allow proprietary information to be exchanged between
	  the HTML and its DOM representation that may be used by scripts. All such custom data are available via the
	  HTMLElement interface of the element the attribute is set on. The HTMLElement.dataset property gives access to them.
	*/
	Data map[string]string `htmlAttr:"data"`

	/*
	  Is an enumerated attribute indicating the directionality of the element's text. It can have the following values:
	  ltr, which means left to right and is to be used for languages that are written from the left to the right (like English);
	  rtl, which means right to left and is to be used for languages that are written from the right to the left (like Arabic);
	  auto, which let the user agent decides. It uses a basic algorithm as it parses the characters inside the element until
	  it finds a character with a strong directionality, then apply that directionality to the whole element.
	*/
	Dir string `htmlAttr:"dir"`

	/*
	  Is an enumerated attribute indicating whether the element can be dragged, using the Drag and Drop API. It can have the
	  following values:
	  true, which indicates that the element may be dragged
	  false, which indicates that the element may not be dragged.
	*/
	Draggable Boolean `htmlAttr:"draggable"`

	/*
	  Is an enumerated attribute indicating what types of content can be dropped on an element, using the Drag and Drop API.
	  It can have the following values:
	  copy, which indicates that dropping will create a copy of the element that was dragged
	  move, which indicates that the element that was dragged will be moved to this new location.
	  link, will create a link to the dragged data.
	*/
	DropZone TypeHtmlDropZone `htmlAttr:"dropzone"`

	/*
	  Is a Boolean attribute indicates that the element is not yet, or is no longer, relevant. For example, it can be used
	  to hide elements of the page that can't be used until the login process has been completed. The browser won't render
	  such elements. This attribute must not be used to hide content that could legitimately be shown.
	*/
	Hidden Boolean `htmlAttrSet:"draggable"`

	/*
	  Defines a unique identifier (ID) which must be unique in the whole document. Its purpose is to identify the element
	  when linking (using a fragment identifier), scripting, or styling (with CSS).
	*/
	Id string `htmlAttr:"id"`

	/*
	  The unique, global identifier of an item.
	*/
	ItemId string `htmlAttr:"itemid"`

	/*
	  Used to add properties to an item. Every HTML element may have an itemprop attribute specified, where an itemprop
	  consists of a name and value pair.
	*/
	ItemDrop string `htmlAttr:"itemdrop"`

	/*
	  Properties that are not descendants of an element with the itemscope attribute can be associated with the item using
	  an itemref. It provides a list of element ids (not itemids) with additional properties elsewhere in the document.
	*/
	ItemRef string `htmlAttr:"itemref"`

	/*
	  itemscope (usually) works along with itemtype to specify that the HTML contained in a block is about a particular
	  item. itemscope creates the Item and defines the scope of the itemtype associated with it. itemtype is a valid URL of
	  a vocabulary (such as schema.org) that describes the item and its properties context.
	*/
	ItemScope string `htmlAttr:"itemscope"`

	/*
	  Specifies the URL of the vocabulary that will be used to define itemprops (item properties) in the data structure.
	  itemscope is used to set the scope of where in the data structure the vocabulary set by itemtype will be active.
	*/
	ItemType string `htmlAttr:"itemtype"`

	/*
	  Participates in defining the language of the element, the language that non-editable elements are written in or the
	  language that editable elements should be written in. The tag contains one single entry value in the format defined in
	  the Tags for Identifying Languages (BCP47) IETF document. xml:lang has priority over it.
	*/
	Lang string `htmlAttr:"lang"`

	/*
	  Assigns a slot in a shadow DOM shadow tree to an element: An element with a slot attribute is assigned to the slot
	  created by the <slot> element whose name attribute's value matches that slot attribute's value.
	*/
	Sort string `htmlAttr:"sort"`

	/*
	  Is an enumerated attribute defines whether the element may be checked for spelling errors. It may have the following
	  values:
	  true, which indicates that the element should be, if possible, checked for spelling errors;
	  false, which indicates that the element should not be checked for spelling errors.
	*/
	SpellCheck string `htmlAttr:"spellcheck"`

	/*
	  Contains CSS styling declarations to be applied to the element. Note that it is recommended for styles to be defined
	  in a separate file or files. This attribute and the <style> element have mainly the purpose of allowing for quick
	  styling, for example for testing purposes.
	*/
	Style string `htmlAttr:"style"`

	/*
	  Is an integer attribute indicating if the element can take input focus (is focusable), if it should participate to
	  sequential keyboard navigation, and if so, at what position. It can takes several values:
	  a negative value means that the element should be focusable, but should not be reachable via sequential keyboard
	  navigation;
	  0 means that the element should be focusable and reachable via sequential keyboard navigation, but its relative order
	  is defined by the platform convention;
	  a positive value means that the element should be focusable and reachable via sequential keyboard navigation; the
	  order in which the elements are focused is the increasing value of the tabindex. If several elements share the same
	  tabindex, their relative order follows their relative positions in the document.
	*/
	TabIndex int `htmlAttr:"tabindex"`

	/*
	  Contains a text representing advisory information related to the element it belongs to. Such information can
	  typically, but not necessarily, be presented to the user as a tooltip.
	*/
	Title string `htmlAttr:"title"`

	/*
	  Is an enumerated attribute that is used to specify whether an element's attribute values and the values of its Text
	  node children are to be translated when the page is localized, or whether to leave them unchanged. It can have the
	  following values:
	  empty string and "yes", which indicates that the element will be translated.
	  "no", which indicates that the element will not be translated.
	*/
	Translate Boolean `htmlAttr:"translate"`

	Extra map[string]interface{} `htmlAttr:"extra"`

	OnAbort             string `htmlAttr:"onabort"`
	OnAutoComplete      string `htmlAttr:"onautocomplete"`
	OnAutoCompleteError string `htmlAttr:"onautocompleteerror"`
	OnBlur              string `htmlAttr:"onblur"`
	OnCancel            string `htmlAttr:"oncancel"`
	OnCanPlay           string `htmlAttr:"oncanplay"`
	OnCanPlayThrough    string `htmlAttr:"oncanplaythrough"`
	OnChange            string `htmlAttr:"onchange"`
	OnClick             string `htmlAttr:"onclick"`
	OnClose             string `htmlAttr:"onclose"`
	OnContextMenu       string `htmlAttr:"oncontextmenu"`
	OnCueChange         string `htmlAttr:"oncuechange"`
	OnDblClick          string `htmlAttr:"ondblclick"`
	OnDrag              string `htmlAttr:"ondrag"`
	OnDragEnd           string `htmlAttr:"ondragend"`
	OnDragEnter         string `htmlAttr:"ondragenter"`
	OnDragExit          string `htmlAttr:"ondragexit"`
	OnDragLeave         string `htmlAttr:"ondragleave"`
	OnDragOver          string `htmlAttr:"ondragover"`
	OnDragStart         string `htmlAttr:"ondragstart"`
	OnDrop              string `htmlAttr:"ondrop"`
	OnDurationChange    string `htmlAttr:"ondurationchange"`
	OnEmptied           string `htmlAttr:"onemptied"`
	OnEnded             string `htmlAttr:"onended"`
	OnError             string `htmlAttr:"onerror"`
	OnFocus             string `htmlAttr:"onfocus"`
	OnInput             string `htmlAttr:"oninput"`
	OnInvalid           string `htmlAttr:"oninvalid"`
	OnKeyDown           string `htmlAttr:"onkeydown"`
	OnKeyPress          string `htmlAttr:"onkeypress"`
	OnKeyUp             string `htmlAttr:"onkeyup"`
	OnLoad              string `htmlAttr:"onload"`
	OnLoadedData        string `htmlAttr:"onloadeddata"`
	OnLoadedMetadata    string `htmlAttr:"onloadedmetadata"`
	OnLoadStart         string `htmlAttr:"onloadstart"`
	OnMouseDown         string `htmlAttr:"onmousedown"`
	OnMouseEnter        string `htmlAttr:"onmouseenter"`
	OnMouseLeave        string `htmlAttr:"onmouseleave"`
	OnMouseMove         string `htmlAttr:"onmousemove"`
	OnMouseOut          string `htmlAttr:"onmouseout"`
	OnMouseOver         string `htmlAttr:"onmouseover"`
	OnMouseUp           string `htmlAttr:"onmouseup"`
	OnMouseWheel        string `htmlAttr:"onmousewheel"`
	OnPause             string `htmlAttr:"onpause"`
	OnPlay              string `htmlAttr:"onplay"`
	OnPlaying           string `htmlAttr:"onplaying"`
	OnProgress          string `htmlAttr:"onprogress"`
	OnRateChange        string `htmlAttr:"onratechange"`
	OnReset             string `htmlAttr:"onreset"`
	OnResize            string `htmlAttr:"onresize"`
	OnScroll            string `htmlAttr:"onscroll"`
	OnSeeked            string `htmlAttr:"onseeked"`
	OnSeeking           string `htmlAttr:"onseeking"`
	OnSelect            string `htmlAttr:"onselect"`
	OnShow              string `htmlAttr:"onshow"`
	OnSort              string `htmlAttr:"onsort"`
	OnStalled           string `htmlAttr:"onstalled"`
	OnSubmit            string `htmlAttr:"onsubmit"`
	OnSuspend           string `htmlAttr:"onsuspend"`
	OnTimeUpdate        string `htmlAttr:"ontimeupdate"`
	OnToggle            string `htmlAttr:"ontoggle"`
	OnVolumeChange      string `htmlAttr:"onvolumechange"`
	OnWaiting           string `htmlAttr:"onwaiting"`

	*ToJavaScriptConverter `htmlAttr:"-"`
}

func (el *HtmlGlobalAttributes) ToHtml() []byte {
	var buffer bytes.Buffer

	if el.Id != "-" {
		if el.Id == "" {
			el.Id = GetAutoId()
		}
	} else {
		el.Id = ""
	}

	element := reflect.ValueOf(el).Elem()
	data := el.ToJavaScriptConverter.ToTelerikHtml(element)
	buffer.Write(data)

	return buffer.Bytes()
}
func (el *HtmlGlobalAttributes) GetId() []byte {
	var buffer bytes.Buffer

	if el.Id == "" {
		el.Id = GetAutoId()
	}

	buffer.WriteString(el.Id)

	return buffer.Bytes()
}
