package repo

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	id        int64
	username  string
	firstName string
	lastName  string
	mobile    int32
	dob       int64
	doj       int64
	created   int64
	updated   int64
}
