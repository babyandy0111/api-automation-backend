package models

import "time"

type Account struct {
	Id        int64     `json:"id" xorm:"not null pk autoincr INT(11)"`
	Account   string    `json:"account" xorm:"not null comment('帳號') VARCHAR(64)"`
	Password  string    `json:"password" xorm:"not null VARCHAR(255)"`
	Name      string    `json:"name" xorm:"not null comment('名字') VARCHAR(64)"`
	IsDisable int       `json:"is_disable" xorm:"not null comment('1:停用') TINYINT(4)"`
	CreatedAt time.Time `json:"created_at" xorm:"created"` // created_at
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"` // updated_at
}
