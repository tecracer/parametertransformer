package parametertransformer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)
// ReadParameter ...
func ReadParameter( input CloudFormationDetails)(error){
	sep := "\""
	kvsep := " : "
	for _, tag := range input.Stacks[0].Parameters {
		fmt.Println(sep+tag.ParameterKey+sep+kvsep+sep+tag.ParameterValue+sep)
	}
	return nil
}

// Read a json file
func Read(infile string) (CloudFormationDetails) {
	testFile, err := os.Open(infile)
	if err != nil {
		log.Println("opening config file: "+ err.Error())
	}

	var settings *CloudFormationDetails
	jsonParser := json.NewDecoder(testFile)
	err = jsonParser.Decode(&settings)
	if err != nil {
		log.Println("parsing config file:"+ err.Error())
	}
	return *settings
}
