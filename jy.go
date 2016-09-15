package yj

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"log"
)

func Jy(in []byte) []byte {
	var data map[string]interface{}
	err := json.Unmarshal(in, &data)
	if err != nil {
		log.Fatal("Invalid JSON: ", err)
	}
	out, err := yaml.Marshal(data)
	if err != nil {
		log.Fatal("Error JSON to YAML: ", err)
	}
	return out
}
