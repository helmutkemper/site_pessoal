package telerik

type KendoWeekDay int

var kendoWeekDays = [...]string{
	"",
	"su",
	"mo",
	"tu",
	"we",
	"th",
	"fr",
	"sa",
}

func (el KendoWeekDay) String() string {
	return kendoWeekDays[el]
}

const (
	WEEK_DAY_SUNDAY KendoWeekDay = iota + 1
	WEEK_DAY_MONDAY
	WEEK_DAY_TUESDAY
	WEEK_DAY_WEDNESDAY
	WEEK_DAY_THURSDAY
	WEEK_DAY_FRIDAY
	WEEK_DAY_SATURDAY
)
