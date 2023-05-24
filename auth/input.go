package auth

type TokenUserInput struct {
	ID    int    `json:"id" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Nama  string `json:"nama" binding:"required"`
}
