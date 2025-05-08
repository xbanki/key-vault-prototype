package main

import "github.com/buger/jsonparser"

func CreateGroup(name string) *Group {
	r := new(Group)
	r.Passwords = make([]Password, 0)
	r.Policies = make([]Policy, 0)
	r.Name = name
	return r
}

func CreateGroupFromJSON(json []byte) (*Group, error) {
	name, err := jsonparser.GetString(json, "name")
	if err != nil {
		return nil, err
	}
	group := CreateGroup(name)
	println(name)
	return group, nil
}
