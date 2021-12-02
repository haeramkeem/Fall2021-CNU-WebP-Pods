package service

import (
	"context"
	"time"
    "pods/domain"
    e "pods/modules/errors"

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

func FetchProblemByDifficulty(levelIdx string) (*domain.ProblemDB, error) {
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
    if err := e.Check(err); err != nil {
        return nil, err;
    }

    idx, err := getRand(len(paths))
    if err := e.Check(err); err != nil {
        return nil, err;
    }

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
        if err := e.Check(err); err != nil {
            return nil, err;
        }
    }

    return &domain.ProblemDB{
        Date: nowStrDate(),
        Category: levelIdxToConst[levelIdx],
        Title: title,
        Link: BASE + path,
        Description: html,
    }, nil
}
