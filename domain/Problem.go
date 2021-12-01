package domain

import (
    "gorm.io/gorm"
)

type Problem struct {
    gorm.Model
    Date        int
    Category    string
    Title       string
    Link        string
    Description string
}

type ProblemJSON struct {
    Title       string  `json:"title"`
    Link        string  `json:"link"`
    Description string  `json:"description"`
}
