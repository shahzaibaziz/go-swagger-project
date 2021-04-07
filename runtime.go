package runtime

import (
	wraperrors "github.com/pkg/errors"

	"github.com/shahzaibaziz/GOProject/db"
	"github.com/shahzaibaziz/GOProject/db/mysql"
	"github.com/shahzaibaziz/GOProject/service"
)

// Runtime initializes values for entry point to our application
type Runtime struct {
	svc *service.Service
}

// NewRuntime creates a new runtime
func NewRuntime() (*Runtime, error) {
	client, err := mysql.NewClient(db.Option{})
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to connect with database")
	}

	return &Runtime{svc: service.NewService(client)}, wraperrors.Wrap(err, "failed to connect with database")
}

// Service returns pointer to our service variable
func (r Runtime) Service() *service.Service {
	return r.svc
}
