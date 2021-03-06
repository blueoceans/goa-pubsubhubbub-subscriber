// API "goa-pubsubhubbub-subscriber": Application Contexts
//
// Code generated by goagen v1.1.0-dirty, DO NOT EDIT.
//
// Command:
// $ goagen
// --design=github.com/blueoceans/goa-pubsubhubbub-subscriber/design
// --out=$(GOPATH)/src/github.com/blueoceans/goa-pubsubhubbub-subscriber
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
)

// NotifyHubContext provides the hub notify action context.
type NotifyHubContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewNotifyHubContext parses the incoming request URL and body, performs validations and creates the
// context used by the hub controller notify action.
func NewNotifyHubContext(ctx context.Context, r *http.Request, service *goa.Service) (*NotifyHubContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := NotifyHubContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *NotifyHubContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// VerifiyHubContext provides the hub verifiy action context.
type VerifiyHubContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	HubChallenge    *string
	HubLeaseSeconds *float64
	HubMode         string
	HubReason       *string
	HubTopic        string
}

// NewVerifiyHubContext parses the incoming request URL and body, performs validations and creates the
// context used by the hub controller verifiy action.
func NewVerifiyHubContext(ctx context.Context, r *http.Request, service *goa.Service) (*VerifiyHubContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := VerifiyHubContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramHubChallenge := req.Params["hub.challenge"]
	if len(paramHubChallenge) > 0 {
		rawHubChallenge := paramHubChallenge[0]
		rctx.HubChallenge = &rawHubChallenge
	}
	paramHubLeaseSeconds := req.Params["hub.lease_seconds"]
	if len(paramHubLeaseSeconds) > 0 {
		rawHubLeaseSeconds := paramHubLeaseSeconds[0]
		if hubLeaseSeconds, err2 := strconv.ParseFloat(rawHubLeaseSeconds, 64); err2 == nil {
			tmp1 := &hubLeaseSeconds
			rctx.HubLeaseSeconds = tmp1
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("hub.lease_seconds", rawHubLeaseSeconds, "number"))
		}
		if rctx.HubLeaseSeconds != nil {
			if *rctx.HubLeaseSeconds < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError(`hub.lease_seconds`, *rctx.HubLeaseSeconds, 0, true))
			}
		}
	}
	paramHubMode := req.Params["hub.mode"]
	if len(paramHubMode) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("hub.mode"))
	} else {
		rawHubMode := paramHubMode[0]
		rctx.HubMode = rawHubMode
		if !(rctx.HubMode == "denied" || rctx.HubMode == "subscribe" || rctx.HubMode == "unsubscribe") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`hub.mode`, rctx.HubMode, []interface{}{"denied", "subscribe", "unsubscribe"}))
		}
	}
	paramHubReason := req.Params["hub.reason"]
	if len(paramHubReason) > 0 {
		rawHubReason := paramHubReason[0]
		rctx.HubReason = &rawHubReason
	}
	paramHubTopic := req.Params["hub.topic"]
	if len(paramHubTopic) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("hub.topic"))
	} else {
		rawHubTopic := paramHubTopic[0]
		rctx.HubTopic = rawHubTopic
		if err2 := goa.ValidateFormat(goa.FormatURI, rctx.HubTopic); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`hub.topic`, rctx.HubTopic, goa.FormatURI, err2))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *VerifiyHubContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}
