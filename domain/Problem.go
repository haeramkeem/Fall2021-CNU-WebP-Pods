package domain

import (
    "gorm.io/gorm"
)

type ProblemDB struct {
    gorm.Model
    Date        string
    Category    uint
    Title       string
    Link        string
    Description string
}

type ProblemJSON struct {
    Title       string  `json:"title"`
    Link        string  `json:"link"`
    Description string  `json:"description"`
}
