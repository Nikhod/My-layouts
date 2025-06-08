package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type (
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}

	Db struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	}

	Configs struct {
		Server Server `json:"server"`
		Db     Db     `json:"db`
	}

	Note struct {
		Id        int       `json:"id" gorm:"id"`
		Content   []byte    `json:"content" gorm:"content"`
		UserId    int       `json:"user_id" gorm:"user_id"`
		Active    bool      `json:"active" gorm:"active"`
		CreatedAt time.Time `json:"created_at" gorm:"created_at"`
		UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
		DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"`
	}

	Users struct {
		Id        int       `json:"id" gorm:"id"`
		Name      string    `json:"name" gorm:"name"`
		Login     string    `json:"login" gorm:"login"`
		Password  string    `json:"password" gorm:"password"`
		Active    bool      `json:"active" gorm:"active"`
		CreatedAt time.Time `json:"created_at" gorm:"created_at"`
		UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
		DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"`
		Token     string    `json:"token"`
	}

	Answer struct {
		Date           time.Time
		ResponseAnswer string
	}

	TokenClaims struct {
		jwt.StandardClaims
		Login string `json:"login"`
	}

	SendToken struct {
		Date  time.Time
		Token string
	}
)

type AmountLimits struct {
	Id        int
	Min       int
	Max       int
	ServiceId int
	AgentId   int
	CurrName  string
	Active    bool
}
