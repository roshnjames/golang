package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type club struct {
	Name   string `xml:"Name_of_the_club"` //xml to struct conversion field naming
	League string
	Cups   int
}

func main() {
	//-----------struct to xml-------------------------------------------------------------------------------------
	data := club{Name: "Man City", League: "Premier League", Cups: 5}
	jsd, err := xml.MarshalIndent(data, "||| ", "-- ")
	if err != nil {
		panic(err)
	}
	jsn := string(jsd)
	fmt.Println("\n-----------\nStruct to xml ::::\n", jsn)

	//------------xml to struct------------------------------------------------------------------------------------
	var sdata club
	err = xml.Unmarshal([]byte(jsn), &sdata) //converts xml object "js" into byte slice and stores the same into sdata of struct club type
	//err = xml.Unmarshal(jsd, &sdata)
	if err != nil {
		log.Fatal(":::error::::", err)
	}
	fmt.Println("\n-------------------\nxml to Struct:::::::", sdata)

}
