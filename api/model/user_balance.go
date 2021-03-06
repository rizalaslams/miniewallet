package model

import (
	"time"
)

type UserBalance struct {
	ID             uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID         uint64    `gorm:"size:255;not null;unique" json:"user_id"`
	Balance        int       `gorm:"size:100;not null;unique" json:"email"`
	BalanceAchieve int       `gorm:"size:100;not null;" json:"password"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserBalanceTable interface {
	TableName() string
}

func (UserBalance) TableName() string {
	return "user_balance"
}

func (s *Server) UpdateUserBalance(userBalance *UserBalance) *UserBalance {
	topUp := s.DB.Debug().Model(&userBalance).Where("user_id = ?", userBalance.UserID).Updates(UserBalance{
		Balance:        userBalance.Balance,
		BalanceAchieve: userBalance.Balance,
	})
	if topUp != nil {
		return userBalance
	}

	return nil
}

func (s *Server) GetUserBalance(userid uint64) (*UserBalance, error) {
	balance := &UserBalance{}
	err := s.DB.Debug().Where("user_id = ?", userid).Take(&balance).Error
	if err != nil {
		return nil, err
	}
	return balance, nil
}
