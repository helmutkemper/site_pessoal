package telerik

import (
	//"fmt"
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

type Content []interface{}

func (el Content) ToHtmlSupport() []byte {
	var buffer bytes.Buffer

	keys := make([]int, 0)
	for k := range el {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		switch outConverted := el[k].(type) {
		case KendoUiMultiSelect:
			buffer.Write(outConverted.ToHtml())

			if !reflect.DeepEqual(outConverted.Dialog, KendoUiDialog{}) {
				buffer.Write(outConverted.Dialog.ToHtml())
			}
		}
	}

	return buffer.Bytes()
}

func (el Content) ToHtml() []byte {
	var buffer bytes.Buffer

	keys := make([]int, 0)
	for k := range el {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		switch outConverted := el[k].(type) {
		case string:
			buffer.WriteString(outConverted)
		case *AceEditor:
			buffer.Write(outConverted.ToHtml())
		case *KendoUiWindow:
			buffer.Write(outConverted.ToHtml())
		case *KendoUiGrid:
			buffer.Write(outConverted.ToHtml())
		case *HtmlInputSubmit:
			buffer.Write(outConverted.ToHtml())
		case *HtmlElementDiv:
			buffer.Write(outConverted.ToHtml())
		case *HtmlInputHidden:
			buffer.Write(outConverted.ToHtml())
		case *HtmlInputGeneric:
			buffer.Write(outConverted.ToHtml())
		case *HtmlInputNumber:
			buffer.Write(outConverted.ToHtml())
		case *HtmlElementSpan:
			buffer.Write(outConverted.ToHtml())
		case *HtmlElementFormLabel:
			buffer.Write(outConverted.ToHtml())
		case *HtmlInputText:
			buffer.Write(outConverted.ToHtml())
		case *HtmlElementFormLegend:
			buffer.Write(outConverted.ToHtml())
		case *HtmlElementFormFieldSet:
			buffer.Write(outConverted.ToHtml())
		case *HtmlElementForm:
			buffer.Write(outConverted.ToHtml())
		case *HtmlElementFormSelect:
			buffer.Write(outConverted.ToHtml())
		case *HtmlElementFormTextArea:
			buffer.Write(outConverted.ToHtml())
		case *HtmlElementFormMeter:
			buffer.Write(outConverted.ToHtml())
		case *HtmlElementFormButton:
			buffer.Write(outConverted.ToHtml())
		case *HtmlElementFormDataList:
			buffer.Write(outConverted.ToHtml())
		case *KendoUiAutoComplete:
			buffer.Write(outConverted.ToHtml())
		case *KendoUiMobileSwitch:
			buffer.Write(outConverted.ToHtml())
		case *KendoUiNumericTextBox:
			buffer.Write(outConverted.ToHtml())
		case *KendoUiComboBox:
			buffer.Write(outConverted.ToHtml())
		case *KendoUiDialog:
			buffer.Write(outConverted.ToHtml())
		case *HtmlElementScript:
			buffer.Write(outConverted.ToHtml())
		case *KendoUiDropDownList:
			buffer.Write(outConverted.ToHtml())
		case *KendoUiButton:
			buffer.Write(outConverted.ToHtml())
		case *KendoUiMultiSelect:
			buffer.Write(outConverted.ToHtml())

			if !reflect.DeepEqual(outConverted.Dialog, KendoUiDialog{}) {
				buffer.Write(outConverted.Dialog.ToHtml())
			}

		case *HtmlScriptType:
			buffer.WriteString(outConverted.String())
		case *HtmlElementLi:
			buffer.Write(outConverted.ToHtml())
		case *HtmlElementUl:
			buffer.Write(outConverted.ToHtml())
		case *KendoDataSource:
			//do noting

		case *HtmlInputCheckBox:
			buffer.Write(outConverted.ToHtml())

		/*case HtmlElementFormOutput:
		    buffer.Write( outConverted.ToHtml() )
		  case HtmlElementFormProgress:
		    buffer.Write( outConverted.ToHtml() )*/

		default:
			fmt.Printf("\n\n\nToHtml() - typeContent.go: error: type %T not found in typeContent\n\n\n", outConverted)
		}
	}

	return buffer.Bytes()
}
func (el *Content) ToJavaScript() []byte {
	var buffer bytes.Buffer
	//buffer.WriteString( "\n" )

	keys := make([]int, 0)
	for k := range *el {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		switch outConverted := (*el)[k].(type) {
		case string:
			buffer.WriteString(outConverted)
			//buffer.WriteString( "\n" )
		case *KendoUiGrid:
			buffer.Write(outConverted.ToJavaScript())
		case *KendoUiWindow:
			buffer.Write(outConverted.ToJavaScript())
		case *AceEditor:
			buffer.Write(outConverted.ToJavaScript())
		case *KendoUiDialog:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiDateTimePicker:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiDatePicker:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiDateInput:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiContextMenu:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiConfirm:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiComboBox:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiColorPicker:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiColorPalette:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiCalendar:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiButton:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiMultiSelect:
			buffer.Write(outConverted.ToJavaScript())

			if !reflect.DeepEqual(outConverted.Dialog, KendoUiDialog{}) {
				buffer.Write(outConverted.Dialog.ToJavaScript())
			}

			//buffer.WriteString( "\n" )
		case *KendoUiAutoComplete:
			buffer.Write(outConverted.ToJavaScript())
		case *KendoUiMobileSwitch:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiDraggable:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiDropDownList:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiDropTarget:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiDropTargetArea:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *HtmlElementDiv:
			buffer.Write(outConverted.ToJavaScript())
		//  //buffer.WriteString( "\n" )
		case *HtmlElementFormLabel:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoUiNumericTextBox:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *HtmlElementFormSelect:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *HtmlElementScript:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *HtmlElementFormButton:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *HtmlElementSpan:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		case *KendoDataSource:
			buffer.Write(outConverted.ToJavaScript())
			//buffer.WriteString( "\n" )
		//case KendoSchema:
		//buffer.Write( outConverted.ToJavaScript() )
		//buffer.WriteString( "\n" )

		case *HtmlElementFormTextArea:
		case *HtmlInputCheckBox:
		case *HtmlInputColor:
		case *HtmlInputDate:
		case *HtmlInputDateTimeLocal:
		case *HtmlInputEmail:
			buffer.Write(outConverted.ToJavaScript())
		case *HtmlInputSubmit:
			buffer.Write(outConverted.ToJavaScript())
		case *HtmlInputFile:
		case *HtmlInputGeneric:
			buffer.Write(outConverted.ToJavaScript())
		case *HtmlInputHidden:
		case *HtmlInputImage:
		case *HtmlInputMonth:
		case *HtmlInputNumber:
		case *HtmlInputPassword:
			buffer.Write(outConverted.ToJavaScript())
		case *HtmlInputRadio:
		case *HtmlInputRange:
		case *HtmlInputSearch:
			buffer.Write(outConverted.ToJavaScript())
		case *HtmlInputTel:
			buffer.Write(outConverted.ToJavaScript())
		case *HtmlInputText:
			buffer.Write(outConverted.ToJavaScript())
		case *HtmlInputTime:
		case *HtmlInputUrl:
			buffer.Write(outConverted.ToJavaScript())
		case *HtmlElementForm:
			buffer.Write(outConverted.ToJavaScript())
		case *HtmlInputWeek:

		default:
			fmt.Printf("\n\n\nToJavaScript() - typeContent.go: error: type %T not found in typeContent\n\n\n", outConverted)
		}
	}

	var formElements = el.FilterFormElements()
	for _, v := range formElements {
		switch v.(type) {

		case *KendoUiWindow:

			if reflect.DeepEqual((*v.(*KendoUiWindow)), KendoUiWindow{}) {
				continue
			}

			buffer.Write([]byte("        "))
			buffer.Write([]byte((*v.(*KendoUiWindow)).Html.Global.GetId()))
			buffer.Write([]byte("Widget = $('#"))
			buffer.Write([]byte((*v.(*KendoUiWindow)).Html.Global.GetId()))
			buffer.Write([]byte("').data('kendoWindow');\n"))

			//$("#dialog").data("kendoWindow")

		case *KendoUiGrid:

			if reflect.DeepEqual((*v.(*KendoUiGrid)), KendoUiGrid{}) {
				continue
			}

			buffer.Write([]byte("        "))
			buffer.Write([]byte((*v.(*KendoUiGrid)).Html.Global.GetId()))
			buffer.Write([]byte("Widget = $('#"))
			buffer.Write([]byte((*v.(*KendoUiGrid)).Html.Global.GetId()))
			buffer.Write([]byte("').data('kendoWindow');\n"))

			//$("#dialog").data("kendoWindow")

		case **KendoUiMultiSelect:

			if !reflect.DeepEqual((*(*v.(**KendoUiMultiSelect))).DataSource, KendoDataSource{}) {

				switch (*(*v.(**KendoUiMultiSelect))).DataSource.(type) {
				case KendoDataSource:
					// Widget
					buffer.Write([]byte("        "))
					buffer.Write([]byte((*(*v.(**KendoUiMultiSelect))).Html.GetId()))
					buffer.Write([]byte("Widget = $('#"))
					buffer.Write([]byte((*(*v.(**KendoUiMultiSelect))).Html.GetId()))
					buffer.Write([]byte("').getKendoMultiSelect();\n"))

					// DataSource
					buffer.Write([]byte("        "))
					buffer.Write([]byte((*(*v.(**KendoUiMultiSelect))).Html.GetId()))
					buffer.Write([]byte("DataSource = "))
					buffer.Write([]byte((*(*v.(**KendoUiMultiSelect))).Html.GetId()))
					buffer.Write([]byte("Widget.dataSource;\n"))
				}
			}

		default:
			continue
		}
	}

	return buffer.Bytes()
}
func (el *Content) FilterFormElements() []interface{} {

	var contentProcessedList = make([]interface{}, 0)
	var contentUnprocessedList = make([]interface{}, 0)
	var contentFoundList = make([]interface{}, 0)

	el.processContent(&contentProcessedList, &contentUnprocessedList, &contentFoundList, *el)
	for {
		if len(contentUnprocessedList) == 0 {
			break
		}

		popElement := contentUnprocessedList[0]
		contentUnprocessedList = contentUnprocessedList[1:]

		el.processContent(&contentProcessedList, &contentUnprocessedList, &contentFoundList, popElement.(Content))
	}

	/*for _, v := range contentFoundList{
	  fmt.Printf( "%T\n", v )
	}*/

	return contentFoundList
}
func (el *Content) processContent(contentProcessedList, contentUnprocessedList, contentFoundList *[]interface{}, content Content) {
	*contentProcessedList = append(*contentProcessedList, content)

	for _, contentElement := range content {
		el.addToUnprocessedList(contentUnprocessedList, contentFoundList, contentElement)
	}
}
func (el *Content) addToUnprocessedList(contentUnprocessedList, contentFoundList *[]interface{}, content interface{}) {
	switch converted := content.(type) {
	case string:
	case *Content:
	case *KendoUiGrid:
		*contentFoundList = append(*contentFoundList, converted)
	case *KendoUiWindow:
		*contentFoundList = append(*contentFoundList, converted)

		if converted.Content == nil {
			return
		}

		*contentUnprocessedList = append(*contentUnprocessedList, converted.Content)
	case *HtmlElementSpan:
		if converted.Content == nil {
			return
		}

		*contentUnprocessedList = append(*contentUnprocessedList, converted.Content)
	case *HtmlElementDiv:
		if converted.Content == nil {
			return
		}

		*contentUnprocessedList = append(*contentUnprocessedList, converted.Content)
	case *HtmlElementForm:
		if converted.Content == nil {
			return
		}

		*contentUnprocessedList = append(*contentUnprocessedList, converted.Content)

		// Elementos de formulário que necessitam de javascript - início

	case *HtmlElementScript:
		*contentFoundList = append(*contentFoundList, converted)
	case *KendoUiNumericTextBox:
		*contentFoundList = append(*contentFoundList, converted)
	case **KendoUiComboBox:
		*contentFoundList = append(*contentFoundList, &converted)
	case *KendoUiComboBox:
		*contentFoundList = append(*contentFoundList, &converted)
	case *KendoUiAutoComplete:
		*contentFoundList = append(*contentFoundList, &converted)
	case *KendoUiButton:
		*contentFoundList = append(*contentFoundList, &converted)
	case *KendoUiCalendar:
		*contentFoundList = append(*contentFoundList, &converted)
	case *KendoUiColorPalette:
		*contentFoundList = append(*contentFoundList, &converted)
	case *KendoUiColorPicker:
		*contentFoundList = append(*contentFoundList, &converted)
	case *KendoUiDateInput:
		*contentFoundList = append(*contentFoundList, &converted)
	case *KendoUiDatePicker:
		*contentFoundList = append(*contentFoundList, converted)
	case *KendoUiDateTimePicker:
		*contentFoundList = append(*contentFoundList, converted)
	case KendoUiDateTimePicker:
		*contentFoundList = append(*contentFoundList, &converted)
	case *KendoUiDropDownList:
		*contentFoundList = append(*contentFoundList, &converted)
	case *KendoUiMultiSelect:
		*contentFoundList = append(*contentFoundList, &converted)

		if reflect.DeepEqual(converted.Dialog, KendoUiDialog{}) == false {

			switch convertedInternal := converted.Dialog.Content.(type) {
			case Content:
				for _, subContent := range convertedInternal.FilterFormElements() {
					el.addToUnprocessedList(contentUnprocessedList, contentFoundList, subContent)
					//*contentFoundList        =append( *contentFoundList, &subContent )
				}
			}
		}

	case *KendoUiMobileSwitch:
		*contentFoundList = append(*contentFoundList, converted)
	case *HtmlElementFormSelect:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlElementFormTextArea:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputCheckBox:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputColor:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputDate:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputDateTimeLocal:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputEmail:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputFile:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputGeneric:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputHidden:
		*contentFoundList = append(*contentFoundList, &converted)
	case **HtmlInputHidden: //fixme: isto está certo?
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputImage:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputMonth:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputNumber:
		*contentFoundList = append(*contentFoundList, converted)
	case *HtmlInputPassword:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputRadio:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputRange:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputSearch:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputTel:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputText:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputTime:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputUrl:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlElementFormButton:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputWeek:
		*contentFoundList = append(*contentFoundList, &converted)
	case *HtmlInputSubmit:
		*contentFoundList = append(*contentFoundList, &converted)
	case *KendoDataSource:
		*contentFoundList = append(*contentFoundList, &converted)
	case *AceEditor:
		*contentFoundList = append(*contentFoundList, &converted)
	case *KendoUiDialog:
		*contentFoundList = append(*contentFoundList, &converted)

		// Elementos de formulário que necessitam de javascript - fim

	case *HtmlElementFormLabel:
	default:
		fmt.Printf("HtmlElementForm.addToUnprocessedList().type: %T\n", converted)
	}
}
func (el *Content) GetNamesAndIds() []map[string][]byte {
	var pass = false
	var name, id []byte
	var ret = make([]map[string][]byte, 0)
	var formElements = el.FilterFormElements()
	for _, v := range formElements {
		pass = false
		switch converted := v.(type) {
		case **HtmlInputText:
			pass = true
			name = []byte((*(*v.(**HtmlInputText))).Name)
			id = (*(*v.(**HtmlInputText))).GetId()
		case **HtmlInputHidden:
			pass = true
			name = []byte((*(*v.(**HtmlInputHidden))).Name)
			id = (*(*v.(**HtmlInputHidden))).GetId()
		case **KendoUiComboBox:
			pass = true
			name = []byte((*(*v.(**KendoUiComboBox))).Html.Name)
			id = (*(*v.(**KendoUiComboBox))).GetId()
		case **KendoUiAutoComplete:
			pass = true
			name = []byte((*(*v.(**KendoUiAutoComplete))).Html.Name)
			id = (*(*v.(**KendoUiAutoComplete))).GetId()
		case **KendoUiButton:
			//pass = true
			//name = []byte( (*(*v.(**KendoUiButton))).Html.Name )
			//id = (*(*v.(**KendoUiButton))).GetId()
		case **KendoUiCalendar:
			pass = true
			name = []byte((*(*v.(**KendoUiCalendar))).Html.Name)
			id = (*(*v.(**KendoUiCalendar))).GetId()
		case **KendoUiColorPalette:
			pass = true
			name = []byte((*(*v.(**KendoUiColorPalette))).Html.Name)
			id = (*(*v.(**KendoUiColorPalette))).GetId()
		case **KendoUiColorPicker:
			pass = true
			name = []byte((*(*v.(**KendoUiColorPicker))).Html.Name)
			id = (*(*v.(**KendoUiColorPicker))).GetId()
		case **KendoUiNumericTextBox:
			pass = true
			name = []byte((*(*v.(**KendoUiNumericTextBox))).Html.Name)
			id = (*(*v.(**KendoUiNumericTextBox))).GetId()
		case **KendoUiDateInput:
			pass = true
			name = []byte((*(*v.(**KendoUiDateInput))).Html.Name)
			id = (*(*v.(**KendoUiDateInput))).GetId()
		case **KendoUiDatePicker:
			pass = true
			name = []byte((*(*v.(**KendoUiDatePicker))).Html.Name)
			id = (*(*v.(**KendoUiDatePicker))).GetId()
		case **KendoUiDateTimePicker:
			pass = true
			name = []byte((*(*v.(**KendoUiDateTimePicker))).Html.Name)
			id = (*(*v.(**KendoUiDateTimePicker))).GetId()
		case **KendoUiDropDownList:
			pass = true
			name = []byte((*(*v.(**KendoUiDropDownList))).Html.Name)
			id = (*(*v.(**KendoUiDropDownList))).GetId()
		case **KendoUiMultiSelect:
			pass = true
			name = []byte((*(*v.(**KendoUiMultiSelect))).Html.Name)
			id = (*(*v.(**KendoUiMultiSelect))).GetId()
		case **HtmlElementFormSelect:
			pass = true
			name = []byte((*(*v.(**HtmlElementFormSelect))).Name)
			id = (*(*v.(**HtmlElementFormSelect))).GetId()
		case **HtmlElementFormTextArea:
			pass = true
			name = []byte((*(*v.(**HtmlElementFormTextArea))).Name)
			id = (*(*v.(**HtmlElementFormTextArea))).GetId()
		case **HtmlInputCheckBox:
			pass = true
			name = []byte((*(*v.(**HtmlInputCheckBox))).Name)
			id = (*(*v.(**HtmlInputCheckBox))).GetId()
		case **HtmlInputColor:
			pass = true
			name = []byte((*(*v.(**HtmlInputColor))).Name)
			id = (*(*v.(**HtmlInputColor))).GetId()
		case **HtmlInputDate:
			pass = true
			name = []byte((*(*v.(**HtmlInputDate))).Name)
			id = (*(*v.(**HtmlInputDate))).GetId()
		case **HtmlInputDateTimeLocal:
			pass = true
			name = []byte((*(*v.(**HtmlInputDateTimeLocal))).Name)
			id = (*(*v.(**HtmlInputDateTimeLocal))).GetId()
		case **HtmlInputEmail:
			pass = true
			name = []byte((*(*v.(**HtmlInputEmail))).Name)
			id = (*(*v.(**HtmlInputEmail))).GetId()
		case **HtmlInputFile:
			pass = true
			name = []byte((*(*v.(**HtmlInputFile))).Name)
			id = (*(*v.(**HtmlInputFile))).GetId()
		case **HtmlInputGeneric:
			pass = true
			name = []byte((*(*v.(**HtmlInputGeneric))).Name)
			id = (*(*v.(**HtmlInputGeneric))).GetId()
		case **HtmlInputImage:
			pass = true
			name = []byte((*(*v.(**HtmlInputImage))).Name)
			id = (*(*v.(**HtmlInputImage))).GetId()
		case **HtmlInputMonth:
			pass = true
			name = []byte((*(*v.(**HtmlInputMonth))).Name)
			id = (*(*v.(**HtmlInputMonth))).GetId()
		case **HtmlInputNumber:
			pass = true
			name = []byte((*(*v.(**HtmlInputNumber))).Name)
			id = (*(*v.(**HtmlInputNumber))).GetId()
		case **HtmlInputPassword:
			pass = true
			name = []byte((*(*v.(**HtmlInputPassword))).Name)
			id = (*(*v.(**HtmlInputPassword))).GetId()
		case **HtmlInputRadio:
			pass = true
			name = []byte((*(*v.(**HtmlInputRadio))).Name)
			id = (*(*v.(**HtmlInputRadio))).GetId()
		case **HtmlInputRange:
			pass = true
			name = []byte((*(*v.(**HtmlInputRange))).Name)
			id = (*(*v.(**HtmlInputRange))).GetId()
		case **HtmlInputSearch:
			pass = true
			name = []byte((*(*v.(**HtmlInputSearch))).Name)
			id = (*(*v.(**HtmlInputSearch))).GetId()
		case **HtmlInputTel:
			pass = true
			name = []byte((*(*v.(**HtmlInputTel))).Name)
			id = (*(*v.(**HtmlInputTel))).GetId()
		case **HtmlInputTime:
			pass = true
			name = []byte((*(*v.(**HtmlInputTime))).Name)
			id = (*(*v.(**HtmlInputTime))).GetId()
		case **HtmlInputUrl:
			pass = true
			name = []byte((*(*v.(**HtmlInputUrl))).Name)
			id = (*(*v.(**HtmlInputUrl))).GetId()
		case **HtmlInputWeek:
			pass = true
			name = []byte((*(*v.(**HtmlInputWeek))).Name)
			id = (*(*v.(**HtmlInputWeek))).GetId()
		case **HtmlInputSubmit:
			pass = true
			name = []byte((*(*v.(**HtmlInputSubmit))).Name)
			id = (*(*v.(**HtmlInputSubmit))).GetId()
		case **KendoUiMobileSwitch:
			pass = true
			name = []byte((*(*v.(**KendoUiMobileSwitch))).Html.Name)
			id = (*(*v.(**KendoUiMobileSwitch))).GetId()
		case *HtmlInputText:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputHidden:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *KendoUiComboBox:
			pass = true
			name = []byte(converted.Html.Name)
			id = converted.GetId()
		case *KendoUiAutoComplete:
			pass = true
			name = []byte(converted.Html.Name)
			id = converted.GetId()
		case *KendoUiButton:
			//pass = true
			//name = []byte( converted.Html.Name )
			//id = converted.GetId()
		case *KendoUiCalendar:
			pass = true
			name = []byte(converted.Html.Name)
			id = converted.GetId()
		case *KendoUiColorPalette:
			pass = true
			name = []byte(converted.Html.Name)
			id = converted.GetId()
		case *KendoUiColorPicker:
			pass = true
			name = []byte(converted.Html.Name)
			id = converted.GetId()
		case *KendoUiNumericTextBox:
			pass = true
			name = []byte(converted.Html.Name)
			id = converted.GetId()
		case *KendoUiDateInput:
			pass = true
			name = []byte(converted.Html.Name)
			id = converted.GetId()
		case *KendoUiDatePicker:
			pass = true
			name = []byte(converted.Html.Name)
			id = converted.GetId()
		case *KendoUiDateTimePicker:
			pass = true
			name = []byte(converted.Html.Name)
			id = converted.GetId()
		case *KendoUiDropDownList:
			pass = true
			name = []byte(converted.Html.Name)
			id = converted.GetId()
		case *KendoUiMultiSelect:
			pass = true
			name = []byte(converted.Html.Name)
			id = converted.GetId()
		case *HtmlElementFormSelect:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlElementFormTextArea:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputCheckBox:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputColor:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputDate:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputDateTimeLocal:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputEmail:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputFile:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputGeneric:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputImage:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputMonth:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputNumber:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputPassword:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputRadio:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputRange:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputSearch:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputTel:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputTime:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputUrl:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputWeek:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *HtmlInputSubmit:
			pass = true
			name = []byte(converted.Name)
			id = converted.GetId()
		case *KendoUiMobileSwitch:
			pass = true
			name = []byte(converted.Html.Name)
			id = converted.GetId()

		case *Content:
		case *KendoDataSource:
		default:
			fmt.Printf("\n\n\nGetNamesAndIds.id.type.%T.not.found.\n\n\n", converted)
		}

		if pass == true {
			var element = make(map[string][]byte)
			element["name"] = name
			element["id"] = id
			ret = append(ret, element)
		}
	}

	return ret
}
func (el *Content) MakeJsObject() []byte {
	var pass = false
	var buffer bytes.Buffer
	var key, jsCode []byte
	var formElements = el.FilterFormElements()
	// fixme: mfalta um getName() ou algo parecido
	// fixme: KendoUiCalendar e KendoUiColorPalette devem ter name como obrigatórios

	for _, v := range formElements {
		switch v.(type) {

		case *KendoUiGrid:
			if reflect.DeepEqual(*v.(*KendoUiGrid), KendoUiGrid{}) {
				continue
			}

			buffer.Write([]byte("      var "))
			buffer.Write([]byte((*v.(*KendoUiGrid)).Html.Global.GetId()))
			buffer.Write([]byte("Widget;\n"))

		case *KendoUiWindow:
			if reflect.DeepEqual(*v.(*KendoUiWindow), KendoUiWindow{}) {
				continue
			}

			buffer.Write([]byte("      var "))
			buffer.Write([]byte((*v.(*KendoUiWindow)).Html.Global.GetId()))
			buffer.Write([]byte("Widget;\n"))

		case **KendoUiMultiSelect:

			if !reflect.DeepEqual((*(*v.(**KendoUiMultiSelect))).DataSource, KendoDataSource{}) {

				switch (*(*v.(**KendoUiMultiSelect))).DataSource.(type) {
				case KendoDataSource:
					// Widget
					buffer.Write([]byte("      var "))
					buffer.Write([]byte((*(*v.(**KendoUiMultiSelect))).Html.GetId()))
					buffer.Write([]byte("Widget;\n"))

					// DataSource
					buffer.Write([]byte("      var "))
					buffer.Write([]byte((*(*v.(**KendoUiMultiSelect))).Html.GetId()))
					buffer.Write([]byte("DataSource;\n"))
				}
			}

		default:
			continue
		}
	}

	for _, v := range formElements {
		switch converted := v.(type) {

		case **KendoUiMultiSelect:

			var pass = false
			var contentToFind []map[string][]byte
			if !reflect.DeepEqual((*(*v.(**KendoUiMultiSelect))).Dialog, KendoUiDialog{}) {
				var mainElementDataSourceDataKeyId = []byte((*(*v.(**KendoUiMultiSelect))).DataValueField)
				for _, action := range (*(*v.(**KendoUiMultiSelect))).Dialog.Actions {

					if action.ButtonType == BUTTON_TYPE_ADD_AND_CLOSE || action.ButtonType == BUTTON_TYPE_ADD {

						buffer.Write([]byte("      function "))
						buffer.Write([]byte((*(*v.(**KendoUiMultiSelect))).GetId()))

						if action.ButtonType == BUTTON_TYPE_ADD_AND_CLOSE {

							buffer.Write([]byte("AddAndCloseButton() {\n"))

						} else if action.ButtonType == BUTTON_TYPE_ADD {

							buffer.Write([]byte("AddButton() {\n"))

						}

						var elementId []byte
						var dataSourceName []byte
						switch convertedContent := (*(*v.(**KendoUiMultiSelect))).Dialog.Content.(type) {

						case Content:

							for _, element := range convertedContent {

								pass = false
								switch convertedElement := element.(type) {

								case *HtmlElementForm:
									pass = true
									elementId = []byte(convertedElement.Global.Id)
									dataSourceName = (*converted).Html.GetId()
									contentToFind = convertedElement.Content.GetNamesAndIds()

								case *HtmlElementDiv:
									elementId = []byte(convertedElement.Global.Id)
									dataSourceName = (*converted).Html.GetId()
									contentToFind = convertedElement.Content.GetNamesAndIds()
									pass = true

								case *HtmlElementSpan:
									elementId = []byte(convertedElement.Global.Id)
									dataSourceName = (*converted).Html.GetId()
									contentToFind = convertedElement.Content.GetNamesAndIds()
									pass = true

								}

								if pass == true {
									buffer.Write([]byte("        if( $('#"))
									buffer.Write(elementId)
									buffer.Write([]byte("').kendoValidator().data('kendoValidator').validate() == false ){\n"))
									buffer.Write([]byte("          "))
									buffer.Write([]byte("return false;\n"))
									buffer.Write([]byte("        "))
									buffer.Write([]byte("}\n"))

									buffer.Write([]byte("        "))
									buffer.Write(dataSourceName)
									buffer.Write([]byte("DataSource.add({\n"))

									for _, element := range contentToFind {
										buffer.Write([]byte("          '"))
										buffer.Write(element["name"])
										buffer.Write([]byte("': getValueById('id:"))
										buffer.Write(element["id"])
										buffer.Write([]byte("'),\n"))
									}

									buffer.Write([]byte("        });\n"))

									buffer.Write([]byte("        "))
									buffer.Write(dataSourceName)
									buffer.Write([]byte("DataSource.one('requestEnd', function (args) {\n"))

									buffer.Write([]byte("          "))
									buffer.Write([]byte("if (args.type !== 'create') {\n"))
									buffer.Write([]byte("            "))
									buffer.Write([]byte("return;\n"))
									buffer.Write([]byte("          "))
									buffer.Write([]byte("}\n"))
									buffer.Write([]byte("          "))
									buffer.Write([]byte("var newValue = args.response.Objects[ 0 ]."))
									buffer.Write(mainElementDataSourceDataKeyId)
									buffer.Write([]byte("\n"))

									buffer.Write([]byte("          "))
									buffer.Write(dataSourceName)
									buffer.Write([]byte("DataSource.one('sync', function () {\n"))

									buffer.Write([]byte("            "))
									buffer.Write(dataSourceName)
									buffer.Write([]byte("Widget.value("))
									buffer.Write(dataSourceName)
									buffer.Write([]byte("Widget.value().concat([newValue]));\n"))

									buffer.Write([]byte("          "))
									buffer.Write([]byte("});\n"))

									buffer.Write([]byte("        "))
									buffer.Write([]byte("});\n"))

									buffer.Write([]byte("        "))
									buffer.Write(dataSourceName)
									buffer.Write([]byte("DataSource.sync();\n"))

									if action.ButtonType == BUTTON_TYPE_ADD_AND_CLOSE {

										buffer.Write([]byte("        "))
										buffer.Write(dataSourceName)
										buffer.Write([]byte("Widget.close();\n"))

									} else if action.ButtonType == BUTTON_TYPE_ADD {

										buffer.Write([]byte("        "))
										buffer.Write([]byte("return false;\n"))

									}
								}
							}

						}

						buffer.Write([]byte("      }\n"))

					}

				}

			}

		default:
			continue
		}
	}

	// fixme: isto está constante
	buffer.Write([]byte("      function addNewItemToKendoDataSource( id ){\n"))
	buffer.Write([]byte("        switch( id ){\n"))
	for _, v := range formElements {
		switch converted := v.(type) {
		case *KendoUiMultiSelect:
			if reflect.DeepEqual(converted.Dialog, KendoUiDialog{}) {
				continue
			}

			key = []byte(converted.Html.GetId())
			jsCode = []byte(`$('#` + string(converted.Dialog.GetId()) + `').data('kendoDialog').open()`)
			buffer.Write([]byte("          case 'id:"))
			buffer.Write(key)
			buffer.Write([]byte("': "))
			buffer.Write(jsCode)
			buffer.Write([]byte(";\n"))

			switch convertedFromInterface := converted.Dialog.Content.(type) {
			case Content:
				buffer.Write([]byte("            "))
				buffer.Write(convertedFromInterface.ToJavaScript())
				buffer.Write([]byte("       break;\n"))
			}

			//buffer.Write( []byte( ";\n" ) )

		case **KendoUiMultiSelect:
			if reflect.DeepEqual((*(*v.(**KendoUiMultiSelect))).Dialog, KendoUiDialog{}) {
				continue
			}

			key = []byte((*(*v.(**KendoUiMultiSelect))).Html.GetId())
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiMultiSelect))).Dialog.GetId()) + `').data('kendoDialog').open()`)
			buffer.Write([]byte("          case 'id:"))
			buffer.Write(key)
			buffer.Write([]byte("': "))
			buffer.Write(jsCode)
			buffer.Write([]byte(";\n"))

			switch convertedFromInterface := (*(*v.(**KendoUiMultiSelect))).Dialog.Content.(type) {
			case Content:
				buffer.Write([]byte("            "))
				buffer.Write(convertedFromInterface.ToJavaScript())
				buffer.Write([]byte("       break;\n"))
			}

			//buffer.Write( []byte( ";\n" ) )

		default:
			continue
		}
	}
	buffer.Write([]byte("        }\n"))
	buffer.Write([]byte("      }\n"))

	buffer.Write([]byte("      function getValueById( id ){\n"))
	buffer.Write([]byte("        switch( id ){\n"))
	for _, v := range formElements {
		pass = false
		switch converted := v.(type) {
		case ***AceEditor:
			pass = true
			key = []byte((*(*(*v.(***AceEditor)))).Html.Name)
			jsCode = []byte(`ace.edit("` + string((*(*(*v.(***AceEditor)))).GetId()) + `").getValue()`)
		case ***HtmlElementScript:
			pass = true
			key = []byte((*(*(*v.(***HtmlElementScript)))).GetId())
			jsCode = []byte(`kendo.template($("#` + string((*(*(*v.(***HtmlElementScript)))).GetId()) + `").html())`)
		case ***KendoUiComboBox:
			pass = true
			key = []byte((*(*(*v.(***KendoUiComboBox)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiComboBox)))).GetId()) + `').data('kendoComboBox').value()`)
		case ***HtmlInputText:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputText)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputText)))).GetId()) + `').val()`)
		case ***HtmlInputHidden:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputHidden)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputHidden)))).GetId()) + `').val()`)
		case ***KendoUiDropDownList:
			pass = true
			key = []byte((*(*(*v.(***KendoUiDropDownList)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiDropDownList)))).GetId()) + `').data('kendoDropDownList').value()`)
		case ***KendoUiMultiSelect: // fixme: colocar o conteúdo da janela aqui também
			pass = true
			key = []byte((*(*(*v.(***KendoUiMultiSelect)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiMultiSelect)))).GetId()) + `').data('kendoMultiSelect').value()`)
		case ***KendoUiAutoComplete:
			pass = true
			key = []byte((*(*(*v.(***KendoUiAutoComplete)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiAutoComplete)))).GetId()) + `').data('kendoAutoComplete').value()`)
		case ***KendoUiButton:
			//pass = true
			//key = []byte( (*(*(*v.(***KendoUiButton)))).Html.Name )
			//jsCode = []byte( `$('#` + string( (*(*(*v.(***KendoUiButton)))).GetId() ) + `').data('kendoButton').value()` )
		case ***KendoUiCalendar:
			pass = true
			key = []byte((*(*(*v.(***KendoUiCalendar)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiCalendar)))).GetId()) + `').data('kendoCalendar').value()`)
		case ***KendoUiColorPalette:
			pass = true
			key = []byte((*(*(*v.(***KendoUiColorPalette)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiColorPalette)))).GetId()) + `').data('kendoColorPalette').value()`)
		case ***KendoUiColorPicker:
			pass = true
			key = []byte((*(*(*v.(***KendoUiColorPicker)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiColorPicker)))).GetId()) + `').data('kendoColorPicker').value()`)
		case ***KendoUiNumericTextBox:
			pass = true
			key = []byte((*(*(*v.(***KendoUiNumericTextBox)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiNumericTextBox)))).GetId()) + `').data('kendoNumericTextBox').value()`)
		case ***KendoUiDateInput:
			pass = true
			key = []byte((*(*(*v.(***KendoUiDateInput)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiDateInput)))).GetId()) + `').data('kendoDateInput').value()`)
		case ***KendoUiDatePicker:
			pass = true
			key = []byte((*(*(*v.(***KendoUiDatePicker)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiDatePicker)))).GetId()) + `').data('kendoDatePicker').value()`)
		case ***KendoUiDateTimePicker:
			pass = true
			key = []byte((*(*(*v.(***KendoUiDateTimePicker)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiDateTimePicker)))).GetId()) + `').data('kendoDateTimePicker').value()`)
		case ***HtmlElementFormSelect:
			pass = true
			key = []byte((*(*(*v.(***HtmlElementFormSelect)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlElementFormSelect)))).GetId()) + `').val()`)
		case ***HtmlElementFormTextArea:
			pass = true
			key = []byte((*(*(*v.(***HtmlElementFormTextArea)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlElementFormTextArea)))).GetId()) + `').val()`)
		case ***HtmlInputCheckBox:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputCheckBox)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputCheckBox)))).GetId()) + `').val()`)
		case ***HtmlInputColor:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputColor)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputColor)))).GetId()) + `').val()`)
		case ***HtmlInputDate:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputDate)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputDate)))).GetId()) + `').val()`)
		case ***HtmlInputDateTimeLocal:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputDateTimeLocal)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputDateTimeLocal)))).GetId()) + `').val()`)
		case ***HtmlInputEmail:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputEmail)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputEmail)))).GetId()) + `').val()`)
		case ***HtmlInputFile:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputFile)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputFile)))).GetId()) + `').val()`)
		case ***HtmlInputGeneric:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputGeneric)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputGeneric)))).GetId()) + `').val()`)
		case ***HtmlInputImage:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputImage)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputImage)))).GetId()) + `').val()`)
		case ***HtmlInputMonth:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputMonth)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputMonth)))).GetId()) + `').val()`)
		case ***HtmlInputNumber:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputNumber)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputNumber)))).GetId()) + `').val()`)
		case ***HtmlInputPassword:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputPassword)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputPassword)))).GetId()) + `').val()`)
		case ***HtmlInputRadio:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputRadio)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputRadio)))).GetId()) + `').val()`)
		case ***HtmlInputRange:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputRange)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputRange)))).GetId()) + `').val()`)
		case ***HtmlInputSearch:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputSearch)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputSearch)))).GetId()) + `').val()`)
		case ***HtmlInputTel:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputTel)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputTel)))).GetId()) + `').val()`)
		case ***HtmlInputTime:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputTime)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputTime)))).GetId()) + `').val()`)
		case ***HtmlInputUrl:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputUrl)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputUrl)))).GetId()) + `').val()`)
		case ***HtmlInputWeek:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputWeek)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputWeek)))).GetId()) + `').val()`)
		case ***HtmlInputSubmit:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputSubmit)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputSubmit)))).GetId()) + `').val()`)
		case ***KendoUiMobileSwitch:
			pass = true
			key = []byte((*(*(*v.(***KendoUiMobileSwitch)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiMobileSwitch)))).GetId()) + `').val()`)

		case **AceEditor:
			pass = true
			key = []byte((*(*v.(**AceEditor))).Html.Name)
			jsCode = []byte(`ace.edit("` + string((*(*v.(**AceEditor))).GetId()) + `").getValue()`)
		case **HtmlElementScript:
			pass = true
			key = []byte((*(*v.(**HtmlElementScript))).GetId())
			jsCode = []byte(`kendo.template($("#` + string((*(*v.(**HtmlElementScript))).GetId()) + `").html())`)
		case **KendoUiComboBox:
			pass = true
			key = []byte((*(*v.(**KendoUiComboBox))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiComboBox))).GetId()) + `').data('kendoComboBox').value()`)
		case **HtmlInputText:
			pass = true
			key = []byte((*(*v.(**HtmlInputText))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputText))).GetId()) + `').val()`)
		case **HtmlInputHidden:
			pass = true
			key = []byte((*(*v.(**HtmlInputHidden))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputHidden))).GetId()) + `').val()`)
		case **KendoUiDropDownList:
			pass = true
			key = []byte((*(*v.(**KendoUiDropDownList))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiDropDownList))).GetId()) + `').data('kendoDropDownList').value()`)
		case **KendoUiMultiSelect: // fixme: colocar o conteúdo da janela aqui também
			pass = true
			key = []byte((*(*v.(**KendoUiMultiSelect))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiMultiSelect))).GetId()) + `').data('kendoMultiSelect').value()`)
		case **KendoUiAutoComplete:
			pass = true
			key = []byte((*(*v.(**KendoUiAutoComplete))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiAutoComplete))).GetId()) + `').data('kendoAutoComplete').value()`)
		case **KendoUiButton:
			//pass = true
			//key = []byte( (*(*v.(**KendoUiButton))).Html.Name )
			//jsCode = []byte( `$('#` + string( (*(*v.(**KendoUiButton))).GetId() ) + `').data('kendoButton').value()` )
		case **KendoUiCalendar:
			pass = true
			key = []byte((*(*v.(**KendoUiCalendar))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiCalendar))).GetId()) + `').data('kendoCalendar').value()`)
		case **KendoUiColorPalette:
			pass = true
			key = []byte((*(*v.(**KendoUiColorPalette))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiColorPalette))).GetId()) + `').data('kendoColorPalette').value()`)
		case **KendoUiColorPicker:
			pass = true
			key = []byte((*(*v.(**KendoUiColorPicker))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiColorPicker))).GetId()) + `').data('kendoColorPicker').value()`)
		case **KendoUiNumericTextBox:
			pass = true
			key = []byte((*(*v.(**KendoUiNumericTextBox))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiNumericTextBox))).GetId()) + `').data('kendoNumericTextBox').value()`)
		case **KendoUiDateInput:
			pass = true
			key = []byte((*(*v.(**KendoUiDateInput))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiDateInput))).GetId()) + `').data('kendoDateInput').value()`)
		case **KendoUiDatePicker:
			pass = true
			key = []byte((*(*v.(**KendoUiDatePicker))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiDatePicker))).GetId()) + `').data('kendoDatePicker').value()`)
		case **KendoUiDateTimePicker:
			pass = true
			key = []byte((*(*v.(**KendoUiDateTimePicker))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiDateTimePicker))).GetId()) + `').data('kendoDateTimePicker').value()`)
		case **HtmlElementFormSelect:
			pass = true
			key = []byte((*(*v.(**HtmlElementFormSelect))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlElementFormSelect))).GetId()) + `').val()`)
		case **HtmlElementFormTextArea:
			pass = true
			key = []byte((*(*v.(**HtmlElementFormTextArea))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlElementFormTextArea))).GetId()) + `').val()`)
		case **HtmlInputCheckBox:
			pass = true
			key = []byte((*(*v.(**HtmlInputCheckBox))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputCheckBox))).GetId()) + `').val()`)
		case **HtmlInputColor:
			pass = true
			key = []byte((*(*v.(**HtmlInputColor))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputColor))).GetId()) + `').val()`)
		case **HtmlInputDate:
			pass = true
			key = []byte((*(*v.(**HtmlInputDate))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputDate))).GetId()) + `').val()`)
		case **HtmlInputDateTimeLocal:
			pass = true
			key = []byte((*(*v.(**HtmlInputDateTimeLocal))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputDateTimeLocal))).GetId()) + `').val()`)
		case **HtmlInputEmail:
			pass = true
			key = []byte((*(*v.(**HtmlInputEmail))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputEmail))).GetId()) + `').val()`)
		case **HtmlInputFile:
			pass = true
			key = []byte((*(*v.(**HtmlInputFile))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputFile))).GetId()) + `').val()`)
		case **HtmlInputGeneric:
			pass = true
			key = []byte((*(*v.(***HtmlInputGeneric))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputGeneric))).GetId()) + `').val()`)
		case **HtmlInputImage:
			pass = true
			key = []byte((*(*v.(**HtmlInputImage))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputImage))).GetId()) + `').val()`)
		case **HtmlInputMonth:
			pass = true
			key = []byte((*(*v.(**HtmlInputMonth))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputMonth))).GetId()) + `').val()`)
		case **HtmlInputNumber:
			pass = true
			key = []byte((*(*v.(**HtmlInputNumber))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputNumber))).GetId()) + `').val()`)
		case **HtmlInputPassword:
			pass = true
			key = []byte((*(*v.(**HtmlInputPassword))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputPassword))).GetId()) + `').val()`)
		case **HtmlInputRadio:
			pass = true
			key = []byte((*(*v.(**HtmlInputRadio))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputRadio))).GetId()) + `').val()`)
		case **HtmlInputRange:
			pass = true
			key = []byte((*(*v.(**HtmlInputRange))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputRange))).GetId()) + `').val()`)
		case **HtmlInputSearch:
			pass = true
			key = []byte((*(*v.(**HtmlInputSearch))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputSearch))).GetId()) + `').val()`)
		case **HtmlInputTel:
			pass = true
			key = []byte((*(*v.(**HtmlInputTel))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputTel))).GetId()) + `').val()`)
		case **HtmlInputTime:
			pass = true
			key = []byte((*(*v.(**HtmlInputTime))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputTime))).GetId()) + `').val()`)
		case **HtmlInputUrl:
			pass = true
			key = []byte((*(*v.(**HtmlInputUrl))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputUrl))).GetId()) + `').val()`)
		case **HtmlInputWeek:
			pass = true
			key = []byte((*(*v.(**HtmlInputWeek))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputWeek))).GetId()) + `').val()`)
		case **HtmlInputSubmit:
			pass = true
			key = []byte((*(*v.(**HtmlInputSubmit))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputSubmit))).GetId()) + `').val()`)
		case **KendoUiMobileSwitch:
			pass = true
			key = []byte((*(*v.(**KendoUiMobileSwitch))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiMobileSwitch))).GetId()) + `').val()`)

		case *AceEditor:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`ace.edit("` + string(converted.GetId()) + `").getValue()`)
		case *HtmlElementScript:
			pass = true
			key = []byte(converted.GetId())
			jsCode = []byte(`kendo.template($("#` + string(converted.GetId()) + `").html())`)
		case *KendoUiComboBox:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoComboBox').value()`)
		case *HtmlInputText:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputHidden:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *KendoUiDropDownList:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoDropDownList').value()`)
		case *KendoUiMultiSelect: // fixme: colocar o conteúdo da janela aqui também
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoMultiSelect').value()`)
		case *KendoUiAutoComplete:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoAutoComplete').value()`)
		case *KendoUiButton:
			//pass = true
			//key = []byte( converted.Html.Name )
			//jsCode = []byte( `$('#` + string( converted.GetId() ) + `').data('kendoButton').value()` )
		case *KendoUiCalendar:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoCalendar').value()`)
		case *KendoUiColorPalette:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoColorPalette').value()`)
		case *KendoUiColorPicker:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoColorPicker').value()`)
		case *KendoUiNumericTextBox:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoNumericTextBox').value()`)
		case *KendoUiDateInput:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoDateInput').value()`)
		case *KendoUiDatePicker:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoDatePicker').value()`)
		case *KendoUiDateTimePicker:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoDateTimePicker').value()`)
		case *HtmlElementFormSelect:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlElementFormTextArea:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputCheckBox:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputColor:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputDate:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputDateTimeLocal:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputEmail:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputFile:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputGeneric:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputImage:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputMonth:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputNumber:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputPassword:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputRadio:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputRange:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputSearch:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputTel:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputTime:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputUrl:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputWeek:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *HtmlInputSubmit:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)
		case *KendoUiMobileSwitch:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val()`)

		case *Content:
		case *KendoDataSource:
		default:
			fmt.Printf("\n\n\n-getValueById.id.type.%T.not.found.\n\n\n", converted)
		}

		if pass == true {
			buffer.Write([]byte("          case 'id:"))
			buffer.Write(key)
			buffer.Write([]byte("': return "))
			buffer.Write(jsCode)
			buffer.Write([]byte(";\n"))
		}
	}
	buffer.Write([]byte("        }\n"))
	buffer.Write([]byte("      }\n"))

	buffer.Write([]byte("      function setValueById( id, value ){\n"))
	buffer.Write([]byte("        switch( id ){\n"))
	for _, v := range formElements {
		pass = false
		switch converted := v.(type) {
		case ***AceEditor:
			pass = true
			key = []byte((*(*(*v.(***AceEditor)))).Html.Name)
			jsCode = []byte(`ace.edit("` + string((*(*(*v.(***AceEditor)))).GetId()) + `").setValue( value )`)
		case ***KendoUiComboBox:
			pass = true
			key = []byte((*(*(*v.(***KendoUiComboBox)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiComboBox)))).GetId()) + `').data('kendoComboBox').value( value )`)
		case ***HtmlInputText:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputText)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputText)))).GetId()) + `').val( value )`)
		case ***HtmlInputHidden:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputHidden)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputHidden)))).GetId()) + `').val( value )`)
		case ***KendoUiDropDownList:
			pass = true
			key = []byte((*(*(*v.(***KendoUiDropDownList)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiDropDownList)))).GetId()) + `').data('kendoDropDownList').value( value )`)
		case ***KendoUiMultiSelect: // fixme: colocar o conteúdo da janela aqui também
			pass = true
			key = []byte((*(*(*v.(***KendoUiMultiSelect)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiMultiSelect)))).GetId()) + `').data('kendoMultiSelect').value( value )`)
		case ***KendoUiAutoComplete:
			pass = true
			key = []byte((*(*(*v.(***KendoUiAutoComplete)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiAutoComplete)))).GetId()) + `').data('kendoAutoComplete').value( value )`)
		case ***KendoUiButton:
			//pass = true
			//key = []byte( (*(*(*v.(***KendoUiButton)))).Html.Name )
			//jsCode = []byte( `$('#` + string( (*(*(*v.(***KendoUiButton)))).GetId() ) + `').data('kendoButton').value( value )` )
		case ***KendoUiCalendar:
			pass = true
			key = []byte((*(*(*v.(***KendoUiCalendar)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiCalendar)))).GetId()) + `').data('kendoCalendar').value( value )`)
		case ***KendoUiColorPalette:
			pass = true
			key = []byte((*(*(*v.(***KendoUiColorPalette)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiColorPalette)))).GetId()) + `').data('kendoColorPalette').value( value )`)
		case ***KendoUiColorPicker:
			pass = true
			key = []byte((*(*(*v.(***KendoUiColorPicker)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiColorPicker)))).GetId()) + `').data('kendoColorPicker').value( value )`)
		case ***KendoUiNumericTextBox:
			pass = true
			key = []byte((*(*(*v.(***KendoUiNumericTextBox)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiNumericTextBox)))).GetId()) + `').data('kendoNumericTextBox').value( value )`)
		case ***KendoUiDateInput:
			pass = true
			key = []byte((*(*(*v.(***KendoUiDateInput)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiDateInput)))).GetId()) + `').data('kendoDateInput').value( value )`)
		case ***KendoUiDatePicker:
			pass = true
			key = []byte((*(*(*v.(***KendoUiDatePicker)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiDatePicker)))).GetId()) + `').data('kendoDatePicker').value( value )`)
		case ***KendoUiDateTimePicker:
			pass = true
			key = []byte((*(*(*v.(***KendoUiDateTimePicker)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiDateTimePicker)))).GetId()) + `').data('kendoDateTimePicker').value( value )`)
		case ***HtmlElementFormSelect:
			pass = true
			key = []byte((*(*(*v.(***HtmlElementFormSelect)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlElementFormSelect)))).GetId()) + `').val( value )`)
		case ***HtmlElementFormTextArea:
			pass = true
			key = []byte((*(*(*v.(***HtmlElementFormTextArea)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlElementFormTextArea)))).GetId()) + `').val( value )`)
		case ***HtmlInputCheckBox:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputCheckBox)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputCheckBox)))).GetId()) + `').val( value )`)
		case ***HtmlInputColor:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputColor)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputColor)))).GetId()) + `').val( value )`)
		case ***HtmlInputDate:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputDate)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputDate)))).GetId()) + `').val( value )`)
		case ***HtmlInputDateTimeLocal:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputDateTimeLocal)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputDateTimeLocal)))).GetId()) + `').val( value )`)
		case ***HtmlInputEmail:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputEmail)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputEmail)))).GetId()) + `').val( value )`)
		case ***HtmlInputFile:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputFile)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputFile)))).GetId()) + `').val( value )`)
		case ***HtmlInputGeneric:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputGeneric)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputGeneric)))).GetId()) + `').val( value )`)
		case ***HtmlInputImage:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputImage)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputImage)))).GetId()) + `').val( value )`)
		case ***HtmlInputMonth:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputMonth)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputMonth)))).GetId()) + `').val( value )`)
		case ***HtmlInputNumber:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputNumber)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputNumber)))).GetId()) + `').val( value )`)
		case ***HtmlInputPassword:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputPassword)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputPassword)))).GetId()) + `').val( value )`)
		case ***HtmlInputRadio:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputRadio)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputRadio)))).GetId()) + `').val( value )`)
		case ***HtmlInputRange:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputRange)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputRange)))).GetId()) + `').val( value )`)
		case ***HtmlInputSearch:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputSearch)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputSearch)))).GetId()) + `').val( value )`)
		case ***HtmlInputTel:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputTel)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputTel)))).GetId()) + `').val( value )`)
		case ***HtmlInputTime:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputTime)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputTime)))).GetId()) + `').val( value )`)
		case ***HtmlInputUrl:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputUrl)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputUrl)))).GetId()) + `').val( value )`)
		case ***HtmlInputWeek:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputWeek)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputWeek)))).GetId()) + `').val( value )`)
		case ***HtmlInputSubmit:
			pass = true
			key = []byte((*(*(*v.(***HtmlInputSubmit)))).Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***HtmlInputSubmit)))).GetId()) + `').val( value )`)
		case ***KendoUiMobileSwitch:
			pass = true
			key = []byte((*(*(*v.(***KendoUiMobileSwitch)))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*(*v.(***KendoUiMobileSwitch)))).GetId()) + `').val( value )`)
		case **AceEditor:
			pass = true
			key = []byte((*(*v.(**AceEditor))).Html.Name)
			jsCode = []byte(`ace.edit("` + string((*(*v.(**AceEditor))).GetId()) + `").setValue( value )`)
		case **KendoUiComboBox:
			pass = true
			key = []byte((*(*v.(**KendoUiComboBox))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiComboBox))).GetId()) + `').data('kendoComboBox').value( value )`)
		case **HtmlInputText:
			pass = true
			key = []byte((*(*v.(**HtmlInputText))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputText))).GetId()) + `').val( value )`)
		case **HtmlInputHidden:
			pass = true
			key = []byte((*(*v.(**HtmlInputHidden))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputHidden))).GetId()) + `').val( value )`)
		case **KendoUiDropDownList:
			pass = true
			key = []byte((*(*v.(**KendoUiDropDownList))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiDropDownList))).GetId()) + `').data('kendoDropDownList').value( value )`)
		case **KendoUiMultiSelect: // fixme: colocar o conteúdo da janela aqui também
			pass = true
			key = []byte((*(*v.(**KendoUiMultiSelect))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiMultiSelect))).GetId()) + `').data('kendoMultiSelect').value( value )`)
		case **KendoUiAutoComplete:
			pass = true
			key = []byte((*(*v.(**KendoUiAutoComplete))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiAutoComplete))).GetId()) + `').data('kendoAutoComplete').value( value )`)
		case **KendoUiButton:
			//pass = true
			//key = []byte( (*(*v.(**KendoUiButton))).Html.Name )
			//jsCode = []byte( `$('#` + string( (*(*v.(**KendoUiButton))).GetId() ) + `').data('kendoButton').value( value )` )
		case **KendoUiCalendar:
			pass = true
			key = []byte((*(*v.(**KendoUiCalendar))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiCalendar))).GetId()) + `').data('kendoCalendar').value( value )`)
		case **KendoUiColorPalette:
			pass = true
			key = []byte((*(*v.(**KendoUiColorPalette))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiColorPalette))).GetId()) + `').data('kendoColorPalette').value( value )`)
		case **KendoUiColorPicker:
			pass = true
			key = []byte((*(*v.(**KendoUiColorPicker))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiColorPicker))).GetId()) + `').data('kendoColorPicker').value( value )`)
		case **KendoUiNumericTextBox:
			pass = true
			key = []byte((*(*v.(**KendoUiNumericTextBox))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiNumericTextBox))).GetId()) + `').data('kendoNumericTextBox').value( value )`)
		case **KendoUiDateInput:
			pass = true
			key = []byte((*(*v.(**KendoUiDateInput))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiDateInput))).GetId()) + `').data('kendoDateInput').value( value )`)
		case **KendoUiDatePicker:
			pass = true
			key = []byte((*(*v.(**KendoUiDatePicker))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiDatePicker))).GetId()) + `').data('kendoDatePicker').value( value )`)
		case **KendoUiDateTimePicker:
			pass = true
			key = []byte((*(*v.(**KendoUiDateTimePicker))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiDateTimePicker))).GetId()) + `').data('kendoDateTimePicker').value( value )`)
		case **HtmlElementFormSelect:
			pass = true
			key = []byte((*(*v.(**HtmlElementFormSelect))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlElementFormSelect))).GetId()) + `').val( value )`)
		case **HtmlElementFormTextArea:
			pass = true
			key = []byte((*(*v.(**HtmlElementFormTextArea))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlElementFormTextArea))).GetId()) + `').val( value )`)
		case **HtmlInputCheckBox:
			pass = true
			key = []byte((*(*v.(**HtmlInputCheckBox))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputCheckBox))).GetId()) + `').val( value )`)
		case **HtmlInputColor:
			pass = true
			key = []byte((*(*v.(**HtmlInputColor))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputColor))).GetId()) + `').val( value )`)
		case **HtmlInputDate:
			pass = true
			key = []byte((*(*v.(**HtmlInputDate))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputDate))).GetId()) + `').val( value )`)
		case **HtmlInputDateTimeLocal:
			pass = true
			key = []byte((*(*v.(**HtmlInputDateTimeLocal))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputDateTimeLocal))).GetId()) + `').val( value )`)
		case **HtmlInputEmail:
			pass = true
			key = []byte((*(*v.(**HtmlInputEmail))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputEmail))).GetId()) + `').val( value )`)
		case **HtmlInputFile:
			pass = true
			key = []byte((*(*v.(**HtmlInputFile))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputFile))).GetId()) + `').val( value )`)
		case **HtmlInputGeneric:
			pass = true
			key = []byte((*(*v.(***HtmlInputGeneric))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputGeneric))).GetId()) + `').val( value )`)
		case **HtmlInputImage:
			pass = true
			key = []byte((*(*v.(**HtmlInputImage))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputImage))).GetId()) + `').val( value )`)
		case **HtmlInputMonth:
			pass = true
			key = []byte((*(*v.(**HtmlInputMonth))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputMonth))).GetId()) + `').val( value )`)
		case **HtmlInputNumber:
			pass = true
			key = []byte((*(*v.(**HtmlInputNumber))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputNumber))).GetId()) + `').val( value )`)
		case **HtmlInputPassword:
			pass = true
			key = []byte((*(*v.(**HtmlInputPassword))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputPassword))).GetId()) + `').val( value )`)
		case **HtmlInputRadio:
			pass = true
			key = []byte((*(*v.(**HtmlInputRadio))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputRadio))).GetId()) + `').val( value )`)
		case **HtmlInputRange:
			pass = true
			key = []byte((*(*v.(**HtmlInputRange))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputRange))).GetId()) + `').val( value )`)
		case **HtmlInputSearch:
			pass = true
			key = []byte((*(*v.(**HtmlInputSearch))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputSearch))).GetId()) + `').val( value )`)
		case **HtmlInputTel:
			pass = true
			key = []byte((*(*v.(**HtmlInputTel))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputTel))).GetId()) + `').val( value )`)
		case **HtmlInputTime:
			pass = true
			key = []byte((*(*v.(**HtmlInputTime))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputTime))).GetId()) + `').val( value )`)
		case **HtmlInputUrl:
			pass = true
			key = []byte((*(*v.(**HtmlInputUrl))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputUrl))).GetId()) + `').val( value )`)
		case **HtmlInputWeek:
			pass = true
			key = []byte((*(*v.(**HtmlInputWeek))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputWeek))).GetId()) + `').val( value )`)
		case **HtmlInputSubmit:
			pass = true
			key = []byte((*(*v.(**HtmlInputSubmit))).Name)
			jsCode = []byte(`$('#` + string((*(*v.(**HtmlInputSubmit))).GetId()) + `').val( value )`)
		case **KendoUiMobileSwitch:
			pass = true
			key = []byte((*(*v.(**KendoUiMobileSwitch))).Html.Name)
			jsCode = []byte(`$('#` + string((*(*v.(**KendoUiMobileSwitch))).GetId()) + `').val( value )`)

		case *AceEditor:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`ace.edit("` + string(converted.GetId()) + `").setValue( value )`)
		case *KendoUiComboBox:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoComboBox').value( value )`)
		case *HtmlInputText:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputHidden:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *KendoUiDropDownList:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoDropDownList').value( value )`)
		case *KendoUiMultiSelect: // fixme: colocar o conteúdo da janela aqui também
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoMultiSelect').value( value )`)
		case *KendoUiAutoComplete:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoAutoComplete').value( value )`)
		case *KendoUiButton:
			//pass = true
			//key = []byte( converted.Html.Name )
			//jsCode = []byte( `$('#` + string( converted.GetId() ) + `').data('kendoButton').value( value )` )
		case *KendoUiCalendar:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoCalendar').value( value )`)
		case *KendoUiColorPalette:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoColorPalette').value( value )`)
		case *KendoUiColorPicker:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoColorPicker').value( value )`)
		case *KendoUiNumericTextBox:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoNumericTextBox').value( value )`)
		case *KendoUiDateInput:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoDateInput').value( value )`)
		case *KendoUiDatePicker:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoDatePicker').value( value )`)
		case *KendoUiDateTimePicker:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').data('kendoDateTimePicker').value( value )`)
		case *HtmlElementFormSelect:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlElementFormTextArea:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputCheckBox:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputColor:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputDate:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputDateTimeLocal:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputEmail:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputFile:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputGeneric:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputImage:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputMonth:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputNumber:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputPassword:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputRadio:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputRange:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputSearch:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputTel:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputTime:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputUrl:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputWeek:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *HtmlInputSubmit:
			pass = true
			key = []byte(converted.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)
		case *KendoUiMobileSwitch:
			pass = true
			key = []byte(converted.Html.Name)
			jsCode = []byte(`$('#` + string(converted.GetId()) + `').val( value )`)

			//fixme: set value - start
		case *KendoUiGrid:
		case *HtmlElementScript:
		case *KendoUiWindow:
			//fixme: set value - end

		case *Content:
		case *KendoDataSource:
		default:
			fmt.Printf("\n\n\n-setValueById.id.type.%T.not.found.\n\n\n", converted)
		}

		if pass == true {
			buffer.Write([]byte("          case 'id:"))
			buffer.Write(key)
			buffer.Write([]byte("': return "))
			buffer.Write(jsCode)
			buffer.Write([]byte(";\n"))
		}
	}
	buffer.Write([]byte("        }\n"))
	buffer.Write([]byte("      }\n"))

	return buffer.Bytes()
}
func (el *Content) MakeJsScript() []byte {
	var buffer bytes.Buffer

	var elementList = el.FilterFormElements()
	for keyElementList := range elementList {
		//switch converted := elementList[ keyElementList ].(type) {
		//default:
		//  fmt.Printf( "%T\n\n\n\n", converted )
		//}
		switch elementList[keyElementList].(type) {
		case **KendoUiMultiSelect:

			(*(*elementList[keyElementList].(**KendoUiMultiSelect))).Html.Global.Id = string((*(*elementList[keyElementList].(**KendoUiMultiSelect))).GetId())

			if reflect.DeepEqual((*(*elementList[keyElementList].(**KendoUiMultiSelect))).DataSource, KendoDataSource{}) == false && reflect.DeepEqual((*(*elementList[keyElementList].(**KendoUiMultiSelect))).Dialog, KendoUiDialog{}) == false {

				switch (*(*elementList[keyElementList].(**KendoUiMultiSelect))).NoDataTemplate.(type) {
				case HtmlElementScript:

					for k := range (*(*elementList[keyElementList].(**KendoUiMultiSelect))).NoDataTemplate.(HtmlElementScript).Content {

						switch (*(*elementList[keyElementList].(**KendoUiMultiSelect))).NoDataTemplate.(HtmlElementScript).Content[k].(type) {
						case *HtmlElementFormButton:

							if (*(*elementList[keyElementList].(**KendoUiMultiSelect))).NoDataTemplate.(HtmlElementScript).Content[k].(*HtmlElementFormButton).ButtonType == BUTTON_TYPE_ADD_IN_TEMPLATE {
								//el.NoDataTemplate.(HtmlElementScript).Content[ k ].(HtmlElementFormButton).Global.OnClick = "addNewItemToKendoDataSource('id:#: instance.element[0].id #')"

								//reflect.ValueOf(&el.NoDataTemplate.(HtmlElementScript).Content[ k ].(*HtmlElementFormButton).Global).FieldByName("OnClick").SetString("esta vivo")
								p := reflect.ValueOf(&(*(*elementList[keyElementList].(**KendoUiMultiSelect))).NoDataTemplate.(HtmlElementScript).Content[k].(*HtmlElementFormButton).Global.OnClick)
								p.Elem().SetString("addNewItemToKendoDataSource('id:#: instance.element[0].id #')")
								//fmt.Printf( "------%v\n\n\n\n\n\n\n\n", p.Elem().CanSet() )
								//fmt.Printf( "------%v--------\n\n\n\n\n\n\n\n", (*(*elementList[ keyElementList ].(**KendoUiMultiSelect))).NoDataTemplate.(HtmlElementScript).Content[ k ].(*HtmlElementFormButton).Global.OnClick )
								//el.NoDataTemplate.(HtmlElementScript).Content[ k ].(HtmlElementFormButton).Global.OnClick = "esta vivo"
							}

						}

					}

				}
			}

			switch tplConverted := (*(*elementList[keyElementList].(**KendoUiMultiSelect))).FooterTemplate.(type) {
			case HtmlElementScript:
				if tplConverted.Type == SCRIPT_TYPE_KENDO_TEMPLATE {
					buffer.Write(tplConverted.ToElementScriptTag())
				}
			}
			switch tplConverted := (*(*elementList[keyElementList].(**KendoUiMultiSelect))).NoDataTemplate.(type) {
			case HtmlElementScript:
				if tplConverted.Type == SCRIPT_TYPE_KENDO_TEMPLATE {
					buffer.Write(tplConverted.ToElementScriptTag())
				}
			}
			switch tplConverted := (*(*elementList[keyElementList].(**KendoUiMultiSelect))).HeaderTemplate.(type) {
			case HtmlElementScript:
				if tplConverted.Type == SCRIPT_TYPE_KENDO_TEMPLATE {
					buffer.Write(tplConverted.ToElementScriptTag())
				}
			}
		}
	}

	return buffer.Bytes()
}
