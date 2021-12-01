package tests

import (
    "fmt"
    "testing"

    "pods/service"
)

func TestFetchProblemByDifficulty(t *testing.T) {
    res, err := service.FetchProblemByDifficulty("easy")
    check(err)

    if len(res) == 0 {
        t.Error("Want HTML, got nothing")
    }

    fmt.Println("----- EASY -----")
    fmt.Println(res)

    res, err = service.FetchProblemByDifficulty("medium")
    check(err)

    if len(res) == 0 {
        t.Error("Want HTML, got nothing")
    }

    fmt.Println("----- MEDIUM -----")
    fmt.Println(res)

    res, err = service.FetchProblemByDifficulty("hard")
    check(err)

    if len(res) == 0 {
        t.Error("Want HTML, got nothing")
    }

    fmt.Println("----- HARD -----")
    fmt.Println(res)
}
