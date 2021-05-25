// Package health contains handler implementations of the ./internal/api/restapi/operations/health
package health

import (
	"context"
)

// RepoInterface health repository behavior
type RepoInterface interface {
	Health(ctx context.Context) error
}
