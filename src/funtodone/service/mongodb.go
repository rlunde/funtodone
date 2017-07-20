package service

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

type SessionConfig struct {
	mongoSession        *mgo.Session
	mongoCollection     *mgo.Collection
	mongoHost           string
	mongoDatabase       string
	mongoCollectionName string
}

//TODO: read the host and database from a config file
func GetSessionConfig(mgr *SessionManager) {
	sessionConfig := SessionConfig{
		mongoHost:           "127.0.0.1",
		mongoDatabase:       "funtodone",
		mongoCollectionName: "sessions",
	}
	mgr.sessionConfig = sessionConfig
}

//GetDatabaseConnection - open a Mongo database for storing sessions
func GetDatabaseConnection() (mgr *SessionManager, err error) {

	mgr = globalSessionManager
	mongoSession, err := mgo.Dial(mgr.sessionConfig.mongoHost)
	if err != nil {
		fmt.Printf("Could not open mongo database session: %s", err.Error())
		return nil, err
	}
	mgr.sessionConfig.mongoSession = mongoSession
	mgr.sessionConfig.mongoSession.SetMode(mgo.Monotonic, true)

	mongoCollection := mgr.sessionConfig.mongoSession.DB(mgr.sessionConfig.mongoDatabase).C(mgr.sessionConfig.mongoCollectionName)
	mgr.sessionConfig.mongoCollection = mongoCollection
	return
}

func SessionInit(mgr *SessionManager, sid string) (session *Session, err error) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	//TODO: store the session in mongodb
	return nil, nil // TODO: return a session
}

// func (pder *Provider) SessionRead(sid string) (Session, error) {
func SessionRead(mgr *SessionManager, sid string) (session *Session, err error) {
	//TODO: retrieve the session from mongodb if it's there

	return nil, nil
}

func SessionDestroy(mgr *SessionManager, sid string) error {
	// TODO: delete from mongodb

	return nil
}

func SessionUpdate(mgr *SessionManager, sid string) error {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	// TODO: update in mongodb

	return nil
}

func init() {

}
