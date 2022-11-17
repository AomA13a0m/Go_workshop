package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Data        []DogsRes `json:"data"`
	Name        string    `json:"name"`
	Count       int       `json:"count"`
	Sum_red     int       `json:"sum_red"`
	Sum_green   int       `json:"sum_green"`
	Sum_pink    int       `json:"sum_pink"`
	Sum_nocolor int       `json:"sum_nocolor"`
}

type Company struct {
	gorm.Model
	Company_Name    string `json:"company_name"`
	Company_Address string `json:"company_address"`
	Company_Mail    string `json:"company_mail"`
	Company_Phone   int    `json:"company_phone"`
}

type Users struct {
	gorm.Model
	Employee_id string `json:"employee_id"`
	Name        string `json:"name"`
	LastName    string `json:"lastname"`
	Birthday    string `json:"birthday"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	Tel         int    `json:"tel"`
}

type UsersRes struct {
	gorm.Model
	Employee_id string `json:"employee_id"`
	Name        string `json:"name"`
	LastName    string `json:"lastname"`
	Birthday    string `json:"birthday"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	Tel         int    `json:"tel"`
	Generation  string `json:"generation"`
}

type User_Generation struct {
	Data          []UsersRes `json:"data"`
	Count         int         `json:"count"`
	GenZ          int         `json:"genZ"`
	GenY          int         `json:"genY"`
	GenX          int         `json:"genX"`
	Baby_Boomer   int         `json:"baby_boomer"`
	GI_Generation int         `json:"gi_generation"`
}
