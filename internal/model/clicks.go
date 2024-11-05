package model

import "time"

type ClickMetadata struct {
	ID              int       `json:"id"`
	UrlUid          int       `json:"url_uid"`
	ClickedAt       time.Time `json:"clicked_at"`
	IpAddress       string    `json:"ip_address"`
	DeviceType      string    `json:"device_type"`
	OperatingSystem string    `json:"operating_system"`
	Referrer        string    `json:"referrer"`
	Browser         string    `json:"browser"`
}