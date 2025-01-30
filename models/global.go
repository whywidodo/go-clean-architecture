package models

import "time"

type Response struct {
	StatusCode       string          `json:"statusCode"`
	Success          bool            `json:"success"`
	ResponseDatetime time.Time       `json:"responseDatetime"`
	Result           interface{}     `json:"result"`
	Messages         MessageResponse `json:"messages"`
}

type RequestList struct {
	Order   string `json:"order"`
	OrderBy string `json:"orderBy"`
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	Keyword string `json:"keyword"`
}

type MessageResponse struct {
	Id string `json:"id"`
	En string `json:"en"`
}
