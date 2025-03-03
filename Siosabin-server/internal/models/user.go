package models

import "time"

type Data struct {
	ID        int       `json:"id"`
	DeviceID  int       `json:"device_id"`
	Location  string    `json:"location"`
	InTank    int       `json:"in_tank"`
	OutTank   int       `json:"out_tank"`
	Salinity  int       `json:"salinity"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
