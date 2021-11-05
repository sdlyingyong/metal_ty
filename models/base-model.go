package models

import "time"

//default response style
type BaseModel struct {
	Id		int 	`json:"id"`
	CreatedAt		time.Time 	`json:"created_at"`
	UpdatedAt		time.Time 	`json:"updated_at"`
}
