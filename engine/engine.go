package engine

import (
	"fmt"
)

type Engine struct {
	Scheduler
	WorkCount int
	Delay     int
}

func (e *Engine) Run(seeds ...Request) {
	out := e.Scheduler.Run(e.WorkCount, e.Delay)
	for _, req := range seeds {
		e.Scheduler.Submit(req)
	}

	for {
		parseReq := <-out
		for _, item := range parseReq.Items {
			fmt.Println("Got item: ", item)
		}
		for _, req := range parseReq.Requests {
			e.Scheduler.Submit(req)
		}
	}
}
