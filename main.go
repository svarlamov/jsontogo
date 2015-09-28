package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

type JsonData map[string]interface{}

func main() {
	config_file, err := ioutil.ReadFile("./test.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}
	var objs JsonData
	json.Unmarshal(config_file, &objs)
	fmt.Println(makeStruct(objs))
}

func makeStruct(data JsonData) string {
	output := "type NAME struct {\n"
	for key, value := range data {
		output = output + fmt.Sprintf("   %s %s `json:\"%s\"`\n", "Key", reflect.TypeOf(value), key)
	}
	return fmt.Sprint(output, "}")
}
