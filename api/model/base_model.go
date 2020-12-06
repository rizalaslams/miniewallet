package model

import (
	"fmt"
	"miniewallet/api/auth"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB *gorm.DB
}

var (
	Model modelInterface = &Server{}
)

type modelInterface interface {
	//db initialization
	Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error)

	//user methods
	ValidateEmail(string) error
	VerifyPassword(string, string) error
	CreateUser(*User) (*User, error)
	GetUserByEmail(string) (*User, error)

	//auth methods:
	FetchAuth(*auth.AuthDetails) (*Auth, error)
	DeleteAuth(*auth.AuthDetails) error
	CreateAuth(uint64) (*Auth, error)
}

func (s *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	s.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}

	var users = []User{
		{
			Username: "User 1",
			Email:    "user1@test.com",
			Password: "112233aa",
		},
		{
			Username: "User 2",
			Email:    "user2@test.com",
			Password: "112233aa",
		},
	}

	for i, _ := range users {
		users[i].Password = "adsad"
		err = s.DB.Debug().Model(User{}).FirstOrCreate(&users[i]).Error
	}

	return s.DB, nil
}
