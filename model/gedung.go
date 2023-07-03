package model

import "time"

type Gedung struct {
	Id        string     `gorm:"type:uuid;primary_key" json:"id"`
	Nama      string     `json:"nama" form:"nama" validate:"required"`
	Kapasitas int        `json:"kapasitas" form:"kapasitas" validate:"required"`
	Harga     int        `json:"harga" form:"harga" validate:"required"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (Gedung) TableName() string {
	return "gedung"
}

type PutGedung struct {
	Id        string `json:"id" form:"id" validate:"required"`
	Nama      string `json:"nama" form:"nama" validate:"required"`
	Kapasitas int    `json:"kapasitas" form:"kapasitas" validate:"required"`
	Harga     int    `json:"harga" form:"harga" validate:"required"`
}

type DelGedung struct {
	Id        string     `json:"id" form:"id" validate:"required"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (DelGedung) TableName() string {
	return "gedung"
}
