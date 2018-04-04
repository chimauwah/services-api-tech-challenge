package model

import "time"

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
	ID                    int       `json:"id"`
	EnterTs               time.Time `json:"enterts"`
	LastChangeTs          time.Time `json:"lastchangets"`
	Active                bool      `json:"active"`
	FirstName             string    `json:"firstname"`
	LastName              string    `json:"lastname"`
	Address1              string    `json:"address1"`
	Address2              string    `json:"address2"`
	City                  string    `json:"city"`
	State                 string    `json:"state"`
	Zip                   string    `json:"zip"`
	CellPhone             string    `json:"cellphone"`
	HomePhone             string    `json:"homephone"`
	Picture               []byte    `json:"picture"`
	Title                 string    `json:"title"`
	Role                  int       `json:"role"`
	IPPhone               string    `json:"ipphone"`
	Samaccountname        string    `json:"samaccountname"`
	Mail                  string    `json:"mail"`
	PrimaryPa             string    `json:"primarypa"`
	SecondaryPa           string    `json:"secondarypa"`
	Office                string    `json:"office"`
	ManagerDn             string    `json:"managerdn"`
	TravelPref            string    `json:"travelpref"`
	ManagerSamaccountname string    `json:"managersamaccountname"`
	LastHash              int       `json:"lasthash"`
	ImageHash             string    `json:"imagehash"`
	NickName              string    `json:"nickname"`
	ClientLoc             string    `json:"clientloc"`
}

// Device struct
type Device struct {
	ID             int       `json:"id"`
	EnterTs        time.Time `json:"enterts"`
	LastChangeTs   time.Time `json:"lastchangets"`
	Active         bool      `json:"active"`
	DeviceType     string    `json:"devicetype"`
	NotificationID string    `json:"notifcationid"`
	DeviceID       string    `json:"deviceId"`
	EmployeeID     int       `json:"employeeid"`
}

// CoreSkill struct
type CoreSkill struct {
	ID           int       `json:"id"`
	EnterTs      time.Time `json:"enterts"`
	LastChangeTs time.Time `json:"lastchangets"`
	Active       bool      `json:"active"`
	Skill        string    `json:"skill"`
	Sequence     int       `json:"sequence"`
	EmployeeID   int       `json:"employeeid"`
	Proficiency  string    `json:"proficiency"`
}

// Checkin struct
type Checkin struct {
	ID             int       `json:"id"`
	EnterTs        time.Time `json:"enterts"`
	LastChangeTs   time.Time `json:"lastchangets"`
	Active         bool      `json:"active"`
	GeoID          string    `json:"geoid"`
	LAT            float32   `json:"lat"`
	LNG            float32   `json:"lng"`
	Name           string    `json:"name"`
	Distance       float32   `json:"distance"`
	NotifRequested int       `json:"notifrequested"`
	NotiExpiration int       `json:"notifrequsted"`
	EmployeeID     int       `json:"employeeid"`
}

// PracticeArea struct
type PracticeArea struct {
	ID           int       `json:"id"`
	EnterTs      time.Time `json:"enterts"`
	LastChangeTs time.Time `json:"lastchangets"`
	Active       bool      `json:"active"`
	Code         string    `json:"code"`
	Description  string    `json:"description"`
	PrimaryOnly  bool      `json:"primaryonly"`
	Sequence     int       `json:"sequence"`
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
