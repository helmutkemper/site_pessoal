package telerik

import (
	"bytes"
	log "github.com/helmutkemper/seelog"
	"reflect"
	"time"
)

type KendoUiCalendar struct {
	Html HtmlElementDiv `jsObject:"-"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-culture

	  Specifies the culture info used by the widget.

	  Example - specify German culture internationalization
	   <div id="calendar"></div>
	   <script>
	       $("#calendar").kendoCalendar({
	           culture: "de-DE"
	       });
	   </script>
	*/
	Culture string `jsObject:"culture"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-dates

	  Specifies a list of dates, which will be passed to the month template.

	  Example - specify a list of dates
	   <div id="calendar"></div>
	   <script>
	       $("#calendar").kendoCalendar({
	           value: new Date(2000, 10, 1),
	           dates: [
	               new Date(2000, 10, 10, 10, 0, 0),
	               new Date(2000, 10, 10, 30, 0)
	           ] //can manipulate month template depending on this array.
	       });
	   </script>
	*/
	Dates []time.Time `jsObject:"dates"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-depth

	  Specifies the navigation depth. The following settings are available for the <strong>depth</strong> value:
	  Note the option will not be applied if <strong>start</strong> option is <em>lower</em> than <strong>depth</strong>. Always set both and <strong>start</strong> and <strong>depth</strong> options.

	  Example - set navigation depth of the calendar
	   <div id="calendar"></div>
	   <script>
	       $("#calendar").kendoCalendar({
	           depth: "year"
	       });
	   </script>
	*/
	Depth KendoTimeDepth `jsObject:"depth"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-disableDates

	  An array or function that will be used to determine which dates to be disabled in the calendar.
	  you can also pass a function that will be dynamically resolved for each date of the calendar. Note that when the function returns true, the date will be disabled.
	  note that a check for an empty <b><u>date</u></b> is needed, as the widget can work with a null value as well.
	  This functionality was added with the Q1 release of 2016.

	  Example - specify an array of days to be disabled
	   <div id="calendar"></div>
	   <script>
	   $("#calendar").kendoCalendar({
	       value: new Date(),
	       disableDates: ["we", "th"],
	   });
	   </script>
	*/
	DisableDates KendoWeekDays `jsObject:"disableDates"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-footer

	  The <a href="/kendo-ui/api/javascript/kendo#methods-template">template</a> which renders the footer. If false, the footer will not be rendered.

	  Example - specify footer template as a function
	   <div id="calendar"></div>
	   <script id="footer-template" type="text/x-kendo-template">
	       Today - #: kendo.toString(data, "d") #
	   </script>
	   <script>
	       $("#calendar").kendoCalendar({
	           footer: kendo.template($("#footer-template").html())
	       });
	   </script>
	*/
	Footer string `jsObject:"footer"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-format

	  Specifies the format, which is used to parse value set with value() method.

	  Example - specify a custom date format
	   <div id="calendar"></div>
	   <script>
	       $("#calendar").kendoCalendar({
	           format: "yyyy/MM/dd"
	       });
	   </script>
	*/
	Format string `jsObject:"format"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-max

	  Specifies the maximum date, which the calendar can show.

	  Example - specify the maximum date
	   <div id="calendar"></div>
	   <script>
	       $("#calendar").kendoCalendar({
	           max: new Date(2013, 0, 1) // set the max date to Jan 1st, 2013
	       });
	   </script>
	*/
	Max time.Time `jsObject:"max"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-messages

	  Allows localization of the strings that are used in the widget.

	  Example
	   <div id="calendar"></div>
	   <script>
	   $("#calendar").kendoCalendar({
	       "weekNumber": true,
	       "messages": {
	           "weekColumnHeader": "W"
	       }
	    })
	   </script>
	*/
	Messages KendoCalendarMessages `jsObject:"messages"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-min

	  Specifies the minimum date, which the calendar can show.

	  Example - specify the minimum date
	   <div id="calendar"></div>
	   <script>
	       // set the min date to Jan 1st, 2011
	       $("#calendar").kendoCalendar({
	           min: new Date(2011, 0, 1)
	       });
	   </script>
	*/
	Min time.Time `jsObject:"min"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-month

	  Templates for the cells rendered in "month" view.

	  Example - specify cell template as a string
	  <style>
	    .exhibition{
	      background-color: #9DD0E0;
	      color:black;
	    }
	    .party{
	      color: red;
	      background-color: #ccc;
	    }
	  </style>
	  <body>

	  <div id="calendar"></div>
	  <script id="cell-template" type="text/x-kendo-template">
	      <div class="#= data.value < 10 ? 'exhibition' : 'party' #">
	      #= data.value #
	    </div>
	  </script>
	  <script>
	    $("#calendar").kendoCalendar({
	      month: {
	        content: $("#cell-template").html()
	      }
	    });
	  </script>
	*/
	Month KendoMonth `jsObject:"month"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-selectable

	  By default user is able to select a single date. The property can also be set to "multiple" in order the multiple date selection to be enabled. More information about multiple selection can be found in the <a href="/kendo-ui/controls/scheduling/calendar/overview#selection">Selection</a> article.
	  Check <a href="http://demos.telerik.com/kendo-ui/calendar/selection">Selection</a> for a live demo.

	  Example - enable the multiple selection
	   <div id="calendar"></div>
	   <script>
	       $("#calendar").kendoCalendar({
	           selectable: "multiple"
	       });
	   </script>
	*/
	Selectable string `jsObject:"selectable"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-selectDates

	  Specifies which dates to be selected when the calendar is initialized.
	  <strong>Important:</strong> This configuration option requires the <a href="/kendo-ui/api/javascript/ui/calendar#configuration-selectable">selectable</a>: "multiple" option to be set.
	  Check <a href="http://demos.telerik.com/kendo-ui/calendar/selection">Selection</a> for a live demo.

	  Example - set two dates to be selected upon calendar initialization
	   <div id="calendar"></div>
	   <script>
	       $("#calendar").kendoCalendar({
	          selectDates: [new Date(2013, 10, 10), new Date(2015, 10, 10)]
	       });
	   </script>
	*/
	SelectDates []time.Time `jsObject:"selectDates"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-weekNumber

	  If set to <b><u>true</u></b> a week of the year will be shown on the left side of the calendar.

	  Example - enable the week of the year option
	   <div id="calendar"></div>
	   <script>
	       $("#calendar").kendoCalendar({
	           weekNumber: true
	       });
	   </script>
	*/
	WeekNumber Boolean `jsObject:"weekNumber"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-start

	  Specifies the start view. The following settings are available for the <strong>start</strong> value:

	  Example - specify the initial view, which calendar renders
	   <div id="calendar"></div>
	   <script>
	       $("#calendar").kendoCalendar({
	           start: "year"
	       });
	   </script>
	*/
	Start string `jsObject:"start"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar#configuration-value

	  Specifies the selected date.

	  Example - specify the selected value of the widget
	   <div id="calendar"></div>
	   <script>
	       $("#calendar").kendoCalendar({
	           value: new Date(2012, 0, 1)
	       });
	   </script>
	*/
	Value time.Time `jsObject:"value"`

	*ToJavaScriptConverter `jsObject:"-"`
}

func (el *KendoUiCalendar) ToJavaScript() []byte {
	var ret bytes.Buffer
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}

	element := reflect.ValueOf(el).Elem()
	data, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoUiCalendar.Error: %v", err.Error())
		return []byte{}
	}

	ret.Write([]byte(`$("#` + el.Html.Global.Id + `").kendoCalendar({`))
	ret.Write(data)
	ret.Write([]byte(`});`))
	ret.Write([]byte{0x0A})

	return ret.Bytes()
}
func (el *KendoUiCalendar) ToHtml() []byte {
	return el.Html.ToHtml()
}
func (el *KendoUiCalendar) GetId() []byte {
	if el.Html.Global.Id == "" {
		el.Html.Global.Id = GetAutoId()
	}
	return []byte(el.Html.Global.Id)
}
func (el *KendoUiCalendar) GetName() []byte {
	if el.Html.Name == "" {
		el.Html.Name = GetAutoId()
	}
	return []byte(el.Html.Name)
}
