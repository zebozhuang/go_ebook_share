package sqlconn_test

import "testing"
import ."libs/sqlconn"
import "conf"

func Test_GetConn(t *testing.T) {
	conn := SqlConn{
			User: conf.DB_User,
			Password: conf.DB_Password, 
			Host: conf.DB_Host, 
			Port: conf.DB_Port, 
			DB: conf.DB_Name,
		}

	_, err := conn.GetConn()
	if err != nil {
		t.Error("fail to get conn:", err)
	}
}

func Test_Query(t *testing.T) {
	conn := SqlConn{
			User: conf.DB_User,
			Password: conf.DB_Password, 
			Host: conf.DB_Host, 
			Port: conf.DB_Port, 
			DB: conf.DB_Name,
		}

	_, err := conn.Query("show tables;")
	if err != nil {
		t.Error("fail to query", err)
	} 
}

func Test_Execute(t *testing.T){
	conn := SqlConn{
			User: conf.DB_User,
			Password: conf.DB_Password, 
			Host: conf.DB_Host, 
			Port: conf.DB_Port, 
			DB: conf.DB_Name,
		}

	var err error 

	_, err = conn.Execute("create table if not exists `t`(id int);")
	if err != nil {
		t.Error("fail to create table:", err)
	}
	
	_, err = conn.Execute("drop table `t`;")
	if err != nil {
		t.Error("fail to drop table:", err)
	}
}
