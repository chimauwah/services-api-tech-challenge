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
	if err != nil {
		fmt.Println("Error occurred loading (%s): (%s)\n", schemaFile, err)
		return nil, err
	}
	fmt.Printf("Successfully loaded sql queries from file: (%s)\n", schemaFile)
	return dot, nil

}

// executeScript executes the given script on given database.
func executeScript(db *sql.DB, dot *dotsql.DotSql, scriptName string) {
	_, err := dot.Exec(db, scriptName)
	if err != nil {
		fmt.Printf("error occurred executing (%s): (%s)\n", scriptName, err)
	} else {
		// fmt.Printf("Executed (%s) script successfully\n", scriptName)
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

// AddEmployee adds an Employee to the datastore and returns the Employee with the generated id
func AddEmployee(emp model.Employee) (e model.Employee) {
	q := "INSERT INTO Employee (id, enter_ts, last_change_ts, active, first_name, last_name, " +
		"address1,address2,city,state,zip,cell_phone,home_phone,picture,title,role,ip_phone," +
		"samaccountname,mail,primary_pa,secondary_pa,office,manager_dn,travel_pref," +
		"manager_samaccountname,last_hash,image_hash,nick_name,client_loc) VALUES " +
		"(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	emp.ID = getEmployeeNextSequence()
	_, err := getDB().Exec(q, emp.ID, emp.EnterTs, emp.LastChangeTs, emp.Active, emp.FirstName,
		emp.LastName, emp.Address1, emp.Address2, emp.City, emp.State, emp.Zip, emp.CellPhone,
		emp.HomePhone, emp.Picture, emp.Title, emp.Role, emp.IPPhone, emp.Samaccountname,
		emp.Mail, emp.PrimaryPa, emp.SecondaryPa, emp.Office, emp.ManagerDn, emp.TravelPref,
		emp.ManagerSamaccountname, emp.LastHash, emp.ImageHash, emp.NickName, emp.ClientLoc)
	if err != nil {
		log.Fatal(err)
	}
	return emp
}

// FindAllEmployees returns all employees
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

// DeleteEmployee deletes the employee with given id
func DeleteEmployee(id int) (int64, error) {
	q := "DELETE FROM Employee WHERE ID = ?"
	res, err := getDB().Exec(q, id)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	v, _ := res.RowsAffected()
	return v, nil
}

func getEmployeeNextSequence() (seq int) {
	q := "SELECT MAX(ID) FROM Employee"
	row := getDB().QueryRow(q)
	err := row.Scan(&seq)
	if err != nil {
		panic(err)
	}
	return seq + 1
}
