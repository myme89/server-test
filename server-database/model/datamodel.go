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
	Id       string `bson:"_id,omitempty" json:"_id"`
	UserName string `bson:"user_name,omitempty" json:"user_name"`
	Password string `bson:"password," json:"password"`
	LastName string `bson:"last_name,omitempty" json:"last_name"`
	FistName string `bson:"first_name,omitempty" json:"first_name"`
	FullName string `bson:"full_name,omitempty" json:"full_name"`
}
