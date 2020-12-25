package parse

import (
	"dbbook/engine"
	"regexp"
)

const listReg = `<a href="([^"]+)" title="([^"]+)" [^<]+</a>`

func ParseBookList(content []byte) engine.ParseRequest {
	reg := regexp.MustCompile(listReg)
	res := reg.FindAllSubmatch(content, -1)

	pRequest := engine.ParseRequest{}
	for _, m := range res {
		pRequest.Items = append(pRequest.Items, string(m[2]))
		pRequest.Requests = append(pRequest.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: NilParse,
		})
	}
	return pRequest
}
