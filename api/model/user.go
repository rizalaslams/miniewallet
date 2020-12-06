package model

import (
	"errors"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Username    string `gorm:"size:255;not null;unique" json:"username"`
	Email       string `gorm:"size:100;not null;unique" json:"email"`
	Password    string `gorm:"size:100;not null;" json:"password"`
	UserBalance []UserBalance
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type TopUp struct {
	Nominal   int    `json:"nominal"`
	IdBank    uint64 `json:"id_bank"`
	Type      string `json:"type"`
	IP        string `json:"ip"`
	Location  string `json:"location"`
	UserAgent string `json:"user_agent"`
	Author    string `json:"author"`
}

type Transfer struct {
	Nominal   int    `json:"nominal"`
	ToUser    uint64 `json:"to_user"`
	Type      string `json:"type"`
	IP        string `json:"ip"`
	Location  string `json:"location"`
	UserAgent string `json:"user_agent"`
	Author    string `json:"author"`
}

type UserTable interface {
	TableName() string
}

func (User) TableName() string {
	return "user"
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (s *Server) VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *Server) ValidateEmail(email string) error {
	if email == "" {
		return errors.New("required email")
	}
	if email != "" {
		if err := checkmail.ValidateFormat(email); err != nil {
			return errors.New("invalid email")
		}
	}
	return nil
}

func (s *Server) CreateUser(user *User) (*User, error) {
	emailErr := s.ValidateEmail(user.Email)
	if emailErr != nil {
		return nil, emailErr
	}
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	create := s.DB.Debug().Create(&user).Error
	if create != nil {
		return nil, create
	}
	return user, nil
}

func (s *Server) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	err := s.DB.Debug().Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
