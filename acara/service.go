package acara

import (
	"fmt"
)

type Service interface {
	GetAllAcaras() ([]Acara, error)
	GetAcaraByID(input GetAcaraDetailInput) (Acara, error)
	CreateAcara(input CreateAcaraInput) (Acara, error)
	UpdateAcara(inputID GetAcaraDetailInput, inputData CreateAcaraInput) (Acara, error)
	DeleteAcara(inputID GetAcaraDetailInput) (Acara, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllAcaras() ([]Acara, error) {
	acaras, err := s.repository.FindAll()
	if err != nil {
		return acaras, err
	}

	return acaras, nil
}

func (s *service) GetAcaraByID(input GetAcaraDetailInput) (Acara, error) {
	acara, err := s.repository.FindByID(input.ID)

	if err != nil {
		return acara, err
	}

	return acara, nil
}

func (s *service) CreateAcara(input CreateAcaraInput) (Acara, error) {
	acara := Acara{}
	acara.Nama_acara = input.NamaAcara
	acara.Harga = input.Harga
	acara.Tanggal_acara = input.TanggalAcara
	acara.Deskripsi = input.Deskripsi
	acara.ImgURL = input.ImgUrl

	newAccara, err := s.repository.Save(acara)
	fmt.Println(newAccara, "masuk")
	if err != nil {
		return newAccara, err
	}

	return newAccara, nil
}

func (s *service) UpdateAcara(inputID GetAcaraDetailInput, inputData CreateAcaraInput) (Acara, error) {
	acara, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return acara, err
	}

	acara.Nama_acara = inputData.NamaAcara
	acara.Harga = inputData.Harga
	acara.Tanggal_acara = inputData.TanggalAcara
	acara.Deskripsi = inputData.Deskripsi
	acara.ImgURL = inputData.ImgUrl

	updatedCampaign, err := s.repository.Update(acara)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (s *service) DeleteAcara(inputID GetAcaraDetailInput) (Acara, error) {
	acara, err := s.repository.FindByID(inputID.ID)
	newacara, err := s.repository.Delete(acara)
	return newacara, err
}
