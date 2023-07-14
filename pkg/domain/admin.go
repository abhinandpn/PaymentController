package domain

type Admin struct {
	Id       uint   `json:"id" gorm:"primary key"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
