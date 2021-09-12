package repositories

import (
	"context"

	newrelic "github.com/newrelic/go-agent"
)

func CreateB(ctx context.Context) {
	s := newrelic.DatastoreSegment{
		Product:            newrelic.DatastoreMySQL,
		Collection:         "users",
		Operation:          "INSERT",
		ParameterizedQuery: "INSERT INTO users (name, age) VALUES ($1, $2)",
		QueryParameters: map[string]interface{}{
			"name": "Dracula",
			"age":  439,
		},
		Host:         "mysql-server-1",
		PortPathOrID: "3306",
		DatabaseName: "my_database",
	}

	s = startSegment(ctx, s)
	// ... make the datastore call
	s.End()
}
