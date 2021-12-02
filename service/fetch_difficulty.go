package service

import (
	"context"
	"time"
    "pods/domain"
    "pods/modules/errors"

	cdp "github.com/chromedp/chromedp"
)

var levelIdxToDifficulty = map[string]string{
    "0":    "EASY",
    "1":    "MEDIUM",
    "2":    "HARD",
}

var levelIdxToConst = map[string]uint{
    "0":    EASY,
    "1":    MEDIUM,
    "2":    HARD,
}

func FetchProblemByDifficulty(levelIdx string) *domain.ProblemDB {
	// create context
	ctx, cancel := cdp.NewContext(context.Background())
	defer cancel()

    ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    // Get an url path
    paths := make([]map[string]string, 0)
    level := levelIdxToDifficulty[levelIdx]

	err := cdp.Run(ctx,
		cdp.Navigate(BASE + `/problemset/all/?difficulty=` + level + `&page=1`),
        cdp.WaitVisible(`nav>button.pointer-events-none`, cdp.NodeVisible, cdp.ByQuery),
        cdp.AttributesAll(`div[class="jsx-784799233 overflow-hidden"]:last-child a[href]`, &paths, cdp.NodeVisible, cdp.ByQueryAll),
	)
    errors.Check(err)

    idx, err := getRand(len(paths))
    errors.Check(err)

    path := paths[idx]["href"]

    // Fetch problem content
    html := ""
    title := ""
    for html == "" || title == "" {
        err = cdp.Run(ctx,
            cdp.Navigate(BASE + path),
            cdp.InnerHTML(`div.content__u3I1`, &html, cdp.NodeVisible, cdp.ByQuery),
            cdp.Text(`div[data-cy="question-title"]`, &title, cdp.NodeVisible, cdp.ByQuery),
        )
        errors.Check(err)
    }

    return &domain.ProblemDB{
        Date: nowStrDate(),
        Category: levelIdxToConst[levelIdx],
        Title: title,
        Link: BASE + path,
        Description: html,
    }
}
