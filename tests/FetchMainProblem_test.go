package tests

import (
    "fmt"
    "testing"

    "pods/service"
)

func TestFetchMainProblem(t *testing.T) {
    res, err := service.FetchMain("211201")
    check(err)

    if res == "" {
        t.Error("Want HTML, got nothing")
    }

    fmt.Println(res)
}
