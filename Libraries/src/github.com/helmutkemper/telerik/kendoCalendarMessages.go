package telerik

import (
	log "github.com/helmutkemper/seelog"
	"reflect"
)

type KendoCalendarMessages struct {
	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/calendar/configuration/messages

	  Allows localization of the strings that are used in the widget. (default: "")

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

	  Allows customization of the week column header text. Set the value to make the widget compliant with web accessibility standards. (default: "")

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
	WeekColumnHeader string `jsObject:"weekColumnHeader"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput/configuration/messages#messages.year

	  The placeholder for the years part.
	  (default: "year")

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
	Year string `jsObject:"year"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput/configuration/messages#messages.month

	  The placeholder for the months part.
	  (default: "month")

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
	Month string `jsObject:"month"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput/configuration/messages#messages.day

	  The placeholder for the day of the month part.
	  (default: "day")

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
	Day string `jsObject:"day"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput/configuration/messages#messages.weekday

	  The placeholder for the day of the week part.
	  (default: "day of the week")

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
	Weekday string `jsObject:"weekday"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput/configuration/messages#messages.hour

	  The placeholder for the hours part.
	  (default: "hours")

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
	Hour string `jsObject:"hour"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput/configuration/messages#messages.minute

	  The placeholder for the minutes part.
	  (default: "minutes")

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
	Minute string `jsObject:"minute"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput/configuration/messages#messages.second

	  The placeholder for the seconds part.
	  (default: "seconds")

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
	Second string `jsObject:"second"`

	/*
	  @see https://docs.telerik.com/kendo-ui/api/javascript/ui/dateinput/configuration/messages#messages.dayperiod

	  The placeholder for the AM/PM part.
	  (default: "AM/PM")

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
	DayPeriod string `jsObject:"dayperiod"`

	*ToJavaScriptConverter
}

func (el *KendoCalendarMessages) ToJavaScript() []byte {
	element := reflect.ValueOf(el).Elem()
	ret, err := el.ToJavaScriptConverter.ToTelerikJavaScript(element)
	if err != nil {
		log.Criticalf("KendoCalendarMessages.Error: %v", err.Error())
		return []byte{}
	}

	return ret
}
