package internal

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrUserNameExists    = errors.New("Username already exists")
	ErrProjectNameExists = errors.New("Project name already exists")
)

type MongoDbStore struct {
	url     string
	session *mgo.Session
}

func ensureIndexes(session *mgo.Session) error {
	var err error

	// users
	userIndex := mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}
	users := session.DB("weekproject").C("users")
	err = users.EnsureIndex(userIndex)
	if err != nil {
		return err
	}

	// projects
	projectIndex := mgo.Index{
		Key:    []string{"username", "name"},
		Unique: true,
	}
	projects := session.DB("weekproject").C("projects")
	err = projects.EnsureIndex(projectIndex)
	if err != nil {
		return err
	}

	return nil
}

func NewMongoDbStore(url string) (*MongoDbStore, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	// get the collections and ensure unique indexes on name
	ensureIndexes(session)
	if err != nil {
		return nil, err
	}

	// create the store and return it
	store := MongoDbStore{
		url:     url,
		session: session,
	}
	return &store, nil
}

func (s *MongoDbStore) GetUser(name string) (*User, error) {
	users := s.session.DB("weekproject").C("users")

	user := User{}
	err := users.Find(bson.M{"name": name}).One(&user)
	if err != nil {
		return nil, nil
	}

	return &user, nil
}

func (s *MongoDbStore) InsUser(user User) error {
	users := s.session.DB("weekproject").C("users")
	err := users.Insert(user)
	if mgo.IsDup(err) {
		return ErrUserNameExists
	}
	return err
}

func (s *MongoDbStore) GetProject(name string) (*Project, error) {
	projects := s.session.DB("weekproject").C("projects")

	project := Project{}
	err := projects.Find(bson.M{"name": name}).One(&project)
	if err != nil {
		return nil, nil
	}

	return &project, nil
}

func (s *MongoDbStore) InsProject(project Project) error {
	projects := s.session.DB("weekproject").C("projects")
	err := projects.Insert(project)
	if mgo.IsDup(err) {
		return ErrProjectNameExists
	}
	return err
}

func (s *MongoDbStore) Close() {
	s.session.Close()
}
