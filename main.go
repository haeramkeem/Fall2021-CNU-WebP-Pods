/************************************************************************************
 * Use https://github.com/chromedp/chromedp to render JS                            *
 *                                                                                  *
 * Code Reference: https://github.com/chromedp/examples/blob/master/text/main.go    *
 ************************************************************************************/

package main

import (
	"context"
    "fmt"

	"github.com/chromedp/chromedp"
)

func check(err error) {
    if err != err {
        panic(err)
    }
}

// func minify(html string) string {
//     html = strings.TrimSpace(html)
//     acc := make([]rune, 0)
//     for _, rn := range html {
//         if rn != '\t' && rn != '\n' {
//             acc = append(acc, rn)
//         }
//     }
//     return string(acc)
// }

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://leetcode.com/problems/delete-node-in-a-bst/`),
		chromedp.InnerHTML(`.content__u3I1`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	)
    check(err)

    fmt.Println(res)
}
