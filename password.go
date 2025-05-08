package main

import "github.com/buger/jsonparser"

func CreatePassword(username string, password string, domain string) *Password {
	r := new(Password)
	r.Policies = make([]Policy, 0)
	r.Username = username
	r.Password = password
	r.Domain = domain
	return r
}

func CreatePasswordFromJSON(json []byte) (*Password, error) {
	username, err := jsonparser.GetString(json, "username")
	if err != nil {
		return nil, err
	}
	password, err := jsonparser.GetString(json, "password")
	if err != nil {
		return nil, err
	}
	domain, err := jsonparser.GetString(json, "domain")
	if err != nil {
		return nil, err
	}
	return CreatePassword(username, password, domain), nil
}

func (pw *Password) AddPolicies(policies ...Policy) {
	for _, policy := range policies {
		pw.Policies = append(pw.Policies, policy)
	}
}
