package javascript

func cronGetRoute(cron, url, method string, header map[string]string) {
	CronSupport.AddFunc(cron, GetRouteCron, cron, url, method, header)
}
