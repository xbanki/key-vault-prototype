package main

const (
	// Unknown policy type
	PolicyTypeUnknown = PolicyType(iota)

	// Passkey policy types
	PolicyTypePassword
	PolicyTypePin
)

type PolicyType uint8

type PolicyMethods interface {
	Execute(input []byte, password *Password, policy *Policy) (bool, error)
}

type DBOptFunc func(db *Database)

type DatabaseOptions struct {
	FilePath string
	Hydrate  bool
}

type Database struct {
	Options   DatabaseOptions `json:"-"`
	Passwords []Password      `json:"passwords"`
	Groups    []Group         `json:"groups"`
}

type Password struct {
	Policies []Policy `json:"policies"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Domain   string   `json:"domain"`
}

type PolicyPassword struct{}
type PolicyPin struct{}

type Policy struct {
	Discriminator []byte     `json:"discriminator"`
	Type          PolicyType `json:"type"`
	PolicyMethods `json:"-"`
}

type Group struct {
	Passwords []Password `json:"passwords"`
	Policies  []Policy   `json:"policies"`
	Name      string     `json:"name"`
}
