package service

import (
	"context"
	"fmt"
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
    paths := make([]map[string]string, 0)
    topic = urlify[topic]
    fmt.Println(BASE + `/tag/` + topic)
	err := chromedp.Run(ctx,
		chromedp.Navigate(BASE + `/tag/` + topic),
        chromedp.AttributesAll(`div.title-cell__ZGos>a[href]`, &paths, chromedp.NodeVisible, chromedp.ByQueryAll),
	)
    if err != nil { return "", err }
    idx, err := getRand(int64(len(paths)))
    if err != nil { return "", err }
    path := paths[idx]["href"]

    fmt.Println(path)

    // Fetch problem content
    var html string
    err = chromedp.Run(ctx,
        chromedp.Navigate(BASE + path),
        chromedp.InnerHTML(`div.content__u3I1`, &html, chromedp.NodeVisible, chromedp.ByQuery),
    )
    if err != nil { return "", err }

    return html, nil
}
