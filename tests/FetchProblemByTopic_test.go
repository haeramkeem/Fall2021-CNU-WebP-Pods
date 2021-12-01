package tests

import (
    "fmt"
    "testing"

    "pods/service"
)

func TestFetchProblemByTopic(t *testing.T) {
    res := service.FetchProblemByTopic("0")

    fmt.Println(res.Title)
    fmt.Println(res.Link)
    fmt.Println(res.Description)
}
