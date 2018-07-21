package main

import (
	"go_projects/go_crawler_in_action/crawler/engine"
	"go_projects/go_crawler_in_action/crawler/scheduler"
	"go_projects/go_crawler_in_action/crawler/zhenai/parser"
	itemsaver "go_projects/go_crawler_in_action/crawler_distributed/persist/client"
	"fmt"
	"go_projects/go_crawler_in_action/crawler_distributed/config"
	worker "go_projects/go_crawler_in_action/crawler_distributed/worker/client"
	"net/rpc"
	"go_projects/go_crawler_in_action/crawler_distributed/rpcsupport"
	"log"
	"errors"
	"flag"
	"strings"
)

var (
	itemSaverHost = flag.Int("itemsaver_host", 0, "itemsaver host")

	workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()

	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", *itemSaverHost))
	if err != nil {
		panic(err)
	}

	pool, err := createClientPool(strings.Split(*workerHosts, ","))
	if err != nil {
		panic(err)
	}

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) (chan *rpc.Client, error) {
	var clients []*rpc.Client

	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	if len(clients) == 0 {
		return nil, errors.New("no connections available")
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out, nil
}
