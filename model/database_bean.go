package model

type Wiki struct {
	ID       string `json:"id"`       //
	Content  string `json:"content"`  //
	Status   string `json:"status"`   //
	PageType string `json:"pageType"` //
	ModelTime
}

type WikiImage struct {
	ID       int    `json:"id"`       //
	PageType string `json:"pageType"` //
	ImageUrl string `json:"imageUrl"` //
	ModelTime
}
