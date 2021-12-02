package tests

import (
    "testing"
    repo "pods/repository"
    . "pods/domain"
)

func TestCreateProblem(t *testing.T) {
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
}
