// +build integration

package repo_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/proplants/plantbook/internal/api/models"
	"github.com/proplants/plantbook/internal/api/repo"
	"github.com/proplants/plantbook/internal/api/restapi/operations/refplant"
)

const (
	dbURLTemplate      string        = "postgres://plantbook_admin:mypassword@%s/plantbook_admin?sslmode=disable"
	dbDefaultHostPort  string        = "localhost:5435"
	dbHOSTAndPortEnv   string        = "DB_HOST_PORT"
	defaultTestTimeout time.Duration = 30 * time.Second
)

// initRepo creates Repo instance by using environment variable DB_HOST_PORT
// for define host and port db server.
func initRepo(ctx context.Context) (*repo.PG, error) {
	dbHostPort := os.Getenv(dbHOSTAndPortEnv)
	if dbHostPort == "" {
		dbHostPort = dbDefaultHostPort
	}
	return repo.NewPG(ctx, fmt.Sprintf(dbURLTemplate, dbHostPort), false)
}

func TestPG_GetRefPlantByID(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
	defer cancel()
	pg, err := initRepo(ctx)
	if err != nil {
		t.Fatalf("initRepo failed, %s", err)
	}
	titleTestRefPlant := "Абутилон, комнатный клен, канатник (Abutilon)"
	shortInfoTestRefPlant := models.ShortInfo{
		Kind:              "кустарник",
		RecommendPosition: "южные окна, западные и/или восточные окна",
		RegardToLight:     "светолюбивое",
		RegardToMoisture:  "влаголюбивое, предпочитает умеренное увлажнение",
		Hight:             "высокое (выше 100 см), среднее (50-100 см), низкое (10-50 см)",
		FloweringTime:     "весна, лето, осень",
		Classifiers:       "красивоцветущее, декоративнолиственное"}
	tests := []struct {
		name    string
		id      int64
		wantGet *models.RefPlant
		wantErr bool
	}{
		{
			"Getting plant by id from reference table.",
			1,
			&models.RefPlant{ID: 1, Category: 1, Title: &titleTestRefPlant, ShortInfo: &shortInfoTestRefPlant,
				ModifiedAt: strfmt.NewDateTime(), CreatedAt: strfmt.NewDateTime()},
			false,
		},
		{
			"Getting plant by wrong id from reference table.",
			23,
			&models.RefPlant{ID: 1, Category: 1, Title: &titleTestRefPlant, ShortInfo: &shortInfoTestRefPlant},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pg.GetRefPlantByID(ctx, tt.id)
			if err != nil {
				t.Errorf("TestPG_GetRefPlantByID got error = %v", err)
				return
			}
			if got != nil {
				got.Creater = ""
				got.Modifier = nil
				got.Images = nil
				got.Infos = nil
				got.CreatedAt = strfmt.NewDateTime()
				got.ModifiedAt = strfmt.NewDateTime()
			}
			if !reflect.DeepEqual(got, tt.wantGet) != tt.wantErr {
				t.Errorf("TestPG_GetRefPlantByID DeepEqual error: \ngot = %v, \nwant = %v", got, tt.wantGet)
			}
		})
	}
}

func TestPG_GetRefPlants(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
	defer cancel()
	pg, err := initRepo(ctx)
	if err != nil {
		t.Fatalf("initRepo failed, %s", err)
	}

	tests := []struct {
		name    string
		params  refplant.GetRefPlantsParams
		wantErr bool
	}{
		{
			"Getting all plants from reference table.",
			refplant.GetRefPlantsParams{HTTPRequest: &http.Request{}, Limit: 20},
			false,
		},
		{
			"Getting plants by category_id from reference table.",
			refplant.GetRefPlantsParams{HTTPRequest: &http.Request{}, Category: 1, Limit: 20},
			false,
		},
		{
			"Getting plants by category_id and some parameters from reference table.",
			refplant.GetRefPlantsParams{HTTPRequest: &http.Request{}, Category: 1,
				Kind: "кустарник", Hight: "высокое", Limit: 20},
			false,
		},
		{
			"Getting plants by some parameters from reference table.",
			refplant.GetRefPlantsParams{HTTPRequest: &http.Request{},
				Kind: "кустарник", Hight: "высокое", Limit: 20},
			false,
		},
		{
			"Getting plants by non-existent category_id from reference table.",
			refplant.GetRefPlantsParams{HTTPRequest: &http.Request{}, Category: 99, Limit: 20},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotCount, _, err := pg.GetRefPlants(ctx, tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestPG_GetRefPlants got error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (gotCount <= 0) != tt.wantErr {
				t.Errorf("TestPG_GetRefPlants, gotCount = %v, wantErr %v", gotCount, tt.wantErr)
				return
			}

		})
	}
}
