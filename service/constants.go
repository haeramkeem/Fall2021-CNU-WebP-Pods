package service

const (
    MAIN            = 0
    EASY
    MEDIUM
    HARD
    DFS
    BFS
    DP
    TWO_POINTER
    SORT
    BINARY_SEARCH
    TREE
    HEAP
    STRING
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
    "5":    BINARY_SEARCH,
    "6":    TREE,
    "7":    HEAP,
    "8":    STRING,
}
