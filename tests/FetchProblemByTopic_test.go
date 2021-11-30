package tests

import (
    "fmt"
    "testing"

    "pods/service"
)

func TestFetchProblemByTopic(t *testing.T) {
    res, err := service.FetchProblemByTopic("bfs")
    check(err)
    if res == "" {
        t.Error("Want HTML, got nothing")
    } else {
        fmt.Println("BFS passed")
    }

    res, err = service.FetchProblemByTopic("dfs")
    check(err)
    if res == "" {
        t.Error("Want HTML, got nothing")
    } else {
        fmt.Println("DFS passed")
    }

    res, err = service.FetchProblemByTopic("dp")
    check(err)
    if res == "" {
        t.Error("Want HTML, got nothing")
    } else {
        fmt.Println("DP passed")
    }

    res, err = service.FetchProblemByTopic("tp")
    check(err)
    if res == "" {
        t.Error("Want HTML, got nothing")
    } else {
        fmt.Println("Two Pointer passed")
    }

    res, err = service.FetchProblemByTopic("sort")
    check(err)
    if res == "" {
        t.Error("Want HTML, got nothing")
    } else {
        fmt.Println("Sorting passed")
    }

    res, err = service.FetchProblemByTopic("bsearch")
    check(err)
    if res == "" {
        t.Error("Want HTML, got nothing")
    } else {
        fmt.Println("Binary Search passed")
    }

    res, err = service.FetchProblemByTopic("tree")
    check(err)
    if res == "" {
        t.Error("Want HTML, got nothing")
    } else {
        fmt.Println("Tree passed")
    }

    res, err = service.FetchProblemByTopic("heap")
    check(err)
    if res == "" {
        t.Error("Want HTML, got nothing")
    } else {
        fmt.Println("Heap passed")
    }

    res, err = service.FetchProblemByTopic("string")
    check(err)
    if res == "" {
        t.Error("Want HTML, got nothing")
    } else {
        fmt.Println("String passed")
    }
}
