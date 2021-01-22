package database

import (
	"go-fiber-gorm/config"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_time"` // CreatedAt gorm will generate the value when insert
	UpdatedAt time.Time `json:"updated_time"` // UpdatedAt gorm will generate the value when insert
}

func db() *gorm.DB {
	return config.DB
}
