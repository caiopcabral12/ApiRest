package models

type Presidents struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Biography string `json:"biography"`
	Birth     string `json:"birth"`
}

var President []Presidents
