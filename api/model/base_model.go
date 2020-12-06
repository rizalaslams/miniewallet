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

	//bank methods
	GetBankByID(uint64) (*BlanceBank, error)
	UpdateBankBalance(*BlanceBank) *BlanceBank
	CreateBankBalanceHistory(*BlanceBankHistory) *BlanceBankHistory

	//topup methods:
	GetUserBalance(uint64) (*UserBalance, error)
	UpdateUserBalance(*UserBalance) *UserBalance
	CreateUserBalanceHistory(*UserBalanceHistory) *UserBalanceHistory
	CountUserBalanceHistory(uint64) int64

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
			ID:       1,
			Username: "User 1",
			Email:    "user1@test.com",
			Password: "112233aa",
		},
		{
			ID:       2,
			Username: "User 2",
			Email:    "user2@test.com",
			Password: "112233aa",
		},
	}

	var users_balance = []UserBalance{
		{
			UserID:         1,
			Balance:        0,
			BalanceAchieve: 0,
		},
		{
			UserID:         2,
			Balance:        0,
			BalanceAchieve: 0,
		},
	}

	var blance_bank = []BlanceBank{
		{
			ID:             1,
			Balance:        10000000,
			BalanceAchieve: 10000000,
			Code:           "BRI",
			Enable:         "TRUE",
		},
		{
			ID:             2,
			Balance:        10000000,
			BalanceAchieve: 10000000,
			Code:           "BCA",
			Enable:         "TRUE",
		},
	}

	for i, _ := range users {
		users[i].Password = "adsad"
		err = s.DB.Debug().Model(User{}).FirstOrCreate(&users[i]).Error
	}

	for i, _ := range users_balance {
		err = s.DB.Debug().Model(UserBalance{}).FirstOrCreate(&users_balance[i]).Error
	}

	for i, _ := range blance_bank {
		err = s.DB.Debug().Model(BlanceBank{}).FirstOrCreate(&blance_bank[i]).Error
	}

	return s.DB, nil
}
