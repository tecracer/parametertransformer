package parametertransformer_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"tecracer/parametertransformer"
)

const infile = "input.json"




func TestReadParameter( t *testing.T){
	 assert.NotPanics(t,func() { parametertransformer.ReadParameter(parametertransformer.Read(infile))}, "the code should  not panic")
}