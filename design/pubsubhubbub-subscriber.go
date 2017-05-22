package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

const (
	appID = "goa-pubsubhubbub-subscriber"
)

var _ = API(appID, func() {
	Title(appID)
	Docs(func() {
		URL("https://pubsubhubbub.github.io/PubSubHubbub/pubsubhubbub-core-0.4.html")
	})
	Consumes("application/atom+xml", func() {
		Package("github.com/goadesign/goa")
		Function("NewXMLDecoder")
	})
	Consumes("application/rss+xml", func() {
		Package("github.com/goadesign/goa")
		Function("NewXMLDecoder")
	})
	Produces("application/json")
})

var _ = Resource("hub", func() {
	BasePath("/subscriber")
	// 5.2. Subscription Validation
	// 5.3. Hub Verifies Intent of the Subscriber
	Action("verifiy", func() {
		Routing(GET(""))
		Params(func() {
			Attribute("hub.mode", String, func() {
				Enum("denied", "subscribe", "unsubscribe")
			})
			Attribute("hub.topic", String, func() {
				Format("uri")
			})
			Attribute("hub.reason", String)
			Attribute("hub.challenge", String)
			Attribute("hub.lease_seconds", Number, func() {
				Minimum(0)
			})
			Required("hub.mode", "hub.topic")
		})
		Response(OK)
	})
	// 7. Content Distribution
	Action("notify", func() {
		Routing(POST(""))
		Response(OK)
	})
})
