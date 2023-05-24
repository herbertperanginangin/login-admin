package transaction

import (
	"umiEvient/acara"
	"umiEvient/user"
)

type Transaksi struct {
	ID      int
	UserID  int
	AcaraID int
	Total   int
	Status  string
	User    user.User
	Acara   acara.Acara
}
