package Vehicle

import (
	"database/sql"
	"fmt"
)

type Store struct {
	Db *sql.DB
}
type Car struct {
	Id         int
	Name       string
	Model      string
	EngineType string
}

func (s Store) Set(c Car) bool {
	if c.Id > 0 {
		_, err := s.Db.Query("insert into car values(?,?,?,?);", c.Id, c.Name, c.Model, c.EngineType)
		if err != nil {
			fmt.Println(err)
		}
		return true
	} else {
		return false
	}
}

func (s Store) Get(Id int) (c Car) {
	if Id <= 0 {
		c.Id = 0
		c.Name = ""
		c.Model = ""
		c.EngineType = ""
		return
	}
	rows, err := s.Db.Query("select * from car where Id=?;", Id)
	rows.Next()
	rows.Scan(&c.Id, &c.Name, &c.Model, &c.EngineType)
	rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (s Store) Delete(Id int) bool {
	row, _ := s.Db.Exec("delete from car where id=?", Id)
	aff, _ := row.RowsAffected()
	if aff == 0 {
		return false
	}
	return true
}
