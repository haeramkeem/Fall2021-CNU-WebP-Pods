package tests

import (
    "fmt"
    "testing"

    "pods/service"
)

func TestFetchProblemByTopic(t *testing.T) {
    res, err := service.FetchProblemByTopic("0")

    if err != nil {
        t.Error(err)
    }

    fmt.Println(res.Title)
    fmt.Println(res.Link)
    fmt.Println(res.Description)
}
