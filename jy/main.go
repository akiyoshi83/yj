package main

import (
	"flag"
	"fmt"
	"github.com/akiyoshi83/yj"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var fp *os.File
	var err error

	flag.Parse()

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

	out := yj.Jy(in)

	fmt.Println(string(out))
}
