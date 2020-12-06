package model

import (
	"time"
)

type BlanceBank struct {
	ID             uint64    `json:"id"`
	Balance        uint64    `json:"balance"`
	BalanceAchieve uint64    `json:"balance_achieve"`
	Code           string    `json:"code"`
	Enable         string    `json:"enable"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type BlanceBankTable interface {
	TableName() string
}

func (BlanceBank) TableName() string {
	return "blance_bank"
}

func (s *Server) GetBankByID(bankid uint64) (*BlanceBank, error) {
	bank := &BlanceBank{}
	err := s.DB.Debug().Where("id = ?", bankid).Take(&bank).Error
	if err != nil {
		return nil, err
	}

	return bank, nil
}

func (s *Server) UpdateBankBalance(blanceBank *BlanceBank) *BlanceBank {
	topUp := s.DB.Debug().Model(&blanceBank).Update("balance", blanceBank.Balance)
	if topUp != nil {
		return blanceBank
	}

	return nil
}
