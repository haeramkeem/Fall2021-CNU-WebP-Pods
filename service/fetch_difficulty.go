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
		chromedp.Navigate(BASE + `/problemset/all/?difficulty=` + diff + `&page=1`),
        chromedp.WaitVisible(`nav>button.pointer-events-none`, chromedp.NodeVisible, chromedp.ByQuery),
        chromedp.AttributesAll(`div[class="jsx-784799233 overflow-hidden"]:last-child a[href]`, &paths, chromedp.NodeVisible, chromedp.ByQueryAll),
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
