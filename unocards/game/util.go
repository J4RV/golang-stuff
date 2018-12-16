package main

import (
	"encoding/json"
	"fmt"
)

//from https://stackoverflow.com/questions/24512112/how-to-print-struct-variables-in-console
func prettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}
