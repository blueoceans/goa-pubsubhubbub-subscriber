package backend

import (
	"net/http"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"

	"github.com/blueoceans/goa-pubsubhubbub-subscriber/app"
)

func init() {
	service := goa.New("appengine")
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	app.MountItemsController(service, newItemsController(service))

	http.HandleFunc("/", service.Mux.ServeHTTP)
	//http.HandleFunc("/", handleFunc)
}
