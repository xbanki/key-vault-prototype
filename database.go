package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

func ensureFileExistence(fpath string) (*os.File, error) {
	if _, err := os.Stat(fpath); err != nil {
		return os.Create(fpath)
	}
	return os.OpenFile(fpath, os.O_RDONLY|os.O_WRONLY, 0644)
}

func hydrateDatabaseFromDisk(db *Database) error {
	file, err := ensureFileExistence(db.Options.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	contents := bytes.NewBuffer(nil)
	if _, err := io.Copy(contents, file); err != nil {
		return err
	}
	// Do JSON decoding here
	return nil
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
		hydrateDatabaseFromDisk(db)
	}
	return db
}

func (db *Database) Read() error {
	return hydrateDatabaseFromDisk(db)
}

func (db *Database) Write() error {
	return writeHydrationData(db)
}
