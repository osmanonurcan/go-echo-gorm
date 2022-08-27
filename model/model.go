package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// type Plan struct {
// 	gorm.Model
// 	Name       string `json:"name"`
// 	Day        string `json:"day"`
// 	StartTime  string `json:"start_time"`
// 	FinishTime string `json:"finish_time"`
// 	State      string `json:"state"`
// }

type Plan struct {
	gorm.Model
	Name       string    `json:"name"`
	StartTime  time.Time `json:"start_time"`
	FinishTime time.Time `json:"finish_time"`
	State      string    `json:"state"`
}
