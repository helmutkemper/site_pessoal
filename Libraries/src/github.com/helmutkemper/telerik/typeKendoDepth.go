package telerik

type KendoTimeDepth int

var kendoTimeDepths = [...]string{
	"",
	"day",
	"month",
	"year",
	"decade",
	"century",
}

func (el KendoTimeDepth) String() string {
	return kendoTimeDepths[el]
}

const (
	TIME_DEPTH_DAY KendoTimeDepth = iota + 1

	// Shows the days of the month.
	TIME_DEPTH_MONTH

	// Shows the months of the year.
	TIME_DEPTH_YEAR

	// Shows the years of the decade.
	TIME_DEPTH_DECADE

	// Shows the decades from the century.
	TIME_DEPTH_CENTURY
)
