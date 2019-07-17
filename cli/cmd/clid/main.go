package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/vijayshukla30/NexthoughtsPortal/user/proto"
)

func main() {
	service := micro.NewService(micro.Name("nexthoughts.client"))

	nexthoughts := user.NewAccountService("go.micro.srv.user", service.Client())

	rsp, err := nexthoughts.Create(context.Background(), &user.CreateRequest{
		User: &user.User{
			Username:  "vijay1",
			FirstName: "Vijay",
			LastName:  "Shukla",
		},
		Password: "123456",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Message)
}
