package parametertransformer

// CloudFormationDetails output from aws cloudformation desxribe-stacks
type CloudFormationDetails struct{
	Stacks []StackType `json:"Stacks"`		
}

// StackType ...
type StackType struct{
	Parameters []ParameterType `json:"Parameters"`
} 

// ParameterType single Tag
type ParameterType struct{
	ParameterKey string `json:"ParameterKey"`
	ParameterValue string `json:"ParameterValue"`
}

