package tests

import (
    "testing"
    "fmt"

    "pods/service"
)

func TestFetchMainDiscuss(t *testing.T) {
    res, err := service.FetchDiscuss("/problems/accounts-merge/", "")
    check(err)

    if len(res) == 0 {
        t.Error("Want discussion PATHs, got nothing")
    }

    fmt.Println("----- All -----")
    fmt.Println(res)

    res, err = service.FetchDiscuss("/problems/accounts-merge/", "cpp")
    check(err)

    if len(res) == 0 {
        t.Error("Want discussion PATHs, got nothing")
    }

    fmt.Println("----- C++ -----")
    fmt.Println(res)

    res, err = service.FetchDiscuss("/problems/accounts-merge/", "python")
    check(err)

    if len(res) == 0 {
        t.Error("Want discussion PATHs, got nothing")
    }

    fmt.Println("----- Python -----")
    fmt.Println(res)

    res, err = service.FetchDiscuss("/problems/accounts-merge/", "java")
    check(err)

    if len(res) == 0 {
        t.Error("Want discussion PATHs, got nothing")
    }

    fmt.Println("----- Java -----")
    fmt.Println(res)
}
