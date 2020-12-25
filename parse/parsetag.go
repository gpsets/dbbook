package parse

import (
	"dbbook/engine"
	"regexp"
)

func ParseTag(b []byte) engine.ParseRequest {
	expstr := `<a href="(/tag/[^"]+)">([^<]+)</a>`
	reg := regexp.MustCompile(expstr)
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
