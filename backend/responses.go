package main

type SimpleResponse struct {
	Data string `json:"data"`
}

type OpResponse struct {
	Success bool `json:"success"`
}
