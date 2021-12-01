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

	"github.com/chromedp/chromedp"
)

func getDataValue(strDate string) string {
    t := ParseStrDate(strDate)
    prefix := t.Format(time.RubyDate)[:10]
    return fmt.Sprintf("a[data-value=\"%s %d 00:00:00 GMT+0900 (Korean Standard Time)\"]", prefix, t.Year())
}

func FetchMain(strDate string) (string, error) {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
    ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    // Get an url path
    var path string
    for ok := false; !ok; {
        err := chromedp.Run(ctx,
            chromedp.Navigate(BASE + `/problemset/all`),
            chromedp.AttributeValue(getDataValue(strDate), `href`, &path, &ok, chromedp.NodeVisible, chromedp.ByQuery),
        )
        if err != nil { return "", err }
    }

    // Fetch problem content
    var html string
    err := chromedp.Run(ctx,
        chromedp.Navigate(BASE + path),
        chromedp.InnerHTML(`div.content__u3I1`, &html, chromedp.NodeVisible, chromedp.ByQuery),
    )
    if err != nil { return "", err }

    return html, nil
}
