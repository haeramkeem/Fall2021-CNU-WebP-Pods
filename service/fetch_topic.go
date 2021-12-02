package service

import (
	"context"
	"time"
    "pods/domain"
    e "pods/modules/errors"

	cdp "github.com/chromedp/chromedp"
)

var topicIdxToPath = map[string]string{
    "0":    "breadth-first-search/",
    "1":    "depth-first-search/",
    "2":    "dynamic-programming/",
    "3":    "two-pointers/",
    "4":    "sorting/",
    "5":    "binary-search/",
    "6":    "tree/",
    "7":    "heap-priority-queue/",
    "8":    "string/",
}

var topicIdxToConst = map[string]uint{
    "0":    BFS,
    "1":    DFS,
    "2":    DP,
    "3":    TWO_POINTER,
    "4":    SORT,
    "5":    BINARY_SEARCH,
    "6":    TREE,
    "7":    HEAP,
    "8":    STRING,
}

func FetchProblemByTopic(topicIdx string) (*domain.ProblemDB, error) {
	// create context
	ctx, cancel := cdp.NewContext(context.Background())
	defer cancel()
    ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    // Get an url path
    nodes := make([]map[string]string, 0)
    tag := topicIdxToPath[topicIdx]

	err := cdp.Run(ctx,
		cdp.Navigate(BASE + `/tag/` + tag),
        cdp.AttributesAll(`div.title-cell__ZGos>*:last-child`, &nodes, cdp.NodeVisible, cdp.ByQueryAll),
	)
    if err := e.Check(err); err != nil {
        return nil, err
    }

    path := ""
    for exs := false; !exs || path == ""; {
        idx, err := getRand(len(nodes))
        if err := e.Check(err); err != nil {
            return nil, err
        }
        path, exs = nodes[idx]["href"]
    }

    // Fetch problem content
    html := ""
    title := ""
    for html == "" || title == "" {
        err = cdp.Run(ctx,
            cdp.Navigate(BASE + path),
            cdp.InnerHTML(`div.content__u3I1`, &html, cdp.NodeVisible, cdp.ByQuery),
            cdp.Text(`div[data-cy="question-title"]`, &title, cdp.NodeVisible, cdp.ByQuery),
        )
        if err := e.Check(err); err != nil {
            return nil, err
        }
    }

    return &domain.ProblemDB{
        Date: nowStrDate(),
        Category: topicIdxToConst[topicIdx],
        Title: title,
        Link: BASE + path,
        Description: html,
    }, nil
}
