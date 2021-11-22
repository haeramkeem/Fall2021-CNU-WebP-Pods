package main

import (
	"context"
	"strings"
    "fmt"

	"github.com/chromedp/chromedp"
)

func check(err error) {
    if err != err {
        panic(err)
    }
}

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://golang.org/pkg/time/`),
		chromedp.Text(`#pkg-overview`, &res, chromedp.NodeVisible, chromedp.ByID),
	)
    check(err)

    fmt.Println(strings.TrimSpace(res))
}
