package tests

import (
    "fmt"
    "testing"

    "pods/service"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func TestFetchMainProblem(t *testing.T) {
    res, err := service.FetchMain()
    check(err)

    if res == "" {
        t.Error("Want HTML, got nothing")
    }

    fmt.Println(res)
}

func TestFetchMainDiscuss(t *testing.T) {
    res, err := service.FetchMainDiscuss("/problems/accounts-merge/", "")
    check(err)

    if len(res) == 0 {
        t.Error("Want discussion PATHs, got nothing")
    }

    fmt.Println("----- All -----")
    fmt.Println(res)

    res, err = service.FetchMainDiscuss("/problems/accounts-merge/", "cpp")
    check(err)

    if len(res) == 0 {
        t.Error("Want discussion PATHs, got nothing")
    }

    fmt.Println("----- C++ -----")
    fmt.Println(res)

    res, err = service.FetchMainDiscuss("/problems/accounts-merge/", "python")
    check(err)

    if len(res) == 0 {
        t.Error("Want discussion PATHs, got nothing")
    }

    fmt.Println("----- Python -----")
    fmt.Println(res)

    res, err = service.FetchMainDiscuss("/problems/accounts-merge/", "java")
    check(err)

    if len(res) == 0 {
        t.Error("Want discussion PATHs, got nothing")
    }

    fmt.Println("----- Java -----")
    fmt.Println(res)
}
