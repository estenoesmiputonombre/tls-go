package test

var (
	//ErrProtoType ... Error when a bad protocol is passed by params.
	ErrProtoType = &ProtocolError{ErrorString: "Error passing http protocol by params"}
	//ErrProtoNotFound ... Error when the protocol exists, but the func dont handle it.
	ErrProtoNotFound = &ProtocolError{ErrorString: "Error passing http protocol by params"}
)

//Target ... represents the URL used by Request
type Target struct {
	protocol,
	domain,
	path string
	params map[string]string
	port   int
}

//NewTarget ... Returns an object of the URL to attack
func NewTarget(protocol, domain, path string, params map[string]string, port int) (Target, error) {

	return Target{}, nil
}

//ProtocolError ... Represents an HTTP protocol error
type ProtocolError struct {
	ErrorString string
}

func (p ProtocolError) Error() string {
	return p.ErrorString
}
