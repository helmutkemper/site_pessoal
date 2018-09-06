package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
	"time"
)

type KendoUiDateInput struct {
	Html HtmlInputDate `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput#configuration-format

	  Specifies the format, which is used to format the value of the DateInput displayed in the input. The format also will be used to parse the input.

	  Example - specify a custom date format
	   <input id="dateinput" />
	   <script>
	   $("#dateinput").kendoDateInput({
	       format: "yyyy/MM/dd"
	   });
	   </script>
	*/
	Format string `jsObject:"format"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput#configuration-max

	  Specifies the maximum date which can be entered in the input.

	  Example - specify the maximum date
	   <input id="dateinput" />
	   <script>
	   $("#dateinput").kendoDateInput({
	       max: new Date(2013, 0, 1) // sets max date to Jan 1st, 2013
	   });
	   </script>
	*/
	Max time.Time `jsObject:"max"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput#configuration-min

	  Specifies the minimum date that which be entered in the input.

	  Example - specify the minimum date
	   <input id="dateinput" />
	   <script>
	   $("#dateinput").kendoDateInput({
	       min: new Date(2011, 0, 1) // sets min date to Jan 1st, 2011
	   });
	   </script>
	*/
	Min time.Time `jsObject:"min"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput#configuration-value

	  Specifies the selected date.

	  Example
	   <input id="dateinput" />
	   <script>
	   $("#dateinput").kendoDateInput({
	       value: new Date(2011, 0, 1)
	   });
	   </script>
	*/
	Value time.Time `jsObject:"value"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput#configuration-messages

	  The messages that DateInput uses. Use it to customize or localize the placeholders of each date/time part.

	  Example - customize column menu messages
	   <input id="dateinput" />
	   <script>
	   $("#dateinput").kendoDateInput({
	       messages:{
	           "year": "year",
	           "month": "month",
	           "day": "day",
	           "weekday": "day of the week",
	           "hour": "hours",
	           "minute": "minutes",
	           "second": "seconds",
	           "dayperiod": "AM/PM"
	       }
	   });
	   </script>
	*/
	Messages KendoCalendarMessages `jsObject:"messages"`

	*ToJavaScriptConverter
}

func (el *KendoUiDateInput) ToJavaScript() []byte {
	var ret bytes.Buffer

	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("kendoDateInput.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoDateInput({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiDateInput) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiDateInput) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiDateInput) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
