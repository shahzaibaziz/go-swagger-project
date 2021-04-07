package handlers

import (
	"context"

	"github.com/go-openapi/loads"

	runtime "github.com/shahzaibaziz/GOProject"
	"github.com/shahzaibaziz/GOProject/gen/restapi/operations"
)

// Handler replaces swagger handler
type Handler *operations.GOProjectAPI

// NewHandler overrides swagger api handlers
func NewHandler(ctx context.Context, rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewGOProjectAPI(spec)

	handler.TestTestEndpointHandler = NewDemoHandler(ctx, rt)

	return handler
}
