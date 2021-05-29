package repo

import (
	"context"
	"strconv"

	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/pkg/errors"
)

func (pg *PG) GetPlants(ctx context.Context, plant *models.Plant, limit int, offset int) ([]*models.Plant, error) {
	query := `select id, title, category_id, short_info::jsonb, notes::jsonb,
			img_links::jsonb, creator, created_at, modifier, modified_at
			from reference.plants`
	limitoffset := " limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset) + ";"
	params := " where"
	shorInfo := plant.ShortInfo
	var plants []*models.Plant

	if shorInfo.Hight == "" && shorInfo.Kind == "" && shorInfo.RecommendPosition == "" && shorInfo.RegardToLight == "" &&
		shorInfo.RegardToMoisture == "" && shorInfo.FloweringTime == "" && shorInfo.Classifiers == "" && plant.Category == 0 {
		if plant.Category == 0 {
			query = query + limitoffset
		} else {
			query = query + params + " category_id = " + strconv.Itoa(int(plant.Category)) + limitoffset
		}
	}
	if shorInfo.Hight != "" || shorInfo.Kind != "" || shorInfo.RecommendPosition != "" || shorInfo.RegardToLight != "" ||
		shorInfo.RegardToMoisture != "" || shorInfo.FloweringTime != "" || shorInfo.Classifiers != "" {
		if plant.Category == 0 {
			params = params + " to_tsvector('russian', short_info) @@ plainto_tsquery('russian','" + shorInfo.Hight + " " + shorInfo.Kind + " " + shorInfo.RecommendPosition + " " +
				shorInfo.RegardToLight + " " + shorInfo.RegardToMoisture + " " + shorInfo.FloweringTime + " " + shorInfo.Classifiers + "')"
			query = query + params + limitoffset
		} else {
			params = params + " to_tsvector('russian', short_info) @@ plainto_tsquery('russian','" + shorInfo.Hight + " " + shorInfo.Kind + " " + shorInfo.RecommendPosition + " " +
				shorInfo.RegardToLight + " " + shorInfo.RegardToMoisture + " " + shorInfo.FloweringTime + " " + shorInfo.Classifiers + "')"
			query = query + params + " and and category_id = " + strconv.Itoa(int(plant.Category)) + limitoffset
		}
	}

	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		return nil, errors.WithMessage(err, "select rows failed")
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&plant.ID, &plant.Title, &plant.Category, &plant.ShortInfo, &plant.Infos, &plant.Images, &plant.Creater, &plant.CreatedAt, &plant.Modifier)
		if err != nil {
			return nil, errors.WithMessage(err, "scan rows failed")
		}
		plants = append(plants, plant)
	}
	return plants, nil
}
