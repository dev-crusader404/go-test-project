package models

type RequestID string

type SearchRequest struct {
	Title string `json:"title,omitempty"`
	Year  string `json:"year,omitempty"`
}

type SearchResponse struct {
	MovieTitle   string   `json:"movieTitle,omitempty"`
	Year         string   `json:"year,omitempty"`
	Description  string   `json:"description,omitempty"`
	Rating       float32  `json:"rating,omitempty"`
	Genre        []string `json:"genre,omitempty"`
	ReleasedDate string   `json:"releasedDate,omitempty"`
	GrossIncome  int64    `json:"grossIncome,omitempty"`
}
