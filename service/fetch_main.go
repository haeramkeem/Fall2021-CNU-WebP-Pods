package service

import (
	"context"
    "fmt"
    "time"
    "pods/domain"
    "pods/modules/errors"

	cdp "github.com/chromedp/chromedp"
)

func getDataValue(strDate string) string {
    t := parseStrDate(strDate)
    prefix := t.Format(time.RubyDate)[:10]
    return fmt.Sprintf("a[data-value=\"%s %d 00:00:00 GMT+0900 (Korean Standard Time)\"]", prefix, t.Year())
}

func FetchMain(strDate string) *domain.ProblemDB {
	// create context
	ctx, cancel := cdp.NewContext(context.Background())
	defer cancel()
    ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    // Get an url path
    path := ""
    for ok := false; !ok && path == ""; {
        err := cdp.Run(ctx,
            cdp.Navigate(BASE + `/problemset/all`),
            cdp.AttributeValue(getDataValue(strDate), `href`, &path, &ok, cdp.NodeVisible, cdp.ByQuery),
        )
        errors.Check(err)
    }

    // Fetch problem content
    html := ""
    title := ""
    for html == "" || title == "" {
        err := cdp.Run(ctx,
            cdp.Navigate(BASE + path),
            cdp.InnerHTML(`div.content__u3I1`, &html, cdp.NodeVisible, cdp.ByQuery),
            cdp.Text(`div[data-cy="question-title"]`, &title, cdp.NodeVisible, cdp.ByQuery),
        )
        errors.Check(err)
    }

    return &domain.ProblemDB{
        Date: strDate,
        Category: MAIN,
        Title: title,
        Link: BASE + path,
        Description: html,
    }
}
