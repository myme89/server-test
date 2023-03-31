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

type FileInfoUpload struct {
	Id              string  `bson:"_id,omitempty" json:"_id"`
	FileName        string  `bson:"file_name,omitempty" json:"file_name"`
	Size            float64 `bson:"size," json:"size"`
	TypeFile        string  `bson:"type_file,omitempty" json:"type_file"`
	CreateAt        string  `bson:"create_at,omitempty" json:"create_at"`
	Status          string  `bson:"status,omitempty" json:"status"`
	IdUser          string  `bson:"id_user,omitempty" json:"id_user"`
	Link            string  `bson:"link,omitempty" json:"link"`
	StatusProcesing string  `bson:"status_processing,omitempty" json:"status_processing"`
}

type TemplateInfoPerson struct {
	Id          string `bson:"_id,omitempty" json:"_id"`
	FistName    string `bson:"first_name,omitempty" json:"first_name"`
	LastName    string `bson:"last_name,omitempty" json:"last_name"`
	FullName    string `bson:"full_name,omitempty" json:"full_name"`
	PhoneNumber string `bson:"phone_number,omitempty" json:"phone_number"`
	Address     string `bson:"address,omitempty" json:"address"`
}

type InfoServiceUpload struct {
	Id   string `bson:"id_service,omitempty" json:"id_service"`
	Name string `bson:"name_service,omitempty" json:"name_service"`
}

type InfoServiceProcess struct {
	Id              string `bson:"id_service,omitempty" json:"id_service"`
	IdServiceUpload string `bson:"id_service_upload,omitempty" json:"id_service_upload"`
	Name            string `bson:"name_service,omitempty" json:"name_service"`
}

type InfoFunctionProcess struct {
	Id               string `bson:"id_function,omitempty" json:"id_function"`
	IdServiceProcess string `bson:"id_service_process,omitempty" json:"id_service_process"`
	Name             string `bson:"name_funtion,omitempty" json:"name_funtion"`
}
