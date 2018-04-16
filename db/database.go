package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

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
	executeScript(db, dot, "update-employee-nulls")

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

// AddEmployee adds an Employee to the datastore and returns the Employee with the generated id
func AddEmployee(emp model.Employee) (model.Employee, error) {
	q := `INSERT INTO Employee (id, enter_ts, active, first_name, last_name, cell_phone, title, 
		samaccountname, mail, primary_pa, office, manager_dn, manager_samaccountname, travel_pref) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	emp.ID = getEmployeeNextSequence()
	_, err := getDB().Exec(q, emp.ID, time.Now(), 1, emp.FirstName, emp.LastName, emp.CellPhone,
		emp.Title, emp.Samaccountname, emp.Mail, emp.PrimaryPa, emp.Office, emp.ManagerDn,
		emp.ManagerSamaccountname, emp.TravelPref)
	if err != nil {
		return emp, err
	}
	return emp, nil
}

// FindEmployee returns employee with given id
func FindEmployee(id int) (model.Employee, error) {
	var emp model.Employee
	q := `SELECT id, enter_ts, active, first_name, last_name, cell_phone, title, samaccountname, 
	mail, primary_pa, office, manager_dn, manager_samaccountname, travel_pref 
	FROM Employee where id = ?`
	err := getDB().QueryRow(q, id).Scan(&emp.ID, &emp.EnterTs, &emp.Active, &emp.FirstName,
		&emp.LastName, &emp.CellPhone, &emp.Title, &emp.Samaccountname, &emp.Mail, &emp.PrimaryPa,
		&emp.Office, &emp.ManagerDn, &emp.ManagerSamaccountname, &emp.TravelPref)
	switch {
	case err == sql.ErrNoRows:
		return emp, err
	case err != nil:
		return emp, err
	default:
		return emp, nil
	}

}

// FindEmployeeDetails returns details of employee with given id
func FindEmployeeDetails(id int) (model.EmployeeDetail, error) {
	var det model.EmployeeDetail
	var err error
	// retrieve employee
	det.Employee, err = FindEmployee(id)
	switch {
	case err == sql.ErrNoRows:
		return det, err
	case err != nil:
		return det, err
	}

	// retrieve all core skills for employee
	q := `SELECT skill, proficiency FROM CoreSkill where employee_id = ?`
	rows, err := getDB().Query(q, id)
	if err != nil {
		return det, err
	}
	defer rows.Close()
	for rows.Next() {
		var skill model.CoreSkill
		err := rows.Scan(&skill.Skill, &skill.Proficiency)
		if err != nil {
			return det, err
		}
		det.Skills = append(det.Skills, skill)
	}
	err = rows.Err()
	if err != nil {
		return det, err
	}
	return det, nil

}

// FindAllEmployees returns all employees
func FindAllEmployees() ([]model.Employee, error) {
	var res []model.Employee
	q := `SELECT id, enter_ts, active, first_name, last_name, cell_phone, title, samaccountname, 
	mail, primary_pa, office, manager_dn, manager_samaccountname, travel_pref 
	FROM Employee`
	rows, err := getDB().Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var emp model.Employee
		err := rows.Scan(&emp.ID, &emp.EnterTs, &emp.Active, &emp.FirstName,
			&emp.LastName, &emp.CellPhone, &emp.Title, &emp.Samaccountname, &emp.Mail,
			&emp.PrimaryPa, &emp.Office, &emp.ManagerDn, &emp.ManagerSamaccountname,
			&emp.TravelPref)
		if err != nil {
			return res, err
		}
		res = append(res, emp)
	}
	err = rows.Err()
	if err != nil {
		return res, err
	}
	return res, nil
}

// SearchEmployees returns all employee that match provided criteria
func SearchEmployees(args url.Values) ([]model.Employee, error) {
	var emps []model.Employee
	q := `SELECT id, enter_ts, active, first_name, last_name, cell_phone, title, samaccountname, 
	mail, primary_pa, office, manager_dn, manager_samaccountname, travel_pref 
	FROM Employee`
	buildWhereClause(&q, args)
	rows, err := getDB().Query(q)
	log.Println(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var emp model.Employee
		err := rows.Scan(&emp.ID, &emp.EnterTs, &emp.Active, &emp.FirstName,
			&emp.LastName, &emp.CellPhone, &emp.Title, &emp.Samaccountname, &emp.Mail,
			&emp.PrimaryPa, &emp.Office, &emp.ManagerDn, &emp.ManagerSamaccountname,
			&emp.TravelPref)
		if err != nil {
			return emps, err
		}
		emps = append(emps, emp)
	}
	err = rows.Err()
	if err != nil {
		return emps, err
	}
	return emps, nil
}

// DeleteEmployee deletes the employee with given id
func DeleteEmployee(id int) (int64, error) {
	q := "DELETE FROM Employee WHERE ID = ?"
	res, err := getDB().Exec(q, id)
	if err != nil {
		return 0, err
	}

	cnt, _ := res.RowsAffected()
	return cnt, nil
}

// UpdateEmployee updates the employee with given id, with given values
func UpdateEmployee(id int, emp model.Employee) (int64, error) {
	q := `UPDATE Employee SET 
	active = coalesce(?,active), first_name = coalesce(?,first_name), last_name = coalesce(?,last_name),
	cell_phone = coalesce(?,cell_phone), title = coalesce(?,title), samaccountname = coalesce(?,samaccountname), 
	mail = coalesce(?,mail), primary_pa = coalesce(?,primary_pa), office = coalesce(?,office), 
	manager_dn = coalesce(?,manager_dn), manager_samaccountname = coalesce(?,manager_samaccountname),
	travel_pref = coalesce(?,travel_pref) WHERE id = ?`
	res, err := getDB().Exec(q, emp.Active, emp.FirstName, emp.LastName, emp.CellPhone,
		emp.Title, emp.Samaccountname, emp.Mail, emp.PrimaryPa, emp.Office, emp.ManagerDn,
		emp.ManagerSamaccountname, emp.TravelPref, id)
	if err != nil {
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

func buildWhereClause(q *string, args url.Values) {
	if len(args) > 0 {
		*q += ` WHERE `
		for k, v := range args {
			*q += k + ` = "` + v[0] + `" and `
		}
		*q += `1 = 1`
	}

}

// You can use interface types as arguments, in which case you can call the
// function with any type that implements the given interface. In Go types
// automatically implement any interfaces if they have the interface's
// methods. So if you want to accept all possible types, you can use empty
// interface (interface{}) since all types implement that
// func handleValueType(v interface{}) string {
// 	switch v := v.(type) {
// 	case string:
// 		return "`" + v + "`"
// 	case int:
// 		return v
// 	}
// }
