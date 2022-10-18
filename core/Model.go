package core

import (
  "fmt"
  "log"
    "database/sql"
    "github.com/go-sql-driver/mysql"
)
var db *sql.DB

type Model struct {
  Table string
}

func (m Model) fetch_count() {

}

func (m Model) connect(){
    cfg := mysql.Config{
        //User:   os.Getenv("DBUSER"),
        User:   "root",
        //Passwd: os.Getenv("DBPASS"),
        Passwd: "123456",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "days7",
        AllowNativePasswords: true,
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}

func (m Model) FetchRow(id int) (*sql.Row)  {
    m.connect()
    sql := fmt.Sprintf("SELECT id,ename,name FROM `%s` WHERE `id` = '%d'", m.Table, id)
    row := db.QueryRow(sql)
    return row
}
