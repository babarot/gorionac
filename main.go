package main

import (
	"flag"
	"log"
	"os"

	"github.com/cheggaaa/pb"
	"github.com/k0kubun/pp"
)

type failed []string

var update = flag.Bool("u", false, "Update")

func usage() {
}

func main() {
	y, err := readYaml()
	if err != nil {
		log.Fatal(err)
	}

	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		usage()
	}

	var f failed = y.Start()

	if ok := f.Check(); !ok {
		pp.Printf("failed package(s): \n%s\n", f)
		os.Exit(1)
	}

}

func (y *Yaml) Start() (failed []string) {
	bar := pb.StartNew(len(y.Package))
	bar.SetWidth(80)
	for _, p := range y.Package {
		if err := p.Get(); err != nil {
			failed = append(failed, p.Name)
		}
		bar.Increment()
	}
	bar.Finish()

	return failed
}

func (f failed) Check() bool {
	return len(f) == 0
}
