package model

import (
	"time"
)

type Model struct {
	CreateAt time.Time	`json:"create_at"`
	UpdateAt time.Time	`json:"update_at"`
}
