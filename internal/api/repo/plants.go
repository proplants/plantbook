package repo

import (
	"context"
)

func (pg *PG) GetPlants(ctx context.Context, plant *models.Plant) (*models.Plants, error) {
	query := `select id, title, category_id, short_info, notes,
			img_links, creator, created_at, modifier, modified_at
			from reference.plants where category_id = $1 and to_tsvector('russian', title) ||
			to_tsvector('russian', short_info) @@
			plainto_tsquery('russian','$2')
			limit 10 offset 0;`

}
