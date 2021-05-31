package repo

import (
	"context"
	"strconv"

	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/plant"
	"github.com/kaatinga/plantbook/pkg/logging"
	"github.com/pkg/errors"
)

func (pg *PG) GetPlants(ctx context.Context, params plant.GetPlantsParams) ([]*models.Plant, error) {
	query := `select id, title, category_id, short_info::jsonb, notes::jsonb,
			img_links::jsonb, creator, created_at, modifier, modified_at
			from reference.plants`
	var plants []*models.Plant
	si := models.ShortInfo{
		Kind:              *params.Kind,
		Hight:             *params.Hight,
		RecommendPosition: *params.RecommendPosition,
		RegardToLight:     *params.RegardToLight,
		RegardToMoisture:  *params.RegardToMoisture,
		FloweringTime:     *params.FloweringTime,
		Classifiers:       *params.Classifiers,
	}
	log := logging.FromContext(params.HTTPRequest.Context())
	limitoffset := " limit " + strconv.Itoa(int(params.Limit)) + " offset " + strconv.Itoa(int(params.Offset)) + ";"
	where := " where"
	// Если параметры пустые, получаем весь списко с limit and offset
	if params.Hight == nil && params.Kind == nil && params.RecommendPosition == nil && params.RegardToLight == nil &&
		params.RegardToMoisture == nil && params.FloweringTime == nil && params.Classifiers == nil && params.Category == nil {
		// Учитываем категорию, есть или нет
		if params.Category == nil {
			query = query + limitoffset
		} else {
			query = query + where + " category_id = " + strconv.Itoa(int(*params.Category)) + limitoffset
			log.Errorf("error handle: %v", query)
		}
	}
	// Если есть параметры, ищем по параметрам, добавляем к строке
	if params.Hight != nil || params.Kind != nil || si.RecommendPosition != "" || params.RegardToLight != nil ||
		params.RegardToMoisture != nil || params.FloweringTime != nil || params.Classifiers != nil {
		// Учитываем категории
		if params.Category == nil {
			where = where + " to_tsvector('russian', short_info) @@ plainto_tsquery('russian','" + si.Hight + " " + si.Kind + " " + si.RecommendPosition + " " +
				si.RegardToLight + " " + si.RegardToMoisture + " " + si.FloweringTime + " " + si.Classifiers + "')"
			query = query + where + limitoffset
		} else {
			where = where + " to_tsvector('russian', short_info) @@ plainto_tsquery('russian','" + *params.Hight + " " + *params.Kind + " " + *params.RecommendPosition + " " +
				*params.RegardToLight + " " + *params.RegardToMoisture + " " + *params.FloweringTime + " " + *params.Classifiers + "')"
			query = query + where + " and and category_id = " + strconv.Itoa(int(*params.Category)) + limitoffset
		}
	}
	// Получение строк с базы
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		return nil, errors.WithMessage(err, "select rows failed")
	}
	defer rows.Close()
	for rows.Next() {
		var plant models.Plant
		err = rows.Scan(&plant.ID, &plant.Title, &plant.Category, &plant.ShortInfo, &plant.Infos,
			&plant.Images, &plant.Creater, &plant.CreatedAt, &plant.Modifier, &plant.ModifiedAt)
		if err != nil {
			return nil, errors.WithMessage(err, "scan rows failed")
		}
		plants = append(plants, &plant)
	}

	return plants, err
}
