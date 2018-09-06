package telerik

import "strconv"

type MapStringArr map[string]interface{}

func (el MapStringArr) String() string {
	var ret string
	var length = len(el) - 1
	var counter = 0

	for k, v := range el {
		counter += 1

		switch v.(type) {
		case string:
			ret += `"` + k + `": "` + v.(string) + `"`
		case int:
			ret += `"` + k + `": "` + strconv.Itoa(v.(int)) + `"`
		case int64:
			ret += `"` + k + `": "` + strconv.FormatInt(v.(int64), 64) + `"`
		case float64:
			ret += `"` + k + `": "` + strconv.FormatFloat(v.(float64), 'E', -1, 64) + `"`
		case float32:
			ret += `"` + k + `": "` + strconv.FormatFloat(float64(v.(float32)), 'E', -1, 32) + `"`
		}

		if length != counter {
			ret += `,`
		}
	}

	return ret
}
