package models

import "fmt"

type BookDetail struct {
	Author string
	Pub    string
	Year   string
	Pages  int
	Prices string
	Scores string
}

func (b BookDetail) String() string {
	return fmt.Sprintf("作者:%s;  出版社:%s;  出版年:%s;  页数:%d;  价格:%s;  评分:%s", b.Author, b.Pub, b.Year, b.Pages, b.Prices, b.Scores)
}
