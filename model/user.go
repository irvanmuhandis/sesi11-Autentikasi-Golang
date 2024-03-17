package model

import (
	"sesi11/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Fullname string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your Name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your Email is required, email~Invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your Password is required,minstringlength(6)~Password have to minimal lenth of 6 char"`
	Products []Product `gorm:"constaint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
