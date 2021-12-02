package service

import (
    . "pods/domain"
)

func GetMainProb(date string) (*ProblemJSON, error) {
    return &ProblemJSON{
        Title: "testing title",
        Link: "testing link",
        Description: "testing description",
    }, nil
}

func GetDifficultyProb(levelIdx string) (*ProblemJSON, error) {
    return &ProblemJSON{
        Title: "testing title",
        Link: "testing link",
        Description: "testing description",
    }, nil
}

func GetTopicProb(topicIdx string) (*ProblemJSON, error) {
    return &ProblemJSON{
        Title: "testing title",
        Link: "testing link",
        Description: "testing description",
    }, nil
}

func GetDiscussion(probpath, lang string) (*DiscussionJSON, error) {
    return FetchDiscuss(`/problems/` + probpath, lang)
}
