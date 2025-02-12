package database

import (
	"fmt"

	data "github.com/maseiler/avh-booking-system/server/data"
)

// getItemsByQuery returns list of items as requested in string
func getItemsByQuery(query string) []data.Item {
	items := []data.Item{}
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		item := data.Item{}
		err := rows.Scan(&item.ID, &item.Name, &item.Type, &item.Size, &item.Unit, &item.Price)
		items = append(items, item)
		HandleDatabaseError(err)
	}
	defer rows.Close()
	err = rows.Err()
	HandleDatabaseError(err)
	// fmt.Printf("Performed item query: \"%s\"\n", query)
	return items
}

// GetAllItems returns all items from database
func GetAllItems() []data.Item {
	query := "SELECT * FROM items;"
	return getItemsByQuery(query)
}

// GetItemsOfColumnWithValue returns all items where value matches in specific column
// func GetItemsOfColumnWithValue(column string, value string) []data.Item {
// 	//TODO string conversion ?!
// 	query := ""
// 	if column == "ItemId" {
// 		intValue, _ := strconv.Atoi(value)
// 		query = fmt.Sprintf("SELECT * FROM items WHERE %s = %d;", column, intValue)
// 	} else if column == "Price" || column == "Size" {
// 		floatValue, _ := strconv.ParseFloat(value, 32)
// 		query = fmt.Sprintf("SELECT * FROM items WHERE %s = %f;", column, floatValue)
// 	} else if column == "Name" || column == "Unit" || column == "Type" {
// 		query = fmt.Sprintf("SELECT * FROM items WHERE %s = \"%s\";", column, value)
// 	} else {
// 		panic("Invalid column")
// 	}
// 	return getItemsByQuery(query)
// }

// NewItemExists returns true if item exists in database (based on Name and Size)
func NewItemExists(newItem data.Item) bool {
	queryString := fmt.Sprintf("SELECT * FROM items WHERE name = \"%s\" AND size = %.2f;", newItem.Name, newItem.Size)
	items := getItemsByQuery(queryString)
	if len(items) == 0 {
		return false
	}
	return true
}

// ItemExists returns true if item with same ItemID exists in database
func ItemExists(item data.Item) bool {
	queryString := fmt.Sprintf("SELECT * FROM items WHERE id = %d;", item.ID)
	items := getItemsByQuery(queryString)
	if len(items) == 0 {
		return false
	}
	return true
}

// AddItem adds a new user to database and prints info
func AddItem(newItem data.Item) {
	// todo: get info from input
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("INSERT INTO items(name, price, size, unit, type) VAlUES(?, ?, ?, ?, ?)")
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec(newItem.Name, newItem.Price, newItem.Size, newItem.Unit, newItem.Type)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
	stmt.Close()
}

// ModifyItem replaces all values of a item
func ModifyItem(item data.Item) {
	query := fmt.Sprintf("UPDATE items SET name = \"%s\", price = \"%f\", size = \"%f\", unit = \"%s\", type = \"%s\" WHERE id = %d;", item.Name, item.Price, item.Size, item.Unit, item.Type, item.ID)
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	fmt.Println(rows)
}

// DeleteItem deletes a item with corresponding ID from database
func DeleteItem(item data.Item) {
	query := fmt.Sprintf("DELETE FROM items WHERE id = %d;", item.ID)
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	fmt.Println(rows)
}
