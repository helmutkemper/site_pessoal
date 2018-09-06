package javascript

import "github.com/helmutkemper/cron"

var CronSupport *cron.Cron

func init() {
	CronSupport = cron.New()
}
