package hmackeyfinder

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type config struct {
	MaxGoroutines int
	SearchInOrder bool
	Hashcode      string
	Message       string
}

//Config for this shit
var Config config

func init() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println(err)
		initDefault()
	} else {
		initFromFile(file)
	}
	log.Println("Config:", Config)
}

func initDefault() {
	Config = config{
		MaxGoroutines: 8,
		SearchInOrder: false,
		Hashcode:      "f3c2ae334dc98a387601c85ef83c77360943023a",
		Message:       "341567891 487654 500",
	}
	log.Println("Using default configuration:", Config)
}

func initFromFile(file []byte) {
	err := json.Unmarshal([]byte(file), &Config)
	if err != nil {
		log.Println("Found error while parsing config file:", err)
		initDefault()
	}
}
