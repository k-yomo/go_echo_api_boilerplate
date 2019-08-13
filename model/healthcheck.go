package model

// Ping checks if db connection is valid
func Ping(db Queryer) error {
	err := db.Ping()
	return err
}

