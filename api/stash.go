package api

type Stash struct {
	AccountName       string `json:"accountName"`
	LastCharacterName string `json:"lastCharacterName"`
	Id                string `json:"id"`
	Label             string `json:"stash"`
	Type              string `json:"stashType"`
	Items             []Item `json:"items"`
	IsPublic          bool   `json:"public"`
}
