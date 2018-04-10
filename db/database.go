package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/chimauwah/services-api-tech-challenge/model"
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
	dot, err := loadSQLFile("db/schema.sql")

	// run schema scripts
	executeScript(db, dot, "create-employee-table")
	executeScript(db, dot, "create-coreskill-table")

	// load seed date from file
	dot, err = loadSQLFile("db/seed.sql")

	// run seed data scripts
	executeScript(db, dot, "insert-employee")
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
		fmt.Printf("Error occurred loading (%s): (%s)\n", schemaFile, err)
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
func AddEmployee(emp model.Employee) (model.Employee, error) {
	q := `INSERT INTO Employee (id, enter_ts, last_change_ts, active, first_name, last_name, 
		address1,address2,city,state,zip,cell_phone,home_phone,picture,title,role,ip_phone,
		samaccountname,mail,primary_pa,secondary_pa,office,manager_dn,travel_pref,
		manager_samaccountname,last_hash,image_hash,nick_name,client_loc) VALUES
		(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	emp.ID = getEmployeeNextSequence()
	_, err := getDB().Exec(q, emp.ID, emp.EnterTs, emp.LastChangeTs, emp.Active, emp.FirstName,
		emp.LastName, emp.Address1, emp.Address2, emp.City, emp.State, emp.Zip, emp.CellPhone,
		emp.HomePhone, emp.Picture, emp.Title, emp.Role, emp.IPPhone, emp.Samaccountname,
		emp.Mail, emp.PrimaryPa, emp.SecondaryPa, emp.Office, emp.ManagerDn, emp.TravelPref,
		emp.ManagerSamaccountname, emp.LastHash, emp.ImageHash, emp.NickName, emp.ClientLoc)
	if err != nil {
		log.Println(err)
		return emp, err
	}
	return emp, nil
}

// FindEmployee returns employee with given id
func FindEmployee(id int) (model.Employee, error) {
	var emp model.Employee
	q := "SELECT * FROM Employee where id = ?"
	err := getDB().QueryRow(q, id).Scan(&emp.ID, &emp.EnterTs, &emp.LastChangeTs,
		&emp.Active, &emp.FirstName, &emp.LastName, &emp.Address1, &emp.Address2,
		&emp.City, &emp.State, &emp.Zip, &emp.CellPhone, &emp.HomePhone, &emp.Picture,
		&emp.Title, &emp.Role, &emp.IPPhone, &emp.Samaccountname, &emp.Mail,
		&emp.PrimaryPa, &emp.SecondaryPa, &emp.Office, &emp.ManagerDn,
		&emp.TravelPref, &emp.ManagerSamaccountname, &emp.LastHash,
		&emp.ImageHash, &emp.NickName, &emp.ClientLoc)
	switch {
	case err == sql.ErrNoRows:
		log.Println(err)
		return emp, err
	case err != nil:
		log.Println(err)
		return emp, err
	default:
		return emp, nil
	}

}

// FindAllEmployees returns all employees
func FindAllEmployees() ([]model.Employee, error) {
	var res []model.Employee
	q := "SELECT * FROM Employee"
	rows, err := getDB().Query(q)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
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
			log.Println(err)
			return res, err
		}
		res = append(res, emp)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
		return res, err
	}
	return res, nil
}

// DeleteEmployee deletes the employee with given id
func DeleteEmployee(id int) (int64, error) {
	q := "DELETE FROM Employee WHERE ID = ?"
	res, err := getDB().Exec(q, id)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	cnt, _ := res.RowsAffected()
	return cnt, nil
}

// UpdateEmployee updates the employee with given id, with given values
func UpdateEmployee(id int, emp model.Employee) (int64, error) {
	q := `UPDATE Employee SET 
	active = coalesce(?,active), first_name = coalesce(?,first_name), last_name = coalesce(?,last_name),
	address1 = coalesce(?,address1), address2 = coalesce(?,address2), city = coalesce(?,city), 
	state = coalesce(?,state), zip = coalesce(?,zip), cell_phone = coalesce(?,cell_phone), 
	home_phone = coalesce(?,home_phone), picture = coalesce(?,picture), title = coalesce(?,title), 
	role = coalesce(?,role), ip_phone = coalesce(?,ip_phone), samaccountname = coalesce(?,samaccountname), 
	mail = coalesce(?,mail), primary_pa = coalesce(?,primary_pa), secondary_pa = coalesce(?,secondary_pa), 
	office = coalesce(?,office), manager_dn = coalesce(?,manager_dn), travel_pref = coalesce(?,travel_pref), 
	manager_samaccountname = coalesce(?,manager_samaccountname), last_hash = coalesce(?,last_hash), 
	image_hash = coalesce(?,image_hash), nick_name = coalesce(?,nick_name), client_loc = coalesce(?,client_loc)
	WHERE id = ?`
	res, err := getDB().Exec(q, emp.Active, emp.FirstName,
		emp.LastName, emp.Address1, emp.Address2, emp.City, emp.State, emp.Zip, emp.CellPhone,
		emp.HomePhone, emp.Picture, emp.Title, emp.Role, emp.IPPhone, emp.Samaccountname,
		emp.Mail, emp.PrimaryPa, emp.SecondaryPa, emp.Office, emp.ManagerDn, emp.TravelPref,
		emp.ManagerSamaccountname, emp.LastHash, emp.ImageHash, emp.NickName, emp.ClientLoc, id)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	cnt, _ := res.RowsAffected()
	return cnt, nil
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
