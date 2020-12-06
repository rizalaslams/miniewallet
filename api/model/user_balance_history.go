package model

import (
	"time"
)

type UserBalanceHistory struct {
	ID            uint64    `json:"id"`
	UserBalanceID uint64    `json:"user_balance_id"`
	BalanceBefore int       `json:"balance_before"`
	BalanceAfter  int       `json:"balance_after"`
	Activity      string    `json:"activity"`
	Type          string    `json:"type"`
	IP            string    `json:"ip"`
	Location      string    `json:"location"`
	UserAgent     string    `json:"user_agent"`
	Author        string    `json:"author"`
	CreatedAt     time.Time `json:"created_at"`
}

type UserBalanceHistoryTable interface {
	TableName() string
}

func (UserBalanceHistory) TableName() string {
	return "user_balance_history"
}

func (s *Server) CreateUserBalanceHistory(userBalanceHistory *UserBalanceHistory) *UserBalanceHistory {
	err := s.DB.Debug().Create(&userBalanceHistory).Error
	if err != nil {
		return userBalanceHistory
	}

	return nil
}

func (s *Server) CountUserBalanceHistory(userId uint64) int64 {
	ub := &UserBalance{}
	count := s.DB.Debug().Where("user_id = ?", userId).Find(&ub).RowsAffected

	return count

}
