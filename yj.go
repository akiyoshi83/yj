package yj

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"log"
)

func Yj(in []byte) []byte {
	var data map[string]interface{}
	err := yaml.Unmarshal(in, &data)
	if err != nil {
		log.Fatal("Invalid YAML: ", err)
	}
	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("Error YAML to JSON: ", err)
	}
	return out
}

func YjMinified(in []byte) []byte {
	var data map[string]interface{}
	err := yaml.Unmarshal(in, &data)
	if err != nil {
		log.Fatal("Invalid YAML: ", err)
	}
	out, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error YAML to JSON: ", err)
	}
	return out
}
