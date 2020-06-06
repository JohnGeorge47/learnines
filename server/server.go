package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/JohnGeorge47/learnines/server/esqueries"
	"log"
	"net/http"
)

type searchResult struct {
	Data []esqueries.Game `json:"data"`
}

func main() {
	cl := esqueries.EsManager
	ctx := context.Background()
	str, err := cl.PingEs("http://127.0.0.1:9200", ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*str)
	http.HandleFunc("/", testFunc)
	http.HandleFunc("/details", queryFunc)
	fmt.Println("server running on port 9090")
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hey ya"))
}

func queryFunc(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("name")
	es := esqueries.EsManager
	fmt.Println(query)
	ctx := context.Background()
	res, err := es.Search(query, 0, 5, ctx)
	if err != nil {
		fmt.Println(err)
	}
	resJson := *res
	fmt.Println(resJson)
	finalJson := searchResult{Data: resJson}
	tosend, err := json.Marshal(finalJson)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(tosend)
}
