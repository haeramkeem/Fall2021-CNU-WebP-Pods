package service

import (
	. "pods/domain"
	"pods/modules/errors"
	"pods/modules/logging"
	repo "pods/repository"
)

/**
 * Get JSON-like structure of main problem from DB or Leetcode.com
 * @params date string: requested problem's serving date
 * @return *ProblemJSON: JSON-like structure
 * @return error: generated error
 */
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

/**
 * Get JSON-like structure of problem by difficulty from DB or Leetcode.com
 * @params levelIdx string: level of difficulty. for more information, check service/constants.go
 * @return *ProblemJSON: JSON-like structure
 * @return error: generated error
 */
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

/**
 * Get JSON-like structure of problem by algorithm topic from DB or Leetcode.com
 * @params topicIdx string: algorithm topic index. for more information, check service/constants.go
 * @return *ProblemJSON: JSON-like structure
 * @return error: generated error
 */
func GetTopicProb(topicIdx string) (*ProblemJSON, error) {
    logging.Log("Request Problem of Topic " + topicIdx)

    // Get Problem from DB
    prob, err := repo.SelectProblem(nowStrDate(), topicIdxToConst[topicIdx])
    if err := errors.Check(err); err != nil {
        return nil, err
    }

    if prob == nil {
        // Fetch Problem
        prob, err = FetchProblemByTopic(topicIdx)
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

/**
 * Get JSON-like structure of discussions by problem path from Leetcode.com
 * @params probpath string: problem path of Leetcode.com
 * @params lang string: requested programming language
 * @return *ProblemJSON: JSON-like structure
 * @return error: generated error
 */
func GetDiscussion(probpath, lang string) (*DiscussionJSON, error) {
    return FetchDiscuss(`/problems/` + probpath, lang)
}
