package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Exercise struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Bodypart    string `json:"bodypart"`
	Musclegroup string `json:"musclegroup"`
	Content     string `json:"content"`
}

func main() {
	response, err := http.Get("http://api.celdserv.nl/exercise?key=7d480e33ad8c4f1eac6cb6e95c4a4e3e&id=2")

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject []Exercise
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println(responseObject)
}
