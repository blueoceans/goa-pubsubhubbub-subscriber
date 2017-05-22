package backend

import (
	"io/ioutil"
	"net/http"

	u "github.com/blueoceans/go-common/logutil"
	"github.com/goadesign/goa"

	"github.com/blueoceans/goa-pubsubhubbub-subscriber/app"
)

type hubController struct {
	*goa.Controller
}

func newHubController(service *goa.Service) *hubController {
	return &hubController{Controller: service.NewController("hubController")}
}

func (c *hubController) Verifiy(ctx *app.VerifiyHubContext) error {
	return ctx.OK([]byte(*ctx.HubChallenge))
}

func (c *hubController) Notify(ctx *app.NotifyHubContext) error {
	ac := appengineContext(ctx.RequestData.Request)
	u.Debugf(ac, "%q", ctx.RequestData.Request.Header.Get("Content-Type"))
	body, err := ioutil.ReadAll(ctx.RequestData.Request.Body)
	switch err {
	case nil:
		u.Debugf(ac, "%q", body)
	default:
		u.Debugf(ac, "%q", err)
	}
	return ctx.OK(nil)
}

func handleFunc(
	w http.ResponseWriter,
	r *http.Request,
) {
	return
}
