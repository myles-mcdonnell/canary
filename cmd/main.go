package main

import (
	"canary/env"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var conf *env.Conf

func main() {
	conf = env.Parse()

	testMongoConnect(os.Stdout)

	http.HandleFunc("/mongoconnect", mongoConnect)
	http.HandleFunc("/ping", ping)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mongoConnect(w http.ResponseWriter, r *http.Request) {
	testMongoConnect(w)
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pong")
}

func testMongoConnect(w io.Writer) {
	mongoConnStr := conf.MongoConnStr()
	fmt.Fprintf(w, "Connecting to Mongo %v\r\n", mongoConnStr)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongoConnStr),
		options.Client().SetConnectTimeout(time.Second*3))
	defer func() {
		if client == nil {
			return
		}
		if err = client.Disconnect(ctx); err != nil {
		}
	}()

	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	fmt.Fprintln(w, "Mongo Connection OK")
}
