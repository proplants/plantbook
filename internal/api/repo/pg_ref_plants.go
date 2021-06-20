package repo

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
	"github.com/proplants/plantbook/internal/api/models"
)

// GetRefPlans - get all reference plants by parametrs.
func (pg *PG) GetRefPlants(ctx context.Context, category int32, limit, offset int64, classifier, floweringTime,
	hight, kind, recommendPosition, regardToLight, regardToMoisture string) ([]*models.RefPlant, int64, int64, error) {
	query := `SELECT id, title, category_id, short_info::jsonb, notes::jsonb,
			img_links::jsonb, creator, created_at, modifier, modified_at
			FROM reference.plants`
	totalquery := `select count(1) as cnt from reference.plants`
	var refPlants []*models.RefPlant
	var tsquery string
	if hight != "" {
		tsquery = tsquery + " " + hight
	}
	if kind != "" {
		tsquery = tsquery + " " + kind
	}
	if recommendPosition != "" {
		tsquery = tsquery + " " + recommendPosition
	}
	if regardToLight != "" {
		tsquery = tsquery + " " + regardToLight
	}
	if regardToMoisture != "" {
		tsquery = tsquery + " " + regardToMoisture
	}
	if floweringTime != "" {
		tsquery = tsquery + " " + floweringTime
	}
	if classifier != "" {
		tsquery = tsquery + " " + classifier
	}
	if tsquery != "" {
		query = query + " WHERE to_tsvector('russian', short_info) @@ plainto_tsquery('russian', '" + tsquery + "')"
		totalquery = totalquery + " WHERE to_tsvector('russian', short_info) @@ plainto_tsquery('russian', '" + tsquery + "')"
		if category != 0 {
			query = query + " AND category_id = " + strconv.Itoa(int(category))
			totalquery = totalquery + " AND category_id = " + strconv.Itoa(int(category))
		}
	} else {
		if category != 0 {
			query = query + " WHERE category_id = " + strconv.Itoa(int(category))
			totalquery = totalquery + " WHERE category_id = " + strconv.Itoa(int(category))
		}
	}
	totalquery = totalquery + ";"
	query = query + " ORDER BY title LIMIT $1 OFFSET $2;"
	rows, err := pg.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, 0, errors.WithMessage(err, "Select rows error: ")
	}
	defer rows.Close()
	for rows.Next() {
		var refPlant models.RefPlant
		err = rows.Scan(&refPlant.ID, &refPlant.Title, &refPlant.Category, &refPlant.ShortInfo, &refPlant.Infos,
			&refPlant.Images, &refPlant.Creater, &refPlant.CreatedAt, &refPlant.Modifier, &refPlant.ModifiedAt)
		if err != nil {
			return nil, 0, 0, errors.WithMessage(err, "Scan rows error: ")
		}
		refPlants = append(refPlants, &refPlant)
	}
	var total, count int64
	count = rows.CommandTag().RowsAffected()
	err = pg.db.QueryRow(ctx, totalquery).Scan(&total)
	if err != nil {
		return nil, 0, 0, errors.WithMessage(err, "Scan total rows error: ")
	}
	return refPlants, count, total, err
}

// GetRefPlantByID extracts reference.plant by specified id.
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
