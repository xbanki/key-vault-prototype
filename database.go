package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/buger/jsonparser"
)

func ensureFileExistence(fpath string) (*os.File, error) {
	if _, err := os.Stat(fpath); err != nil {
		return os.Create(fpath)
	}
	return os.OpenFile(fpath, os.O_RDWR, 0644)
}

func parsePolicyJSON(policies *[]Policy, cberr *error) func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
	return func(value []byte, vtype jsonparser.ValueType, _ int, err error) {
		if err != nil || vtype != jsonparser.Object {
			*cberr = err
			return
		}
		policy, err := CreatePolicyFromJSON(value)
		if err != nil {
			*cberr = err
			return
		}
		*policies = append(*policies, *policy)
	}
}

func parsePasswordJSON(passwords *[]Password, cberr *error) func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
	return func(value []byte, vtype jsonparser.ValueType, _ int, err error) {
		if err != nil || vtype != jsonparser.Object {
			*cberr = err
			return
		}
		password, err := CreatePasswordFromJSON(value)
		if err != nil {
			*cberr = err
			return
		}
		jsonparser.ArrayEach(value, parsePolicyJSON(&password.Policies, cberr), "policies")
		*passwords = append(*passwords, *password)
	}
}

func hydrateDatabaseFromDisk(db *Database) error {
	var cberr error = nil
	file, err := ensureFileExistence(db.Options.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	contents := bytes.NewBuffer(nil)
	if _, err := io.Copy(contents, file); err != nil {
		return err
	}
	jsonparser.ArrayEach(contents.Bytes(), parsePasswordJSON(&db.Passwords, &cberr), "passwords")
	if cberr != nil {
		return cberr
	}
	jsonparser.ArrayEach(contents.Bytes(), func(value []byte, vtype jsonparser.ValueType, _ int, err error) {
		if err != nil || vtype != jsonparser.Object {
			cberr = err
			return
		}
		group, err := CreateGroupFromJSON(value)
		if err != nil {
			cberr = err
			return
		}
		jsonparser.ArrayEach(value, parsePasswordJSON(&group.Passwords, &cberr), "passwords")
		jsonparser.ArrayEach(value, parsePolicyJSON(&group.Policies, &cberr), "policies")
		db.Groups = append(db.Groups, *group)
	}, "groups")
	return cberr
}

func writeHydrationData(db *Database) error {
	file, err := ensureFileExistence(db.Options.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	contents, err := json.Marshal(db)
	if err != nil {
		return err
	}
	if _, err := file.Write(contents); err != nil {
		return err
	}
	return nil
}

func WithFilepath(fpath string) DBOptFunc {
	return func(db *Database) {
		db.Options.FilePath = fpath
	}
}

func WithHydrate(value bool) DBOptFunc {
	return func(db *Database) {
		db.Options.Hydrate = value
	}
}

func CreateDatabase(opts ...DBOptFunc) *Database {
	db := new(Database)
	db.Passwords = make([]Password, 0)
	db.Options = *new(DatabaseOptions)
	db.Groups = make([]Group, 0)
	for _, opt := range opts {
		opt(db)
	}
	if db.Options.Hydrate && len(db.Options.FilePath) != 0 {
		if err := hydrateDatabaseFromDisk(db); err != nil {
			println()
		}
	}
	return db
}

func (db *Database) Read() error {
	return hydrateDatabaseFromDisk(db)
}

func (db *Database) Write() error {
	return writeHydrationData(db)
}
