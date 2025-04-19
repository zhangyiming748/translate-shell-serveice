package model

import (
	"time"
	"github.com/zhangyiming748/translate-server/storage"
)

type History struct {
	Id        int64     `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	Src       string    `xorm:"comment('原文') VARCHAR(255)"`
	Dst       string    `xorm:"comment('译文') VARCHAR(255)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}


func init() {
	if mysql.UseMysql() {
		mysql.GetMysql().Sync(History{})
	}
}

/*
插入
*/
func (h *History) InsertOne() (int64, error) {
	return mysql.GetMysql().InsertOne(h)
}

/*
根据原文查询译文
*/
func (h *History) FindBySrc() (bool, error) {
	return mysql.GetMysql().Where("src = ?", h.Src).Get(h)
}

