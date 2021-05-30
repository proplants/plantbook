package repo

import (
	"context"
	"strconv"

	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/plant"
	"github.com/pkg/errors"
)

func (pg *PG) GetPlants(ctx context.Context, params plant.GetPlantsParams) ([]*models.Plant, error) {
	query := `select id, title, category_id, short_info::jsonb, notes::jsonb,
			img_links::jsonb, creator, created_at, modifier, modified_at
			from reference.plants`
	limitoffset := " limit " + strconv.Itoa(int(params.Limit)) + " offset " + strconv.Itoa(int(params.Offset)) + ";"
	where := " where"
	var plants []*models.Plant
	var plant *models.Plant
	if *params.Hight == "" && *params.Kind == "" && *params.RecommendPosition == "" && *params.RegardToLight == "" &&
		*params.RegardToMoisture == "" && *params.FloweringTime == "" && *params.Classifiers == "" && plant.Category == 0 {
		if plant.Category == 0 {
			query = query + limitoffset
		} else {
			query = query + where + " category_id = " + strconv.Itoa(int(plant.Category)) + limitoffset
		}
	}
	if *params.Hight != "" || *params.Kind != "" || *params.RecommendPosition != "" || *params.RegardToLight != "" ||
		*params.RegardToMoisture != "" || *params.FloweringTime != "" || *params.Classifiers != "" {
		if plant.Category == 0 {
			where = where + " to_tsvector('russian', short_info) @@ plainto_tsquery('russian','" + *params.Hight + " " + *params.Kind + " " + *params.RecommendPosition + " " +
				*params.RegardToLight + " " + *params.RegardToMoisture + " " + *params.FloweringTime + " " + *params.Classifiers + "')"
			query = query + where + limitoffset
		} else {
			where = where + " to_tsvector('russian', short_info) @@ plainto_tsquery('russian','" + *params.Hight + " " + *params.Kind + " " + *params.RecommendPosition + " " +
				*params.RegardToLight + " " + *params.RegardToMoisture + " " + *params.FloweringTime + " " + *params.Classifiers + "')"
			query = query + where + " and and category_id = " + strconv.Itoa(int(plant.Category)) + limitoffset
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
