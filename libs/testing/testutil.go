package testing

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"backend/libs/databases"

	"github.com/jmoiron/sqlx"
)

var db databases.Database

func NewTestDB(databaseURL string) databases.Database {
	if databaseURL == "" {
		log.Fatalf("The database url is not defined: %#v", databaseURL)
	}
	db = databases.NewPostgres(databaseURL)
	return db
}

func Seed(db databases.Database, schema string) {
	current, _ := os.Getwd()
	parent := filepath.Dir(current)
	seed := parent + "/../../../schemas/" + schema
	content, err := ioutil.ReadFile(seed)
	if err != nil {
		log.Fatal("Cannot read the seed file: ", seed)
	}
	data := string(content)
	re := regexp.MustCompile(`(\r\n|\n|\r|\t)`)
	sql := re.ReplaceAllString(data, "")
	conn := db.GetConnection()
	if conn == nil {
		log.Fatal("I could not connect to database so I did not seed.", err)
	}
	_, err = conn.Exec(sql)
	if err != nil {
		log.Fatal("Seed error: ", err)
	}
}

func Exec(query string) {
	conn := db.GetConnection()
	_, err := conn.Exec(query)
	if err != nil {
		log.Fatal("Failed to execute query:", query)
	}
}

func Get(query string) map[string]interface{} {
	conn := db.GetConnection()
	var result map[string]interface{}
	err := conn.Get(&result, query)
	if err != nil {
		log.Fatal("Failed to execute Get with query:", query)
	}
	return result
}

func Select(dest interface{}, query string, args ...interface{}) error {
	conn := db.GetConnection()
	return conn.Select(dest, query, args)
}

func QueryRow(query string, args ...interface{}) *sqlx.Row {
	return db.GetConnection().QueryRowx(query)
}
