// Package gardens contains http Handlers for gardens
package gardens

import (
	"context"

	"github.com/kaatinga/plantbook/internal/api/models"
)

// RepoInterface gardens repository behavior
type RepoInterface interface {
	StoreGarden(ctx context.Context, garden *models.Garden) (*models.Garden, error)
}
