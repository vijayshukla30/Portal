package repo

import (
	"github.com/micro/go-micro/errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vijayshukla30/NexthoughtsPortal/user/proto"
)

type UserRepository struct {
	driver     string
	connection string
}

func NewDBRepository(driver string, connection string) *UserRepository {
	return &UserRepository{driver: driver, connection: connection}
}

func (r *UserRepository) Create(user *user.User) (message string, err error) {
	user.Created = time.Now().Unix()
	user.Updated = time.Now().Unix()

	db, err := gorm.Open(r.driver, r.connection)

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	var newUser = User{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Mobile:    user.Mobile,
		Dob:       user.Dob,
		Doj:       user.Doj,
	}

	var checkUser User
	db.First(&checkUser, "username=?", user.Username)

	log.Println(checkUser.ID)
	if checkUser.ID > 0 {
		return "User creation failed", errors.BadRequest("", "User Already Exist")
	}

	db.Create(&newUser)

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
