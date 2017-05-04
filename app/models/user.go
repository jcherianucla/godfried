package model

import (
	"fmt"
	utils "github.com/jcherianucla/godfried/app/helpers"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Username  string
	Password  []byte
	Email     string
}

func (user *User) HashPassword(pass string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	utils.CheckError(err)
	user.Password = hash
}
