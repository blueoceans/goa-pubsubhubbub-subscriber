// +build appengine

package backend

import (
	"net/http"
	"time"

	"appengine"
)

const (
	deadline = time.Duration(60) * time.Second
)

func appengineContext(
	r *http.Request,
) appengine.Context {
	return appengine.Timeout(appengine.NewContext(r), deadline)
}
