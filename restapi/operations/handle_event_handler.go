package operations

import (
	middleware "github.com/go-openapi/runtime/middleware"

	"fmt"
	"github.com/bpcreech/hugohook/impl"
	"github.com/davecgh/go-spew/spew"
)

func HandleEventHandlerImpl(params HandleEventParams) middleware.Responder {
	fmt.Printf("%+v\n", params)
	fmt.Printf("%+v\n", params.HTTPRequest.Body)
	spew.Dump(*params.Event)

	if params.XGitHubEvent != "push" {
		return NewHandleEventOK()
	}

	if params.Event.Ref != "refs/heads/master" {
		return NewHandleEventOK()
	}

	err := impl.CloneAndRun(params.Event.Repository.URL)
	if err != nil {
		return NewHandleEventInternalServerError().WithPayload(err.Error())
	}

	return NewHandleEventOK()
}
