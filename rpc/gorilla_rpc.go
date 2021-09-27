package main

import (
	jsonparse "encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RpcArgs struct {
	Title string
}

type Movie struct {
	Title string `"json:title,omitempty"`
	Rated string `"json:rated,omitempty"`
}
type JSONServer struct {
}

func (t *JSONServer) MovieDetails(r *http.Request, args *RpcArgs, reply *Movie) error {
	log.Printf("Received Title: %s", args.Title)
	var movies []Movie
	// Read JSON from request
	raw, err := ioutil.ReadFile("./movies.json")
	if err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}

	// Parse Json movies data into array
	marshalmovies := jsonparse.Unmarshal(raw, &movies)
	if marshalmovies != nil {
		log.Println("error:", marshalmovies)
		os.Exit(1)
	}

	// Loop over movies
	for _, movie := range movies {
		if movie.Title == args.Title {
			log.Printf("Matching movie found %s", movie)
			*reply = movie
			break
		}
	}

	return nil
}

func main() {
	// Create a Gorilla RPC server
	s := rpc.NewServer()
	// Register JSON Serializer for incoming requests
	s.RegisterCodec(json.NewCodec(), "application/json")
	// Create a new json server which can serve requests over http
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":9000", r)
}
