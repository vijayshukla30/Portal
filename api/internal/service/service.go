package service

import (
	"context"
	stderrors "errors"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"github.com/vijayshukla30/NexthoughtsPortal/user/proto"
)

const (
	userService = "go.micro.srv.user"
)

type CommerceService struct {
	userService user.AccountService
}

type userResults struct {
	createResponse *user.CreateResponse
	err            error
}

func NewCommerceService(c client.Client) *CommerceService {
	return &CommerceService{
		userService: user.NewAccountService(userService, c),
	}
}

func (cs *CommerceService) CreateUser(request *restful.Request, response *restful.Response) {
	log.Println("I am inside Create User")
	usr := &user.User{}
	err := request.ReadEntity(usr)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	ctx := context.Background()

	userCh := cs.createUser(ctx, usr)
	userReply := <-userCh

	if userReply.err != nil {
		writeError(response, userReply.err)
	}

	message := userReply.createResponse.Message

	response.WriteEntity(map[string]string{
		"message": message,
	})

}

func (cs *CommerceService) createUser(ctx context.Context, usr *user.User) chan userResults {
	ch := make(chan userResults)
	log.Println(usr.Username)
	go func() {
		res, err := cs.userService.Create(ctx, &user.CreateRequest{
			User: usr,
		})
		log.Print("Response ")
		log.Println(res)
		log.Println("")
		log.Print("Error ")
		log.Println(err)
		log.Println("")

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
