package repo

import (
	"context"
	"strconv"

	"github.com/kaatinga/plantbook/internal/api/models"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/refplant"
	"github.com/pkg/errors"
)

// GetRefPlants получение растений из справочника по параметрам или просто все, обязательные поля limit and offset.
func (pg *PG) GetRefPlants(ctx context.Context, params refplant.GetRefPlantsParams) ([]*models.RefPlant, error) {
	query := `select id, title, category_id, short_info::jsonb, notes::jsonb,
			img_links::jsonb, creator, created_at, modifier, modified_at
			from reference.plants`
	var refPlants []*models.RefPlant
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

// GetRefPlantByID extracts reference.plant by specified id.
func (pg *PG) GetRefPlantByID(ctx context.Context, id int64) (*models.RefPlant, error) {
	query := `select id, title, category_id, short_info::jsonb, notes::jsonb,
			img_links::jsonb, creator, created_at, modifier, modified_at
			from reference.plants where id= `
	query = query + strconv.Itoa(int(id)) + ";"
	var refPlant models.RefPlant
	err := pg.db.QueryRow(ctx, query).Scan(&refPlant.ID, &refPlant.Title, &refPlant.Category,
		&refPlant.ShortInfo, &refPlant.Infos, &refPlant.Images, &refPlant.Creater, &refPlant.CreatedAt,
		&refPlant.Modifier, &refPlant.ModifiedAt)
	if err != nil {
		return nil, errors.WithMessage(err, "QueryRow.Scan error")
	}

	return &refPlant, err
}
