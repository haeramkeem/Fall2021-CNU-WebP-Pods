package tests

import (
    "fmt"
    "testing"

    "pods/service"
)

func TestFetchProblemByDifficulty(t *testing.T) {
    res, err := service.FetchProblemByDifficulty("0")

    if err != nil { t.Error(err) }

    fmt.Println(res.Date);
    fmt.Println(res.Title)
    fmt.Println(res.Link)
    fmt.Println(res.Description)
}
