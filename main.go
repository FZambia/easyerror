package main

import (
	"io/ioutil"

	"github.com/FZambia/easyerror/easy"
)

func main() {
	data, err := ioutil.ReadFile("./sample.json")
	if err != nil {
		panic(err)
	}
	input := []byte("{\"data\": ")
	input = append(input, data...)
	input = append(input, []byte("}")...)
	easy.Main(input)
}
