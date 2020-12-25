package parse

import (
	"dbbook/engine"
	"dbbook/models"
	"regexp"
	"strconv"
)

const authorReg = `<span class="pl"> 作者</span>:\s*<a class="" href="[^"]+">([^<]+)</a>\s*</span><br/>`
const pubReg = `<span class="pl">出版社:</span>([^<]+)<br/>`
const yearReg = `<span class="pl">出版年:</span>([^<]+)<br/>`
const pagesReg = `<span class="pl">页数:</span>([^<]+)<br/>`
const pricesReg = `<span class="pl">定价:</span>([^<]+)<br/>`
const scoresReg = `<strong class="ll rating_num " property="v:average">([^<]+)</strong>`

func ParseBookDetail(content []byte) engine.ParseRequest {
	detail := models.BookDetail{}
	detail.Author = matchDetail(content, authorReg)
	detail.Prices = matchDetail(content, pricesReg)
	detail.Pub = matchDetail(content, pubReg)
	detail.Scores = matchDetail(content, scoresReg)
	detail.Year = matchDetail(content, yearReg)

	var pstr = matchDetail(content, pagesReg)
	pages, err := strconv.Atoi(pstr)
	if err != nil {
		detail.Pages = pages
	}

	pRequest := engine.ParseRequest{}
	pRequest.Items = append(pRequest.Items, detail)
	return pRequest
}

func matchDetail(content []byte, regStr string) string {
	reg := regexp.MustCompile(regStr)
	res := reg.FindSubmatch(content)
	if len(res) <= 0 || len(res[0]) <= 2 {
		return ""
	}
	return string(res[1])
}
