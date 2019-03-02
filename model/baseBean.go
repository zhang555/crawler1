package model

import "time"

type ModelTime struct {
	CreateTime JSONTime `json:"createTime"`
	UpdateTime JSONTime `json:"updateTime"`
}

type JSONTime time.Time
