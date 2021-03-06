package esqueries

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Search struct {
	esClient *elastic.Client
}

type Game struct {
	Rank         int    `json:"rank"`
	Name         string `json:"name"`
	Platform     string `json:"platform"`
	Year         string `json:"year"`
	Genre        string `json:"genre"`
	Publisher    string `json:"publisher"`
	NA_Sales     string `json:"na_sales"`
	EU_Sales     string `json:"eu_sales"`
	JP_Sales     string `json:"jp_sales"`
	Other_Sales  string `json:"other_sales"`
	Global_Sales string `json:"global_sales"`
}

func (s Search) PingEs(conn string, ctx context.Context) (*string, error) {
	info, code, err := s.esClient.Ping("http://localhost:9200").Do(ctx)
	if err != nil {
		return nil, err
	}
	str := fmt.Sprintf("Es returned with code %d and info %s", code, info)
	return &str, err
}

func (s Search) Search(tosearch string, start int, end int, ctx context.Context) (*[]Game, error) {
	res := make([]Game, 5)
	query := elastic.NewWildcardQuery("name", tosearch+"*")
	search, err := s.esClient.Search().
		Index("gameinfo").
		Query(query).
		From(start).Size(5).Pretty(true).Do(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("Query took milliseconds", search.TookInMillis)
	for i, hits := range search.Hits.Hits {
		var g Game
		err := json.Unmarshal(hits.Source, &g)
		if err != nil {
			fmt.Println(err)
		}
		res[i] = g
	}
	return &res, nil
}
