package main

import (
	"go_projects/go_crawler_in_action/crawler/engine"
	"go_projects/go_crawler_in_action/crawler/scheduler"
	"go_projects/go_crawler_in_action/crawler/zhenai/parser"
	"go_projects/go_crawler_in_action/crawler/persist"
)

func main() {
	// single task
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	// concurrent
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
