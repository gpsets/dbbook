package engine

type ParseRequest struct {
	Items    []interface{}
	Requests []Request
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseRequest
}
