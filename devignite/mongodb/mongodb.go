package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"strconv"
	"time"
)

func main() {

}

type MongoStruct struct {
}

type Movie struct {
	Title string `bson:"title"`
	Rated string `bson:"rated"`
}

func (t *MongoStruct) GetMovieCount() int64 {
	db, ctx := new(MongoStruct).GetDatabase()
	coll := db.Collection("movies")
	count, _ := coll.CountDocuments(ctx, bson.D{})

	log.Printf("Movies count: %d", count)
	return count
}

func (t *MongoStruct) GetMovies() *mongo.Cursor {
	db, ctx := new(MongoStruct).GetDatabase()
	coll := db.Collection("movies")
	movies, _ := coll.Find(ctx, Movie{})
	return movies
}

func (t *MongoStruct) GetMoviesBytes() []byte {
	data, _ := json.Marshal(t.GetMovies())
	return data
}

func GetMovieCount(r *restful.Request, resp *restful.Response) {
	var count = MongoStruct{}
	resp.Write([]byte(strconv.Itoa(int(count.GetMovieCount()))))
}

func GetMovies(r *restful.Request, resp *restful.Response) {
	var mongo = MongoStruct{}
	data, _ := json.Marshal(mongo.GetMovies())
	resp.Write(data)
}

func (t *MongoStruct) GetDatabase() (*mongo.Database, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error in client connection: %s", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Connection Error: %s", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databases)

	return client.Database("sample_mflix"), ctx
}
