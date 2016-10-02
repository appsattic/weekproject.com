package internal

import "time"

type User struct {
	Name     string
	Logins   int
	Inserted time.Time
	Updated  time.Time
}

type Project struct {
	Name     string
	Title    string
	Owner    string
	Inserted time.Time
	Updated  time.Time
}
