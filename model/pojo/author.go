package pojo

type Author struct {
	Base
	Name     string `json:"name;not null"`
	Email    string `json:"email;not null"`
	Password string `json:"password;not null"`
}

func (Author) TableName() string {
	return "author"
}
