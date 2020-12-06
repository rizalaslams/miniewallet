package model

import (
	"time"
)

type BlanceBankHistory struct {
	ID            uint64    `json:"id"`
	BalanceBankID uint64    `json:"balance_bank_id"`
	BalanceBefore uint64    `json:"balance_before"`
	BalanceAfter  uint64    `json:"balance_after"`
	Activity      string    `json:"activity"`
	Type          string    `json:"type"`
	IP            string    `json:"ip"`
	Location      string    `json:"location"`
	UserAgent     string    `json:"user_agent"`
	Author        string    `json:"author"`
	CreatedAt     time.Time `json:"created_at"`
}

type BlanceBankHistoryTable interface {
	TableName() string
}

func (BlanceBankHistory) TableName() string {
	return "blance_bank_history"
}

func (s *Server) CreateBankBalanceHistory(blanceBankHistory *BlanceBankHistory) *BlanceBankHistory {
	err := s.DB.Debug().Create(&blanceBankHistory).Error
	if err != nil {
		return blanceBankHistory
	}

	return nil
}
