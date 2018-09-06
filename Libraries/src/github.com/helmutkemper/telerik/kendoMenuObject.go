package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type MenuObject struct {
	Text           string               `jsObject:"text"`
	CssClass       string               `jsObject:"cssClass"`
	Attr           interface{}          `jsObject:"attr"`
	Url            string               `jsObject:"url"`
	AllowHtml      Boolean              `jsObject:"encoded"`
	Content        interface{}          `jsObject:"content" jsType:"JavaScript,string"`
	ContentAttr    HtmlGlobalAttributes `jsObject:"contentAttr"`
	ImageAttr      MenuImageAttr        `jsObject:"imageAttr"`
	ImageUrl       string               `jsObject:"imageUrl"`
	Items          []MenuObject         `jsObject:"items"`
	SpriteCssClass string               `jsObject:"spriteCssClass"`
	Select         JavaScript           `jsObject:"select"`

	*ToJavaScriptConverter
}

func (el *MenuObject) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("MenuObject.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}

type MenuImageAttr struct {
	/*
	  The name of the control, which is submitted with the form data.
	  @see typeNamesForAutocomplete.go
	  Ex.: const NAMES_FOR_AUTOCOMPLETE_NAME
	*/
	Name string `jsObject:"name"`

	/*
	  The initial value of the control. This attribute is optional except when the value of the type attribute is radio or
	  checkbox.
	  Note that when reloading the page, Gecko and IE will ignore the value specified in the HTML source, if the value was
	  changed before the reload.
	*/
	Value string `jsObject:"value"`

	/*
	  The form element that the input element is associated with (its form owner). The value of the attribute must be an id
	  of a <form> element in the same document. If this attribute is not specified, this <input> element must be a
	  descendant of a <form> element. This attribute enables you to place <input> elements anywhere within a document, not
	  just as descendants of their form elements. An input can only be associated with one form.
	*/
	Form string `jsObject:"form"`

	/*
	  The alt attribute provides alt text for the image, so screen reader users can get a better idea of what the button is
	  used for. It will also display if the image can't be shown for any reason (for example if the path is misspelled). If
	  possible, use text which matches the label you'd use if you were using a standard submit button.
	*/
	Alt string `jsObject:"alt"`

	/*
	  The src attribute is used to specify the path to the image you want to display in the button.
	*/
	Src string `jsObject:"src"`

	/*
	  The width and height attributes are used to specify the width and height the image should be shown at, in pixels. The
	  button is the same size as the image; if you need the button's hit area to be bigger than the image, you will need to
	  use CSS (e.g. padding). Also, if you specify only one dimension, the other is automatically adjusted so that the image
	  maintains its original aspect ratio.
	*/
	Width int `jsObject:"width"`

	/*
	  The width and height attributes are used to specify the width and height the image should be shown at, in pixels. The
	  button is the same size as the image; if you need the button's hit area to be bigger than the image, you will need to
	  use CSS (e.g. padding). Also, if you specify only one dimension, the other is automatically adjusted so that the image
	  maintains its original aspect ratio.
	*/
	Height int `jsObject:"height"`

	*ToJavaScriptConverter
}

func (el *MenuImageAttr) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("MenuImageAttr.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
