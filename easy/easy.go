package easy

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mailru/easyjson"
)

type clientCommand struct {
	UID    string          `json:"uid"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

var jsonObject string = `{"method":"connect","params":{"user":"2694","info":"{\"first_name\": \"Alexandr\", \"last_name\": \"Emelin\"}","timestamp":"1459611382","token":"99e00f13188f4397f6ecca34d0f1f174b821d3338241b48dc2b6b82802fde116"},"uid":"1"}`

func init() {
	log.SetFlags(log.Lshortfile)
}

func Main() {
	jsonObjectBytes := []byte(jsonObject)

	var command clientCommand
	err := easyjson.Unmarshal(jsonObjectBytes, &command)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%#v", command)
}
