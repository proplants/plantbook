package repo

import (
	"context"
	"strconv"

	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/plant"
	"github.com/pkg/errors"
)

//  Получение растений из справочника по параметрам или просто все, обязательные поля limit and offset
func (pg *PG) GetPlants(ctx context.Context, params plant.GetPlantsParams) ([]*models.Plant, error) {
	query := `select id, title, category_id, short_info::jsonb, notes::jsonb,
			img_links::jsonb, creator, created_at, modifier, modified_at
			from reference.plants`
	var plants []*models.Plant
	// log := logging.FromContext(params.HTTPRequest.Context())
	limitoffset := " limit " + strconv.Itoa(int(params.Limit)) + " offset " + strconv.Itoa(int(params.Offset)) + ";"

	//  Проверяем, есть ли параметры поиска и заполняем переменную tsquery
	var tsquery string
	if params.Hight != nil {
		tsquery = tsquery + " " + *params.Hight
	}
	if params.Kind != nil {
		tsquery = tsquery + " " + *params.Kind
	}
	if params.RecommendPosition != nil {
		tsquery = tsquery + " " + *params.RecommendPosition
	}
	if params.RegardToLight != nil {
		tsquery = tsquery + " " + *params.RegardToLight
	}
	if params.RegardToMoisture != nil {
		tsquery = tsquery + " " + *params.RegardToMoisture
	}
	if params.FloweringTime != nil {
		tsquery = tsquery + " " + *params.FloweringTime
	}
	if params.Classifiers != nil {
		tsquery = tsquery + " " + *params.Classifiers
	}
	// Проверяем заполнена ли переменная tsquery
	if tsquery != "" {
		// Проверяем заполнена ли категория
		if params.Category != nil {
			query = query + " where to_tsvector('russian', short_info) @@ plainto_tsquery('russian','" +
				tsquery + "') and category_id = " + strconv.Itoa(int(*params.Category))
		}
	} else {
		if params.Category != nil {
			query = query + " where category_id = " + strconv.Itoa(int(*params.Category))
		}
	}

	// Добавляем лимит и оффсет
	query += limitoffset

	// Получение строк с базы
	rows, err := pg.db.Query(ctx, query)
	if err != nil {
		return nil, errors.WithMessage(err, "Select rows error: ")
	}
	defer rows.Close()
	for rows.Next() {
		var plant models.Plant
		err = rows.Scan(&plant.ID, &plant.Title, &plant.Category, &plant.ShortInfo, &plant.Infos,
			&plant.Images, &plant.Creater, &plant.CreatedAt, &plant.Modifier, &plant.ModifiedAt)
		if err != nil {
			return nil, errors.WithMessage(err, "Scan rows error: ")
		}
		plants = append(plants, &plant)
	}

	return plants, err
}

func (pg *PG) GetPlantByID(ctx context.Context, id int64) (*models.Plant, error) {
	query := `select id, title, category_id, short_info::jsonb, notes::jsonb,
			img_links::jsonb, creator, created_at, modifier, modified_at
			from reference.plants where id= `
	query = query + strconv.Itoa(int(id)) + ";"
	var plant models.Plant
	err := pg.db.QueryRow(ctx, query).Scan(&plant.ID, &plant.Title, &plant.Category, &plant.ShortInfo, &plant.Infos,
		&plant.Images, &plant.Creater, &plant.CreatedAt, &plant.Modifier, &plant.ModifiedAt)
	if err != nil {
		return nil, errors.WithMessage(err, "QueryRow.Scan error")
	}

	return &plant, err
}
