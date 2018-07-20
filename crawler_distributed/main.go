package main

import (
	"go_projects/go_crawler_in_action/crawler/engine"
	"go_projects/go_crawler_in_action/crawler/scheduler"
	"go_projects/go_crawler_in_action/crawler/zhenai/parser"
	"go_projects/go_crawler_in_action/crawler_distributed/persist/client"
	"fmt"
	"go_projects/go_crawler_in_action/crawler_distributed/config"
)

func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
