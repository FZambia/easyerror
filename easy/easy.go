package easy

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mailru/easyjson"
)

type Sample struct {
	Data json.RawMessage `json:"data"`
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func Main(data []byte) {

	var sample Sample
	err := easyjson.Unmarshal(data, &sample)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("success")
}
