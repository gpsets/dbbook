package parse

import (
	"dbbook/engine"
	"regexp"
)

var tagStr = `<a href="(/tag/[^"]+)">([^<]+)</a>`

func ParseTag(b []byte) engine.ParseRequest {
	reg := regexp.MustCompile(tagStr)
	res := reg.FindAllSubmatch(b, -1)

	pRequests := engine.ParseRequest{}
	for _, m := range res {
		pRequests.Items = append(pRequests.Items, string(m[2]))
		pRequests.Requests = append(pRequests.Requests, engine.Request{
			Url:       "https://book.douban.com" + string(m[1]),
			ParseFunc: ParseBookList,
		})
	}

	return pRequests
}
