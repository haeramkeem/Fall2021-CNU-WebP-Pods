package service

import (
    "pods/modules/errors"
    . "pods/domain"
    repo "pods/repository"
)

func GetMainProb(date string) (*ProblemJSON, error) {

    // Get Problem from DB
    prob, err := repo.SelectProblem(date, MAIN)
    if err := errors.Check(err); err != nil {
        return nil, err
    }

    if prob == nil {
        // Fetch Problem
        prob, err = FetchMain(date)
        if err := errors.Check(err); err != nil {
            return nil, err
        }

        // Save to DB
        err = repo.CreateProblem(prob)
        if err := errors.Check(err); err != nil {
            return nil, err
        }
    }

    return &ProblemJSON{
        Title: prob.Title,
        Link: prob.Link,
        Description: prob.Description,
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
