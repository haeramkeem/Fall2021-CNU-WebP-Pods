package tests

import (
    "fmt"
    "testing"

    "pods/service"
)

func TestFetchMain(t *testing.T) {
    res, err := service.FetchMain()

    if err != nil {
        t.Fatal(err)
    } else if res == "" {
        t.Error("Want HTML, got nothing")
    }

    fmt.Println(res)
}
