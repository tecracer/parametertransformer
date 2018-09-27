package main

import (
	"tecracer/parametertransformer"
	"flag"
	//"os"
)

func main(){
	var inputFileName = flag.String("filename","input.json","json output from cloudformation")
	flag.Parse()
	if inputFileName == nil {
		println("Usage")
		flag.PrintDefaults()
		println("---")
		panic("No file input given")
	}
	// if _, err := os.Stat(*inputFileName); os.IsNotExist(err) {
	// 	// path/to/whatever does not exist
	//   }
	parametertransformer.ReadParameter(parametertransformer.Read(*inputFileName))
}

