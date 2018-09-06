package javascript

import (
	"errors"
	"fmt"
	"github.com/helmutkemper/go-candyjs"
	"github.com/helmutkemper/mgo"
	"github.com/helmutkemper/mongodb"
	"runtime"
	"time"
)

func RunJavaScriptCode(name, code string) error {
	var err error
	start := time.Now()
	var elapsed time.Duration

	var dbLog = mongodb.DbStt{}
	dbLog.Init("main")
	dbLog.Collection("bigData_javascriptRunLog")
	if dbLog.HasIndex("delete_after_5_days") == false {
		dbLog.IndexMake(mgo.Index{
			Background:  true,
			ExpireAfter: time.Hour * 24 * 5,
			Key: []string{
				"start",
			},
			Name: "delete_after_5_days",
		})
		err = dbLog.GetLastError()
		if err != nil {
			fmt.Printf("runJsCode.dbLog.IndexMake.error: %v\n", err.Error())

			return err
		}
	}

	if name == "" {
		err = errors.New("code name not found")

		dbLog.Insert(map[string]interface{}{"start": start, "end": time.Now(), "elapsed_ms": float64(elapsed / 1000000), "codeName": "", "error": err.Error()})

		fmt.Printf("javascript error: %v\n", err.Error())

		return err
	}

	if code == "" {
		err = errors.New("code not found")

		dbLog.Insert(map[string]interface{}{"start": start, "end": time.Now(), "elapsed_ms": float64(elapsed / 1000000), "codeName": name, "error": err.Error()})

		fmt.Printf("javascript error: %v\n", err.Error())

		return err
	}

	var ctx *candyjs.Context
	ctx = candyjs.NewContext()

	_, err = ctx.PushGlobalGoFunction("dbMongo", func() *mongodb.DbStt { return &mongodb.DbStt{} })
	if err != nil {
		elapsed = time.Since(start)
		dbLog.Insert(map[string]interface{}{"start": start, "end": time.Now(), "elapsed_ms": float64(elapsed / 1000000), "codeName": "push dbMongo{}", "error": err.Error()})

		fmt.Printf("javascript error: %v\n", err.Error())

		return err
	}

	for funcName, funcPointer := range ExternalFunc.Get() {
		_, err = ctx.PushGlobalGoFunction(funcName, funcPointer)
		if err != nil {
			elapsed = time.Since(start)
			dbLog.Insert(map[string]interface{}{"start": start, "end": time.Now(), "elapsed_ms": float64(elapsed / 1000000), "codeName": "push " + funcName + "()", "error": err.Error()})

			fmt.Printf("javascript error: %v\n", err.Error())

			return err
		}
	}

	_, err = ctx.PushGlobalGoFunction("dbQuery", func() *mongodb.QueryStt { return &mongodb.QueryStt{} })
	if err != nil {
		elapsed = time.Since(start)
		dbLog.Insert(map[string]interface{}{"start": start, "end": time.Now(), "elapsed_ms": float64(elapsed / 1000000), "codeName": "push dbQuery{}", "error": err.Error()})

		fmt.Printf("javascript error: %v\n", err.Error())

		return err
	}

	_, err = ctx.PushGlobalStruct("cron", CronSupport)
	if err != nil {
		elapsed = time.Since(start)
		dbLog.Insert(map[string]interface{}{"start": start, "end": time.Now(), "elapsed_ms": float64(elapsed / 1000000), "codeName": "push cron()", "error": err.Error()})

		fmt.Printf("javascript error: %v\n", err.Error())

		return err
	}

	_, err = ctx.PushGlobalGoFunction("cronGetRoute", cronGetRoute)
	if err != nil {
		elapsed = time.Since(start)
		dbLog.Insert(map[string]interface{}{"start": start, "end": time.Now(), "elapsed_ms": float64(elapsed / 1000000), "codeName": "push cronGetRoute()", "error": err.Error()})

		fmt.Printf("javascript error: %v\n", err.Error())

		return err
	}

	_, err = ctx.PushGlobalGoFunction("getRoute", GetRoute)
	if err != nil {
		elapsed = time.Since(start)
		dbLog.Insert(map[string]interface{}{"start": start, "end": time.Now(), "elapsed_ms": float64(elapsed / 1000000), "codeName": "push cronGetRoute()", "error": err.Error()})

		fmt.Printf("javascript error: %v\n", err.Error())

		return err
	}

	_, err = ctx.PushGlobalGoFunction("untreatedData", func() *mongodb.DbStt { return &mongodb.DbStt{} })
	if err != nil {
		elapsed = time.Since(start)
		dbLog.Insert(map[string]interface{}{"start": start, "end": time.Now(), "elapsed_ms": float64(elapsed / 1000000), "codeName": "push untreatedData()", "error": err.Error()})

		fmt.Printf("javascript error: %v\n", err.Error())

		return err
	}

	_, err = ctx.PushGlobalGoFunction("query", func() *mongodb.QueryStt { return &mongodb.QueryStt{} })
	if err != nil {
		elapsed = time.Since(start)
		dbLog.Insert(map[string]interface{}{"start": start, "end": time.Now(), "elapsed_ms": float64(elapsed / 1000000), "codeName": "push query{}", "error": err.Error()})

		fmt.Printf("javascript error: %v\n", err.Error())

		return err
	}

	_, err = ctx.PushGlobalGoFunction("require", ctx.MyRequire)
	if err != nil {
		elapsed = time.Since(start)
		dbLog.Insert(map[string]interface{}{"start": start, "end": time.Now(), "elapsed_ms": float64(elapsed / 1000000), "codeName": "push require()", "error": err.Error()})

		fmt.Printf("javascript error: %v\n", err.Error())

		return err
	}

	err = ctx.PevalString(code)
	if err != nil {
		elapsed = time.Since(start)
		dbLog.Insert(map[string]interface{}{"start": start, "end": time.Now(), "elapsed_ms": float64(elapsed / 1000000), "codeName": name, "error": err.Error()})

		fmt.Printf("javascript error: %v\n", err.Error())

		return err
	}

	ctx.Destroy()

	runtime.GC()

	elapsed = time.Since(start)
	fmt.Printf("[javascript name '%v'] total time %s\n", name, elapsed)

	dbLog.Insert(map[string]interface{}{"start": start, "end": time.Now(), "elapsed_ms": float64(elapsed / 1000000), "codeName": name, "error": ""})

	return nil
}
