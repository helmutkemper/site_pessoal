package telerik

import (
	"fmt"
	"time"
)

func ExampleKendoUiDatePicker_ToHtml() {
	html := KendoUiDatePicker{
		Html: HtmlInputText{
			Global: HtmlGlobalAttributes{
				Id: "datepicker",
			},
		},
		Animation: KendoAnimation{
			Open: KendoOpen{
				Duration: 300,
				Effects:  EFFECT_EXPAND_OUT,
			},
			Close: KendoClose{
				Duration: 300,
				Effects:  EFFECT_EXPAND_IN,
			},
		},
		ARIATemplate: "Date: #=kendo.toString(data.current, 'G')#",
		Culture:      "de-DE",
		DateInput:    TRUE,
		Dates: []time.Time{
			time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC),
			time.Date(2000, 10, 30, 0, 0, 0, 0, time.UTC),
		},
		Depth: TIME_DEPTH_MONTH,
		DisableDates: KendoWeekDays{
			WEEK_DAY_SATURDAY,
			WEEK_DAY_SUNDAY,
		},
		Format:     "yyyy/MM/dd",
		Max:        time.Date(2013, 0, 1, 0, 0, 0, 0, time.UTC),
		Min:        time.Date(2011, 0, 1, 0, 0, 0, 0, time.UTC),
		WeekNumber: TRUE,
		ParseFormats: []string{
			"MMMM yyyy",
		},
		Start: TIME_DEPTH_YEAR,
		Value: time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	fmt.Printf("%s", html.ToJavaScript())

	// Output:
	// $("#datepicker").kendoDatePicker({animation: { close: { effects: "expand:in",duration: 300,},open: { effects: "expand:out",duration: 300,},},ARIATemplate: "Date: #=kendo.toString(data.current, 'G')#",culture: "de-DE",dateInput: true,dates: [new Date(2000, 10, 10),new Date(2000, 10, 30)],depth: "month",disableDates: ["sa","su",],format: "yyyy/MM/dd",max: new Date(2012, 12, 1),min: new Date(2010, 12, 1),weekNumber: true,parseFormats: ["MMMM yyyy",],start: "year",value: new Date(2011, 1, 1),});
}
