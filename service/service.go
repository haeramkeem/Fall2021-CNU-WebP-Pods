package service

import (
	. "pods/domain"
	"pods/modules/errors"
	"pods/modules/logging"
	repo "pods/repository"
)

func GetMainProb(date string) (*ProblemJSON, error) {
    logging.Log("Request Main Problem of " + date)

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
    logging.Log("Request Problem of Difficulty " + levelIdx)

    // Get Problem from DB
    prob, err := repo.SelectProblem(nowStrDate(), levelIdxToConst[levelIdx])
    if err := errors.Check(err); err != nil {
        return nil, err
    }

    if prob == nil {
        // Fetch Problem
        prob, err = FetchProblemByDifficulty(levelIdx)
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
