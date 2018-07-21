package config

const (
	// es
	ElasticIndex = "dating_profile"

	// RPC endpoints
	ItemSaverRpc      = "ItemSaverService.Save"
	CrawlerServiceRpc = "CrawlerService.Process"

	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	// Rate limit
	QPS = 20
)
