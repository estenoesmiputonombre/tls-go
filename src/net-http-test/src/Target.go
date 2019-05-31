package test

import (
	"fmt"
	"strings"
)

const (
	//ProtocolDefault ... Protocol to use by default in Target
	ProtocolDefault = "HTTP/1.0"
	//DomainDefault ... Domain to use by default in Target
	DomainDefault = "example.com"
	//PathDefault ... Path to attack by default
	PathDefault = "/"
	//PortDefault ... Port to attack by default
	PortDefault = 80
)

//Target ... represents the URL used by Request
type Target struct {
	proto Proto
	domain,
	path string
	params map[string]string
	port   int
}

//NewTarget ... Returns an object of the URL to attack
func NewTarget(protocol, domain, path string, params map[string]string, port int) (Target, error) {
	proto, err := NewProto(protocol)
	if err != nil {
		return Target{}, nil
	}
	if IsEmpty(domain) {
		domain = DomainDefault
	}
	if IsEmpty(path) {
		path = PathDefault
	}
	if !IsPort(port) {
		port = PortDefault
	}
	return Target{
		proto:  proto,
		domain: domain,
		path:   path,
		params: params,
		port:   port,
	}, nil
}

//NewDefaultTarget ... Return an object of a target with values by default
func NewDefaultTarget() Target {
	return Target{
		proto:  NewDefaultProto(),
		domain: DomainDefault,
		path:   PathDefault,
		params: make(map[string]string, 0),
		port:   PortDefault,
	}
}

func (t *Target) String() string {
	if len(t.params) == 0 {
		return fmt.Sprintf("%s://%s:%d/%s", t.proto.String(), t.domain, t.port, t.path)
	}
	return fmt.Sprintf("%s://%s:%d/%s%s", t.proto.String(), t.domain, t.port, t.path, t.AddParams()) // {protocol}://{domain}:{port}/{path}/?{queries}
}

//AddParams ... Add params to the main path
func (t *Target) AddParams() string {
	var str = "?"
	for key, value := range t.params {
		str = fmt.Sprintf("%s%s%s&", str, key, value)
	}
	return str[:len(str)-1]
}

//IsPort ... Return if the int value(port) is between 0 and 65535
func IsPort(port int) bool {
	return port >= 0 && port <= 65535
}

//IsEmpty ... Return a bool value in case of empty string
func IsEmpty(str string) bool {
	return strings.Trim(str, " ") == "" || len(strings.Trim(str, " ")) == 0
}
