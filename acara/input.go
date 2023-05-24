package acara

import (
	"time"
)

type GetAcaraDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateAcaraInput struct {
	NamaAcara    string    `form:"nama_acara" binding:"required"`
	Harga        int       `form:"harga" binding:"required"`
	TanggalAcara time.Time `form:"tanggal_acara" binding:"required" time_format:"2006-01-02"`
	Deskripsi    string    `form:"deskripsi" binding:"required"`
	ImgUrl       string    `form:"img_url" binding:"required"`
}

type FormCreateAcaraInput struct {
	NamaAcara    string    `form:"nama_acara" binding:"required"`
	Harga        int       `form:"harga" binding:"required"`
	TanggalAcara time.Time `form:"tanggal_acara" binding:"required" time_format:"2006-01-02"`
	Deskripsi    string    `form:"deskripsi" binding:"required"`
}

type FormCreateAcaraTemplateData struct {
	FormCreateAcaraInput
	Error string
}

type FormUpdateAcaraInput struct {
	ID           int
	NamaAcara    string    `form:"nama_acara" binding:"required"`
	Harga        int       `form:"harga" binding:"required"`
	TanggalAcara time.Time `form:"tanggal_acara" binding:"required" time_format:"2006-01-02"`
	Deskripsi    string    `form:"deskripsi" binding:"required"`
	ImgUrl       string    `form:"img_url"`
	Error        error
}
type FormUpdateAcaraTemplateData struct {
	FormUpdateAcaraInput
	Error string
}
