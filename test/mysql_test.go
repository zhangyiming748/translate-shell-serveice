package test

import (
	"testing"


	mysql "github.com/zhangyiming748/translate-server/storage"
		"github.com/zhangyiming748/translate-server/model"
)

func TestMysql(t *testing.T) {
	mysql.SetMysql()
	mysql.GetMysql().Sync(model.History{})
	s:=new(model.History)
	s.Src="hello"
	s.Dst="你好"
	if i,err:=s.InsertOne();err!=nil{
		t.Error(err)
	}else{
		t.Log(i)
	}
}