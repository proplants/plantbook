package model

import (
	"encoding/json"
	"fmt"
)

// Plant ...
//easyjson:json
type Plant struct {
	Source    string    `json:"source"`
	Title     string    `json:"title"`
	Category  string    `json:"category"`
	ShortInfo ShortInfo `json:"short_info"`
	Images    Images    `json:"images"`
	Info      Infos     `json:"info"`
	Metadata  Metadata  `json:"metadata"`
}

// String implement stringer interface.
func (p *Plant) String() string {
	bts, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		return fmt.Sprintf("failed %s", err)
	}
	return string(bts)
}

//easyjson:json
// Images ...
type Images []string

//easyjson:json
// Plants ...
type Plants []Plant

//easyjson:json
// Info ...
type Info struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

//easyjson:json
// Infos ...
type Infos []Info

//easyjson:json
// Metadata ...
type Metadata struct {
	DateCollect string `json:"date_collect"`
	Target      string `json:"target"`
}

// ShortInfo ...
type ShortInfo struct {
	Kind              string `json:"kind,omitempty"`
	RecommendPosition string `json:"recommend_position,omitempty"`
	RegardToLight     string `json:"regard_to_light,omitempty"`
	RegardToMoisture  string `json:"regard_to_moisture,omitempty"`
	FloweringTime     string `json:"flowering_time,omitempty"`
	Hight             string `json:"hight,omitempty"`
	Classifiers       string `json:"classifiers,omitempty"`
	Ground            string `json:"ground,omitempty"`
	Wintering         string `json:"wintering,omitempty"`
	Decorativeness    string `json:"decorativeness,omitempty"`
	Composition       string `json:"composition,omitempty"`
	Shearing          string `json:"shearing,omitempty"`
	Growing           string `json:"growing,omitempty"`
	Eating            string `json:"eating,omitempty"`
}

// RefPagePlants - .
type RefPagePlants struct {
	Name       string
	URL        string
	FileName   string
	CategoryID int32
}
