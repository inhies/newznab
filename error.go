package newznab

import "encoding/xml"

// An API error. We get an XML error even if we request a JSON response...
type Error struct {
	Code        int    `xml:"code,attr"`
	Description string `xml:"description,attr"`
	//XMLName     string `xml:"error"`
}

var (
	ErrBadCreds     = NewError(100, "Incorrect user credentials")
	ErrSuspended    = NewError(101, "Account suspended")
	ErrBadPrivs     = NewError(102, "Insufficient privileges/not authorized")
	ErrRegDenied    = NewError(103, "Registration denied")
	ErrRegClosed    = NewError(104, "Registrations are closed")
	ErrEmailTaken   = NewError(105, "Invalid registration (Email Address Taken)")
	ErrEmailFormat  = NewError(106, "Invalid registration (Email Address Bad Format)")
	ErrRegFailed    = NewError(107, "Registration Failed (Data error)")
	ErrMissingParam = NewError(200, "Missing parameter")
	ErrBadParam     = NewError(201, "Incorrect parameter")
	ErrNoSuchFunc   = NewError(202, "No such function. (Function not defined in this specification)")
	ErrFuncUnavail  = NewError(203, "Function not available. (Optional function is not implemented)")
	ErrNoItem       = NewError(300, "No such item")
	ErrUnknown      = NewError(900, "Unknown error")
	ErrAPIDisabled  = NewError(910, "API Disabled")
)

func NewError(code int, description string) Error {
	return Error{Code: code, Description: description}

}

// Return the
func (e *Error) AsXML() ([]byte, error) {
	output, err := xml.Marshal(e)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (e *Error) String() string {
	return e.Description
}

// Check to see if the API data unmarshals to a valid Error struct. Even if we
// request a JSON response, an error will be an XML response.
func CheckForError(data []byte) (*Error, error) {
	var e Error
	err := xml.Unmarshal(data, &e)
	if err != nil {
		return nil, err
	}
	if e.Code > 0 {
		return &e, nil
	}
	return nil, nil

}
