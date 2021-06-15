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
	Kind              string `json:"kind"`
	RecommendPosition string `json:"recommend_position"`
	RegardToLight     string `json:"regard_to_light"`
	RegardToMoisture  string `json:"regard_to_moisture"`
	FloweringTime     string `json:"flowering_time"`
	Hight             string `json:"hight"`
	Classifiers       string `json:"classifiers"`
}
