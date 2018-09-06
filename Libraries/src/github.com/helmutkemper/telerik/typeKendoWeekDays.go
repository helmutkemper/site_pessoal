package telerik

type KendoWeekDays []KendoWeekDay

func (el KendoWeekDays) String(prefix, glue, suffix string) string {
	var ret string
	for k, v := range el {
		if k == 0 {
			ret += prefix
		}

		ret += kendoWeekDays[v]

		if k != len(el)-1 {
			ret += glue
		} else {
			ret += suffix
		}
	}

	return ret
}
