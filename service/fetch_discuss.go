package service

import (
	"context"
    "fmt"
    "time"

	"github.com/chromedp/chromedp"
)

func FetchDiscuss(path, query string) ([]string, error) {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
    ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    url := fmt.Sprintf("https://leetcode.com%sdiscuss/?currentPage=1&orderBy=hot&query=%s", path, query)

    // Get all attributes
    res := make([]map[string]string, 0)
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
        chromedp.AttributesAll(`a.title-link__1ay5`, &res, chromedp.NodeVisible, chromedp.ByQueryAll),
	)
    if err != nil { return nil, err }

    ret := make([]string, len(res))
    for idx, el := range res {
        ret[idx] = el["href"]
    }

    return ret, nil
}
