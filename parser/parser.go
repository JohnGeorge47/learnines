package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"io"
	"os"
	"strconv"
)

type Game struct {
	Rank         *int    `json:"rank"`
	Name         *string `json:"name"`
	Platform     *string `json:"platform"`
	Year         *string `json:"year"`
	Genre        *string `json:"genre"`
	Publisher    *string `json:"publisher"`
	NA_Sales     *string `json:"na_sales"`
	EU_Sales     *string `json:"eu_sales"`
	JP_Sales     *string `json:"jp_sales"`
	Other_Sales  *string `json:"other_sales"`
	Global_Sales *string `json:"global_sales"`
}


const mapping =
`{
   "mappings":{
      "game":{
         "properties":{
            "rank":{
               "type":"integer"
            },
            "name":{
               "type":"string"
            },
            "genre":{
               "type":"string"
            },
            "publisher":{
               "type":"string"
            },
			"platform":{
				"type":"string"
			}
         }
      }
   },
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0	
	}
}`

func main() {
	ctx := context.Background()
	client, err := elastic.NewClient()
	info, code, err := client.Ping("http://localhost:9200").Do(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info, code)
	exists, err := client.IndexExists("gameinfo").Do(ctx)
	if err != nil {
		panic("Error while querying" + err.Error())
	}
	if !exists {
		createIndex, err := client.CreateIndex("gameinfo").BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			fmt.Println("Not acknowledged")
		}
	}
	f, err := os.Open("./vgsales.csv")
	if err != nil {
		fmt.Println(err)
	}
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err != nil {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		rank, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println("error while converting string to int")
		} else {
			data := Game{
				Rank:         &rank,
				Name:         &record[1],
				Platform:     &record[2],
				Year:         &record[3],
				Genre:        &record[4],
				Publisher:    &record[5],
				NA_Sales:     &record[6],
				EU_Sales:     &record[7],
				JP_Sales:     &record[8],
				Other_Sales:  &record[9],
				Global_Sales: &record[10],
			}
			fmt.Println(&data.Name)
		}
	}
}


func InsertToEs(ctx context.Context,gameinfo *Game,esclient *elastic.Client)error{
	datajson,err:=json.Marshal(gameinfo)
	if err!=nil{
		return err
	}
	data:=string(datajson)
	idx,err:=esclient.Index().Index("gameinfo").BodyJson(data).Do(ctx)
	if err!=nil{
		return err
	}

}
