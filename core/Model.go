package core

import (
    "database/sql"
    "github.com/go-sql-driver/mysql"
)

type Model struct {
  table string
}

func (m Model) fetch_count() {

}

func (m Model) Init(){
   cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }
}

func (m Model) FetchRow(){
  
}
