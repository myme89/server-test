package model

type DataInfo struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

type DataPost struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}
