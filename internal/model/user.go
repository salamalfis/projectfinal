package model

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username" gorm:"unique" validate:"required"`
	Email     string    `json:"email" gorm:"unique" validate:"email,required"`
	Password  string    `json:"-" validate:"required " gorm:"column:password" `
	Age       int       `json:"age" validate:"required" gorm:"type:int column:age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserMediaSocial struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
// https://gin-gonic.com/docs/examples/binding-and-validation/
type UserSignUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserItem struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (u UserSignUp) Validate() error {
	// check username
	if u.Username == "" {
		return errors.New("invalid username")
	}
	if len(u.Password) < 6 {
		return errors.New("invalid password")
	}
	if u.Age < 8 {
		return errors.New("Minimal age 8 years old")
	}
	return nil
}

func (u User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
