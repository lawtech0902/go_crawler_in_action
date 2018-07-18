
package persist

import (
	"testing"
	"go_projects/go_crawler_in_action/crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
)

func TestSave(t *testing.T) {
	profile := model.Profile{
		Name:       "惠儿",
		Age:        50,
		Height:     156,
		Weight:     0,
		Income:     "3000元以下",
		Gender:     "女",
		Xinzuo:     "魔羯座",
		Marriage:   "离异",
		Education:  "高中及以下",
		Occupation: "销售总监",
		Hukou:      "四川阿坝",
		House:      "租房",
		Car:        "未购车",
	}

	id, err := save(profile)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)

	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	if actual != profile {
		t.Errorf("got %v; expected %v", actual, profile)
	}
}