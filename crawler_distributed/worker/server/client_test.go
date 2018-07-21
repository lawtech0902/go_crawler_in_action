package main

import (
	"testing"
	"go_projects/go_crawler_in_action/crawler_distributed/rpcsupport"
	"go_projects/go_crawler_in_action/crawler_distributed/worker"
	"time"
	"go_projects/go_crawler_in_action/crawler_distributed/config"
	"fmt"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlerService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/107042520",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "笑看人生",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlerServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
