package telerik

import "fmt"

func ExampleKendoWeekDay_String() {
	var wd KendoWeekDay

	wd = WEEK_DAY_SUNDAY
	fmt.Printf("week day: %v\n", wd.String())

	wd = WEEK_DAY_SATURDAY
	fmt.Printf("week day: %v\n", wd.String())

	wd = WEEK_DAY_FRIDAY
	fmt.Printf("week day: %v\n", wd.String())

	wd = WEEK_DAY_MONDAY
	fmt.Printf("week day: %v\n", wd.String())

	wd = WEEK_DAY_THURSDAY
	fmt.Printf("week day: %v\n", wd.String())

	wd = WEEK_DAY_TUESDAY
	fmt.Printf("week day: %v\n", wd.String())

	wd = WEEK_DAY_WEDNESDAY
	fmt.Printf("week day: %v\n", wd.String())

	// Output:
	// week day: su
	// week day: sa
	// week day: fr
	// week day: mo
	// week day: th
	// week day: tu
	// week day: we
}
