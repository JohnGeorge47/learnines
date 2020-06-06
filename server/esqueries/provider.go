package esqueries

import (
	"github.com/olivere/elastic/v7"
	"log"
)

var EsManager Iquery

func init() {
	client, err := elastic.NewClient()
	if err != nil {
		log.Fatal("Failed to connect to elastic searh")
	}
	EsManager = Search{esClient: client}
}
