package service

import (
	"errors"
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

//TODO: move database refs out of config struct to its own struct

/* Note: see article on escape analysis: references are passed into functions, not out.
http://www.agardner.me/golang/garbage/collection/gc/escape/analysis/2015/10/18/go-escape-analysis.html
*/

//SessionConfig -- keep track of all the config data
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
func GetDatabaseConnection(mgr *SessionManager) (err error) {
	//Note: pass globalSessionManager as the argument to this function
	if mgr == nil {
		err = errors.New("GetDatabaseConnection called with nil SessionManager")
		return
	}
	mongoSession, err := mgo.Dial(mgr.sessionConfig.mongoHost)
	if err != nil {
		fmt.Printf("Could not open mongo database session: %s", err.Error())
		return err
	}
	mgr.sessionConfig.mongoSession = mongoSession
	mgr.sessionConfig.mongoSession.SetMode(mgo.Monotonic, true)

	mongoCollection := mgr.sessionConfig.mongoSession.DB(mgr.sessionConfig.mongoDatabase).C(mgr.sessionConfig.mongoCollectionName)
	mgr.sessionConfig.mongoCollection = mongoCollection
	return
}

//SessionInit - create a new session record in MongoDB
func SessionInit(mgr *SessionManager, session *Session) (err error) {
	if mgr == nil {
		err = errors.New("SessionInit called with nil SessionManager")
		return
	}
	if session == nil {
		err = errors.New("SessionInit called with nil Session")
		return
	}
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	//TODO: store the session in mongodb
	return nil // TODO: return a session
}

func SessionRead(mgr *SessionManager, session *Session) (err error) {
	if mgr == nil {
		err = errors.New("SessionRead called with nil SessionManager")
		return
	}
	if session == nil {
		err = errors.New("SessionRead called with nil Session")
		return
	}
	//TODO: retrieve the session from mongodb if it's there

	return nil
}

func SessionDestroy(mgr *SessionManager, session *Session) (err error) {
	if mgr == nil {
		err = errors.New("SessionDestroy called with nil SessionManager")
		return
	}
	if session == nil {
		err = errors.New("SessionDestroy called with nil Session")
		return
	}
	// TODO: delete from mongodb

	return nil
}

func SessionUpdate(mgr *SessionManager, session *Session) (err error) {
	if mgr == nil {
		err = errors.New("SessionUpdate called with nil SessionManager")
		return
	}
	if session == nil {
		err = errors.New("SessionUpdate called with nil Session")
		return
	}
	//TODO: figure out if we still need this lock for anything
	// mgr.lock.Lock()
	// defer mgr.lock.Unlock()
	// TODO: update in mongodb

	return nil
}

func init() {

}
