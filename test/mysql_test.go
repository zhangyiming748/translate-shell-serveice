package test

import (
	"testing"

	"github.com/zhangyiming748/translate-server/model"
	mysql "github.com/zhangyiming748/translate-server/storage"
)

func TestMysql(t *testing.T) {
	mysql.SetMysql()
	mysql.GetMysql().Sync(model.History{})
	s := new(model.History)
	s.Src = "hello"
	s.Dst = "你好"
	if i, err := s.InsertOne(); err != nil {
		t.Error(err)
	} else {
		t.Log(i)
	}
}
