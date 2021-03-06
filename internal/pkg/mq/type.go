package mq

import (
	"context"
	"io"

	"github.com/batazor/shortlink/internal/pkg/mq/query"
	"github.com/batazor/shortlink/internal/pkg/notify"
)

// MQ - common interface of DataBus
type MQ interface { // nolint unused
	// setting
	Init(context.Context) error
	io.Closer // Closer is the interface that wraps the basic Close method.

	// Pub/Sub a pattern
	Publish(context.Context, query.Message) error
	Subscribe(query.Response) error
	UnSubscribe() error
}

// DataBus abstract type
type DataBus struct { // nolint unused
	mq     MQ
	typeMQ string

	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}
