package test

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//Proto ... Struct that handles procotol
type Proto struct {
	proto string
	protoMajor,
	protoMinor int
}

//ProtoError ... Represents an HTTP Proto error
type ProtoError struct {
	ErrorString string
}

const (
	//ProtoMinor ... The minimum version of the protocol HTTP
	ProtoMinor = 0
	//ProtoMajor ... The maximum version of the protocol HTTP
	ProtoMajor = 1
	//ProtoNameDefault ... The protocol name by default
	ProtoNameDefault = "HTTP"
)

var (
	//ErrProtoType ... Error when a bad protocol is passed by params.
	ErrProtoType = &ProtoError{ErrorString: "Error passing http protocol by params"}
	//ErrLength ... Error parsing length of the string. Ex: "    "
	ErrLength = errors.New("Error parsing length of the string")
)

var (
	//ProtoRegex ... The regex to test the protocol and his version
	ProtoRegex = regexp.MustCompile(`^HTTP\/((0\.[1-9])|1\.[0-1])$`)
	//ProtoNumbers ... The regex to get the numbers of the Proto version
	ProtoNumbers = regexp.MustCompile(`[0-9]`)
)

//NewProto ... Creates a new protocol
func NewProto(proto string) (Proto, error) {
	major, minor, err := ParseHTTPVersion(proto)
	if err != nil {
		return Proto{}, err
	}
	return Proto{
		proto:      ProtoNameDefault,
		protoMajor: major,
		protoMinor: minor,
	}, nil
}

//NewDefaultProto ... Creates a new protocol with defaults values
func NewDefaultProto() Proto {
	return Proto{
		proto:      ProtoNameDefault,
		protoMajor: ProtoMajor,
		protoMinor: ProtoMinor,
	}
}

//String ... Returns a readable form of the properties all in one
func (p *Proto) String() string {
	return fmt.Sprintf("%s/%d.%d", p.proto, p.protoMajor, p.protoMinor)
}

//ParseHTTPVersion ... Parse the version of the protocol returning the major, minor and if it formats well
func ParseHTTPVersion(ver string) (int, int, error) {
	if strings.Trim(ver, " ") == "" || len(strings.Trim(ver, " ")) == 0 {
		return -1, -1, ErrLength
	}
	if !ProtoRegex.MatchString(ver) {
		return -1, -1, ErrProtoType
	}
	arr := ProtoNumbers.FindAllString(ver, -1)
	if arr == nil || len(arr) == 0 {
		return -1, -1, ErrProtoType
	}
	m, err := strconv.Atoi(arr[0])
	if err != nil {
		return -1, -1, ErrProtoType
	}
	n, err := strconv.Atoi(arr[1])
	if err != nil {
		return -1, -1, ErrProtoType
	}
	return m, n, nil
}

func (p ProtoError) Error() string {
	return p.ErrorString
}
