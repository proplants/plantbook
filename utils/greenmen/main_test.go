package main

import (
	"testing"

	"github.com/kaatinga/plantbook/utils/greenmen/model"
)

var roomPlant model.Plant = model.Plant{
	Title: "plant",
	ShortInfo: model.ShortInfo{
		Kind:              "kind",
		RecommendPosition: "pos",
		RegardToLight:     "ok",
		RegardToMoisture:  "bed",
		FloweringTime:     "spring",
		Hight:             "ft",
		Classifiers:       "test",
	},
	Images: []string{"http://localhost/plant.png"},
	Info: []model.Info{
		{Title: "ititle", Content: "content"},
	},
}

func Test_makeSQLInsert(t *testing.T) {
	type args struct {
		plant    model.Plant
		template string
	}
	shortInfo := `{"kind":"kind","recommend_position":"pos","regard_to_light":"ok","regard_to_moisture":"bed","flowering_time":"spring","hight":"ft","classifiers":"test"}`
	notes := `[{"title":"ititle","content":"content"}]`
	imagLinks := `["http://localhost/plant.png"]`
	tests := []struct {
		name string
		args args
		want string
	}{
		{"roomplant_insert", args{plant: roomPlant, template: templateRoomPlants}, `INSERT INTO reference.plants (title, category_id, short_info, notes, img_links, created_at, creator) VALUES('plant', 1, '` + shortInfo + `'::jsonb, '` + notes + `'::jsonb, '` + imagLinks + `'::jsonb, CURRENT_TIMESTAMP, CURRENT_USER);`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeSQLInsert(tt.args.plant, tt.args.template); got != tt.want {
				t.Errorf("makeSQLInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
