package tests

import (
    "fmt"
    "testing"

    "pods/service"
)

func TestFetchProblemByDifficulty(t *testing.T) {
    res := service.FetchProblemByDifficulty("easy")

    fmt.Println(res.Date);
    fmt.Println(res.Title)
    fmt.Println(res.Link)
    fmt.Println(res.Description)
}
