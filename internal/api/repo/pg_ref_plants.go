package repo

import (
	"context"

	"github.com/pkg/errors"
	"github.com/proplants/plantbook/internal/api/models"
)

//  Получение растений из справочника по параметрам или просто все, обязательные поля limit and offset
func (pg *PG) GetRefPlants(ctx context.Context, category, limit, offset int32, classifier, floweringTime,
	hight, kind, recommendPosition, regardToLight, regardToMoisture string) ([]*models.RefPlant, error) {
	query := `SELECT id, title, category_id, short_info::jsonb, notes::jsonb,
			img_links::jsonb, creator, created_at, modifier, modified_at from reference.plants 
			WHERE to_tsvector('russian', short_info) @@ plainto_tsquery('russian','|| $2 ||','|| $3 ||','|| $4 ||','
			|| $5 ||','|| $6 ||','|| $7 ||','|| $8 ||') 
			AND category_id = $1
			ORDER BY title 
			LIMIT $9 
			OFFSET $10`
	var refPlants []*models.RefPlant

	// Получение строк с базы
	rows, err := pg.db.Query(ctx, query, category, classifier, floweringTime, hight, kind, recommendPosition,
		regardToLight, regardToMoisture, limit, offset)
	if err != nil {
		return nil, errors.WithMessage(err, "Select rows error: ")
	}
	defer rows.Close()
	for rows.Next() {
		var refPlant models.RefPlant
		err = rows.Scan(&refPlant.ID, &refPlant.Title, &refPlant.Category, &refPlant.ShortInfo, &refPlant.Infos,
			&refPlant.Images, &refPlant.Creater, &refPlant.CreatedAt, &refPlant.Modifier, &refPlant.ModifiedAt)
		if err != nil {
			return nil, errors.WithMessage(err, "Scan rows error: ")
		}
		refPlants = append(refPlants, &refPlant)
	}

	return refPlants, err
}

func (pg *PG) GetRefPlantByID(ctx context.Context, id int64) (*models.RefPlant, error) {
	query := `select id, title, category_id, short_info::jsonb, notes::jsonb,
			img_links::jsonb, creator, created_at, modifier, modified_at
			from reference.plants where id = $1`

	var refPlant models.RefPlant
	err := pg.db.QueryRow(ctx, query, id).Scan(&refPlant.ID, &refPlant.Title, &refPlant.Category,
		&refPlant.ShortInfo, &refPlant.Infos, &refPlant.Images, &refPlant.Creater, &refPlant.CreatedAt,
		&refPlant.Modifier, &refPlant.ModifiedAt)
	if err != nil {
		return nil, errors.WithMessage(err, "QueryRow.Scan error")
	}

	return &refPlant, err
}
