package main

func CreatePassword(username string, password string, domain string) *Password {
	r := new(Password)
	r.Policies = make([]Policy, 0)
	r.Username = username
	r.Password = password
	r.Domain = domain
	return r
}

func (pw *Password) AddPolicies(policies ...Policy) {
	for _, policy := range policies {
		pw.Policies = append(pw.Policies, policy)
	}
}
