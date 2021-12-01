package tests

import (
    "fmt"
    "testing"

    "pods/service"
)

func TestFetchMainProblem(t *testing.T) {
    res := service.FetchMain("211201")

    fmt.Println(res.Title)
    fmt.Println(res.Link)
    fmt.Println(res.Description)
}
