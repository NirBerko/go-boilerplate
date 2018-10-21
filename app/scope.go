package app

import (
	"github.com/jinzhu/gorm"
)

type RequestScope interface {
	Db() *gorm.DB
	SetDB(db *gorm.DB)
	RequestID() string
	SetUserID(userId int)
}

type requestScope struct {
	db        *gorm.DB
	requestID string
	userID    int
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

func (rs *requestScope) SetUserID(userId int) {
	rs.userID = userId
}

func (rs *requestScope) UserID() int {
	return rs.userID
}

func (rs *requestScope) SetDB(db *gorm.DB) {
	rs.db = db
}

func (rs *requestScope) Db() *gorm.DB {
	return rs.db
}
