package javascript

import (
	"github.com/helmutkemper/mongodb"
	"time"
)

func GetRouteCron(cron, url, method string, header map[string]string) {

	resp := GetRoute(url, method, header)

	var dbLog = mongodb.DbStt{}
	dbLog.Init("main")
	dbLog.Collection("bigData_cronLog")
	dbLog.Insert(map[string]interface{}{
		"resp":   resp,
		"cron":   cron,
		"url":    url,
		"method": method,
		"header": header,
		"time":   time.Now(),
	})
}
