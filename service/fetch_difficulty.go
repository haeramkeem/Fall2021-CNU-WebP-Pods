package service

import (
	"context"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func FetchProblemByDifficulty(diff string) (string, error) {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
    ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    // Get an url path
    paths := make([]map[string]string, 0)
    diff = strings.ToUpper(diff)
	err := chromedp.Run(ctx,
        // GOTO problems by difficulty page
		chromedp.Navigate(BASE + `/problemset/all/?difficulty=` + diff + `&page=1`),
        // Waiting for all nodes to visible
        chromedp.WaitVisible(`nav>button.pointer-events-none`, chromedp.NodeVisible, chromedp.ByQuery),
        // Get all hyperlinks
        chromedp.AttributesAll(`button.jsx-784799233>div.truncate>a[href]`, &paths, chromedp.NodeVisible, chromedp.ByQueryAll),
	)
    if err != nil { return "", err }
    idx, err := getRand(len(paths))
    if err != nil { return "", err }
    path := paths[idx]["href"]

    // Fetch problem content
    var html string
    err = chromedp.Run(ctx,
        chromedp.Navigate(BASE + path),
        chromedp.InnerHTML(`div.content__u3I1`, &html, chromedp.NodeVisible, chromedp.ByQuery),
    )
    if err != nil { return "", err }

    return html, nil
}
