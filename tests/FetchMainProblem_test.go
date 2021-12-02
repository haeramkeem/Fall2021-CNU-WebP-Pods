package tests

import (
    "fmt"
    "testing"

    "pods/service"
)

func TestFetchMainProblem(t *testing.T) {
    res, err := service.FetchMain("211201")

    if err != nil { t.Error(err) }

    fmt.Println(res.Title)
    fmt.Println(res.Link)
    fmt.Println(res.Description)
}
