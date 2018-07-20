package client

import (
	"go_projects/go_crawler_in_action/crawler/engine"
	"go_projects/go_crawler_in_action/crawler_distributed/rpcsupport"
	"log"
	"go_projects/go_crawler_in_action/crawler_distributed/config"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+
				"#%d: %v", itemCount, item)
			itemCount++

			// call rpc to save item
			result := ""
			client.Call(config.ItemSaverRpc, item, &result)

			if err != nil {
				log.Printf("Item Saver: error "+
					"saving item %v: %v",
					item, err)
			}
		}
	}()

	return out, nil
}
