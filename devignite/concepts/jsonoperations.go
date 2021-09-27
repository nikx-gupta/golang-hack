package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Movie struct {
	Title  string
	Genres []string
	Rated  string
	Award  Awards
}

type Awards struct {
	Wins        int
	Nominations int
	Text        string
}

func main() {
	//readJSONFile()
	//writingJSONFile()
	//marshalJson()
	ReadJsonAsMap()
}

func readJSONFile() {
	in, err := os.Open("sample_mflix.json")
	if err != nil {
		fmt.Println("Error in Reading file: ", err)
		return
	}
	defer in.Close()

	jsonDecoder := json.NewDecoder(in)
	record := Movie{}
	err = jsonDecoder.Decode(&record)
	if err != nil {
		fmt.Println("Error in parsing file: ", err)
		return
	}

	fmt.Println(record)
}

func writingJSONFile() {

	record := Movie{
		Title:  "Sample Movie",
		Rated:  "Approved",
		Genres: []string{"Horror", "Action"},
		Award: Awards{
			Text:        "1 win",
			Nominations: 2,
			Wins:        1,
		},
	}

	var encoder = json.NewEncoder(os.Stdout)
	encoder.Encode(record)
}

func marshalJson() {
	record := Movie{
		Title:  "Sample Movie",
		Rated:  "Approved",
		Genres: []string{"Horror", "Action"},
		Award: Awards{
			Text:        "1 win",
			Nominations: 2,
			Wins:        1,
		},
	}

	rec, err := json.Marshal(&record)
	if err != nil {
		fmt.Println("Error in Marhsalling: ", err)
		return
	}

	fmt.Println("Marshaled Json \n", string(rec))

	UnmarshalJson(rec)
}

func UnmarshalJson(input []byte) {
	var rec Movie
	err := json.Unmarshal(input, &rec)
	if err != nil {
		fmt.Println("Error in unmarshal json: ", err)
		return
	}

	fmt.Println(rec)
}

func ReadJsonAsMap(){
	fileData, err := os.ReadFile("sample_mflix.json")
	if err != nil {
		fmt.Println("Error in Reading file: ", err)
		return
	}

	var parsedMap map[string]interface{}
	json.Unmarshal(fileData, &parsedMap)

	fmt.Println("Json Map: ", parsedMap)

}
