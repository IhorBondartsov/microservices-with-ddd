package models

import (
	"encoding/json"
)

type Place struct {
	Name        string        `json:"name"`
	City        string        `json:"city"`
	Country     string        `json:"country"`
	Alias       []interface{} `json:"alias"`
	Regions     []interface{} `json:"regions"`
	Coordinates []float64     `json:"coordinates"`
	Province    string        `json:"province"`
	Timezone    string        `json:"timezone"`
	Unlocs      []string      `json:"unlocs"`
	Code        string        `json:"code"`
}

func (p *Place) UnmarshalJSON(b []byte) error{
	type place Place
	internal := &struct {
		*place
	}{
		(*place)(p),
	}
	err := json.Unmarshal(b, internal)
	if err != nil {
		return err
	}
	return nil
}

