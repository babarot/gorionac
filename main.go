package main

import (
	"fmt"
	"log"

	"github.com/cheggaaa/pb"
)

type failed []string

func main() {
	y, err := readYaml()
	if err != nil {
		log.Fatal(err)
	}

	var f failed = y.Start()
	if ok := f.Check(); !ok {
		fmt.Println("miss:", f)
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
	//bar.FinishPrint("")
	bar.Finish()

	return failed
}

func (f failed) Check() bool {
	return len(f) == 0
}
