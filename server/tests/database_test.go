package tests

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	data "../data"
	dbP "../database"
)

var db *sql.DB

func createTestDatabase() {
	_, currFile, _, _ := runtime.Caller(0)
	currDirectory := filepath.Dir(currFile)
	dbLogin := dbP.GetDatabaseLoginFromFile(currDirectory, "testDatabaseLogin.txt")
	loginInfo := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/", dbLogin[0], dbLogin[1])
	var err error
	db, err = sql.Open("mysql", loginInfo)
	dbP.HandleDatabaseError(err)
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE test_db")
	dbP.HandleDatabaseError(err)

	_, err = db.Exec("USE test_db")
	dbP.HandleDatabaseError(err)

	createUserTableQuery := `CREATE TABLE users(
		UserId int not null primary key auto_increment,
		BierName varchar(100),
		FirstName varchar(100),
		LastName varchar(100),
		Status varchar(100),
		Email varchar(100),
		Balance float,
		PhoneNumber varchar(100)
		);`
	_, err = db.Exec(createUserTableQuery)
	dbP.HandleDatabaseError(err)
}

func deleteDatabase() {
	deleteDatabaseQuery := "DROP DATABASE test_db;"
	_, err := db.Exec(deleteDatabaseQuery)
	dbP.HandleDatabaseError(err)
}

// TestAddUser checks if a user was added to the database
func TestAddUserFails(t *testing.T) {
	newUser := data.User{
		UserID:    1234,
		BierName:  "NewUser",
		FirstName: "New",
		LastName:  "User",
		Status:    "AH",
		Email:     "newuser@huette.de",
		Balance:   8,
		Phone:     "0000"}
	createTestDatabase()
	// t.Error(database.AddUser(newUser))
	dbP.AddUser(newUser)
}
