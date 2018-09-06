package telerik

import (
	"fmt"
	"time"
)

func ExampleKendoUiCalendar_ToHtml() {
	calendar := KendoUiCalendar{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "calendar",
			},
		},
		Value: time.Date(2018, 3, 5, 22, 44, 0, 0, time.UTC),
	}

	fmt.Printf(`<!DOCTYPE html>
<html>
<head>
<base href="https://demos.telerik.com/kendo-ui/calendar/index">
<style>html { font-size: 14px; font-family: Arial, Helvetica, sans-serif; }</style>
<title></title>
<link rel="stylesheet" href="https://kendo.cdn.telerik.com/2018.1.221/styles/kendo.default-v2.min.css" />
<script src="https://kendo.cdn.telerik.com/2018.1.221/js/jquery.min.js"></script>
<script src="https://kendo.cdn.telerik.com/2018.1.221/js/kendo.all.min.js"></script>
</head>
<body>
<div id="example">
<div class="demo-section k-content" style="text-align: center;">
<h4>Pick a date</h4>
%s
</div>
<script>
$(document).ready(function() {
// create Calendar from div HTML element
%s
});
</script>
</div>
</body>
</html>`, calendar.ToHtml(), calendar.ToJavaScript())

	//Output:
	//<!DOCTYPE html>
	//<html>
	//<head>
	//<base href="https://demos.telerik.com/kendo-ui/calendar/index">
	//<style>html { font-size: 14px; font-family: Arial, Helvetica, sans-serif; }</style>
	//<title></title>
	//<link rel="stylesheet" href="https://kendo.cdn.telerik.com/2018.1.221/styles/kendo.default-v2.min.css" />
	//<script src="https://kendo.cdn.telerik.com/2018.1.221/js/jquery.min.js"></script>
	//<script src="https://kendo.cdn.telerik.com/2018.1.221/js/kendo.all.min.js"></script>
	//</head>
	//<body>
	//<div id="example">
	//<div class="demo-section k-content" style="text-align: center;">
	//<h4>Pick a date</h4>
	//<div id="calendar"></div>
	//</div>
	//<script>
	//$(document).ready(function() {
	//// create Calendar from div HTML element
	//$("#calendar").kendoCalendar({value: new Date(2018, 3, 5, 22, 44),});
	//});
	//</script>
	//</div>
	//</body>
	//</html>
}

func ExampleKendoUiCalendar_ToTelerikTemplate() {
	calendar := KendoUiCalendar{
		Html: HtmlElementDiv{
			Global: HtmlGlobalAttributes{
				Id: "calendar",
			},
		},
		Culture: "en-US",
		Dates: []time.Time{
			time.Date(2018, 2, 28, 0, 0, 0, 0, time.UTC),
			time.Date(2019, 2, 28, 23, 59, 0, 0, time.UTC),
			time.Date(2019, 2, 28, 23, 59, 59, 0, time.UTC),
		},
		Depth:        TIME_DEPTH_DAY,
		DisableDates: KendoWeekDays{WEEK_DAY_SATURDAY, WEEK_DAY_SUNDAY},
		Footer:       `kendo.template($('#footer-template').html())`,
		Format:       `yyyy/MM/dd`,
		Max:          time.Date(2019, 2, 28, 23, 59, 59, 0, time.UTC),
		Messages: KendoCalendarMessages{
			WeekColumnHeader: "W",
		},
		Min: time.Date(2018, 2, 28, 0, 0, 0, 0, time.UTC),
		Month: KendoMonth{
			Content:    `$('#cell-template').html()`,
			WeekNumber: `$('#week-template').html()`,
			Empty:      `-`,
		},
		Selectable: "multiple",
		SelectDates: []time.Time{
			time.Date(2018, 2, 28, 0, 0, 0, 0, time.UTC),
			time.Date(2019, 2, 28, 0, 0, 0, 0, time.UTC),
		},
		WeekNumber: TRUE,
		Start:      "year",
		Value:      time.Date(2018, 3, 5, 22, 44, 0, 0, time.UTC),
	}
	fmt.Printf(`<!DOCTYPE html>
<html>
<head>
<base href="https://demos.telerik.com/kendo-ui/calendar/index">
<style>html { font-size: 14px; font-family: Arial, Helvetica, sans-serif; }</style>
<title>Telerik Kendo-UI test</title>
<link rel="stylesheet" href="https://kendo.cdn.telerik.com/2018.1.221/styles/kendo.default-v2.min.css" />
<script src="https://kendo.cdn.telerik.com/2018.1.221/js/jquery.min.js"></script>
<script src="https://kendo.cdn.telerik.com/2018.1.221/js/kendo.all.min.js"></script>
</head>
<body>
<div id="example">
<div class="demo-section k-content" style="text-align: center;">
<h4>Pick a date</h4>
%s
</div>
<script>
$(document).ready(function() {
// create Calendar from div HTML element
%s
});
</script>
</div>
</body>
</html>`, calendar.ToHtml(), calendar.ToJavaScript())

	//Output:
	//<!DOCTYPE html>
	//<html>
	//<head>
	//<base href="https://demos.telerik.com/kendo-ui/calendar/index">
	//<style>html { font-size: 14px; font-family: Arial, Helvetica, sans-serif; }</style>
	//<title>Telerik Kendo-UI test</title>
	//<link rel="stylesheet" href="https://kendo.cdn.telerik.com/2018.1.221/styles/kendo.default-v2.min.css" />
	//<script src="https://kendo.cdn.telerik.com/2018.1.221/js/jquery.min.js"></script>
	//<script src="https://kendo.cdn.telerik.com/2018.1.221/js/kendo.all.min.js"></script>
	//</head>
	//<body>
	//<div id="example">
	//<div class="demo-section k-content" style="text-align: center;">
	//<h4>Pick a date</h4>
	//<div id="calendar"></div>
	//</div>
	//<script>
	//$(document).ready(function() {
	//// create Calendar from div HTML element
	//$("#calendar").kendoCalendar({culture: "en-US",dates: [new Date(2018, 2, 28),new Date(2019, 2, 28, 23, 59),new Date(2019, 2, 28, 23, 59, 59)],depth: "day",disableDates: ["sa","su",],footer: "kendo.template($('#footer-template').html())",format: "yyyy/MM/dd",max: new Date(2019, 2, 28, 23, 59, 59),messages: {weekColumnHeader: "W",},min: new Date(2018, 2, 28),month: { content: "$('#cell-template').html()",weekNumber: "$('#week-template').html()",empty: "-",},selectable: "multiple",selectDates: [new Date(2018, 2, 28),new Date(2019, 2, 28)],weekNumber: true,start: "year",value: new Date(2018, 3, 5, 22, 44),});
	//});
	//</script>
	//</div>
	//</body>
	//</html>
}
