package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string    `gorm:"type:varchar(100);not null;unique" json:"username" binding:"required,max=100,min=3"`
	Email     string    `gorm:"type:varchar(100);not null;unique" json:"email" binding:"required,email"`
	Password  string    `gorm:"type:varchar(100);not null" json:"password" binding:"required,max=50,min=8"`
	Age       int       `gorm:"type:int;not null" json:"age" binding:"required,gte=21,lte=60"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

type Login struct {
	UserName string `json:"username" binding:"required,max=100,min=3"`
	Password string `json:"password" binding:"required,max=50,min=8"`
}
type UserOut struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserName string `gorm:"type:varchar(100);not null;unique" json:"username" binding:"required,max=100,min=3"`
	Email    string `gorm:"type:varchar(100);not null;unique" json:"email" binding:"required,email"`
	Age      int    `gorm:"type:int;not null" json:"age" binding:"required,gte=21,lte=60"`
}

func (u *User) CreateUser(db *gorm.DB) (*User, error) {
	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (l *Login) GetUserByUserName(userName string, db *gorm.DB) (*User, error) {
	var user User
	err := db.Debug().Where("user_name = ?", userName).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}
