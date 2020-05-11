package db

import (
	"testing"
)

type todo struct {
	id          string
	title       string
	description string
}

func TestSql(t *testing.T) {
	stmt, e := dbConn.Prepare("select * from todo;")
	defer stmt.Close()
	if e != nil {
		t.Fatal(e)
	}
	rows, e := stmt.Query()
	if e != nil {
		t.Fatal(e)
	}
	var todos []*todo
	for rows.Next() {
		var id, title, description string
		if e := rows.Scan(&id, &title, &description); e != nil {
			t.Fatal(e)
		}
		i := &todo{
			id:          id,
			title:       title,
			description: description,
		}
		todos = append(todos, i)
	}
	for _, value := range todos {
		t.Log(*value)
	}
}
