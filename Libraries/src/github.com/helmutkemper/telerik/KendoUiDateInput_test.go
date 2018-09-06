package telerik

import (
	"fmt"
	"time"
)

func ExampleKendoUiDateInput_ToHtml() {
	html := KendoUiDateInput{
		Html: HtmlInputDate{
			Global: HtmlGlobalAttributes{
				Id: "dateinput",
			},
		},
		Format: "yyyy/MM/dd",
		Max:    time.Date(2018, 12, 31, 23, 59, 59, 0, time.UTC),
		Min:    time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		Value:  time.Date(2018, 02, 01, 0, 0, 0, 0, time.UTC),
		Messages: KendoCalendarMessages{
			Year:      "year",
			Month:     "month",
			Day:       "day",
			Weekday:   "day of the week",
			Hour:      "hour",
			Minute:    "minute",
			Second:    "second",
			DayPeriod: "AM/PM",
		},
	}

	fmt.Printf("%s", html.ToJavaScript())

	// Output:
	// $("#dateinput").kendoDateInput({format: "yyyy/MM/dd",max: new Date(2018, 12, 31, 23, 59, 59),min: new Date(2018, 1, 1),value: new Date(2018, 2, 1),messages: {year: "year",month: "month",day: "day",weekday: "day of the week",hour: "hour",minute: "minute",second: "second",dayperiod: "AM/PM",},});
}
