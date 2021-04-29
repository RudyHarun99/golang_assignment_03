package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
	"encoding/json"
	"fmt"
	"os"
)

// create struct status
type Data struct {
	Status Element	`json:"status"`
}

type Element struct {
	Water	int	`json:"water"`
	Wind	int	`json:"wind"`
}

func init() {
	go AutoReloadJson()
}

func main() {
	http.HandleFunc("/", AutoReloadWeb)
	http.ListenAndServe(":8080", nil)
}

func AutoReloadJson() {
	for {
		// read file json
		// defer file close
		// parse json to struct
		// generate 2 random int
		// update json
		// file write datajson
		// file sync
		file, _ := os.Open("data.json")
		defer file.Close()

		jsonData, _ := ioutil.ReadAll(file)

		var data Data

		_ = json.Unmarshal(jsonData, &data)

		// fmt.Println(data)
		water := RandomNumber()
		wind := RandomNumber()
		newElement := Element{
			water,
			wind,
		}
		newData := Data{newElement}
		// fmt.Println(newData)
		newJson, _ := json.Marshal(newData)
		// fmt.Println(newJson)
		fmt.Println(string(newJson))

		// file.Write(newJson)
		// file.Sync()
		_ = ioutil.WriteFile("data.json", newJson, 0644)
		
		time.Sleep(time.Second * 5)
	}
}

func RandomNumber() int {
	random := rand.Intn(20)
	return random
}

func AutoReloadWeb(w http.ResponseWriter, r *http.Request) {
	// read file json
	// defer file Close
	// logic check status
	// response html to client
	// log html utk website auto reload
	file, _ := os.Open("data.json")
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var data Data

	_ = json.Unmarshal(jsonData, &data)

	fmt.Println(data.Status.Water)
}