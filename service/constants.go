package service

const (
    MAIN            = 0
    EASY            = 1
    MEDIUM          = 2
    HARD            = 3
    DFS             = 4
    BFS             = 5
    DP              = 6
    TWO_POINTER     = 7
    SORT            = 8
    BINARY_SEARCH   = 9
    TREE            = 10
    HEAP            = 11
    STRING          = 12
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
