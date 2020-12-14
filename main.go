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
	engine.Run(seed)
}
