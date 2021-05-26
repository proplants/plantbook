package model

import (
	"encoding/json"
	"fmt"
)

type Plant struct {
	Source    string    `json:"source"`
	Title     string    `json:"title"`
	Category  string    `json:"category"`
	ShortInfo ShortInfo `json:"short_info"`
	Images    []string  `json:"images"`
	Info      []Info    `json:"info"`
	Metadata  Metadata  `json:"metadata"`
}

func (p *Plant) String() string {
	bts, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		return fmt.Sprintf("failed %s", err)
	}
	return string(bts)
}

//easyjson:json
type Plants []Plant

type Info struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Metadata struct {
	DateCollect string `json:"date_collect"`
	Target      string `json:"target"`
}

type ShortInfo struct {
	Kind              string `json:"kind"`
	RecommendPosition string `json:"recommend_position"`
	RegardToLight     string `json:"regard_to_light"`
	RegardToMoisture  string `json:"regard_to_moisture"`
	FloweringTime     string `json:"flowering_time"`
	Hight             string `json:"hight"`
	Classifiers       string `json:"classifiers"`
}
