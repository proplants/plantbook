package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/proplants/plantbook/internal/api/models"
)

//  Получение растений из справочника по параметрам или просто все, обязательные поля limit and offset
func (pg *PG) GetRefPlants(ctx context.Context, category int32, limit, offset int64, classifier, floweringTime,
	hight, kind, recommendPosition, regardToLight, regardToMoisture string) ([]*models.RefPlant, error) {
	query := `SELECT id, title, category_id, short_info::jsonb, notes::jsonb,
			img_links::jsonb, creator, created_at, modifier, modified_at
			FROM reference.plants
			WHERE to_tsvector('russian', short_info) @@ plainto_tsquery('russian', '$1::text')
			and category_id = $2
			ORDER BY title 
			LIMIT $3
			OFFSET $4 ;`
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

	// Получение строк с базы
	rows, err := pg.db.Query(ctx, query, tsquery, category, limit, offset)
	if err != nil {
		return nil, errors.WithMessage(err, "Select rows error: ")
	}
	defer rows.Close()
	for rows.Next() {
		var refPlant models.RefPlant
		err = rows.Scan(&refPlant.ID, &refPlant.Title, &refPlant.Category, &refPlant.ShortInfo, &refPlant.Infos,
			&refPlant.Images, &refPlant.Creater, &refPlant.CreatedAt, &refPlant.Modifier, &refPlant.ModifiedAt)
		fmt.Printf("%v", refPlant)
		if err != nil {
			return nil, errors.WithMessage(err, "Scan rows error: ")
		}
		fmt.Printf("%v", refPlant)
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

func quoteIdentifier(s string) string {
	return `"` + strings.Replace(s, `"`, `""`, -1) + `"`
}
