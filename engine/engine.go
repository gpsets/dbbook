package engine

import (
	fetch "dbbook/fetcher"
	"fmt"
)

func Run(seed ...Request) {
	queue := []Request{}
	for _, r := range seed {
		queue = append(queue, r)
	}

	for len(queue) > 0 {
		task := queue[0]
		queue = queue[1:]
		resp, err := fetch.Fetch(task.Url)
		if err != nil {
			continue
		}

		pRequest := task.ParseFunc(resp)
		for _, item := range pRequest.Items {
			fmt.Printf("Got item: %v", item)
		}

		queue = append(queue, pRequest.Requests...)
	}
}
