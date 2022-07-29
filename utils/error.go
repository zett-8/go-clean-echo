package utils

type Error struct {
	Errors struct {
		Message error `json:"message"`
	} `json:"errors"`
}
