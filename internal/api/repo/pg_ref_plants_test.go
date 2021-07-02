// +build integration

package repo_test

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/proplants/plantbook/internal/api/models"
	"github.com/proplants/plantbook/internal/api/repo"
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
		Hight:             "высокое (выше 100 см), среднее (50-100 см), низкое (10-50 см)",
		FloweringTime:     "весна, лето, осень",
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
			&models.RefPlant{ID: 1, Category: 1, Title: &titleTestRefPlant, ShortInfo: &shortInfoTestRefPlant},
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
			if reflect.DeepEqual(got, tt.wantGet) && tt.wantErr {
				t.Errorf("TestPG_GetRefPlantByID DeepEqual error: \ngot = %v, \nwant = %v", got, tt.wantGet)
			}
		})
	}
}
