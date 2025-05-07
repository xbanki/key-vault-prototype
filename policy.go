package main

func CreatePolicy(ptype PolicyType, pmethods PolicyMethods) *Policy {
	r := new(Policy)
	r.PolicyMethods = pmethods
	r.Type = ptype
	return r
}
