package main

import (
	"bytes"
	"crypto/md5"
	"errors"

	"github.com/buger/jsonparser"
)

var (
	ErrorWrongPassword = errors.New("Password does not match")
	ErrorWrongPin      = errors.New("Pin does not match")
)

func (p *PolicyPassword) Execute(input []byte, password *Password, policy *Policy) (bool, error) {
	hash := md5.New().Sum(input)
	if res := bytes.Equal(hash, policy.Discriminator); res != true {
		return false, ErrorWrongPassword
	}
	return true, nil
}

func (p *PolicyPin) Execute(input []byte, password *Password, policy *Policy) (bool, error) {
	hash := md5.New().Sum(input)
	if res := bytes.Equal(hash, policy.Discriminator); res != true {
		return false, ErrorWrongPin
	}
	return true, nil
}

func CreatePasswordPolicy(password []byte) *Policy {
	r := new(PolicyPassword)
	hash := md5.New()
	return CreatePolicy(PolicyTypePassword, hash.Sum(password), r)
}

func CreatePinPolicy(pin []byte) *Policy {
	r := new(PolicyPin)
	hash := md5.New()
	return CreatePolicy(PolicyTypePin, hash.Sum(pin), r)
}

func CreatePolicyFromJSON(json []byte) (*Policy, error) {
	discriminator, err := jsonparser.GetString(json, "discriminator")
	if err != nil {
		return nil, err
	}
	ptype, err := jsonparser.GetInt(json, "type")
	if err != nil {
		return nil, err
	}
	switch PolicyType(ptype) {
	case PolicyTypePassword:
		return CreatePasswordPolicy([]byte(discriminator)), nil
	case PolicyTypePin:
		return CreatePinPolicy([]byte(discriminator)), nil
	default:
		return nil, errors.New("Invalid policy type")
	}
}

func CreatePolicy(ptype PolicyType, discriminator []byte, pmethods PolicyMethods) *Policy {
	r := new(Policy)
	r.Discriminator = discriminator
	r.PolicyMethods = pmethods
	r.Type = ptype
	return r
}
