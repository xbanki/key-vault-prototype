package main

func hydrateDatabaseFromDisk(db *Database) error {
	return nil
}

func writeHydrationData(db *Database) error {
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
