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

	var minify *bool

	minify = flag.Bool("minify", false, "Minify JSON")
	minify = flag.Bool("m", false, "Minify JSON")

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

	var out []byte
	if *minify {
		out = yj.YjMinified(in)
	} else {
		out = yj.Yj(in)
	}

	fmt.Println(string(out))
}
