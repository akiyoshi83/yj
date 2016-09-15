package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var fp *os.File
	var err error

	yj := flag.Bool("yj", false, "YAML to JSON")
	jy := flag.Bool("jy", false, "JSON to YAML")
	flag.Parse()

	if *yj && *jy {
		log.Fatal("You mean yj or jy?")
	}

	if !*yj && !*jy {
		log.Fatal("You need yj or jy.")
	}

	if flag.NArg() < 1 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()
	}

	in, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}

	var out []byte
	if *jy {
		out = j2y(in)
	} else {
		out = y2j(in)
	}

	fmt.Println(string(out))
}

func y2j(in []byte) []byte {
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

func j2y(in []byte) []byte {
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
