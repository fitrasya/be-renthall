package model

import "time"

type Pemesan struct {
	Id        string     `gorm:"type:uuid;primary_key" json:"id"`
	Nama      string     `json:"nama" form:"nama" validate:"required"`
	Kontak    string     `json:"kontak" form:"kontak" validate:"required"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (Pemesan) TableName() string {
	return "pemesan"
}

type PutPemesan struct {
	Id     string `json:"id" form:"id" validate:"required"`
	Nama   string `json:"nama" form:"nama" validate:"required"`
	Kontak string `json:"kontak" form:"kontak" validate:"required"`
}

type DelPemesan struct {
	Id        string     `json:"id" form:"id" validate:"required"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (DelPemesan) TableName() string {
	return "pemesan"
}
