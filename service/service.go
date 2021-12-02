package service

import (
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

func GetDiscussion(probpath, lang string) (*DiscussionJSON, error) {
    return FetchDiscuss(`/problems/` + probpath, lang)
}
