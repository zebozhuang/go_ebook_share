package sqlconn

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"
import "log"

type SqlConn struct {
	User	 string
	Password string	
	Host	 string	
	Port	 string 
	DB		 string

	conn	 *sql.DB 
	dsn		 string
}    
 
func (conn *SqlConn)GetConn() (*sql.DB, error) {
	if conn.dsn== "" {
		conn.dsn= fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
							   conn.User,
                               conn.Password,
							   conn.Host,
                               conn.Port,
						       conn.DB)
	}
	return sql.Open("mysql", conn.dsn)
}

func (sqlConn *SqlConn)Query(sqlstmt string) (*sql.Rows, error){
	if sqlConn.conn == nil {
		sqlConn.conn, _ = sqlConn.GetConn()
	}

	res, err := sqlConn.conn.Query(sqlstmt)
	if err != nil {
		log.Printf("fail to query:", sqlstmt)
	}
	return res, err
}

func (sqlConn *SqlConn)Execute(sqlstmt string) (int64, error) {
	var affectedRows int64

	if sqlConn.conn == nil {
		sqlConn.conn, _ = sqlConn.GetConn()
	}
	res, err := sqlConn.conn.Exec(sqlstmt)
	if err != nil {
		log.Printf("fail to execute:", sqlstmt)
		return -1, err
	}
	affectedRows, err = res.RowsAffected()

	if err != nil {
		return -1, err
	}
	return affectedRows, err
}

func (sqlConn *SqlConn)ExecuteReturnInsertId(sqlstmt string) (int64, error) {
	var lastInsertId int64

	if sqlConn.conn == nil {
		sqlConn.conn, _ = sqlConn.GetConn()
	}
	res, err := sqlConn.conn.Exec(sqlstmt)
	if err != nil {
		return -1, err
	}
	lastInsertId, err = res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return lastInsertId, err
}
