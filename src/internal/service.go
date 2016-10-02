package internal

type StoreService interface {
	GetUser(name string) (*User, error)
	InsUser(user User) error
	// SaveUser(user User) error

	GetProject(name string) (*Project, error)
	InsProject(project Project) error
	// SaveProject(project Project) error

	Close()
}
