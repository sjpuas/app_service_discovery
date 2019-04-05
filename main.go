package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	serverAddress = ":3030"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ping")

	db := os.Getenv("DB")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + db + ":27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	name, _ := os.Hostname()
	fmt.Fprintf(w, "Hostname: %s\nDatabase: %s", name, db)
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Listening on port ", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}
