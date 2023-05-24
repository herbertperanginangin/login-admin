package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Save(transaction Transaksi) (Transaksi, error)
	FindAll() ([]Transaksi, error)
	FindByID(ID int) (Transaksi, error)
	Update(transaksi Transaksi) (Transaksi, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) Save(transaksi Transaksi) (Transaksi, error) {
	err := r.db.Create(&transaksi).Error

	if err != nil {
		return transaksi, err
	}

	return transaksi, nil
}

func (r *repository) FindAll() ([]Transaksi, error) {
	var transactions []Transaksi

	err := r.db.Preload("User").Preload("Acara").Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) FindByID(ID int) (Transaksi, error) {
	var trx Transaksi
	err := r.db.Find(&trx, ID).Error
	return trx, err
}

func (r *repository) Update(transaksi Transaksi) (Transaksi, error) {
	err := r.db.Save(&transaksi).Error
	if err != nil {
		return transaksi, err
	}
	return transaksi, nil
}
