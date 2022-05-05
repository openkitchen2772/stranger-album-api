package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"stranger-album-api/controller"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// connect to mongo db, defer a cleanup disconnect()
	database := "albumPhotos"
	connectionString := "mongodb+srv://admin:admin@cluster0.sdez4.mongodb.net/" + database + "?retryWrites=true&w=majority"
	client, ctx := connectDB(connectionString)
	defer disconnectDB(client, ctx)

	// use database, client to create a controller
	cc := controller.NewCommentController(client, database)

	// new a julientschmidt/httpRouter
	r := httprouter.New()

	// setup api path using router instance
	r.GET("/hello", hello)
	r.POST("/newComment", cc.NewComment)
	r.GET("/getComments/:photoId", cc.GetCommentByPhotoId)

	// http.listenAndServe("port", router)
	log.Fatal(http.ListenAndServe(":8080", r))
}

// router api testing hello func
func hello(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(rw, "Hello Moto")
}

// connectDB func
func connectDB(URL string) (*mongo.Client, context.Context) {
	var client *mongo.Client
	var err error

	client, err = mongo.NewClient(options.Client().ApplyURI(URL))
	if err != nil {
		log.Fatal("Mongo fails to create connection client", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		log.Fatal("Mongo fails to connect to server", err)
	}

	return client, ctx
}

// disconnectDB func
func disconnectDB(client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		log.Fatal("Mongo fails to disconnect from Server", err)
	}
}
