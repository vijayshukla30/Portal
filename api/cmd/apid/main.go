package main

import (
	"log"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"
	"github.com/vijayshukla30/NexthoughtsPortal/api/internal/service"
)

const (
	serviceName = "go.micro.api.greeter"
)

func main() {
	webService := web.NewService(
		web.Name(serviceName),
		web.Version("(undefined)"),
	)
	webService.Init()

	log.Println("I am inside........")

	handler := service.NewCommerceService(client.DefaultClient)

	wc := restful.NewContainer()
	ws := new(restful.WebService)

	ws.
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)
	log.Println("API URL")
	ws.Path("/greeter")
	ws.Route(ws.POST("/users/create").To(handler.CreateUser)).
		Doc("Create a User")

	wc.Add(ws)
	webService.Handle("/", wc)
	if err := webService.Run(); err != nil {
		log.Fatal(err)
	}
}
