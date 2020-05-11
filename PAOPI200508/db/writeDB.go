package db

import (
	"log"
	"strconv"
	"time"
)

func WrtieDB() {
	sqlStr := "INSERT INTO maip(order_id, end_time) VALUES "
	var vals []interface{}
	data := []map[string]string{
		{"v1": "1", "v2": time.Now().Format("2006-01-02 15:04:05")},
		{"v1": "2", "v2": time.Now().Format("2006-01-02 15:04:05")},
		{"v1": "3", "v2": time.Now().Format("2006-01-02 15:04:05")},
	}
	for i := 30001; i <= 60000; i++ {
		//数字变时间
		data = append(data, map[string]string{"v1": strconv.Itoa(i), "v2": time.Now().Format("2006-01-02 15:04:05")})
	}
	for _, row := range data {
		sqlStr += "(?, ?),"
		vals = append(vals, row["v1"], row["v2"])
	}
	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	log.Println(sqlStr)
	log.Println(vals)
	//prepare the statement
	stmt, e := dbConn.Prepare(sqlStr)
	if e != nil {
		log.Println(e)
	}
	//format all vals at once
	result, e := stmt.Exec(vals...)
	if e != nil {
		log.Println(e)
	}
	log.Println(result)
}
