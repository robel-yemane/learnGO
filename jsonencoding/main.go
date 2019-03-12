package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	//example of encoding basic data types to JSON
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
	// example of encoding slices and maps to JSON
	slcD := []string{"apply", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	// example of encoding custom data types
	// exported fields used in the encoded output  and by default use those
	// names as the JSON keys
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	// example of tags on struct field declaration to customise
	// the encoded JSON key names
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}

	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// **********************************************

	// example of decoding  JSON data into Go values.

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// we need to provide a variable where the JSON package can put the
	// decoded data. This map[string]interface{} will hold a map
	// of strings to arbitrary data types

	var dat map[string]interface{}

	// decode the data
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// in order to use the values in the decoded map,
	// they'll need to be casted to their appropriate type

	num := dat["num"].(float64)
	fmt.Println(num)
	// accessing nested data requires a series of casts.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// decode JSON into custom data types.
	// This has the advantage of adding additional type-safety to your
	// programs and eliminating the need for type assertions
	// when accessing the decoded data

	str := `{"page":1,"fruits":["apple","peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)

	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// in the above examples bytes and strings were used as
	// intermediates between the data and JSON representation on stdout
	// JSON encodings can also be streamed directly to os.Writers like
	// os.Stdout or even HTTP response bodies

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
