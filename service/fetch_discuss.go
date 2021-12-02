package service

import (
	"context"
	"fmt"
	"strings"
	"time"
	. "pods/domain"
    "pods/modules/errors"

	"github.com/chromedp/chromedp"
)

func hrefToTitle(href string) string {
    part := strings.Split(href, "/")
    last := part[len(part) - 1]
    return strings.Join(strings.Split(last, "-"), " ")
}

func FetchDiscuss(probpath, query string) (*DiscussionJSON, error) {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
    ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    if query == "all" { query = "" }
    url := fmt.Sprintf("%s%s/discuss/?currentPage=1&orderBy=hot&query=%s", BASE, probpath, query)

    // Get all attributes
    res := make([]map[string]string, 0)
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
        chromedp.AttributesAll(`a.title-link__1ay5`, &res, chromedp.NodeVisible, chromedp.ByQueryAll),
	)
    if err := errors.Check(err); err != nil {
        return nil, err
    }

    // Set alternative link
    alt := ""
    if len(res) == 0 {
        alt = url
    }

    // Set []DiscussionEntryJSON
    entries := make([]DiscussionEntryJSON, 0)
    for _, el := range res {
        href, exs := el["href"]
        if !exs { continue }
        entry := DiscussionEntryJSON{
            Title: hrefToTitle(href),
            Link: href,
        }
        entries = append(entries, entry)
    }

    return &DiscussionJSON{
        Alt: alt,
        Links: entries,
    }, nil
}
