package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
)

type JsonData map[string]interface{}

func main() {
	fileNameIn := flag.String("i", "./test.json", "Path to the JSON input file")
	fileNameOutput := flag.String("o", "./test.go", "Path to the Go output file")
	fileName := flag.String("n", "FooStruct", "Name for the output struct")
	flag.Parse()

	j, err := ioutil.ReadFile(*fileNameIn)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}
	var objs JsonData
	json.Unmarshal(j, &objs)
	output := makeStruct(objs, *fileName)

	f, err := os.Create(*fileNameOutput)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(output)
	if err != nil {
		panic(err)
	}
	cmd := "gofmt"
	args := []string{"-w", *fileNameOutput}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln("There was an error executing `gofmt` on the output file")
		os.Exit(1)
	}
}

func makeStruct(data JsonData, name string) string {
	header := "package main\n"
	output := fmt.Sprintf("%stype %s struct {\n", header, name)
	for key, value := range data {
		output = output + fmt.Sprintf("   %s %s `json:\"%s\"`\n", "Key", reflect.TypeOf(value), key)
	}
	return fmt.Sprint(output, "}")
}
