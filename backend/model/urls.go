package model

import "time"

type UrlModel struct {
	Uid        int       `json:"uid"`
	User_id    int       `json:"user_id"`
	Url        string    `json:"url"`
	Created_at time.Time `json:"created_at"`
	Clicks     int       `json:"clicks"`
}