package db

import (
	"database/sql"
	"fmt"
	"log"
	"services-api-tech-challenge/model"

	"github.com/gchaincl/dotsql"
)

var db *sql.DB

func setDB(db2 *sql.DB) {
	db = db2
}

func getDB() *sql.DB {
	return db
}

const (
	// name of the db driver.
	sqlite3Str = "sqlite3"
	// name of datasource that gets created as a successful result of open operation.
	memStr = ":memory:"
)

// Init initializes database and loads data into it.
func Init() {
	db, err := sql.Open(sqlite3Str, memStr)
	if err != nil {
		log.Fatalf("Error opening DB (%s)\n", err)
		panic(err)
	}

	//set db instance
	setDB(db)

	// load schema from file
	dot, err := loadSQLFile("schema.sql")

	// run schema scripts
	executeScript(db, dot, "create-employee-table")
	executeScript(db, dot, "create-practicearea-table")
	executeScript(db, dot, "create-checkin-table")
	executeScript(db, dot, "create-coreskill-table")
	executeScript(db, dot, "create-device-table")

	// load seed date from file
	dot, err = loadSQLFile("seed.sql")

	// run seed data scripts
	executeScript(db, dot, "insert-employee")
	executeScript(db, dot, "insert-practicearea")
	executeScript(db, dot, "insert-checkin")
	executeScript(db, dot, "insert-coreskill")
	executeScript(db, dot, "insert-coreskill2")

	// testSelect(db)
	// stmt, err := dot.Prepare(db, "drop-users-table")
	// restul, err := stmt.Exec()
}

// loadSQLFile imports SQL queries from given file.
func loadSQLFile(schemaFile string) (*dotsql.DotSql, error) {
	dot, err := dotsql.LoadFromFile(schemaFile)
	if err == nil {
		fmt.Printf("Successfully loaded sql queries from file: (%s)\n", schemaFile)
		return dot, nil
	}
	fmt.Println("Error occurred loading (%s): (%s)\n", schemaFile, err)
	return nil, err

}

// executeScript executes the given script on given database.
func executeScript(db *sql.DB, dot *dotsql.DotSql, scriptName string) {
	res, err := dot.Exec(db, scriptName)
	if err == nil && res != nil {
		// fmt.Printf("Executed (%s) script successfully\n", scriptName)
	} else {
		fmt.Printf("error occurred executing (%s): (%s)\n", scriptName, err)
	}
}

func testSelect(db *sql.DB) {
	//test reading an item
	stmt := "SELECT id, first_name, last_name FROM Employee WHERE office = ? and active = ?"
	rows, err := db.Query(stmt, "Charlotte", 1)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var (
		id    int
		fname string
		lname string
	)

	cnt := 0
	for rows.Next() {
		err := rows.Scan(&id, &fname, &lname)
		if err != nil {
			fmt.Printf("error occurred running (%s): (%s)\n", stmt, err)
		}
		fmt.Printf("ID: (%d) | NAME: (%s %s) \n", id, fname, lname)
		cnt++
	}
	fmt.Printf("Total rows returned: (%d)\n", cnt)
}

// FindAllEmployees return all employees
func FindAllEmployees() (e []model.Employee) {
	q := "SELECT * FROM Employee"
	rows, err := getDB().Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var res []model.Employee
	for rows.Next() {
		var emp model.Employee
		err := rows.Scan(&emp.ID, &emp.EnterTs, &emp.LastChangeTs, &emp.Active,
			&emp.FirstName, &emp.LastName, &emp.Address1, &emp.Address2, &emp.City,
			&emp.State, &emp.Zip, &emp.CellPhone, &emp.HomePhone, &emp.Picture,
			&emp.Title, &emp.Role, &emp.IPPhone, &emp.Samaccountname, &emp.Mail,
			&emp.PrimaryPa, &emp.SecondaryPa, &emp.Office, &emp.ManagerDn,
			&emp.TravelPref, &emp.ManagerSamaccountname, &emp.LastHash,
			&emp.ImageHash, &emp.NickName, &emp.ClientLoc)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, emp)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return res
}
