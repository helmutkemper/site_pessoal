package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
	"time"
)

type KendoUiDatePicker struct {
	Html HtmlInputText `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-animation

	  Configures the opening and closing animations of the calendar popup. Setting the <b><u>animation</u></b> option to <b><u>false</u></b> will disable the opening and closing animations. As a result the calendar popup will open and close instantly.
	  <b><u>animation:true</u></b> is not a valid configuration.

	  Example - disable open and close animations
	   <input id="datepicker" />
	   <script>
	   $("#datepicker").kendoDatePicker({
	     animation: false
	   });
	   </script>
	*/
	Animation interface{} `jsObject:"animation" jsType:"*KendoAnimation,Boolean"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-ARIATemplate

	  Specifies a template used to populate value of the aria-label attribute.
	  (default: "Current focused date is #=kendo.toString(data.current, 'D')#")

	  Example
	   <input id="datepicker" />
	   <script>
	   $("#datepicker").kendoDatePicker({
	       ARIATemplate: "Date: #=kendo.toString(data.current, 'G')#"
	   });
	   </script>
	*/
	ARIATemplate interface{} `jsObject:"ARIATemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-culture

	  Specifies the culture info used by the widget.

	  Example - specify German culture internationalization
	   <input id="datepicker" />
	   <script>
	   $("#datepicker").kendoDatePicker({
	       culture: "de-DE"
	   });
	   </script>
	*/
	Culture string `jsObject:"culture"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-dateInput

	  Specifies if the DatePicker will use <a href="/kendo-ui/api/javascript/ui/dateinput">DateInput</a> for editing value

	  Example
	   <input id="datepicker" />
	   <script>
	   $("#datepicker").kendoDatePicker({
	       dateInput: true
	   });
	   </script>
	*/
	DateInput Boolean `jsObject:"dateInput"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-dates

	  Specifies a list of dates, which will be passed to the <a href="https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-month.content">month template</a>.

	  Example - specify a list of dates
	   <style>
	     .party{color:red}
	   </style>

	   <input id="datepicker" />

	   <script id="cell-template" type="text/x-kendo-template">
	     <span class="#= isInArray(data.date, data.dates) ? 'party' : '' #">#= data.value #</span>
	   </script>

	   <script>
	     $("#datepicker").kendoDatePicker({
	       value: new Date(2000, 10, 1),
	       month: {
	         content: $("#cell-template").html()
	       },
	       dates: [
	         new Date(2000, 10, 10),
	         new Date(2000, 10, 30)
	       ] //can manipulate month template depending on this array.
	     });

	     function isInArray(date, dates) {
	       for(var idx = 0, length = dates.length; idx < length; idx++) {
	         var d = dates[idx];
	         if (date.getFullYear() == d.getFullYear() &amp;&amp;
	             date.getMonth() == d.getMonth() &amp;&amp;
	             date.getDate() == d.getDate()) {
	           return true;
	         }
	       }

	       return false;
	     }

	   </script>
	*/
	Dates []time.Time `jsObject:"dates"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-depth

	  Specifies the navigation depth. The following settings are available for the <strong>depth</strong> value:
	  Note the option will not be applied if <strong>start</strong> option is <em>lower</em> than <strong>depth</strong>. Always set both and <strong>start</strong> and <strong>depth</strong> options.

	  Example - set navigation depth of the calendar popup
	   <input id="datepicker"/>
	   <script>
	   $("#datepicker").kendoDatePicker({
	       depth: "year",
	       start: "year"
	   });
	   </script>
	*/
	Depth KendoTimeDepth `jsObject:"depth"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-disableDates

	  An array or function that will be used to determine which dates to be disabled for selection by the widget.
	  you can also pass a function that will be dynamically resolved for each date of the calendar. Note that when the function returns true, the date will be disabled.
	  note that a check for an empty <b><u>date</u></b> is needed, as the widget can work with a null value as well.
	  This functionality was added with the Q1 release of 2016.

	  Example - specify an array of days to be disabled
	   <input id="datepicker" />
	   <script>
	   $("#datepicker").kendoDatePicker({
	       value: new Date(),
	       disableDates: ["we", "th"],
	   });
	   </script>
	*/
	DisableDates KendoWeekDays `jsObject:"disableDates"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-footer

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">template</a> which renders the footer of the calendar. If false, the footer will not be rendered.

	  Example - specify footer template as a function
	   <input id="datepicker" />
	   <script id="footer-template" type="text/x-kendo-template">
	       Today - #: kendo.toString(data, "d") #
	   </script>
	   <script>
	   $("#datepicker").kendoDatePicker({
	       footer: kendo.template($("#footer-template").html())
	   });
	   </script>
	*/
	FooterTemplate interface{} `jsObject:"footer" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-format

	  Specifies the format, which is used to format the value of the DatePicker displayed in the input. The format also will be used to parse the input.

	  Example - specify a custom date format
	   <input id="datepicker" />
	   <script>
	   $("#datepicker").kendoDatePicker({
	       format: "yyyy/MM/dd"
	   });
	   </script>
	*/
	Format string `jsObject:"format"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-max

	  Specifies the maximum date, which the calendar can show.

	  Example - specify the maximum date
	   <input id="datepicker" />
	   <script>
	   $("#datepicker").kendoDatePicker({
	       max: new Date(2013, 0, 1) // sets max date to Jan 1st, 2013
	   });
	   </script>
	*/
	Max time.Time `jsObject:"max"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-min

	  Specifies the minimum date that the calendar can show.

	  Example - specify the minimum date
	   <input id="datepicker" />
	   <script>
	   $("#datepicker").kendoDatePicker({
	       min: new Date(2011, 0, 1) // sets min date to Jan 1st, 2011
	   });
	   </script>
	*/
	Min time.Time `jsObject:"min"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-month

	  Templates for the cells rendered in the calendar "month" view.
	*/
	Month KendoMonth `jsObject:"month"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-weekNumber

	  If set to <b><u>true</u></b> a week of the year will be shown on the left side of the calendar. It is possible to define a template in order to customize what will be displayed.

	  Example - enable the week of the year option
	   <input id="datepicker1" />
	   <script>
	       $("#datepicker1").kendoDatePicker({
	           weekNumber: true
	       });
	   </script>
	*/
	WeekNumber Boolean `jsObject:"weekNumber"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-parseFormats

	  Specifies a list of date formats used to parse the value set with <b><u>value()</u></b> method or by direct user input. If not set the value of the format will be used. Note that the <b><u>format</u></b> option is always used during parsing.
	  The order of the provided parse formats is important and it should go from more strict to less strict.

	  Example
	   <input id="datepicker" />
	   <script>
	   $("#datepicker").kendoDatePicker({
	       format: "yyyy/MM/dd",
	       parseFormats: ["MMMM yyyy"] //format also will be added to parseFormats
	   });
	   </script>
	*/
	ParseFormats []string `jsObject:"parseFormats"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-start

	  Specifies the start view. The following settings are available for the <strong>start</strong> value:

	  Example - specify the initial view, which calendar renders
	   <input id="datepicker" />
	   <script>
	       $("#datepicker").kendoDatePicker({
	           start: "year"
	       });
	   </script>
	*/
	Start KendoTimeDepth `jsObject:"start"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datepicker#configuration-value

	  Specifies the selected date.

	  Example
	   <input id="datepicker" />
	   <script>
	   $("#datepicker").kendoDatePicker({
	       value: new Date(2011, 0, 1)
	   });
	   </script>
	*/
	Value time.Time `jsObject:"value"`

	*ToJavaScriptConverter
}

func (el *KendoUiDatePicker) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("kendoDatePicker.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoDatePicker({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiDatePicker) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiDatePicker) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiDatePicker) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
