package acara

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Acara, error)
	FindByID(ID int) (Acara, error)
	Save(acara Acara) (Acara, error)
	Update(acara Acara) (Acara, error)
	Delete(acara Acara) (Acara, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Acara, error) {
	var acaras []Acara

	err := r.db.Find(&acaras).Error
	if err != nil {
		return acaras, err
	}

	return acaras, nil
}

func (r *repository) FindByID(ID int) (Acara, error) {
	var acara Acara

	err := r.db.Where("id = ?", ID).Find(&acara).Error

	if err != nil {
		return acara, err
	}

	return acara, nil
}

func (r *repository) Save(acara Acara) (Acara, error) {
	err := r.db.Create(&acara).Error
	if err != nil {
		return acara, err
	}

	return acara, nil
}

func (r *repository) Update(acara Acara) (Acara, error) {
	err := r.db.Save(&acara).Error

	if err != nil {
		return acara, err
	}

	return acara, nil
}

func (r *repository) Delete(acara Acara) (Acara, error) {
	err := r.db.Delete(&acara).Error
	return acara, err
}
