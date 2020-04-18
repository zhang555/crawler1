package model

type Wiki struct {
	ID      string `json:"id"`      //
	Content string `json:"content"` //
	Status  string `json:"status"`  //
	ModelTime
}
