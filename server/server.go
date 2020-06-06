package main

import (
	"context"
	"fmt"
	"github.com/JohnGeorge47/learnines/server/esqueries"
	"log"
	"net/http"
)

func main() {
	cl := esqueries.EsManager
	ctx := context.Background()
	str, err := cl.PingEs("http://127.0.0.1:9200", ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*str)
	http.HandleFunc("/", testFunc)
	http.HandleFunc("/name", queryFunc)
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
	fmt.Println(query)
}
