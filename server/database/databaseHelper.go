package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

// GetDatabaseLoginFromFile returns array of login information for SQL database from file
/*  file format:
user
password
database name
*/
func GetDatabaseLoginFromFile(dirPath string, fileName string) []string {
	if dirPath == "" {
		_, currFile, _, _ := runtime.Caller(0)
		dirPath = filepath.Dir(currFile)
	}
	filepath := fmt.Sprintf("%s/%s", dirPath, fileName)
	b, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	return strings.Split(str, "\n")
}

// HandleDatabaseError logs error to console
func HandleDatabaseError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// HandleTxError rolls back transaction and logs error to console
func HandleTxError(tx *sql.Tx, err error) {
	// TODO make safer (copies lock value)
	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}
}

// TxRowsAffected handles transaction error and prints affected rows to console.
func TxRowsAffected(res sql.Result, tx *sql.Tx) {
	lastID, err := res.LastInsertId()
	HandleTxError(tx, err)
	rowCnt, err := res.RowsAffected()
	HandleDatabaseError(err)
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
}
