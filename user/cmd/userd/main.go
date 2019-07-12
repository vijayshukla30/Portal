package main

import (
	"fmt"
	"time"

	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/vijayshukla30/NexthoughtsPortal/user/internal/platform/config"
	"github.com/vijayshukla30/NexthoughtsPortal/user/internal/platform/repo"
	"github.com/vijayshukla30/NexthoughtsPortal/user/internal/service"
	user "github.com/vijayshukla30/NexthoughtsPortal/user/proto"
)

func main() {
	svc := grpc.NewService(
		micro.Name(config.ServiceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version(config.Version),
	)

	svc.Init()
	//DB Connection
	dbUser := "root"
	dbPassword := "password"
	dbName := "portal"
	connection := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbName)
	driver := "mysql"

	repoD := repo.NewDBRepository(driver, connection)
	user.RegisterAccountHandler(svc.Server(), service.NewUserService(repoD))
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
