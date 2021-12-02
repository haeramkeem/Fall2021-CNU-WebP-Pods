package service

import (
    "fmt"
    . "pods/domain"
)

func GetMainProb(date string) *ProblemJSON {
    return &ProblemJSON{
        Title: "testing title",
        Link: "testing link",
        Description: "testing description",
    }
}

func GetDifficultyProb(levelIdx string) *ProblemJSON {
    return &ProblemJSON{
        Title: "testing title",
        Link: "testing link",
        Description: "testing description",
    }
}

func GetTopicProb(topicIdx string) *ProblemJSON {
    return &ProblemJSON{
        Title: "testing title",
        Link: "testing link",
        Description: "testing description",
    }
}

func GetDiscussion(probpath, lang string) *DiscussionJSON {
    return &DiscussionJSON{
        Alt: "testing alt",
        Links: []DiscussionEntryJSON{
            { Title: "testing title", Link: fmt.Sprint(probpath, lang) },
        },
    }
}
