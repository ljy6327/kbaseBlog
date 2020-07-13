package pojo

type Category struct {
	Base
	Author_id uint
	Name      string `json:"name;not null"`
}
