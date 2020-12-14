package engine

type ParseRequest struct {
	Items    []string
	Requests []Request
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseRequest
}

func NilParse(b []byte) ParseRequest {
	return ParseRequest{}
}
