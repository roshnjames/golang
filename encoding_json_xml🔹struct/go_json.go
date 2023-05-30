package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type club struct {
	Name   string `json:"Name_of_the_club"` //json to struct conversion field naming
	League string
	Cups   int
}

func main() {
	//-----------struct to JSON-------------------------------------------------------------------------------------
	data := club{Name: "Man City", League: "Premier League", Cups: 5}
	jsd, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	jsn := string(jsd)
	fmt.Println("\n-----------\nStruct to JSON ::::", jsn)

	//------------JSON to struct------------------------------------------------------------------------------------
	var sdata club
	err = json.Unmarshal([]byte(jsn), &sdata) //converts json object "js" into byte slice and stores the same into sdata of struct club type
	//err = json.Unmarshal(jsd, &sdata)
	if err != nil {
		log.Fatal(":::error::::", err)
	}
	fmt.Println("\n-------------------\nJSON to Struct:::::::", sdata)

}
