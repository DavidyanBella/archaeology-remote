package models

type TimeRequest struct {
	AccessToken int    `json:"access_key"`
	Time        string `json:"expedition_date"`
}
