package main

import (
	"dbbook/engine"
	"dbbook/parse"
	"dbbook/scheduler"
)

func main() {
	seed := engine.Request{
		Url:       "https://book.douban.com/tag/",
		ParseFunc: parse.ParseTag,
	}
	e := &engine.Engine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkCount: 10,
		Delay:     3, // 3s
	}

	e.Run(seed)
}
