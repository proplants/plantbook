package repo

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/proplants/plantbook/internal/api/models"
	"github.com/proplants/plantbook/internal/api/restapi/operations/refplant"
)

// GetRefPlans extracts all plants from the reference by parameters.
func (pg *PG) GetRefPlants(ctx context.Context,
	params refplant.GetRefPlantsParams) ([]*models.RefPlant, int64, int64, error) {
	query := `SELECT id, title, category_id, short_info::jsonb, notes::jsonb,
			img_links::jsonb, creator, created_at, modifier, modified_at
			FROM reference.plants`
	totalquery := `select count(1) as cnt from reference.plants`
	var refPlants []*models.RefPlant
	var tsquery string
	paramsArr := [7]string{
		params.Classifiers, params.Hight, params.Kind, params.RecommendPosition,
		params.RegardToLight, params.RegardToMoisture, params.FloweringTime,
	}
	for _, param := range paramsArr {
		if param != "" {
			tsquery += param + " "
		}
	}
	var argsQuery []interface{}
	if tsquery != "" {
		query += " WHERE to_tsvector('russian', short_info) @@ plainto_tsquery('russian', $3)"
		totalquery += " WHERE to_tsvector('russian', short_info) @@ plainto_tsquery('russian', $1)"
		argsQuery = append(argsQuery, tsquery)
		if params.Category != 0 {
			query += " AND category_id = $4"
			totalquery += " AND category_id = $2"
			argsQuery = append(argsQuery, strconv.Itoa(int(params.Category)))
		}
	} else if params.Category != 0 {
		query += " WHERE category_id = $3"
		totalquery += " WHERE category_id = $1"
		argsQuery = append(argsQuery, strconv.Itoa(int(params.Category)))
	}
	totalquery += ";"
	query += " ORDER BY title LIMIT $1 OFFSET $2;"

	var argsQueryAll []interface{}
	argsQueryAll = append(argsQueryAll, params.Limit, params.Offset)
	argsQueryAll = append(argsQueryAll, argsQuery...)

	rows, err := pg.db.Query(params.HTTPRequest.Context(), query, argsQueryAll...)
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
	if count == 0 {
		return nil, 0, 0, errors.Errorf("Not found rows")
	}
	err = pg.db.QueryRow(ctx, totalquery, argsQuery...).Scan(&total)
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.Errorf("Not found rows")
		}
		return nil, errors.WithMessage(err, "QueryRow.Scan error")
	}

	return &refPlant, err
}
