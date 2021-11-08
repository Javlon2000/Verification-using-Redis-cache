package models

import (
	"errors"
	"time"
  
    "gorm.io/gorm"
)

type Check struct {
	Username string `json:"username" binding:"required"`
	Password string 
}

type InsertDatabase struct {
	Username string
	Email string 
	Password string 
}

type User struct {
	Id uint `gorm:"column:user_id"`
	Email string
	Username string
	CreatedAt time.Time
}

func IsNotFound(row *gorm.DB) bool {

  return errors.Is(row.Error, gorm.ErrRecordNotFound)
}
