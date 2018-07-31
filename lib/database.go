package lib

import (
	"database/sql"
	"os"
	"github.com/siddontang/go/log"
)


type DB struct {
	TableName string
	DBName string
	Conn *sql.DB
}

func (client *DB) Connect(){
	db, err := sql.Open("sqlite3", client.DBName)
	if err != nil {
		log.Fatal(err)
	}

	client.Conn = db
}

func (client *DB) CreateDb() {

	os.Remove(client.DBName)

	db, err := sql.Open("sqlite3", client.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "create table " + client.TableName + "(id integer not null primary key, original_url mediumtext, shorten_url tinytext); delete from " + client.TableName + "; "

	log.Info("Creating table...")
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("%q: %s\n", err, sqlStmt)
	}

}

func (client *DB) Save(sql string){
	_, err := client.Conn.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func (client *DB) QueryDb(sql string) string {

	stmt, err := client.Conn.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var url string
	err = stmt.QueryRow().Scan(&url)
	if err != nil{
		return "-1"
	}
	return url
}


/**
tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
 */
