package tests

import (
    "testing"
    "fmt"

    "pods/service"
)

func TestFetchDiscuss(t *testing.T) {
    res, err := service.FetchDiscuss("/problems/accounts-merge", "all")
    if err != nil { t.Error(err) }
    fmt.Println(*res)
}
