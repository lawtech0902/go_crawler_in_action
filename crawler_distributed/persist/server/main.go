package main

import (
	"go_projects/go_crawler_in_action/crawler_distributed/rpcsupport"
	"go_projects/go_crawler_in_action/crawler_distributed/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"fmt"
	"go_projects/go_crawler_in_action/crawler_distributed/config"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
