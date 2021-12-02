package service

const BASE string = `https://leetcode.com`;

const (
    MAIN            = 0
    EASY            = 1
    MEDIUM          = 2
    HARD            = 3
    BFS             = 4
    DFS             = 5
    DP              = 6
    TWO_POINTER     = 7
    SORT            = 8
    GREEDY          = 9
    BINARY_SEARCH   = 10
    TREE            = 11
    HEAP            = 12
    STRING          = 13
)

var levelIdxToConst = map[string]uint{
    "0":    EASY,
    "1":    MEDIUM,
    "2":    HARD,
}

var topicIdxToConst = map[string]uint{
    "0":    BFS,
    "1":    DFS,
    "2":    DP,
    "3":    TWO_POINTER,
    "4":    SORT,
    "5":    GREEDY,
    "6":    BINARY_SEARCH,
    "7":    TREE,
    "8":    HEAP,
    "9":    STRING,
}
