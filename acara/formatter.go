package acara

import (
	"strconv"
	"time"
)

type AcaraResponse struct {
	ID            int       `json:"id"`
	Nama_acara    string    `json:"nama_acara"`
	Img_url       string    `json:"img_url"`
	Harga         string    `json:"harga"`
	Tanggal_acara time.Time `json:"tanggal_acara"`
	Deskripsi     string    `json:"deskripsi"`
}

func FormatAcara(ka Acara) AcaraResponse {
	return AcaraResponse{
		ID:            ka.ID,
		Nama_acara:    ka.Nama_acara,
		Harga:         strconv.Itoa(ka.Harga),
		Tanggal_acara: ka.Tanggal_acara,
		Deskripsi:     ka.Deskripsi,
		Img_url:       ka.ImgURL,
	}
}
