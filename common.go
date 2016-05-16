package gocommon

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"runtime"
	"strings"
)

// GetFileLine is a debug utility that get the file&line of caller.
func GetFileLine() (string, int) {
	_, fn, line, _ := runtime.Caller(1)
	return fn, line
}

// CheckedDeferFunc0 is useful for report error in deferred function without argument.
func CheckedDeferFunc0(f func() error) {
	err := f()
	if err != nil {
		log.Println(err)
		_, fn, line, _ := runtime.Caller(1)
		log.Println(fn, line)
	}
}

// MatchPattern function checks if the str matches specified regular expression pattern.
func MatchPattern(str, pattern string) bool {
	r := regexp.MustCompile(pattern)
	return r.MatchString(str)
}

// GetMySQLConnectionString constructs a connection string based on input.
func GetMySQLConnectionString(host, port, database, username, password string) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	return connectionString
}

// RandomString returns a random string of n length with letters.
func RandomString(n int, letters string) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// InitializeDatabase initials the database, make sure database and table exist
func InitializeDatabase(sqlDriver, connectionString, database, table, sqlCreateTable string) {
	var sqlCreateDatabase = "create database if not exists " + database

	connectionString = strings.TrimSuffix(connectionString, database)
	db, _ := sql.Open(sqlDriver, connectionString)
	_, err := db.Exec(sqlCreateDatabase)
	if err != nil {
		panic(err.Error())
	}

	connectionString += database
	db, err = sql.Open(sqlDriver, connectionString)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec(sqlCreateTable)
	if err != nil {
		panic(err.Error())
	}
}
