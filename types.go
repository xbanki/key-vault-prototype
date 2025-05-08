package main

const (
	// Pin code policies
	PolicyTypePinForeign PolicyType = iota
	PolicyTypePinEveryone

	// Password policies
	PolicyTypePasswordForeign
	PolicyTypePasswordEveryone
)

type PolicyType = uint8

type PolicyMethods interface {
	Execute(pa *Password, po *Policy) (bool, error)
	Discriminate(pa *Password, po *Policy) bool
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

type Policy struct {
	Type          PolicyType `json:"type"`
	PolicyMethods `json:"-"`
}

type Group struct {
	Passwords []Password `json:"passwords"`
	Policies  []Policy   `json:"policies"`
	Name      string     `json:"name"`
}
