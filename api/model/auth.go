package model

import (
	"miniewallet/api/auth"

	"github.com/twinj/uuid"
)

type Auth struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID   uint64 `gorm:";not null;" json:"user_id"`
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
}

type AuthTable interface {
	TableName() string
}

func (Auth) TableName() string {
	return "auth"
}

func (s *Server) FetchAuth(authD *auth.AuthDetails) (*Auth, error) {
	au := &Auth{}
	err := s.DB.Debug().Where("user_id = ? AND auth_uuid = ?", authD.UserId, authD.AuthUuid).Take(&au).Error
	if err != nil {
		return nil, err
	}
	return au, nil
}

func (s *Server) DeleteAuth(authD *auth.AuthDetails) error {
	au := &Auth{}
	db := s.DB.Debug().Where("user_id = ? AND auth_uuid = ?", authD.UserId, authD.AuthUuid).Take(&au).Delete(&au)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (s *Server) CreateAuth(userId uint64) (*Auth, error) {
	au := &Auth{}
	au.AuthUUID = uuid.NewV4().String()
	au.UserID = userId
	s.DB.Debug().Where("user_id = ?", userId).Take(&au).Delete(&au)
	err := s.DB.Debug().Create(&au).Error
	if err != nil {
		return nil, err
	}

	return au, nil
}
