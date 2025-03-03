package models

import "time"

type (
	Features struct {
		ID            int       `json:"id"`
		Year          int       `json:"year"`
		Month         int       `json:"month"`
		Day           int       `json:"day"`
		Hour          int       `json:"hour"`
		Precipitation int       `json:"precipitation"`
		TempMax       float32   `json:"tempMax"`
		TempMin       float32   `json:"tempMin"`
		CreatedAt     time.Time `json:"created_at"`
		UpdateAt      time.Time `json:"update_at"`
	}

	FeaturesDays struct {
		ID            int       `json:"id"`
		Year          int       `json:"year"`
		Month         int       `json:"month"`
		Day           int       `json:"day"`
		Precipitation int       `json:"precipitation"`
		TempMax       float32   `json:"tempMax"`
		TempMin       float32   `json:"tempMin"`
		CreatedAt     time.Time `json:"created_at"`
		UpdateAt      time.Time `json:"update_at"`
	}

	Target struct {
		ID           int       `json:"ID"`
		PeriodUnit   string    `json:"period_unit"`
		FutureOffset int       `json:"future_offset"`
		FutureValue  int       `json:"future_value"`
		CreatedAt    time.Time `json:"created_at"`
		UpdateAt     time.Time `json:"update_at"`
	}
)
