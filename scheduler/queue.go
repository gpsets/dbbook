package scheduler

import (
	"dbbook/engine"
	fetch "dbbook/fetcher"
	"fmt"
	"time"
)

type QueueScheduler struct {
	workChan chan chan engine.Request
	reqChan  chan engine.Request
}

func (q *QueueScheduler) Run(workCount int, delay int) chan engine.ParseRequest {
	out := make(chan engine.ParseRequest)
	for i := 0; i < workCount; i++ {
		q.createWork(delay, out)
	}

	q.workChan = make(chan chan engine.Request)
	q.reqChan = make(chan engine.Request)
	go func() {
		var reqSlice []engine.Request
		var workSlice []chan engine.Request
		for {
			var req engine.Request
			var work chan engine.Request
			if len(reqSlice) > 0 && len(workSlice) > 0 {
				req, work = reqSlice[0], workSlice[0]
			}

			select {
			case r := <-q.reqChan:
				reqSlice = append(reqSlice, r)
			case w := <-q.workChan:
				workSlice = append(workSlice, w)
			case work <- req:
				reqSlice, workSlice = reqSlice[1:], workSlice[1:]
			}
		}
	}()

	return out
}

func (q *QueueScheduler) Submit(r engine.Request) {
	q.reqChan <- r
	return
}

func (q *QueueScheduler) WorkReady(in chan engine.Request) {
	q.workChan <- in
	return
}

func (q *QueueScheduler) createWork(delay int, out chan engine.ParseRequest) {
	go func() {
		in := make(chan engine.Request)
		t := time.NewTicker(time.Duration(delay) * time.Second)
		for {
			q.WorkReady(in)
			req := <-in
			body, err := fetch.Fetch(req.Url)
			if err != nil {
				fmt.Println("Fetch url: ", req.Url)
				continue
			}
			parseReq := req.ParseFunc(body)
			out <- parseReq
			<-t.C
		}
	}()

	return
}
