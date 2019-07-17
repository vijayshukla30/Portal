package service

import (
	"context"
	"log"

	"github.com/micro/go-micro/errors"
	"github.com/vijayshukla30/NexthoughtsPortal/user/proto"
)

type userService struct {
	repo userRepository
}

type userRepository interface {
	Create(user *user.User) (message string, err error)
	UpdateUser(user *user.User) (message string, err error)
	ReadUser(id string) (user *user.User, err error)
	DeleteUser(id string) (message string, err error)
	UserSearchRequest(username string, limit int64, offset int64) (users []*user.User, err error)
	UserExist(username string) (exists bool, err error)
}

func NewUserService(repo userRepository) user.AccountHandler {
	return &userService{repo: repo}
}
func (s *userService) Create(ctx context.Context, request *user.CreateRequest, response *user.CreateResponse) error {
	log.Println(">>>>>>>>>>>>>>>>>>>>")
	log.Println(">>>>>>>>>>>>>>>>>>>>")

	if request == nil {
		return errors.BadRequest("", "Missing User Create Request")
	}
	exists, err := s.repo.UserExist(request.User.Username)
	if err != nil {
		return errors.InternalServerError("", "Failed to check user existence: %s", err)
	}

	if exists {
		return errors.NotFound(request.User.Username, "User Already exists")
	}
	message, err := s.repo.Create(request.User)
	if err != nil {
		return errors.InternalServerError("", "Failed to save new User: %s", err)
	}
	response.Message = message
	return nil
}
func (s *userService) Read(ctx context.Context, request *user.ReadRequest, response *user.ReadResponse) error {
	if request == nil {
		return errors.BadRequest("", "Missing User Read Request")
	}
	return nil
}
func (s *userService) Update(ctx context.Context, request *user.UpdateRequest, response *user.UpdateResponse) error {
	if request == nil {
		return errors.BadRequest("", "Missing User Update Request")
	}
	return nil
}
func (s *userService) Delete(ctx context.Context, request *user.DeleteRequest, response *user.DeleteResponse) error {
	if request == nil {
		return errors.BadRequest("", "Missing User Delete Request")
	}
	return nil
}
func (s *userService) Search(ctx context.Context, request *user.SearchRequest, response *user.SearchResponse) error {
	if request == nil {
		return errors.BadRequest("", "Missing User Search Request")
	}

	return nil
}
func (s *userService) UpdatePassword(ctx context.Context, request *user.UpdatePasswordRequest, response *user.UpdatePasswordResponse) error {
	if request == nil {
		return errors.BadRequest("", "Missing User Update Password Request")
	}

	return nil
}
func (s *userService) Login(ctx context.Context, request *user.LoginRequest, response *user.LoginResponse) error {
	if request == nil {
		return errors.BadRequest("", "Missing User Login Request")
	}
	return nil
}
func (s *userService) Logout(ctx context.Context, request *user.LogoutRequest, response *user.LogoutResponse) error {
	if request == nil {
		return errors.BadRequest("", "Missing User Logout Request")
	}
	return nil
}
func (s *userService) ReadSession(ctx context.Context, request *user.ReadSessionRequest, response *user.ReadSessionResponse) error {
	if request == nil {
		return errors.BadRequest("", "Missing Read Session Request")
	}
	return nil
}
