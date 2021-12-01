package service

import (
    "fmt"
    . "pods/domain"
)

func ProblemMock(arg ...string) *ProblemJSON {
    return &ProblemJSON{
        Title: "testing title",
        Link: "testing link",
        Description: fmt.Sprint(arg),
    }
}

func DiscussionMock(arg ...string) *DiscussionJSON {
    return &DiscussionJSON{
        Alt: "testing alt",
        Links: []DiscussionEntryJSON{
            { Title: "testing title", Link: fmt.Sprint(arg) },
        },
    }
}
