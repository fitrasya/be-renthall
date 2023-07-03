package model

import "time"

type PesananDetail struct {
	Id             string     `json:"id"`
	GedungId       string     `json:"gedung_id"`
	GedungNama     string     `json:"gedung_nama"`
	PemesanId      string     `json:"pemesan_id"`
	PemesanNama    string     `json:"pemesan_nama"`
	PemesanKontak  string     `json:"pemesan_kontak"`
	TanggalMulai   string     `json:"tanggal_mulai"`
	TanggalSelesai string     `json:"tanggal_selesai"`
	HargaDeal      string     `json:"harga_deal"`
	StatusId       string     `json:"status_id"`
	StatusNama     string     `json:"status_nama"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

type Pesanan struct {
	Id             string     `gorm:"type:uuid;primary_key" json:"id"`
	GedungId       string     `json:"gedung_id" form:"gedung_id" validate:"required"`
	PemesanId      string     `json:"pemesan_id" form:"pemesan_id" validate:"required"`
	TanggalMulai   string     `json:"tanggal_mulai" form:"tanggal_mulai" validate:"required"`
	TanggalSelesai string     `json:"tanggal_selesai" form:"tanggal_selesai" validate:"required"`
	HargaDeal      string     `json:"harga_deal" form:"harga_deal" validate:"required"`
	StatusId       string     `json:"status_id" form:"status_id" validate:"required"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

func (Pesanan) TableName() string {
	return "pesanan"
}

type PutPesanan struct {
	Id             string `json:"id" form:"id" validate:"required"`
	GedungId       string `json:"gedung_id" form:"gedung_id" validate:"required"`
	PemesanId      string `json:"pemesan_id" form:"pemesan_id" validate:"required"`
	TanggalMulai   string `json:"tanggal_mulai" form:"tanggal_mulai" validate:"required"`
	TanggalSelesai string `json:"tanggal_selesai" form:"tanggal_selesai" validate:"required"`
	HargaDeal      string `json:"harga_deal" form:"harga_deal" validate:"required"`
	StatusId       string `json:"status_id" form:"status_id" validate:"required"`
}

type DelPesanan struct {
	Id        string     `json:"id" form:"id" validate:"required"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (DelPesanan) TableName() string {
	return "pesanan"
}

type PutPesananStatus struct {
	Id       string `json:"id" form:"id" validate:"required"`
	StatusId string `json:"status_id" form:"status_id" validate:"required"`
}

type Dashboard struct {
	Diajukan  string `json:"diajukan"`
	Disetujui string `json:"disetujui"`
	Ditolak   string `json:"ditolak"`
	Gedung    string `json:"gedung"`
	Pemesan   string `json:"pemesan"`
	Pesanan   string `json:"pesanan"`
}

type Schedule struct {
	Id      string `json:"id"`
	Gedung  string `json:"gedung"`
	Pemesan string `json:"pemesan"`
	Kontak  string `json:"kontak"`
	Mulai   string `json:"mulai"`
	Selesai string `json:"selesai"`
}
