package engine

import (
	fetch "dbbook/fetcher"
	"fmt"
)

type ConcurrentEngine struct {
	Scheduler SimpleScheduler
	WorkCount int
}

func (c *ConcurrentEngine) Run(seed ...Request) {
	in := make(chan Request)
	out := make(chan ParseRequest)
	c.Scheduler.ConfigureChan(in)
	for i := 0; i < c.WorkCount; i++ {
		createWork(in, out)
	}

	for _, r := range seed {
		c.Scheduler.Submit(r)
	}

	for {
		parseR := <-out
		for _, item := range parseR.Items {
			fmt.Println("Got Item: ", item)
		}

		for _, request := range parseR.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

func createWork(in chan Request, out chan ParseRequest) {
	go func() {
		for {
			request := <-in
			parseR, err := work(request)
			if err != nil {
				continue
			}
			// fmt.Println("Fetch URL:", request.Url)
			out <- parseR
		}
	}()
}

func work(r Request) (ParseRequest, error) {
	body, err := fetch.Fetch(r.Url)
	if err != nil {
		return ParseRequest{}, err
	}

	parseR := r.ParseFunc(body)
	return parseR, nil
}
