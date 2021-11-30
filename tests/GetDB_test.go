package tests

import (
    "fmt"
    "testing"

    "pods/repository"
)

func TestGetDB(t *testing.T) {
    fmt.Println("test")
    db, err := repository.GetDB()
    check(err)
    if db == nil {
        t.Error("Cannot connect to DB")
    }
}
