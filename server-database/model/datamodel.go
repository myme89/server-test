package model

type DataInfo struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

// type DataPost struct {
// 	Id       int    `json:"id"`
// 	Name     string `json:"name"`
// 	FullName string `json:"full_name"`
// }

type DataPost struct {
	Id       int    `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:255;not null" json:"name"`
	FullName string `gorm:"size:100;not null" json:"fullname"`
}

type UserInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	LastName string `json:"last_name"`
	FistName string `json:"fist_name"`
	FullName string `json:"full_name"`
}
