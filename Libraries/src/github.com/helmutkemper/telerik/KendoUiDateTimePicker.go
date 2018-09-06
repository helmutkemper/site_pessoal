package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
	"time"
)

type KendoUiDateTimePicker struct {
	Html HtmlInputText `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-animation

	  Configures the opening and closing animations of the popups. Setting the <b><u>animation</u></b> option to <b><u>false</u></b> will disable the opening and closing animations. As a result the popup will open and close instantly.
	  <b><u>animation:true</u></b> is not a valid configuration.

	  Example - disable open and close animations
	   <input id="datetimepicker" />
	   <script>
	   $("#datetimepicker").kendoDateTimePicker({
	     animation: false
	   });
	   </script>
	*/
	Animation interface{} `jsObject:"animation" jsType:"*KendoAnimation,Boolean"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-ARIATemplate

	  Specifies a template used to populate value of the aria-label attribute.

	  Example
	   <input id="datetimepicker" />
	   <script>
	   $("#datetimepicker").kendoDateTimePicker({
	       ARIATemplate: "Date: #=kendo.toString(data.current, 'G')#"
	   });
	   </script>
	*/
	ARIATemplate interface{} `jsObject:"ARIATemplate" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-culture

	  Specifies the culture info used by the widget.

	  Example - specify German culture internationalization
	   <input id="datetimepicker" />
	   <script>
	   $("#datetimepicker").kendoDateTimePicker({
	       culture: "de-DE"
	   });
	   </script>
	*/
	Culture string `jsObject:"culture"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-dateInput

	  Specifies if the DateTimePicker will use DateInput for editing value

	  Example
	   <input id="datetimepicker" />
	   <script>
	   $("#datetimepicker").DateTimePicker({
	       dateInput: true
	   });
	   </script>
	*/
	DateInput Boolean `jsObject:"dateInput"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-dates

	  Specifies a list of dates, which will be passed to the <a href="https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-month.content">month template</a> of the DateView. All dates, which match the date portion of the selected date will be used to re-bind the TimeView.

	  Example - specify a list of dates
	   <style>
	     .party{color:red}
	   </style>

	   <input id="datetimepicker" />

	   <script id="cell-template" type="text/x-kendo-template">
	     <span class="#= isInArray(data.date, data.dates) ? 'party' : '' #">#= data.value #</span>
	   </script>

	   <script>
	     $("#datetimepicker").kendoDateTimePicker({
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
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-depth

	  Specifies the navigation depth of the calendar. The following settings are available for the <strong>depth</strong> value:
	  Note the option will not be applied if <strong>start</strong> option is <em>lower</em> than <strong>depth</strong>. Always set both and <strong>start</strong> and <strong>depth</strong> options.

	  Example - set navigation depth of the calendar popup
	   <input id="datetimepicker"/>
	   <script>
	   $("#datetimepicker").kendoDateTimePicker({
	       depth: "year"
	   });
	   </script>
	*/
	Depth KendoTimeDepth `jsObject:"depth"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-disableDates

	  An array or function that will be used to determine which dates to be disabled for selection by the widget.
	  you can also pass a function that will be dynamically resolved for each date of the calendar. Note that when the function returns true, the date will be disabled.
	  note that a check for an empty <b><u>date</u></b> is needed, as the widget can work with a null value as well.
	  This functionality was added with the Q1 release of 2016.

	  Example - specify an array of days to be disabled
	   <input id="datetimepicker" />
	   <script>
	   $("#datetimepicker").kendoDateTimePicker({
	       value: new Date(),
	       disableDates: ["we", "th"],
	   });
	   </script>
	*/
	DisableDates KendoWeekDays `jsObject:"disableDates"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-footer

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">template</a> which renders the footer of the calendar. If false, the footer will not be rendered.

	  Example - specify footer template as a function
	   <input id="datetimepicker" />
	   <script id="footer-template" type="text/x-kendo-template">
	       Today - #: kendo.toString(data, "d") #
	   </script>
	   <script>
	   $("#datetimepicker").kendoDateTimePicker({
	       footer: kendo.template($("#footer-template").html())
	   });
	   </script>
	*/
	FooterTemplate interface{} `jsObject:"footer" jsType:"*JavaScript,string"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-format

	  Specifies the format, which is used to format the value of the DateTimePicker displayed in the input. The format also will be used to parse the input.

	  Example - specify a custom date format
	   <input id="datetimepicker" />
	   <script>
	   $("#datetimepicker").kendoDateTimePicker({
	       format: "yyyy/MM/dd hh:mm tt"
	   });
	   </script>
	*/
	Format string `jsObject:"format"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-interval

	  Specifies the interval, between values in the popup list, in minutes.

	  Example - specify a time interval
	   <input id="dateTimePicker" />
	   <script>
	   $("#dateTimePicker").kendoDateTimePicker({
	       interval: 15
	   });
	   </script>
	*/
	Interval int `jsObject:"interval"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-max

	  Specifies the maximum date, which the calendar can show.

	  Example - specify the maximum date
	   <input id="datetimepicker" />
	   <script>
	   $("#datetimepicker").kendoDateTimePicker({
	       max: new Date(2013, 0, 1, 22, 0, 0)
	   });
	   </script>
	*/
	Max time.Time `jsObject:"max"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-min

	  Specifies the minimum date that the calendar can show.

	  Example - specify the minimum date
	   <input id="datetimepicker" />
	   <script>
	   $("#datetimepicker").kendoDateTimePicker({
	       min: new Date(2011, 0, 1, 8, 0, 0)
	   });
	   </script>
	*/
	Min time.Time `jsObject:"min"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-month

	  Templates for the cells rendered in the calendar "month" view.
	*/
	Month KendoMonth `jsObject:"month"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-weekNumber

	  If set to <b><u>true</u></b> a week of the year will be shown on the left side of the calendar. It is possible to define a template in order to customize what will be displayed.

	  Example - enable the week of the year option
	   <input id="datetimepicker1" />
	   <script>
	       $("#datetimepicker1").kendoDateTimePicker({
	           weekNumber: true
	       });
	   </script>
	*/
	WeekNumber Boolean `jsObject:"weekNumber"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-parseFormats

	  Specifies the formats, which are used to parse the value set with value() method or by direct input. If not set the value of the <b><u>options.format</u></b> and <b><u>options.timeFormat</u></b> will be used. Note that value of the <b><u>format</u></b> option is always used. The <b><u>timeFormat</u></b> value also will be used if defined.
	  Order of the provided parse formats is important and it should from stricter to less strict.

	  Example
	   <input id="datetimepicker" />
	   <script>
	   $("#datetimepicker").kendoDateTimePicker({
	       format: "yyyy/MM/dd hh:mm tt",
	       parseFormats: ["MMMM yyyy", "HH:mm"] //format also will be added to parseFormats
	   });
	   </script>
	*/
	ParseFormats []string `jsObject:"parseFormats"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-start

	  Specifies the start view of the calendar. The following settings are available for the <strong>start</strong> value:

	  Example - specify the initial view, which calendar renders
	   <input id="datetimepicker" />
	   <script>
	       $("#datetimepicker").kendoDateTimePicker({
	           start: "year"
	       });
	   </script>
	*/
	Start KendoTimeDepth `jsObject:"start"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-timeFormat

	  Specifies the format, which is used to format the values in the time drop-down list.

	  Example
	   <input id="datetimepicker" />
	   <script>
	   $("#datetimepicker").kendoDateTimePicker({
	       timeFormat: "HH:mm" //24 hours format
	   });
	   </script>
	*/
	TimeFormat string `jsObject:"timeFormat"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/datetimepicker#configuration-value

	  Specifies the selected value.

	  Example
	   <input id="datetimepicker" />
	   <script>
	   $("#datetimepicker").kendoDateTimePicker({
	       value: new Date(2011, 0, 1)
	   });
	   </script>
	*/
	Value time.Time `jsObject:"value"`

	*ToJavaScriptConverter
}

func (el *KendoUiDateTimePicker) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("kendoDateTimePicker.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoDateTimePicker({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiDateTimePicker) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiDateTimePicker) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiDateTimePicker) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
