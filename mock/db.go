package main

// db.go
//go:generate mockgen -source=db.go -destination=db_mock.go -package=main
type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}