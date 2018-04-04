package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gchaincl/dotsql"
)

// InitDB initializes database
func InitDB() (*sql.DB, error) {
	db, err := sql.Open(sqlite3Str, memStr)
	if err != nil {
		log.Fatalf("Error opening DB (%s)\n", err)
		return nil, err
	}
	return db, nil
}

// LoadSQLFile imports SQL queries from given file.
func LoadSQLFile(schemaFile string) (*dotsql.DotSql, error) {
	dot, err := dotsql.LoadFromFile(schemaFile)
	if err == nil {
		fmt.Printf("Successfully loaded sql queries from file: (%s)\n", schemaFile)
		return dot, nil
	}
	fmt.Println("Error occurred loading (%s): (%s)\n", schemaFile, err)
	return nil, err

}

// ExecuteScript executes the given script on given database.
func ExecuteScript(db *sql.DB, dot *dotsql.DotSql, scriptName string) {
	res, err := dot.Exec(db, scriptName)
	if err == nil && res != nil {
		fmt.Printf("Executed (%s) script successfully\n", scriptName)
	} else {
		fmt.Printf("error occurred executing (%s): (%s)\n", scriptName, err)
	}
}
