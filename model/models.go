package model

import (
	"database/sql"
	"time"
)

// An EmployeeCreateEmployeeReq request model
//
// This is used for operations that require an Employee as body of the request
//
// swagger:parameters createEmployee
type EmployeeCreateEmployeeReq struct {
	// The Employee to create
	//
	// in: body
	// required: true
	Employee *Employee `json:"employee"`
}

// An EmployeeUpdateEmployeeReq request model
//
// This is used for operations that require an Employee as body of the request
//
// swagger:parameters updateEmployee
type EmployeeUpdateEmployeeReq struct {
	// The Employee date with which to overwrite existing Employee
	//
	// in: body
	// required: true
	Employee *Employee `json:"employee"`
}

// An EmployeeResponse response model
//
// This is used for returning a response with status code and list of retrieved employees as body
//
// swagger:response employeeResponse
type EmployeeResponse struct {
	// in: body
	Body struct {
		// HTTP status code 200 - Status OK
		Code int `json:"code"`
		// Employee model
		Payload []*Employee `json:"employees"`
	}
}

// An Employee model
//
// This is used for operations that involve an Employee
// swagger:parameters
type Employee struct {
	ID                    int            `json:"id"`
	EnterTs               time.Time      `json:"enterts"`
	LastChangeTs          time.Time      `json:"lastchangets"`
	Active                bool           `json:"active"`
	FirstName             sql.NullString `json:"firstname"`
	LastName              sql.NullString `json:"lastname"`
	Address1              sql.NullString `json:"address1"`
	Address2              sql.NullString `json:"address2"`
	City                  sql.NullString `json:"city"`
	State                 sql.NullString `json:"state"`
	Zip                   sql.NullString `json:"zip"`
	CellPhone             sql.NullString `json:"cellphone"`
	HomePhone             sql.NullString `json:"homephone"`
	Picture               []byte         `json:"picture"`
	Title                 sql.NullString `json:"title"`
	Role                  sql.NullInt64  `json:"role"`
	IPPhone               sql.NullString `json:"ipphone"`
	Samaccountname        sql.NullString `json:"samaccountname"`
	Mail                  sql.NullString `json:"mail"`
	PrimaryPa             sql.NullString `json:"primarypa"`
	SecondaryPa           sql.NullString `json:"secondarypa"`
	Office                sql.NullString `json:"office"`
	ManagerDn             sql.NullString `json:"managerdn"`
	TravelPref            sql.NullString `json:"travelpref"`
	ManagerSamaccountname sql.NullString `json:"managersamaccountname"`
	LastHash              sql.NullInt64  `json:"lasthash"`
	ImageHash             sql.NullString `json:"imagehash"`
	NickName              sql.NullString `json:"nickname"`
	ClientLoc             sql.NullString `json:"clientloc"`
}

// Device struct
type Device struct {
	ID             int            `json:"id"`
	EnterTs        time.Time      `json:"enterts"`
	LastChangeTs   time.Time      `json:"lastchangets"`
	Active         bool           `json:"active"`
	DeviceType     sql.NullString `json:"devicetype"`
	NotificationID sql.NullString `json:"notifcationid"`
	DeviceID       sql.NullString `json:"deviceId"`
	EmployeeID     int            `json:"employeeid"`
}

// CoreSkill struct
type CoreSkill struct {
	ID           int            `json:"id"`
	EnterTs      time.Time      `json:"enterts"`
	LastChangeTs time.Time      `json:"lastchangets"`
	Active       bool           `json:"active"`
	Skill        sql.NullString `json:"skill"`
	Sequence     sql.NullInt64  `json:"sequence"`
	EmployeeID   int            `json:"employeeid"`
	Proficiency  sql.NullString `json:"proficiency"`
}

// Checkin struct
type Checkin struct {
	ID             int            `json:"id"`
	EnterTs        time.Time      `json:"enterts"`
	LastChangeTs   time.Time      `json:"lastchangets"`
	Active         bool           `json:"active"`
	GeoID          sql.NullString `json:"geoid"`
	LAT            float32        `json:"lat"`
	LNG            float32        `json:"lng"`
	Name           sql.NullString `json:"name"`
	Distance       float32        `json:"distance"`
	NotifRequested sql.NullInt64  `json:"notifrequested"`
	NotiExpiration sql.NullInt64  `json:"notifrequsted"`
	EmployeeID     int            `json:"employeeid"`
}

// PracticeArea struct
type PracticeArea struct {
	ID           int            `json:"id"`
	EnterTs      time.Time      `json:"enterts"`
	LastChangeTs time.Time      `json:"lastchangets"`
	Active       bool           `json:"active"`
	Code         sql.NullString `json:"code"`
	Description  sql.NullString `json:"description"`
	PrimaryOnly  bool           `json:"primaryonly"`
	Sequence     sql.NullInt64  `json:"sequence"`
}

// Success response
// swagger:response ok
type swaggScsResp struct {
	// in:body
	Body struct {
		// HTTP Status Code 200
		Code int `json:"code"`
	}
}

// Boolean response
// swagger:response bool
type swaggBoolResp struct {
	// in:body
	Body struct {
		// HTTP Status Code 200
		Code int `json:"code"`
		// Boolean true/false
		Data bool `json:"data"`
	}
}

// Error Bad Request
// swagger:response badReq
type swaggErrBadReq struct {
	// in:body
	Body struct {
		// HTTP status code 400 - Status Bad Request
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// Error Forbidden
// swagger:response forbidden
type swaggErrForbidden struct {
	// in:body
	Body struct {
		// HTTP status code 403 - Forbidden
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// Error Not Found
// swagger:response notFound
type swaggErrNotFound struct {
	// in:body
	Body struct {
		// HTTP status code 404 - Not Found
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// Error Conflict
// swagger:response conflict
type swaggErrConflict struct {
	// in:body
	Body struct {
		// HTTP status code 409 - Conflict
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// Error Interval Server
// swagger:response internal
type swaggErrInternal struct {
	// in:body
	Body struct {
		// HTTP status code 500 - Internal server error
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// A ValidationError is a swagger response to represent error
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}
