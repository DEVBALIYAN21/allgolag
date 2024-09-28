package database

import (
	"context"
	"database/sql"
	"fmt"
)

// Database struct to hold the SQL database object
type Database struct {
	SqlDb *sql.DB
}

var dbContext = context.Background()

// CreateReminder adds a new reminder to the database
func (db *Database) CreateReminder(title, description, alias string) error {
    query := "INSERT INTO reminders (title, description, alias) VALUES (?, ?, ?)"
    _, err := db.SqlDb.Exec(query, title, description, alias)
    if err != nil {
        fmt.Printf("Error creating reminder: %v\n", err)
        return err
    }
    fmt.Println("Reminder created successfully")
    return nil
}

// RetrieveReminder fetches a reminder based on the alias
func (db *Database) RetrieveReminder(alias string) {
	query := `SELECT title, description FROM reminders WHERE alias = ?`
	row := db.SqlDb.QueryRow(query, alias)

	var title, description string
	err := row.Scan(&title, &description)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No reminder found for the given alias.")
		} else {
			fmt.Printf("Error retrieving reminder: %v\n", err)
		}
		return
	}

	fmt.Printf("Reminder - Title: %s, Description: %s\n", title, description)
}

// DeleteReminder removes a reminder based on the alias
func (db *Database) DeleteReminder(alias string) {
	query := `DELETE FROM reminders WHERE alias = ?`
	_, err := db.SqlDb.Exec(query, alias)
	if err != nil {
		fmt.Printf("Error deleting reminder: %v\n", err)
	} else {
		fmt.Println("Reminder deleted successfully!")
	}
}
