package repositories

import (
	"context"

	newrelic "github.com/newrelic/go-agent"
)

func Create(ctx context.Context) {
	defer StartDBSegment("users",
		"INSERT",
		"INSERT INTO users (name, age) VALUES ($1, $2)").End()

	// ... make the datastore call

}
