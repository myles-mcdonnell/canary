package main

import (
	"canary/env"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var conf *env.Conf

func main() {
	conf = env.Parse()

	http.HandleFunc("/", defaultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	mongoConnStr := fmt.Sprintf("mongodb://%v:%v", conf.MongoHost, conf.MongoPort)
	fmt.Fprintf(w, "Connecting to Mongo %v\r\n", mongoConnStr)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnStr))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	fmt.Fprintln(w, "Mongo Connection OK")
}