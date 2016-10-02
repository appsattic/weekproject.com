package internal

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDbStore struct {
	url     string
	session *mgo.Session
}

func NewMongoDbStore(url string) (*MongoDbStore, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

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
	return users.Insert(user)
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
	return projects.Insert(project)
}

func (s *MongoDbStore) Close() {
	s.session.Close()
}
