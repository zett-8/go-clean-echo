package models

type Author struct {
	id      int    `json:"id"`
	name    string `json:"name"`
	country string `json:"country"`
}
