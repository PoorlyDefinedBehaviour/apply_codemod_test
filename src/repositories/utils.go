package repositories

import (
	"context"

	newrelic "github.com/newrelic/go-agent"
)

func startSegment(_ context.Context, segment newrelic.DatastoreSegment) newrelic.DatastoreSegment {
	return segment
}
