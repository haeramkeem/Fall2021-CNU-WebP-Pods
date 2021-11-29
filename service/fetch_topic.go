package service

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

var urlify = map[string]string{
    "bfs":      "breadth-first-search/",
    "dfs":      "depth-first-search/",
    "dp":       "dynamic-programming/",
    "tp":       "two-pointers/",
    "sort":     "sorting/",
    "bsearch":  "binary-search/",
    "tree":     "tree/",
    "heap":     "heap-priority-queue/",
    "string":   "string/",
}

func FetchProblemByTopic(topic string) (string, error) {
    const BASE string = `https://leetcode.com`;

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
    ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    // Get an url path
    nodes := make([]map[string]string, 0)
    topic = urlify[topic]

	err := chromedp.Run(ctx,
		chromedp.Navigate(BASE + `/tag/` + topic),
        chromedp.AttributesAll(`div.title-cell__ZGos>*:last-child`, &nodes, chromedp.NodeVisible, chromedp.ByQueryAll),
	)
    if err != nil { return "", err }

    var path string
    exs := false
    for !exs {
        idx, err := getRand(len(nodes))
        if err != nil { return "", err }
        path, exs = nodes[idx]["href"]
    }

    // Fetch problem content
    var html string
    err = chromedp.Run(ctx,
        chromedp.Navigate(BASE + path),
        chromedp.InnerHTML(`div.content__u3I1`, &html, chromedp.NodeVisible, chromedp.ByQuery),
    )
    if err != nil { return "", err }

    return html, nil
}
