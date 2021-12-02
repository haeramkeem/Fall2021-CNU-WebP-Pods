package tests

import (
	"fmt"
	. "pods/domain"
	repo "pods/repository"
	"testing"
)

func TestSelectProblem(t *testing.T) {
    err := repo.CreateProblem(&ProblemDB{
        Date: "test date",
        Category: 100,
        Title: "test title",
        Link: "test link",
        Description: "test description",
    })

    if err != nil {
        t.Error(err)
    }

    res, err := repo.SelectProblem("test date", 100)

    if err != nil {
        t.Error(err)
    }

    if res == nil {
        t.Error("Want *ProblemDB, got nil")
    }
    fmt.Println(res.Title, res.Link, res.Description)

    res, err = repo.SelectProblem("undefined", 101)

    if err != nil {
        t.Error(err)
    }

    if res != nil {
        t.Error("Want nil, got someting")
    }
}
