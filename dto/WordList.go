package dto

type WordList struct {
	Id      int64  `json:"id"`
	Word    string `json:"word"`
	Context string `json:"context"`
}
