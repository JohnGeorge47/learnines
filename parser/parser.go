package main

import (
	"fmt"
	"encoding/csv"
	"io"
	"os"
	elastic "github.com/olivere/elastic/v7"
	"context"
)

func main(){
	ctx := context.Background()
	client,err:=elastic.NewClient()
	info,code,err:=client.Ping("localhost:9200").Do(ctx)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(info,code)
	f,err:=os.Open("../vgsales.csv")
	if err!=nil{
		fmt.Println(err)
	}
	r:=csv.NewReader(f)
	for  {
		record,err:=r.Read()
		if err!=nil{
			panic(err)
		}
		if err==io.EOF{
			break
		}
		fmt.Println(record)
	}
}
