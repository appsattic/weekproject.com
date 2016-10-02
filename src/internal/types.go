package internal

import "time"

type User struct {
	Name     string
	Logins   int
	Inserted time.Time
	Updated  time.Time
}

type Project struct {
	Username string
	Name     string
	Title    string
	Inserted time.Time
	Updated  time.Time
}
