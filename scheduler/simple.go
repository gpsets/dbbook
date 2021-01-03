package scheduler

import (
	"dbbook/engine"
	fetch "dbbook/fetcher"
	"log"
	"time"
)

type SimpleScheduler struct {
	reqChan chan engine.Request
}

func (s *SimpleScheduler) Run(workCount int, delay int) chan engine.ParseRequest {
	s.reqChan = make(chan engine.Request)
	out := make(chan engine.ParseRequest)
	for i := 0; i < workCount; i++ {
		s.createWork(delay, out)
	}
	return out
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.reqChan <- r
	}()
	return
}

func (s *SimpleScheduler) WorkReady(chan engine.Request) {
	return
}

func (s *SimpleScheduler) createWork(delay int, out chan engine.ParseRequest) {
	go func() {
		t := time.NewTicker(time.Duration(delay) * time.Second)
		for {
			r := <-s.reqChan
			body, err := fetch.Fetch(r.Url)
			if err != nil {
				log.Println("Fetch url: ", r.Url)
				continue
			}

			parseReq := r.ParseFunc(body)
			out <- parseReq

			<-t.C
		}
	}()
}
