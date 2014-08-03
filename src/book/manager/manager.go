package manager 

import ."libs/sqlconn"
import "book/model"
import "conf"
import "fmt"

type BookManager struct {
	Table string
	Sql SqlConn
}

type BookInfo struct {
	Id	  int64		`json:"id"`
	Hash  string	`json:"hash"`	
	Fname string	`json:"fname"`
}

func NewBooKManager() *BookManager {
	sqlConn := SqlConn{
			User: conf.DB_User,
			Password: conf.DB_Password, 
			Host: conf.DB_Host, 
			Port: conf.DB_Port, 
			DB: conf.DB_Name,
		}
	sqlConn.Execute(model.TableStruct)
	return &BookManager{model.Table, sqlConn}
}

func (manager *BookManager)Create(hash string, fname string) (*BookInfo, error) {
	sql := fmt.Sprintf("INSERT INTO `%s`(hash,fname) VALUES(unhex('%s'),'%s');",
					    manager.Table, hash, fname)
	insertId, err := manager.Sql.ExecuteReturnInsertId(sql)
	if err != nil {
		return nil, err 
	}
	return &BookInfo{Id: insertId, Hash:hash, Fname:fname}, err
}

func (manager *BookManager)GetFileInfo(id int64) (*BookInfo, error) {
	sql := fmt.Sprintf("SELECT id,hex(hash) hash,fname FROM `%s` WHERE `id`=%d;", 
						manager.Table, id)
	row, err := manager.Sql.Query(sql)
	if err != nil {
		return nil, err
	}
	var len int64	
	var bookInfo BookInfo 

	len = 0
	for row.Next() {
		errInRow := row.Scan(&bookInfo.Id, &bookInfo.Hash, &bookInfo.Fname)
		if errInRow == nil {
			len = len + 1
		}
	}
	if len == 0 {
		return nil, err	
	}
	return &bookInfo, err
}

func (manager *BookManager)GetFileInfos()(*[]BookInfo, error) {
	sql := fmt.Sprintf("SELECT id,hex(hash) hash,fname FROM `%s`;", 
						manager.Table)
	rows, err := manager.Sql.Query(sql)
	if err != nil {
		return nil, err
	}

	var len int64	
	var bookInfos []BookInfo 

	len = 0
	for rows.Next() {
		var bookInfo BookInfo 
		errInRow := rows.Scan(&bookInfo.Id, &bookInfo.Hash, &bookInfo.Fname)
		if errInRow == nil {
			len = len + 1
			bookInfos = append(bookInfos, bookInfo)
		}
	}
	if len == 0 {
		return nil, err	
	}
	return &bookInfos, err
}

var Manager BookManager = *NewBooKManager()
