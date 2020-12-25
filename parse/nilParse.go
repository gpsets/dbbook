package parse

import "dbbook/engine"

func NilParse(content []byte) engine.ParseRequest {
	return engine.ParseRequest{}
}
