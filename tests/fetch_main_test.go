package tests

import (
    "fmt"
    "testing"

    "pods/service"
)

func TestFetchMainProblem(t *testing.T) {
    res, err := service.FetchMain()
    check(err)

    if res == "" {
        t.Error("Want HTML, got nothing")
    }

    fmt.Println(res)
}
