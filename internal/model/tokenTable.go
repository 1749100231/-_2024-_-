package model

type TokenTable struct {
	UserID     int64 `gorm:"primary_key"`
	UserType   int
	UpdateTime int64
	Token      string
}
