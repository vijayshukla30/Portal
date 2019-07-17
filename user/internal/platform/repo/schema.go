package repo

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(100);unique_index"`
	FirstName string
	LastName  string
	Mobile    int32
	Dob       int64
	Doj       int64
	Password string
}
