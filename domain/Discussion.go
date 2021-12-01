package domain

type DiscussionEntryJSON struct {
    Title   string  `json:"title"`
    Link    string  `json:"link"`
}

type DiscussionJSON struct {
    Alt     string                  `json:"alt"`
    Links   []DiscussionEntryJSON   `json:"links"`
}
