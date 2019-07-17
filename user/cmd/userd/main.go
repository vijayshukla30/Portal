package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/vijayshukla30/NexthoughtsPortal/user/internal/platform/config"
	"github.com/vijayshukla30/NexthoughtsPortal/user/internal/platform/repo"
	"github.com/vijayshukla30/NexthoughtsPortal/user/internal/service"
	"github.com/vijayshukla30/NexthoughtsPortal/user/proto"
)

func main() {
	svc := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version(config.Version),
	)

	svc.Init()

	//DB Connection
	dbUser := "root"
	dbPassword := "nextdefault"
	dbName := "portal"
	connection := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbName)
	driver := "mysql"

	repoD := repo.NewDBRepository(driver, connection)

	user.RegisterAccountHandler(svc.Server(), service.NewUserService(repoD))
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
