package models

type User struct {
	Name     string
	Username string
	Email    string
	Password string
}

var User1 = User{
	Name:     "Dwi Wahyudi",
	Username: "dwiw",
	Email:    "dwiwahyudi1996@gmail.com",
	Password: "secret",
}
