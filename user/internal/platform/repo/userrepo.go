package repo

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	user "github.com/vijayshukla30/NexthoughtsPortal/user/proto"
)

type UserRepository struct {
	driver     string
	connection string
}

func NewDBRepository(driver string, connection string) *UserRepository {
	return &UserRepository{driver: driver, connection: connection}
}

func (r *UserRepository) CreateUser(user *user.User) (message string, err error) {
	user.Created = time.Now().Unix()
	user.Updated = time.Now().Unix()

	return "", nil
}

func (r *UserRepository) UpdateUser(user *user.User) (message string, err error) {
	return "", nil
}

func (r *UserRepository) ReadUser(id string) (user *user.User, err error) {
	return nil, nil
}

func (r *UserRepository) DeleteUser(id string) (message string, err error) {
	return "", nil
}

func (r *UserRepository) UserSearchRequest(username string, limit int64, offset int64) (users []*user.User, err error) {
	return nil, nil
}

func (r *UserRepository) UserExist(username string) (exists bool, err error) {
	return false, nil
}
