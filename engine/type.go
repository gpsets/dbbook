package engine

type ParseRequest struct {
	Items    []interface{}
	Requests []Request
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseRequest
}

type Scheduler interface {
	Submit(Request)
	WorkReady(chan Request)
	// WorkChan() chan Request
	Run(int, int) chan ParseRequest
}
