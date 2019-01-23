package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//A type capable of holding the JSON values
type configuration struct {
	Enabled bool
	Path    string
}

func main() {

	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := &configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)
}
