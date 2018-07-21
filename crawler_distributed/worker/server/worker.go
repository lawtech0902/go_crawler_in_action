package main

import (
	"go_projects/go_crawler_in_action/crawler_distributed/rpcsupport"
	"fmt"
	"go_projects/go_crawler_in_action/crawler_distributed/worker"
	"log"
	"flag"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must specify a port")
		return
	}

	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlerService{}))
}
