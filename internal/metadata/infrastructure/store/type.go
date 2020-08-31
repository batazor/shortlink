package meta_store

import (
	"context"

	rpc "github.com/batazor/shortlink/internal/metadata/domain"
	"github.com/batazor/shortlink/internal/notify"
)

type Repository interface {
	Get(context.Context, string) (*rpc.Meta, error)
	Add(context.Context, *rpc.Meta) error
}

// Store abstract type
type MetaStore struct {
	Repository

	typeStore string
	Store     Repository

	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}
