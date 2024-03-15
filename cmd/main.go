package main

import (
	"go-instrument/pkg/util"
)

func main() {
	err := util.InstrumentFunc("main")
	if err != nil {
		panic(err)
	}

}
