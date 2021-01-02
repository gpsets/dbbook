package main

import (
	"dbbook/engine"
	"dbbook/parse"
)

func main() {
	seed := engine.Request{
		Url:       "https://book.douban.com/tag/",
		ParseFunc: parse.ParseTag,
	}
	e := &engine.ConcurrentEngine{
		Scheduler: engine.SimpleScheduler{},
		WorkCount: 10,
	}

	e.Run(seed)
}
