package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetAllMaip() ([]string, error) {
	stmt, e := dbConn.Prepare("select * from maip;")
	defer stmt.Close()
	if e != nil {
		log.Println("数据库prepare出错", e)
		return nil, e
	}
	rows, e := stmt.Query()
	if e != nil {
		log.Println("数据库卡query出错", e)
		return nil, e
	}
	var ids []string
	for rows.Next() {
		var id string
		if e := rows.Scan(&id); e != nil {
			log.Println("数据库scan出错", e)
			return nil, e
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func UpdateTime(vid string) (sql.Result, error) {
	//UPDATE comment c set c.time = DATE_ADD(c.time, INTERVAL 7 DAY)
	stmtDel, err := dbConn.Prepare("update maip m set m.end_time = DATE_ADD(m.end_time,INTERVAL 7 DAY) WHERE order_id=?")
	defer stmtDel.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, error := stmtDel.Exec(vid)
	if error != nil {
		log.Printf("UPDATE Record error: %v", err)
		stmtDel.Close()
		return nil, error
	}
	stmtDel.Close()
	return result, nil
}
