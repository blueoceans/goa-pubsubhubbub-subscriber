# goa-pubsubhubbub-subscriber
A PubSubHubbub subscriber for goa

[![wercker status](https://app.wercker.com/status/ab487a8c5ac51d87e7c83525356fb63c/m/master "wercker status")](https://app.wercker.com/project/byKey/ab487a8c5ac51d87e7c83525356fb63c)

# clone repository

    $ cd `goapp env GOPATH`
    $ mkdir -p src/github.com/blueoceans
    $ cd src/github.com/blueoceans
    $ git clone --depth 1 https://github.com/blueoceans/goa-pubsubhubbub-subscriber.git
    $ cd goa-pubsubhubbub-subscriber

# dep

    $ go get -u github.com/golang/dep/...
    $ dep init
    $ dep ensure

# start server

    $ dev_appserver.py backend

# with [goa](https://github.com/goadesign/goa)

Install goa:

    $ go get -u github.com/goadesign/goa/...

Generate code:

    $ goagen app -d github.com/blueoceans/goa-pubsubhubbub-subscriber/design
    $ goagen swagger -d github.com/blueoceans/goa-pubsubhubbub-subscriber/design

# with swagger

Get swagger-ui:

    $ make swagger-ui

Open [swagger-ui](http://localhost:8080/swagger-ui/?url=http://localhost:8080/swagger.json).

    $ open "http://localhost:8080/swagger-ui/?url=http://localhost:8080/swagger.json"

# Deploy to the Google App Engine

Create a Google App Engine application:

    $ CLOUDSDK_CORE_PROJECT=<APP-ID> gcloud app create

Create TaskQueue and Cron:

    $ CLOUDSDK_CORE_PROJECT=<APP-ID> gcloud app deploy backend/queue.yaml
    $ CLOUDSDK_CORE_PROJECT=<APP-ID> gcloud app deploy backend/cron.yaml

Deploy to Google App Engine:

    $ goapp deploy backend/app.yaml

Get an OAuth 2.0 refresh token:

    $ jq .refresh_token ~/.appcfg_oauth2_tokens

Then set to an `APP_ENGINE_TOKEN` of the application environment variables on the Wercker CI.
