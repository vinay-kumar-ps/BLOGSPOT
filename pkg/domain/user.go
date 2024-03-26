package domain

type User struct {
	ID   int `json:"id"`
	Name string  `json:"name"`
	Email  string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Phone  string `json:"phone"`
}
