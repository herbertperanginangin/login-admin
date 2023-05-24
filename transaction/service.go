package transaction

type service struct {
	repository Repository
}

type Service interface {
	CreateTransaction(input CreateTransactionInput) (Transaksi, error)
	GetAllTransactions() ([]Transaksi, error)
	UpdateStatus(ID int) (Transaksi, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaksi, error) {
	transaction := Transaksi{}
	transaction.AcaraID = input.AcaraID
	transaction.UserID = input.UserID
	transaction.Total = input.Total
	transaction.Status = "pending"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *service) GetAllTransactions() ([]Transaksi, error) {
	transactions, err := s.repository.FindAll()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) UpdateStatus(ID int) (Transaksi, error) {
	transaksi, _ := s.repository.FindByID(ID)
	transaksi.Status = "Success"
	newTransaksi, err := s.repository.Update(transaksi)
	return newTransaksi, err
}
