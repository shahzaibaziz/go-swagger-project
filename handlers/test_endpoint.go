package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/shahzaibaziz/GOProject"
	"github.com/shahzaibaziz/GOProject/gen/restapi/operations/test"
)

// NewAcceptConsentChallenge handler of the accept consent challenge
func NewDemoHandler(ctx context.Context, rt *runtime.Runtime) test.TestEndpointHandler {
	return &demoHandler{
		ctx: ctx,
		rt:  rt,
	}
}

type demoHandler struct {
	ctx context.Context
	rt  *runtime.Runtime
}

func (c *demoHandler) Handle(params test.TestEndpointParams) middleware.Responder {

	return test.NewTestEndpointOK()
}
