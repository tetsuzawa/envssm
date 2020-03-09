package internal

import "bytes"

type Parameter struct {
	envMap    map[string]string
	SSMBuf    *bytes.Buffer
	ValBuf    *bytes.Buffer
	TFVarsBuf *bytes.Buffer
}

func NewParameter(envMap map[string]string) *Parameter {
	return &Parameter{envMap: envMap, SSMBuf: new(bytes.Buffer), ValBuf: new(bytes.Buffer), TFVarsBuf: new(bytes.Buffer)}
}
