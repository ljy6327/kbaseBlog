package pojo

import (
	"github.com/jinzhu/gorm"
)

type Base struct {
	gorm.Model
	State int
}
