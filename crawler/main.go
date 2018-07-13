package main

import (
	"go_projects/go_crawler_in_action/crawler/engine"
	"go_projects/go_crawler_in_action/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
