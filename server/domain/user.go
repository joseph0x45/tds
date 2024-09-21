package domain

type User struct {
	ID       string `json:"id" db:"id"`
	Phone    string `json:"phone" db:"phone"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
