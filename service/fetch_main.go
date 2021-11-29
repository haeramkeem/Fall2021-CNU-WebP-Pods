/************************************************************************************
 * Use https://github.com/chromedp/chromedp to render JS                            *
 *                                                                                  *
 * Code Reference: https://github.com/chromedp/examples/blob/master/text/main.go    *
 ************************************************************************************/

package service

import (
	"context"
    "fmt"
    "time"
    "errors"

	"github.com/chromedp/chromedp"
)

func getDataValue() string {
    now := time.Now()
    prefix := now.Format(time.UnixDate)[:10]
    return fmt.Sprintf("a[data-value=\"%s %d 00:00:00 GMT+0900 (Korean Standard Time)\"]", prefix, now.Year())
}

func FetchMain() (string, error) {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
    ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    // Get an url path
    var path string
    var ok bool
	err := chromedp.Run(ctx,
		chromedp.Navigate(BASE + `/problemset/all`),
        chromedp.AttributeValue(getDataValue(), `href`, &path, &ok, chromedp.NodeVisible, chromedp.ByQuery),
	)
    if err != nil { return "", err }
    if !ok { return "", errors.New("Fail to navigate") }

    // Fetch problem content
    var html string
    err = chromedp.Run(ctx,
        chromedp.Navigate(BASE + path),
        chromedp.InnerHTML(`div.content__u3I1`, &html, chromedp.NodeVisible, chromedp.ByQuery),
    )
    if err != nil { return "", err }

    return html, nil
}
