package service

import (
	"context"
	stderrors "errors"
	"fmt"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"github.com/vijayshukla30/NexthoughtsPortal/user/proto"
)

const (
	userService = "go.micro.srv.greeter"
)

type CommerceService struct {
	userClient user.AccountService
}

type userResults struct {
	createResponse *user.CreateResponse
	err            error
}

func NewCommerceService(c client.Client) *CommerceService {
	return &CommerceService{
		userClient: user.NewAccountService(userService, c),
	}
}

func (cs *CommerceService) Greeting(request *restful.Request, response *restful.Response) {
	log.Println("?????????????????????")
	response.WriteEntity(map[string]string{
		"message": "Success",
	})
}

func (cs *CommerceService) CreateUser(request *restful.Request, response *restful.Response) {
	log.Println("I am inside Create User")
	username, _ := request.BodyParameter("username")
	firstName, _ := request.BodyParameter("firstName")
	lastName, _ := request.BodyParameter("lastName")
	fmt.Println("Username " + username)
	fmt.Println("First Name " + firstName)
	fmt.Println("Last Name " + lastName)

	ctx := context.Background()

	userCh := cs.createUser(ctx, username, firstName, lastName)
	userReply := <-userCh

	if userReply.err != nil {
		writeError(response, userReply.err)
	}

	message := userReply.createResponse.Message

	response.WriteEntity(map[string]string{
		"message": message,
	})

}

func (cs *CommerceService) createUser(ctx context.Context, username string, firstName string, lastName string) chan userResults {
	ch := make(chan userResults, 1)

	go func() {
		res, err := cs.userClient.Create(ctx, &user.CreateRequest{
			User: &user.User{
				Username:  username,
				FirstName: firstName,
				LastName:  lastName,
			},
		})

		ch <- userResults{
			createResponse: res, err: err,
		}
	}()

	return ch
}

func writeError(response *restful.Response, err error) {
	log.Println("Getting Errors....")
	realError := errors.Parse(err.Error())

	if realError != nil {
		response.WriteError(int(realError.Code), stderrors.New(realError.Detail))
	}

	response.WriteError(http.StatusInternalServerError, err)
}
