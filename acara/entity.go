package acara

import "time"

type Acara struct {
	ID            int
	Nama_acara    string
	Harga         int
	Tanggal_acara time.Time
	Deskripsi     string
	ImgURL        string
}
